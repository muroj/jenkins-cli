package jenkins

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/muroj/gojenkins"
	"golang.org/x/mod/semver"
)

type APIClient struct {
	Client    *gojenkins.Jenkins
	Creds     Credentials
	Context   context.Context
	DebugMode bool
}

type Credentials struct {
	Username string
	APIToken string
}

// NewJenkinsClient attempts to connect to the Jenkins instance at the specified URL by using the provided credentials
// if successful, an APIClient is returned, which can be used to make additional API calls
func NewJenkinsClient(jenkinsURL string, creds Credentials, debug bool) *APIClient {
	var jenkinsClient APIClient
	var ctx context.Context

	if debug {
		ctx = context.WithValue(context.Background(), "debug", "debug")
	} else {
		ctx = context.WithValue(context.Background(), "debug", nil)
	}

	jenkins, err := gojenkins.CreateJenkins(nil, jenkinsURL, creds.Username, creds.APIToken).Init(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Jenkins instance at \"%s\"\n %s", jenkinsURL, err)
	}

	jenkinsClient.Client = jenkins
	jenkinsClient.Context = ctx

	return &jenkinsClient
}

func GetBuild(c *APIClient, projectURL string, buildID int64) {
	buildInfo, err := GetBuildInfo(projectURL, buildID, c)

	if err != nil {
		log.Fatalf("Unable to retrieve build information: %s", err)
	}

	buildInfo.PrintBuildInfo()
}

type BuildInfo struct {
	JobName                string
	BuildID                int64
	ScheduledTimestamp     time.Time
	ScheduledTimeUnix      int64
	ExecutionStartTimeUnix int64
	CompletedTimeUnix      int64
	DurationMs             int64
	ExecutionTimeMs        int64
	AgentHostMachine       string
}

func (bi *BuildInfo) PrintBuildInfo() {
	fmt.Printf("Project Name: %s\n", bi.JobName)
	fmt.Printf("  Build ID: #%d\n", bi.BuildID)
	fmt.Printf("  Host: %s\n", bi.AgentHostMachine)
	fmt.Printf("  Scheduled at: %s\n", bi.ScheduledTimestamp.String())
	fmt.Printf("  Began executing at: %s\n", time.Unix(bi.ExecutionStartTimeUnix, 0))
	fmt.Printf("  Ended: %s\n", time.Unix(bi.CompletedTimeUnix, 0))
	fmt.Printf("  Execution Time(s): %d\n", bi.ExecutionTimeMs/int64(1000))
	fmt.Printf("  Total Duration(s): %d\n", bi.DurationMs/int64(1000))
}

func GetBuildInfo(jobURL string, id int64, jc *APIClient) (BuildInfo, error) {
	var buildInfo BuildInfo

	log.Printf("Retrieving job at URL: \"%s\"", jobURL)
	jobName, path, _ := parseJobURL(jobURL)
	job, err := jc.Client.GetJob(jc.Context, jobName, path...)
	if err != nil {
		return buildInfo, fmt.Errorf("unable to retreive job at URL \"%s\": %s", jobURL, err)
	}

	var build *gojenkins.Build
	if id == 0 {
		log.Printf("Retrieving last successful build")
		build, err = job.GetLastSuccessfulBuild(jc.Context)
	} else {
		log.Printf("Retrieving build %d\n", id)
		build, err = job.GetBuild(jc.Context, id)
	}

	if err != nil {
		return buildInfo, fmt.Errorf("failed to retrieve build: %s", err)
	}

	buildInfo.JobName = job.GetDetails().FullName
	buildInfo.BuildID = build.GetBuildNumber()
	buildInfo.ScheduledTimestamp = build.GetTimestamp()
	buildInfo.DurationMs = int64(math.Round(build.GetDuration()))
	buildInfo.ExecutionTimeMs = build.GetExecutionTimeMs()
	buildInfo.CompletedTimeUnix = buildInfo.ScheduledTimestamp.Unix() + (int64(buildInfo.DurationMs) / 1000)
	buildInfo.ExecutionStartTimeUnix = buildInfo.ScheduledTimestamp.Unix() + (int64(buildInfo.DurationMs-buildInfo.ExecutionTimeMs) / 1000)
	buildInfo.AgentHostMachine, err = findBuildHostMachineName(build, jc)

	if err != nil {
		return buildInfo, fmt.Errorf("could not determine build agent name")
	}

	// buildJson, err := json.Marshal(build.Raw)
	// fmt.Printf("%s", buildJson)

	return buildInfo, nil
}

func GetVersion(c *APIClient) {
	fmt.Printf(c.Client.Version)
}

func SafeRestart(c *APIClient) {
	if err := c.Client.SafeRestart(c.Context); err != nil {
		log.Fatalf("restart failed: %s", err)
	}
}

func InstallPlugins(c *APIClient, pluginListJSON string) error {

	type pluginDesc struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	var requestedPlugins []pluginDesc
	err := json.Unmarshal([]byte(pluginListJSON), &requestedPlugins)
	if err != nil {
		log.Fatalf("failed to decode plugin list as JSON: %s", err)
	}

	installed, err := c.Client.GetPlugins(c.Context, 1)
	if err != nil {
		log.Fatalf("failed to retrieve installed plugins from target jenkins server: %s", err)
	}

	for _, p := range requestedPlugins {
		if tp := installed.Contains(p.Name); tp != nil {
			iv := fmt.Sprintf("v%s", tp.Version)
			sv := fmt.Sprintf("v%s", p.Version)

			if semver.Compare(iv, sv) == 0 {
				log.Printf("%s is already at version: %s", p.Name, iv)
			} else if semver.Compare(iv, sv) > 0 {
				log.Printf("more recent version of %s is installed: installed: %s, requested: %s", p.Name, iv, p.Version)
			} else {
				log.Printf("%s needs updating: %s", p.Name, iv)
			}
		} else {
			log.Printf("installing %s:%s", p.Name, p.Version)

			/* InstallPlugin does not indicate whether plugin installation is successful. It doesn't
			   even verify whether the plugin actually exists.  It only POSTs the data and returns.
			   It will return a 500 server error if the plugin version is not valid.
			   The UpdateManager API can be used to track plugin installation job status.

			   Considering I intend to install plugins retrieved from another Jenkins installation,
			   this should be good enough.
			*/
			err := c.Client.InstallPlugin(c.Context, p.Name, p.Version)
			if err != nil {
				log.Fatalf("failed to install plugin %s:%s %s", p.Name, p.Version, err)
			}
		}
	}

	uc, err := c.Client.GetUpdateCenter(c.Context)
	if err != nil {
		log.Fatalf("failed to retrieve update center info: %s", err)
	}

	if uc.RestartRequired() {
		log.Printf("Restart jenkins to finish installing plugins")
	}

	return nil
}

func ListPlugins(c *APIClient) error {
	plugins, err := c.Client.GetPlugins(c.Context, 2)

	if err != nil {
		return fmt.Errorf("unable to list plugins: %s", err)
	}

	pluginsJSON, err := json.Marshal(plugins.Raw.Plugins)

	if err != nil {
		return fmt.Errorf("failed to encode plugin list as JSON: %s", err)
	}

	fmt.Printf("%s", pluginsJSON)

	return nil
}

/*
	Returns a string indicating the hostname of the machine where this build ran.
*/
func findBuildHostMachineName(build *gojenkins.Build, jc *APIClient) (string, error) {
	buildLog := build.GetConsoleOutput(jc.Context)
	r, _ := regexp.Compile(`running on \b([\w]+\b)`)
	m := r.FindStringSubmatch(buildLog)
	if m == nil {
		return "", fmt.Errorf("unable to determine host node: \"Running on <nodeName>\" line not found in build log")
	}

	return m[1], nil
}

/*
	Given a URL for a Jenkins job. Returns the job name and a slice of the parent folder names.
	For example, a job URL of "job/ai-foundation/job/abp-code-scan/job/ghenkins/" will return (ghenkins, [ai-foundation, abp-code-scan]).
	This is the format expected by the golang Jenkins API.
*/
func parseJobURL(jobURL string) (string, []string, error) {
	jobURLTrimmed := strings.TrimRight(strings.TrimSpace(jobURL), "/")
	path, name := path.Split(jobURLTrimmed)
	segments := strings.Split(path, "/")

	parentIds := make([]string, 0)
	for _, s := range segments {
		if s != "job" && s != "" && s != " " {
			parentIds = append(parentIds, s)
		}
	}

	return name, parentIds, nil
}
