package jenkins

import (
	"context"
	"fmt"
	"time"

	"github.com/muroj/gojenkins"
)

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
