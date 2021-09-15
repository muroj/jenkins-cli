package instana

import (
	"context"
	"fmt"
	"log"

	instana "instana/openapi"
)

func NewInstanaClient(instanaURL string, creds InstanaCredentials) *InstanaAPIClient {
	log.Printf("Initializing Instana client")

	conf := instana.NewConfiguration()
	conf.Host = instanaURL
	conf.BasePath = fmt.Sprintf("https://%s", conf.Host)
	conf.Debug = false
	iac := instana.NewAPIClient(conf)

	authCtx := context.WithValue(context.Background(), instana.ContextAPIKey, instana.APIKey{
		Key:    creds.APIKey,
		Prefix: "apiToken",
	})

	var client InstanaAPIClient
	client.Client = iac
	client.Context = authCtx

	return &client
}
