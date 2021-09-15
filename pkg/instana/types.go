package instana

import (
	"context"
	instana "instana/openapi"
)

type InstanaMetric struct {
	Name      string
	Formatter func()
}

type InstanaAPIClient struct {
	Client    *instana.APIClient
	Creds     InstanaCredentials
	Context   context.Context
	DebugMode bool
}

type InstanaCredentials struct {
	APIKey string
}
