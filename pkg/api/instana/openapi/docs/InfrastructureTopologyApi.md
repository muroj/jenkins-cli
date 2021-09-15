# \InfrastructureTopologyApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRelatedHosts**](InfrastructureTopologyApi.md#GetRelatedHosts) | **Get** /api/infrastructure-monitoring/graph/related-hosts/{snapshotId} | Related hosts
[**GetTopology**](InfrastructureTopologyApi.md#GetTopology) | **Get** /api/infrastructure-monitoring/topology | Gets the infrastructure topology



## GetRelatedHosts

> []string GetRelatedHosts(ctx, snapshotId)

Related hosts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string**|  | 

### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopology

> Topology GetTopology(ctx, optional)

Gets the infrastructure topology

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTopologyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTopologyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeData** | **optional.Bool**|  | 

### Return type

[**Topology**](Topology.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

