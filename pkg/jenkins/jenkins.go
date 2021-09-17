package jenkins

import (
	"context"
	"fmt"
	"log"
	"math"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/muroj/gojenkins"
)

type JenkinsAPIClient struct {
	Client    *gojenkins.Jenkins
	Creds     JenkinsCredentials
	Context   context.Context
	DebugMode bool
}

type JenkinsCredentials struct {
	Username string
	APIToken string
}

func NewJenkinsClient(jenkinsURL string, creds JenkinsCredentials, debug bool) *JenkinsAPIClient {
	log.Printf("Initializing Jenkins client")

	var jenkinsClient JenkinsAPIClient
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

func GetBuild(url string, user string, apiToken string, projectUrl string, buildId int64) {

	jenkinsCreds := JenkinsCredentials{
		Username: user,
		APIToken: apiToken,
	}
	jenkinsClient := NewJenkinsClient(url, jenkinsCreds, false)
	buildInfo, err := GetBuildInfo(projectUrl, buildId, jenkinsClient)

	if err != nil {
		log.Fatalf("Unable to retrieve build information: %s", err)
	}

	buildInfo.PrintBuildInfo()
}

type JenkinsBuildInfo struct {
	JobName                string
	BuildId                int64
	ScheduledTimestamp     time.Time
	ScheduledTimeUnix      int64
	ExecutionStartTimeUnix int64
	CompletedTimeUnix      int64
	DurationMs             int64
	ExecutionTimeMs        int64
	AgentHostMachine       string
}

func (bi *JenkinsBuildInfo) PrintBuildInfo() {
	fmt.Printf("Project Name: %s\n", bi.JobName)
	fmt.Printf("  Build ID: #%d\n", bi.BuildId)
	fmt.Printf("  Host: %s\n", bi.AgentHostMachine)
	fmt.Printf("  Scheduled at: %s\n", bi.ScheduledTimestamp.String())
	fmt.Printf("  Began executing at: %s\n", time.Unix(bi.ExecutionStartTimeUnix, 0))
	fmt.Printf("  Ended: %s\n", time.Unix(bi.CompletedTimeUnix, 0))
	fmt.Printf("  Execution Time(s): %d\n", bi.ExecutionTimeMs/int64(1000))
	fmt.Printf("  Total Duration(s): %d\n", bi.DurationMs/int64(1000))
}

func GetBuildInfo(jobURL string, id int64, jc *JenkinsAPIClient) (JenkinsBuildInfo, error) {
	var buildInfo JenkinsBuildInfo

	log.Printf("Retrieving job at URL: \"%s\"", jobURL)
	jobName, path, _ := parseJobURL(jobURL)
	job, err := jc.Client.GetJob(jc.Context, jobName, path...)
	if err != nil {
		return buildInfo, fmt.Errorf("Unable to retreive job at URL \"%s\": %s", jobURL, err)
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
		return buildInfo, fmt.Errorf("Failed to retrieve build: %s", err)
	}

	buildInfo.JobName = job.GetDetails().FullName
	buildInfo.BuildId = build.GetBuildNumber()
	buildInfo.ScheduledTimestamp = build.GetTimestamp()
	buildInfo.DurationMs = int64(math.Round(build.GetDuration()))
	buildInfo.ExecutionTimeMs = build.GetExecutionTimeMs()
	buildInfo.CompletedTimeUnix = buildInfo.ScheduledTimestamp.Unix() + (int64(buildInfo.DurationMs) / 1000)
	buildInfo.ExecutionStartTimeUnix = buildInfo.ScheduledTimestamp.Unix() + (int64(buildInfo.DurationMs-buildInfo.ExecutionTimeMs) / 1000)
	buildInfo.AgentHostMachine, err = findBuildHostMachineName(build, jc)

	if err != nil {
		return buildInfo, fmt.Errorf("Could not determine build agent name.")
	}

	// buildJson, err := json.Marshal(build.Raw)
	// fmt.Printf("%s", buildJson)

	return buildInfo, nil
}

func GetVersion(url string, user string, apiToken string) {

	creds := JenkinsCredentials{
		Username: user,
		APIToken: apiToken,
	}
	c := NewJenkinsClient(url, creds, false)

	fmt.Printf(c.Client.Version)
}

/*
	Returns a string indicating the hostname of the machine where this build ran.
*/
func findBuildHostMachineName(build *gojenkins.Build, jc *JenkinsAPIClient) (string, error) {
	buildLog := build.GetConsoleOutput(jc.Context)
	r, _ := regexp.Compile(`Running on \b([\w]+\b)`)
	m := r.FindStringSubmatch(buildLog)
	if m == nil {
		return "", fmt.Errorf("Unable to determine host node: \"Running on <nodeName>\" line not found in build log.")
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
