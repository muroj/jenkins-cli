package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/muroj/gojenkins"
	"github.ibm.com/jmuro/ghestimator/pkg/instana"
)

/* Command-line arguments */
var (
	jenkinsURL      string
	jenkinsUser     string
	jenkinsAPIToken string
	jobURL          string
	buildNumber     int64
	instanaTenant   string
	instanaUnit     string
	instanaAPIKey   string
	enableDebug     bool
)

func main() {

	jenkinsCreds := JenkinsCredentials{
		Username: jenkinsUser,
		APIToken: jenkinsAPIToken,
	}
	jenkinsClient := newJenkinsClient(jenkinsURL, jenkinsCreds, enableDebug)
	buildInfo, err := getJenkinsBuildInfo(jobURL, buildNumber, jenkinsClient)

	if err != nil {
		log.Fatalf("Failed to retrieve required build information: %s", err)
	}

	buildInfo.printBuildInfo()

	instanaURL := fmt.Sprintf("%s-%s.instana.io", instanaTenant, instanaUnit)
	instanaClient := instana.NewInstanaClient(instanaURL, instana.InstanaCredentials{instanaAPIKey}, enableDebug)

	err = getResourceUsage(&buildInfo, instanaClient)
	if err != nil {
		log.Fatalf("Failed to retrieve required build information: %s", err)
	}
}

/* Parse command-line arguments */
func init() {
	flag.StringVar(&jenkinsURL, "jenkinsURL", "https://ghenkins.bigdatalab.ibm.com/", "URL of the Jenkins host, e.g. \"https://ghenkins.bigdatalab.ibm.com/\"")
	flag.StringVar(&jenkinsUser, "jenkinsUser", "", "Jenkins username")
	flag.StringVar(&jenkinsAPIToken, "jenkinsAPIToken", "", "Jenkins API token")
	flag.StringVar(&jobURL, "jobURL", "", "Path to the Jenkins job to evaluate, e.g.\"job/ai-foundation/job/envctl/job/main/\"")
	flag.Int64Var(&buildNumber, "buildNumber", 0, "ID of the build to evaluate, e.g. 223. Defaults to the last successful build")
	flag.StringVar(&instanaTenant, "instanaTenant", "tron", "Instana Tentant")
	flag.StringVar(&instanaUnit, "instanaUnit", "ibmdataaiwai", "Instana Unit")
	flag.StringVar(&instanaAPIKey, "instanaAPIKey", "", "Instana API key")
	flag.BoolVar(&enableDebug, "debug", false, "Enable debug output")
	flag.Parse()

	if len(jenkinsUser) == 0 {
		log.Fatal("Required parameter not specified: jenkinsUser")
	} else if len(jenkinsAPIToken) == 0 {
		log.Fatal("Required parameter not specified: jenkinsAPIToken")
	} else if len(jobURL) == 0 {
		log.Fatal("Required parameter not specified: jobURL")
	} else if len(instanaAPIKey) == 0 {
		log.Fatal("Required parameter not specified: instanaAPIKey")
	}
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

func (bi *JenkinsBuildInfo) printBuildInfo() {
	fmt.Printf("Project Name: %s\n", bi.JobName)
	fmt.Printf("  ID: #%d\n", bi.BuildId)
	fmt.Printf("  Host: %s\n", bi.AgentHostMachine)
	fmt.Printf("  Scheduled at: %s\n", bi.ScheduledTimestamp.String())
	fmt.Printf("  Began executing at: %s\n", time.Unix(bi.ExecutionStartTimeUnix, 0))
	fmt.Printf("  Ended: %s\n", time.Unix(bi.CompletedTimeUnix, 0))
	fmt.Printf("  Execution Time(s): %d\n", bi.ExecutionTimeMs/int64(1000))
	fmt.Printf("  Total Duration(s): %d\n", bi.DurationMs/int64(1000))

}

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

func newJenkinsClient(jenkinsURL string, creds JenkinsCredentials, debug bool) *JenkinsAPIClient {
	log.Printf("Initializing Jenkins client")

	var jenkinsClient JenkinsAPIClient
	var ctx context.Context

	if enableDebug {
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

func getJenkinsBuildInfo(jobURL string, id int64, jc *JenkinsAPIClient) (JenkinsBuildInfo, error) {
	var buildInfo JenkinsBuildInfo

	log.Printf("Retrieving job at URL: \"%s\"", jobURL)
	jobName, path, _ := parseJobURL(jobURL)
	job, err := jc.Client.GetJob(jc.Context, jobName, path...)
	if err != nil {
		return buildInfo, fmt.Errorf("Failed to retreive job with URL \"%s\"\n %s", jobURL, err)
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

	return buildInfo, nil
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

func getResourceUsage(buildInfo *JenkinsBuildInfo, instanaClient *instana.InstanaAPIClient) error {

	hostMetrics := []string{
		"cpu.used", "load.1min", "memory.used",
	}

	err := instana.GetHostConfiguration(buildInfo.AgentHostMachine, buildInfo.CompletedTimeUnix, buildInfo.ExecutionTimeMs, instanaClient)

	if err != nil {
		return fmt.Errorf("Error retrieving host configuration: %s", err)
	}

	metricsResult, err := instana.GetHostMetrics(buildInfo.AgentHostMachine, buildInfo.CompletedTimeUnix, buildInfo.ExecutionTimeMs, hostMetrics, instanaClient)

	if err != nil {
		return fmt.Errorf("Error retrieving host metrics: %s", err)
	}

	for _, m := range metricsResult {
		m.PrintInstanaHostMetricResult()
	}

	return nil
}
