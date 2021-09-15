package instana

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/antihax/optional"
	"github.ibm.com/jmuro/ghestimator/pkg/api/instana/openapi"
)

func NewInstanaClient(instanaURL string, creds InstanaCredentials, debug bool) *InstanaAPIClient {
	log.Printf("Initializing Instana client")

	conf := openapi.NewConfiguration()
	conf.Host = instanaURL
	conf.BasePath = fmt.Sprintf("https://%s", conf.Host)
	conf.Debug = debug
	iac := openapi.NewAPIClient(conf)

	authCtx := context.WithValue(context.Background(), openapi.ContextAPIKey, openapi.APIKey{
		Key:    creds.APIKey,
		Prefix: "apiToken",
	})

	var client InstanaAPIClient
	client.Client = iac
	client.Context = authCtx

	return &client
}

func GetHostMetrics(hostname string, windowStartTimeUnix int64, windowSizeMs int64, metricIds []string, instanaClient *InstanaAPIClient) ([]InstanaHostMetricResult, error) {
	infraMetricsOpts := openapi.GetInfrastructureMetricsOpts{
		Offline: optional.EmptyBool(),
		GetCombinedMetrics: optional.NewInterface(openapi.GetCombinedMetrics{
			Metrics: metricIds,
			Plugin:  "host",
			Query:   fmt.Sprintf("entity.host.name:%s", hostname),
			TimeFrame: openapi.TimeFrame{
				To:         windowStartTimeUnix * 1000, // Instana API requires nanosecond resolution
				WindowSize: windowSizeMs,
			},
			Rollup: computeRollupValue(windowStartTimeUnix, windowSizeMs),
		}),
	}

	metricsResult, _, err := instanaClient.Client.InfrastructureMetricsApi.GetInfrastructureMetrics(instanaClient.Context, &infraMetricsOpts)

	if err != nil {
		return nil, fmt.Errorf("Error retrieving host metrics: %s", err)
	}

	var results []InstanaHostMetricResult

	for _, m := range metricsResult.Items {
		for k, v := range m.Metrics {
			var sum float32
			var min float32 = math.MaxFloat32
			var max float32 = math.SmallestNonzeroFloat32

			for _, j := range v {
				d := j[1]
				sum += j[1]

				if d > max {
					max = d
				}

				if d < min {
					min = d
				}
			}

			results = append(results, InstanaHostMetricResult{
				Name:    k,
				Min:     min,
				Max:     max,
				Average: sum / float32(len(v)) * 100,
				Data:    v,
			})
		}
	}

	return results, nil
}

func GetHostConfiguration(hostname string, windowStartTimeUnix int64, windowSizeMs int64, instanaClient *InstanaAPIClient) error {

	instanaSnapshotsOpts := openapi.GetSnapshotsOpts{
		Offline:    optional.NewBool(true),
		Plugin:     optional.NewString("host"),
		Query:      optional.NewString(fmt.Sprintf("entity.host.name:%s", hostname)),
		To:         optional.NewInt64(windowStartTimeUnix * 1000), // Instana API requires nanosecond resolution
		WindowSize: optional.NewInt64(windowSizeMs),
		Size:       optional.NewInt32(10),
	}

	snapshots, _, err := instanaClient.Client.InfrastructureResourcesApi.GetSnapshots(instanaClient.Context, &instanaSnapshotsOpts)

	if err != nil {
		return fmt.Errorf("Failed to search snapshots: %s", err)
	}

	instanaSnapshotOpts := openapi.GetSnapshotOpts{
		To:         optional.NewInt64(windowStartTimeUnix * 1000), // Instana API requires nanosecond resolution
		WindowSize: optional.NewInt64(windowSizeMs),
	}
	snapshot, _, err := instanaClient.Client.InfrastructureResourcesApi.GetSnapshot(instanaClient.Context, snapshots.Items[0].SnapshotId, &instanaSnapshotOpts)

	if err != nil {
		return fmt.Errorf("Failed to retrieve snapshot: %s", err)
	}

	nCPUs := int64(snapshot.Data["cpu.count"].(float64))
	memBytes := int64(snapshot.Data["memory.total"].(float64))

	fmt.Printf("CPUs: %d\nMemory (MB): %d\n", nCPUs, memBytes/1024/1024)

	return nil
}

/*
	The number of data points returned per metric is limited to 600. Therefore, the rollup parameter (i.e. granularity) must be adjusted based on the requestd time window.
	See: https://instana.github.io/openapi/#tag/Infrastructure-Metrics

	Host metrics older than 24 hours are limited to a maximum granularity of 1 minute (see https://www.instana.com/docs/policies/#data-retention-policy)
	Metrics collected within the last 24 hours benefit from 1 or 5 second granularity.

*/
func computeRollupValue(windowStartTimeUnix int64, windowSizeMs int64) int32 {
	MaxDataPoints := 600
	var RollUpValuesSeconds []int

	hoursSince := int(math.Floor(time.Now().Sub(time.Unix(windowStartTimeUnix, 0)).Hours()))
	windowSizeSeconds := int(windowSizeMs / 1000)

	if hoursSince < 24 {
		RollUpValuesSeconds = []int{1, 5, 60, 300, 3600}
	} else {
		RollUpValuesSeconds = []int{60, 300, 3600}
	}

	for _, rollup := range RollUpValuesSeconds {
		if int(windowSizeSeconds/rollup) < MaxDataPoints {
			return int32(rollup)
		}
	}
	return 1
}

func (r *InstanaHostMetricResult) PrintInstanaHostMetricResult() {
	fmt.Printf("Metric: %s\n  average=%.2f%%\n  min=%.2f%%\n  max=%.2f%%\n", r.Name, r.Average, r.Min*100, r.Max*100)
}
