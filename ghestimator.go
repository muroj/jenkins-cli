package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	instana "instana/openapi"

	"github.com/antihax/optional"
	"github.com/bndr/gojenkins"
)

var (
	jenkinsURL      string
	jenkinsUser     string
	jenkinsAPIToken string
	jobURL          string
	buildNumber     int64
	instanaTenant   string
	instanaUnit     string
)

func init() {
	flag.StringVar(&jenkinsURL, "jenkinsURL", "https://ghenkins.bigdatalab.ibm.com/", "URL of the Jenkins host, e.g. \"https://ghenkins.bigdatalab.ibm.com/\"")
	flag.StringVar(&jenkinsUser, "jenkinsUser", "", "Jenkins username")
	flag.StringVar(&jenkinsAPIToken, "jenkinsAPIToken", "", "Jenkins API token")
	flag.StringVar(&jobURL, "jobURL", "", "Path to the Jenkins job to evaluate, e.g.\"job/ai-foundation/job/envctl/job/main/\"")
	flag.Int64Var(&buildNumber, "buildNumber", 0, "ID of the build to evaluate, e.g. 223. Defaults to the last successful build")
	flag.StringVar(&instanaTenant, "instanaTenant", "tron", "Instana Tentant")
	flag.StringVar(&instanaUnit, "instanaUnit", "ibmdataaiwai", "Instana Unit")
	flag.Parse()

	if len(jenkinsUser) == 0 {
		log.Fatal("Required parameter not specified: jenkinsUser")
	} else if len(jenkinsAPIToken) == 0 {
		log.Fatal("Required parameter not specified: jenkinsAPIToken")
	} else if len(jobURL) == 0 {
		log.Fatal("Required parameter not specified: jobURL")
	}
}

/*
	Given a URL for a Jenkins job. Returns the job project name, and all parent.
	For example, a job URL of "job/ai-foundation/job/abp-code-scan/job/ghenkins/"
	will return (ghenkins, [ai-foundation, abp-code-scan])
*/
func parseJobURL(jobURL string) (string, []string, error) {
	jobComponents := strings.Split(strings.TrimSpace(jobURL), "/")
	var sanitizedJobPath = make([]string, 0, len(jobComponents))

	for _, s := range jobComponents {
		fmt.Println(s)
		if s != "job" && s != "" && s != " " {
			sanitizedJobPath = append(sanitizedJobPath, s)
		}
	}
	jobName := sanitizedJobPath[len(sanitizedJobPath)-1]

	return jobName, sanitizedJobPath[0 : len(sanitizedJobPath)-1], nil
}

func main() {

	/*
		Output results in JSON
		Print host specs: CPU, Memory, etc
	*/

	// Jenkins API client
	log.Println("Initializing Jenkins client")
	ctx := context.Background()
	jenkins, err := gojenkins.CreateJenkins(nil, jenkinsURL, jenkinsUser, jenkinsAPIToken).Init(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Jenkins instance at \"%s\"\n %s", jenkinsURL, err)
	}

	jobName, sanitizedJobPath, _ := parseJobURL(jobURL)

	log.Printf("Retrieving job \"%s\" at path %v\n", jobName, sanitizedJobPath)
	job, err := jenkins.GetJob(ctx, jobName, sanitizedJobPath...)
	if err != nil {
		log.Fatalf("Failed to retreive job with URL \"%s\"\n %s", jobURL, err)
	}

	var build *gojenkins.Build
	if buildNumber == 0 {
		log.Printf("Retrieving last successful build")
		build, err = job.GetLastSuccessfulBuild(ctx)
	} else {
		log.Printf("Retrieving build %d\n", buildNumber)
		build, err = job.GetBuild(ctx, buildNumber)
	}

	if err != nil {
		log.Fatalf("Failed to retrieve build: %s", err)
	}

	// get build scheduled time
	startTime := build.GetTimestamp()
	fmt.Printf("Build started: %s\n", startTime.String())

	// get build execution time
	//executionTimeMs := lsb.GetDuration()
	//println(executionTimeMs)

	// get build total duration
	duration := build.GetDuration()
	fmt.Printf("Build duration: %f\n", duration)

	buildEndTimeUnix := startTime.Unix() + (int64(duration) / 1000)
	fmt.Printf("Build ended at %s\n", time.Unix(buildEndTimeUnix, 0))

	// Find out which node the build ran on.
	buildLog := build.GetConsoleOutput(ctx)
	r, _ := regexp.Compile(`Running on \b([\w]+\b)`)
	matches := r.FindStringSubmatch(buildLog)

	if matches == nil {
		panic(fmt.Sprintf("Unable to determine host node: \"Running on <nodeName>\" line not found in build log."))
	}
	nodeName := matches[1]
	fmt.Printf("Build ran on: %s\n", nodeName)

	// Instana API client
	configuration := instana.NewConfiguration()
	hostURL := fmt.Sprintf("%s-%s.instana.io", instanaTenant, instanaUnit)
	configuration.Host = hostURL
	configuration.BasePath = fmt.Sprintf("https://%s", hostURL)
	client := instana.NewAPIClient(configuration)

	apiKey := os.Getenv("INSTANA_API_KEY")
	authCtx := context.WithValue(context.Background(), instana.ContextAPIKey, instana.APIKey{
		Key:    apiKey,
		Prefix: "apiToken",
	})

	/*
		Builds older than 24 hours are limited to a granularity of 1 minute (see https://www.instana.com/docs/policies/#data-retention-policy)
		This means if the build is short (e.g. <5 minutes), it may make sense to expand the time frame to get more metrics.
		Builds run within the last 24 hours benefit from 1 or 5 second granularity.
		API calls are limited to 600 data points, so the granularity must be adjusted based on the timewindow.
	*/

	metrics := instana.GetInfrastructureMetricsOpts{
		Offline: optional.EmptyBool(),
		GetCombinedMetrics: optional.NewInterface(instana.GetCombinedMetrics{
			Metrics: []string{
				"cpu.used", "load.1min", "memory.used",
			},
			Plugin: "host",
			Query:  fmt.Sprintf("entity.host.name:%s", nodeName),
			TimeFrame: instana.TimeFrame{
				To:         (buildEndTimeUnix * 1000),
				WindowSize: int64(duration),
			},
			Rollup: 1,
		}),
	}

	metricsResult, httpResp, err := client.InfrastructureMetricsApi.GetInfrastructureMetrics(authCtx, &metrics)
	if err != nil {
		s := bufio.NewScanner(httpResp.Body)
		for s.Scan() {
			fmt.Println(s.Text())
		}
		panic(fmt.Errorf("Error calling the API: %s\n Aborting", err))
	}

	if httpResp.StatusCode == http.StatusOK {
		fmt.Printf("API call returned %s\n", http.StatusText(http.StatusOK))
	}

	for _, metric := range metricsResult.Items {
		fmt.Printf("Metric label: %s\nHost: %s\n", metric.Label, metric.Host)
		for k, v := range metric.Metrics {
			fmt.Printf("%s\n: %v\n", k, v)
		}
	}
}
