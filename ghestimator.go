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

	instana "instana/openapi"

	"github.com/antihax/optional"
	"github.com/bndr/gojenkins"
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

type BuildInfo struct {
	Name               string
	Id                 int64
	scheduledTimestamp time.Time
	scheduledTimeUnix  int64
	completedTimeUnix  int64
	durationMs         int64
	hostNode           string
}

func main() {

	buildInfo, err := getBuildInfo(buildNumber, jobURL, jenkinsURL, jenkinsUser, jenkinsAPIToken)
	if err != nil {
		log.Fatalf("Failed to retrieve required build information: %s", err)
	}
	buildInfo.printBuildInfo()

	hostMetrics := []string{
		"cpu.used", "load.1min", "memory.used",
	}
	instanaURL := fmt.Sprintf("%s-%s.instana.io", instanaTenant, instanaUnit)
	err = getHostMetrics(buildInfo.hostNode, hostMetrics, buildInfo.completedTimeUnix, buildInfo.durationMs, instanaURL, instanaAPIKey)
	if err != nil {
		log.Fatalf("Failed to retrieve required build information: %s", err)
	}
}

func getBuildInfo(id int64, jobURL string, jenkinsURL string, jenkinsUser string, jenkinsAPIToken string) (BuildInfo, error) {
	var bi BuildInfo

	log.Println("Initializing Jenkins client")
	ctx := context.Background()
	j, err := gojenkins.CreateJenkins(nil, jenkinsURL, jenkinsUser, jenkinsAPIToken).Init(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Jenkins instance at \"%s\"\n %s", jenkinsURL, err)
	}

	log.Printf("Retrieving job at URL: \"%s\"", jobURL)
	jobName, path, _ := parseJobURL(jobURL)
	job, err := j.GetJob(ctx, jobName, path...)
	if err != nil {
		return bi, fmt.Errorf("Failed to retreive job with URL \"%s\"\n %s", jobURL, err)
	}

	var b *gojenkins.Build
	if id == 0 {
		log.Printf("Retrieving last successful build")
		b, err = job.GetLastSuccessfulBuild(ctx)
	} else {
		log.Printf("Retrieving build %d\n", id)
		b, err = job.GetBuild(ctx, id)
	}

	if err != nil {
		return bi, fmt.Errorf("Failed to retrieve build: %s", err)
	}

	bi.scheduledTimestamp = b.GetTimestamp()
	bi.durationMs = int64(math.Round(b.GetDuration()))
	// get build execution time
	//executionTimeMs := lsb.GetDuration()
	//println(executionTimeMs)

	bi.completedTimeUnix = bi.scheduledTimestamp.Unix() + (int64(bi.durationMs) / 1000)

	// Determine which node the build ran on.
	bl := b.GetConsoleOutput(ctx)
	r, _ := regexp.Compile(`Running on \b([\w]+\b)`)
	m := r.FindStringSubmatch(bl)
	if m == nil {
		return bi, fmt.Errorf("Unable to determine host node: \"Running on <nodeName>\" line not found in build log.")
	}
	bi.hostNode = m[1]

	return bi, nil
}

func getHostMetrics(hostname string, hostMetrics []string, startTimeUnix int64, durationMs int64, instanaURL string, instanaAPIKey string) error {

	// Create an instana client
	configuration := instana.NewConfiguration()
	configuration.Host = instanaURL
	configuration.BasePath = fmt.Sprintf("https://%s", configuration.Host)
	client := instana.NewAPIClient(configuration)
	authCtx := context.WithValue(context.Background(), instana.ContextAPIKey, instana.APIKey{
		Key:    instanaAPIKey,
		Prefix: "apiToken",
	})

	infraMetricsOpts := instana.GetInfrastructureMetricsOpts{
		Offline: optional.EmptyBool(),
		GetCombinedMetrics: optional.NewInterface(instana.GetCombinedMetrics{
			Metrics: []string{
				"cpu.used", "load.1min", "memory.used",
			},
			Plugin: "host",
			Query:  fmt.Sprintf("entity.host.name:%s", hostname),
			TimeFrame: instana.TimeFrame{
				To:         startTimeUnix * 1000, // API requires nanosecond resolution
				WindowSize: durationMs,
			},
			Rollup: computeRollupPeriod(startTimeUnix, durationMs),
		}),
	}

	metricsResult, _, err := client.InfrastructureMetricsApi.GetInfrastructureMetrics(authCtx, &infraMetricsOpts)
	if err != nil {
		return fmt.Errorf("Error retrieving instana metrics: %s", err)
	}

	for _, metric := range metricsResult.Items {
		fmt.Printf("Metric label: %s\nHost: %s\n", metric.Label, metric.Host)
		for k, v := range metric.Metrics {
			fmt.Printf("%s\n: %v\n", k, v)
		}
	}
	return nil
}

/*
	The number of data points returned per metric is limited to 600. Therefore, the granularity must be adjusted based on the build duration.
	Instana refers to the granularity as the rollup. See: https://instana.github.io/openapi/#tag/Infrastructure-Metrics

	Builds older than 24 hours are limited to a maximum granularity of 1 minute (see https://www.instana.com/docs/policies/#data-retention-policy)
	This means if the build is short (e.g. <5 minutes), it may make sense to expand the time frame to get more metrics.
	Builds run within the last 24 hours benefit from 1 or 5 second granularity.

*/
func computeRollupPeriod(startTimeUnix int64, durationMs int64) int32 {
	MaxDataPoints := 600
	var RollUpPeriodsSeconds []int

	hours := int(math.Floor(time.Now().Sub(time.Unix(startTimeUnix, 0)).Hours()))
	seconds := int(durationMs / 1000)

	if hours < 24 {
		RollUpPeriodsSeconds = []int{1, 5, 60, 300, 3600}
	} else {
		RollUpPeriodsSeconds = []int{60, 300, 3600}
	}

	for _, period := range RollUpPeriodsSeconds {
		if int(seconds/period) < MaxDataPoints {
			return int32(period)
		}
	}
	return 1
}

/*
	Given a URL for a Jenkins job. Returns the job name, and parent folders.
	For example, a job URL of "job/ai-foundation/job/abp-code-scan/job/ghenkins/"
	will return (ghenkins, [ai-foundation, abp-code-scan])
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

func (bi *BuildInfo) printBuildInfo() {
	fmt.Printf("Build started: %s\n", bi.scheduledTimestamp.String())
	fmt.Printf("Build ended: %s\n", time.Unix(bi.completedTimeUnix, 0))
	fmt.Printf("Build duration (ms): %d\n", bi.durationMs)
	fmt.Printf("Build ran on: %s\n", bi.hostNode)
}
