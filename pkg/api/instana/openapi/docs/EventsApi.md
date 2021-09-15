# \EventsApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /api/events/{eventId} | Get a particular event
[**GetEvents**](EventsApi.md#GetEvents) | **Get** /api/events | Get all events



## GetEvent

> EventResult GetEvent(ctx, eventId)

Get a particular event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**|  | 

### Return type

[**EventResult**](EventResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvents

> []EventResult GetEvents(ctx, optional)

Get all events

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **windowSize** | **optional.Int64**|  | 
 **from** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **excludeTriggeredBefore** | **optional.Bool**|  | 

### Return type

[**[]EventResult**](EventResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

