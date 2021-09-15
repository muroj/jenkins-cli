# \InfrastructureMetricsApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInfrastructureMetrics**](InfrastructureMetricsApi.md#GetInfrastructureMetrics) | **Post** /api/infrastructure-monitoring/metrics | Get infrastructure metrics



## GetInfrastructureMetrics

> InfrastructureMetricResult GetInfrastructureMetrics(ctx, optional)

Get infrastructure metrics

- The **offline** parameter is used to allow deeper visibility into snapshots. Set to `false`, the query will return all snapshots that are still available on the given **to** timestamp. However, set to `true`, the query will return all snapshots that have been active within the time window, this must at least include the online result and snapshots terminated within this time.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetInfrastructureMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInfrastructureMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offline** | **optional.Bool**|  | 
 **getCombinedMetrics** | [**optional.Interface of GetCombinedMetrics**](GetCombinedMetrics.md)|  | 

### Return type

[**InfrastructureMetricResult**](InfrastructureMetricResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

