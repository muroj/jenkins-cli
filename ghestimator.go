package main

import (
	"context"
	"flag"
	"fmt"
	"instana/openapi"
	"log"
	"math"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/antihax/optional"
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
)

func main() {

	jenkinsCreds := JenkinsCredentials{
		Username: jenkinsUser,
		APIToken: jenkinsAPIToken,
	}
	jenkinsClient := newJenkinsClient(jenkinsURL, jenkinsCreds)
	buildInfo, err := getJenkinsBuildInfo(jobURL, buildNumber, jenkinsClient)

	if err != nil {
		log.Fatalf("Failed to retrieve required build information: %s", err)
	}

	buildInfo.printBuildInfo()

	instanaURL := fmt.Sprintf("%s-%s.instana.io", instanaTenant, instanaUnit)
	instanaClient := instana.NewInstanaClient(instanaURL, instana.InstanaCredentials{instanaAPIKey})
	hostMetrics := []string{
		"cpu.used", "load.1min", "memory.used",
	}

	err = getResourceUsage(&buildInfo, hostMetrics, instanaClient)
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

func newJenkinsClient(jenkinsURL string, creds JenkinsCredentials) *JenkinsAPIClient {
	log.Printf("Initializing Jenkins client")

	jenkinsClient := JenkinsAPIClient{
		DebugMode: false,
	}
	ctx := context.WithValue(context.Background(), "debug", nil)
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

func getResourceUsage(buildInfo *JenkinsBuildInfo, hostMetrics []string, ic *instana.InstanaAPIClient) error {

	infraMetricsOpts := openapi.GetInfrastructureMetricsOpts{
		Offline: optional.EmptyBool(),
		GetCombinedMetrics: optional.NewInterface(openapi.GetCombinedMetrics{
			Metrics: hostMetrics,
			Plugin:  "host",
			Query:   fmt.Sprintf("entity.host.name:%s", buildInfo.AgentHostMachine),
			TimeFrame: openapi.TimeFrame{
				To:         buildInfo.ScheduledTimeUnix * 1000, // Instana API requires nanosecond resolution
				WindowSize: buildInfo.DurationMs,
			},
			Rollup: computeRollupValue(buildInfo.CompletedTimeUnix, buildInfo.DurationMs),
		}),
	}

	metricsResult, _, err := ic.Client.InfrastructureMetricsApi.GetInfrastructureMetrics(ic.Context, &infraMetricsOpts)

	if err != nil {
		return fmt.Errorf("Error retrieving host metrics: %s", err)
	}

	fmt.Printf("Metrics:\n")
	for _, m := range metricsResult.Items {
		//fmt.Printf("  label: %s\n  host: %s\n", m.Label, m.Host)
		for k, v := range m.Metrics {
			fmt.Printf("%s: \n", k)

			var sum float32
			var min float32 = math.MaxFloat32
			var max float32 = math.SmallestNonzeroFloat32

			for _, j := range v {
				//t := int64(j[0])
				//fmt.Printf(time.Unix(t/1000, 0).String())
				//fmt.Printf("  %.2f ", j[1])

				d := j[1]
				sum += j[1]

				if d > max {
					max = d
				}

				if d < min {
					min = d
				}
			}

			switch k {
			case "cpu.used":
				{
					fmt.Printf("  average=%.2f%% min=%.2f%%, max=%.2f%%\n", sum/float32(len(v))*100, min*100, max*100)
				}
			case "memory.used":
				{
					fmt.Printf("  average=%.2f%% min=%.2f%%, max=%.2f%%\n", sum/float32(len(v))*100, min*100, max*100)
				}
			case "load.1min":
				{
					fmt.Printf("  average=%.2f min=%.2f, max=%.2f\n", sum/float32(len(v)), min, max)
				}
			}

		}
	}

	instanaSnapshotsOpts := openapi.GetSnapshotsOpts{
		Offline:    optional.NewBool(true),
		Plugin:     optional.NewString("host"),
		Query:      optional.NewString(fmt.Sprintf("entity.host.name:%s", buildInfo.AgentHostMachine)),
		To:         optional.NewInt64(buildInfo.CompletedTimeUnix * 1000), // Instana API requires nanosecond resolution
		WindowSize: optional.NewInt64(buildInfo.DurationMs),
		Size:       optional.NewInt32(10),
	}

	snapshots, _, err := ic.Client.InfrastructureResourcesApi.GetSnapshots(ic.Context, &instanaSnapshotsOpts)

	if err != nil {
		return fmt.Errorf("Failed to search snapshots: %s", err)
	}

	fmt.Printf("Snapshots:  ")
	for _, s := range snapshots.Items {
		fmt.Printf("  host: %s\n  label: %s\n  id: %s\n", s.Host, s.Label, s.SnapshotId)
	}

	instanaSnapshotOpts := openapi.GetSnapshotOpts{
		To:         optional.NewInt64(buildInfo.CompletedTimeUnix * 1000), // Instana API requires nanosecond resolution
		WindowSize: optional.NewInt64(buildInfo.DurationMs),
	}
	snapshot, _, err := ic.Client.InfrastructureResourcesApi.GetSnapshot(ic.Context, snapshots.Items[0].SnapshotId, &instanaSnapshotOpts)

	if err != nil {
		return fmt.Errorf("Failed to retrieve snapshot: %s", err)
	}

	nCPUs := int64(snapshot.Data["cpu.count"].(float64))
	memBytes := int64(snapshot.Data["memory.total"].(float64))

	fmt.Printf("CPUs: %d\nMemory: %d", nCPUs, memBytes)

	return nil
}

/*
	The number of data points returned per metric is limited to 600. Therefore, the rollup parameter (i.e. granularity) must be adjusted based on the build duration.
	See: https://instana.github.io/openapi/#tag/Infrastructure-Metrics

	Builds older than 24 hours are limited to a maximum granularity of 1 minute (see https://www.instana.com/docs/policies/#data-retention-policy)
	This means if the build is short (e.g. <5 minutes), it may make sense to expand the time frame to get more metrics.
	Builds run within the last 24 hours benefit from 1 or 5 second granularity.

*/
func computeRollupValue(buildStartTimeUnix int64, buildDurationMs int64) int32 {
	MaxDataPoints := 600
	var RollUpValuesSeconds []int

	hoursSince := int(math.Floor(time.Now().Sub(time.Unix(buildStartTimeUnix, 0)).Hours()))
	buildDurationSeconds := int(buildDurationMs / 1000)

	if hoursSince < 24 {
		RollUpValuesSeconds = []int{1, 5, 60, 300, 3600}
	} else {
		RollUpValuesSeconds = []int{60, 300, 3600}
	}

	for _, rollup := range RollUpValuesSeconds {
		if int(buildDurationSeconds/rollup) < MaxDataPoints {
			return int32(rollup)
		}
	}
	return 1
}
