package jenkins

import (
	"context"
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
