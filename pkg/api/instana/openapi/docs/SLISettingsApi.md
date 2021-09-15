# \SLISettingsApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSli**](SLISettingsApi.md#CreateSli) | **Post** /api/settings/sli | Create SLI Config
[**Delete**](SLISettingsApi.md#Delete) | **Delete** /api/settings/sli/{id} | Delete SLI Config
[**GetSli1**](SLISettingsApi.md#GetSli1) | **Get** /api/settings/sli | Get All SLI Configs
[**GetSli2**](SLISettingsApi.md#GetSli2) | **Get** /api/settings/sli/{id} | Get SLI Config
[**PutSli**](SLISettingsApi.md#PutSli) | **Put** /api/settings/sli/{id} | Update SLI Config



## CreateSli

> SliConfigurationWithLastUpdated CreateSli(ctx, sliConfiguration)

Create SLI Config

This endpoint creates the Service Level Indicator Configuration  ## Mandatory Parameters:  - **id** A unique identifier for each SLI configuration  - **sliName:** Name for the SLI configuration  - **sliEntity:** Entity of the SLI configuration  ### SLI Entity specific parameters  Depending on the chosen `sliType` in the `sliEntity`, there are further required parameters:  #### Application SLI entity  This option can be used to create a Time-Based SLI  - **sliEntity.applicationId:** The Id of the Application Perspective  - **sliEntity.boundaryScope:** Boundary scope of the Application Perspective  - **metricConfiguration.metricName:** The metric name on which to compute the SLI  - **metricConfiguration.metricAggregation:** The aggregation of the metric  - **metricConfiguration.threshold:** Threshold for the metric  #### Availability SLI entity  This opetion can be used to create an Event-Based SLI  - **sliEntity.applicationId:** The Id of the Application Perspective  - **sliEntity.boundaryScope:** Boundary scope of the Application Perspective  ## Deprecated Parameters for Availability SLI entity:  - **sliEntity.serviceId:** The ID if the Service in he context of an Application Perspective  - **sliEntity.endpointId:** The ID of an Endpoint belonging to a Service  - **sliEntity.goodEventFilters:** The list of TagFilters to match good events / calls  - **sliEntity.badEventFilters:** The list of TagFilters to match bad events / calls  All of these filters can be included using the list of TagFilterExpressions via **sliEntity.goodEventFilterExpression** and **sliEntity.badEventFilterExpression**. These parameters will be removed in the upcoming releases. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sliConfiguration** | [**SliConfiguration**](SliConfiguration.md)|  | 

### Return type

[**SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, id)

Delete SLI Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetSli1

> []SliConfigurationWithLastUpdated GetSli1(ctx, )

Get All SLI Configs

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSli2

> SliConfigurationWithLastUpdated GetSli2(ctx, id)

Get SLI Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSli

> SliConfigurationWithLastUpdated PutSli(ctx, id, sliConfiguration)

Update SLI Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**sliConfiguration** | [**SliConfiguration**](SliConfiguration.md)|  | 

### Return type

[**SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

