# \InfrastructureResourcesApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMonitoringState**](InfrastructureResourcesApi.md#GetMonitoringState) | **Get** /api/infrastructure-monitoring/monitoring-state | Monitored host count
[**GetSnapshot**](InfrastructureResourcesApi.md#GetSnapshot) | **Get** /api/infrastructure-monitoring/snapshots/{id} | Get snapshot details
[**GetSnapshots**](InfrastructureResourcesApi.md#GetSnapshots) | **Get** /api/infrastructure-monitoring/snapshots | Search snapshots
[**SoftwareVersions**](InfrastructureResourcesApi.md#SoftwareVersions) | **Get** /api/infrastructure-monitoring/software/versions | Get installed software



## GetMonitoringState

> MonitoringState GetMonitoringState(ctx, )

Monitored host count

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**MonitoringState**](MonitoringState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshot

> SnapshotItem GetSnapshot(ctx, id, optional)

Get snapshot details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***GetSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSnapshotOpts struct


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


## GetSnapshots

> SnapshotResult GetSnapshots(ctx, optional)

Search snapshots

These APIs can be used to retrieve information about hosts, processes, JVMs and other entities that we are calling snapshots. A snapshot represents static information about an entity as it was at a specific point in time. To clarify:  **Static information** is any information which is seldom changing, e.g. process IDs, host FQDNs or a list of host hard disks. The counterpart to static information are metrics which have a much higher change rate, e.g. host CPU usage or JVM garbage collection activity. Snapshots only contain static information.  - Snapshots are **versioned** and represent an entity's state for a specific point in time. While snapshots only contain static information, even that information may change. For example you may add another hard disk to a server. For such a change, a new snapshot would be created.  - The **size** parameter can be used in order to limit the maximum number of retrieved snapshots.  - The **offline** parameter is used to allow deeper visibility into snapshots. Set to `false`, the query will return all snapshots that are still available on the given **to** timestamp. However, set to `true`, the query will return all snapshots that have been active within the time window, this must at least include the online result and snapshots terminated within this time.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSnapshotsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSnapshotsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**|  | 
 **to** | **optional.Int64**|  | 
 **windowSize** | **optional.Int64**|  | 
 **size** | **optional.Int32**|  | 
 **plugin** | **optional.String**|  | 
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


## SoftwareVersions

> []SoftwareVersion SoftwareVersions(ctx, optional)

Get installed software

Retrieve information about the software you are running. This includes runtime and package manager information.  The `name`, `version`, `origin` and `type` parameters are optional filters that can be used to reduce the result data set.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwareVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SoftwareVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time** | **optional.Int64**|  | 
 **origin** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **version** | **optional.String**|  | 

### Return type

[**[]SoftwareVersion**](SoftwareVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

