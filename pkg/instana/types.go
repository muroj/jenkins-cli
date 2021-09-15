package instana

import (
	"context"

	"github.ibm.com/jmuro/ghestimator/pkg/api/instana/openapi"
)

type InstanaMetric struct {
	Name      string
	Formatter func()
}

type InstanaAPIClient struct {
	Client    *openapi.APIClient
	Creds     InstanaCredentials
	Context   context.Context
	DebugMode bool
}

type InstanaCredentials struct {
	APIKey string
}

type InstanaHostMetricResult struct {
	Name    string
	Min     float32
	Max     float32
	Average float32
	Data    [][]float32
}
