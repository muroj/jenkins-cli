package instana

import (
	"context"
	"fmt"

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

func (r *InstanaHostMetricResult) PrintInstanaHostMetricResult() {
	fmt.Printf("Metric: %s\n  average=%.2f%%\n  min=%.2f%%\n  max=%.2f%%\n", r.Name, r.Average*100, r.Min*100, r.Max*100)
}
