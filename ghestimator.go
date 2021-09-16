package main

import (
	"flag"
	"fmt"
	"log"

	"github.ibm.com/jmuro/ghestimator/pkg/instana"
	"github.ibm.com/jmuro/ghestimator/pkg/jenkins"
)

func main() {

	jenkinsCreds := jenkins.JenkinsCredentials{
		Username: jenkinsUser,
		APIToken: jenkinsAPIToken,
	}
	jenkinsClient := jenkins.NewJenkinsClient(jenkinsURL, jenkinsCreds, enableDebug)
	buildInfo, err := jenkins.GetBuildInfo(jobURL, buildNumber, jenkinsClient)

	if err != nil {
		log.Fatalf("Failed to retrieve required build information: %s", err)
	}

	buildInfo.PrintBuildInfo()

	instanaURL := fmt.Sprintf("%s-%s.instana.io", instanaTenant, instanaUnit)
	instanaClient := instana.NewInstanaClient(instanaURL, instana.InstanaCredentials{instanaAPIKey}, enableDebug)

	err = getResourceUsage(&buildInfo, instanaClient)
	if err != nil {
		log.Fatalf("Failed to retrieve required build information: %s", err)
	}
}

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

func getResourceUsage(buildInfo *jenkins.JenkinsBuildInfo, instanaClient *instana.InstanaAPIClient) error {

	timeWindow := instana.TimeWindow{
		StartTimeUnix:    buildInfo.CompletedTimeUnix,
		WindowDurationMs: buildInfo.ExecutionTimeMs,
	}

	err := instana.GetHostConfiguration(buildInfo.AgentHostMachine, timeWindow, instanaClient)

	if err != nil {
		return fmt.Errorf("Error retrieving host configuration: %s", err)
	}

	hostMetrics := []string{
		"cpu.used", "load.1min", "memory.used",
	}

	metricsResult, err := instana.GetHostMetrics(buildInfo.AgentHostMachine, timeWindow, hostMetrics, instanaClient)

	if err != nil {
		return fmt.Errorf("Error retrieving host metrics: %s", err)
	}

	for _, m := range metricsResult {
		m.PrintInstanaHostMetricResult()
	}

	return nil
}
