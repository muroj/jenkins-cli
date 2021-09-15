# \ApplicationTopologyApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServicesMap**](ApplicationTopologyApi.md#GetServicesMap) | **Get** /api/application-monitoring/topology/services | Gets the service topology



## GetServicesMap

> ServiceMap GetServicesMap(ctx, optional)

Gets the service topology

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetServicesMapOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetServicesMapOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **windowSize** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 

### Return type

[**ServiceMap**](ServiceMap.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

