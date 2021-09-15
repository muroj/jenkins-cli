# \HostAgentApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentSnapshot**](HostAgentApi.md#GetAgentSnapshot) | **Get** /api/host-agent/{id} | Get host agent snapshot details
[**GetLogs**](HostAgentApi.md#GetLogs) | **Get** /api/host-agent/{hostId}/logs | Agent download logs
[**Search**](HostAgentApi.md#Search) | **Get** /api/host-agent | Query host agent snapshots
[**Update**](HostAgentApi.md#Update) | **Post** /api/host-agent/{hostId}/update | Agent update
[**UpdateConfigurationByHost**](HostAgentApi.md#UpdateConfigurationByHost) | **Post** /api/host-agent/{hostId}/configuration | Update agent configuration by host
[**UpdateConfigurationByQuery**](HostAgentApi.md#UpdateConfigurationByQuery) | **Post** /api/host-agent/configuration | Update agent configuration by query



## GetAgentSnapshot

> SnapshotItem GetAgentSnapshot(ctx, id, optional)

Get host agent snapshot details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***GetAgentSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAgentSnapshotOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **optional.Int64**|  | 
 **windowSize** | **optional.Int64**|  | 

### Return type

[**SnapshotItem**](SnapshotItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogs

> GetLogs(ctx, hostId, file, optional)

Agent download logs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string**|  | 
**file** | [**[]string**](string.md)|  | 
 **optional** | ***GetLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **download** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> SnapshotResult Search(ctx, optional)

Query host agent snapshots

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**|  | 
 **to** | **optional.Int64**|  | 
 **windowSize** | **optional.Int64**|  | 
 **size** | **optional.Int32**|  | 
 **offline** | **optional.Bool**|  | 

### Return type

[**SnapshotResult**](SnapshotResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(ctx, hostId)

Agent update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigurationByHost

> UpdateConfigurationByHost(ctx, hostId, optional)

Update agent configuration by host

This endpoint can be used to initialize or change configuration management settings for a specific host agent.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string**|  | 
 **optional** | ***UpdateConfigurationByHostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConfigurationByHostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentConfigurationUpdate** | [**optional.Interface of AgentConfigurationUpdate**](AgentConfigurationUpdate.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigurationByQuery

> UpdateConfigurationByQuery(ctx, optional)

Update agent configuration by query

This endpoint can be used to initialize or change configuration management settings for all agents selected by the given Dynamic Focus Query.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateConfigurationByQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConfigurationByQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**|  | 
 **to** | **optional.Int64**|  | 
 **windowSize** | **optional.Int64**|  | 
 **size** | **optional.Int32**|  | 
 **offline** | **optional.Bool**|  | 
 **agentConfigurationUpdate** | [**optional.Interface of AgentConfigurationUpdate**](AgentConfigurationUpdate.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

