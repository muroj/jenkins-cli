# \ApplicationAlertConfigurationApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationAlertConfig**](ApplicationAlertConfigurationApi.md#CreateApplicationAlertConfig) | **Post** /api/events/settings/application-alert-configs | Create Application Alert Config
[**DeleteApplicationAlertConfig**](ApplicationAlertConfigurationApi.md#DeleteApplicationAlertConfig) | **Delete** /api/events/settings/application-alert-configs/{id} | Delete Application Alert Config
[**DisableApplicationAlertConfig**](ApplicationAlertConfigurationApi.md#DisableApplicationAlertConfig) | **Put** /api/events/settings/application-alert-configs/{id}/disable | Disable Application Alert Config
[**EnableApplicationAlertConfig**](ApplicationAlertConfigurationApi.md#EnableApplicationAlertConfig) | **Put** /api/events/settings/application-alert-configs/{id}/enable | Enable Application Alert Config
[**FindActiveApplicationAlertConfigs**](ApplicationAlertConfigurationApi.md#FindActiveApplicationAlertConfigs) | **Get** /api/events/settings/application-alert-configs | All Application Alert Configs
[**FindApplicationAlertConfig**](ApplicationAlertConfigurationApi.md#FindApplicationAlertConfig) | **Get** /api/events/settings/application-alert-configs/{id} | Get Application Alert Config
[**FindApplicationAlertConfigVersions**](ApplicationAlertConfigurationApi.md#FindApplicationAlertConfigVersions) | **Get** /api/events/settings/application-alert-configs/{id}/versions | Get versions of Application Alert Config
[**UpdateApplicationAlertConfig**](ApplicationAlertConfigurationApi.md#UpdateApplicationAlertConfig) | **Post** /api/events/settings/application-alert-configs/{id} | Update Application Alert Config



## CreateApplicationAlertConfig

> ApplicationAlertConfigWithMetadata CreateApplicationAlertConfig(ctx, applicationAlertConfig)

Create Application Alert Config

This API endpoint creates the Application Alert Configuration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationAlertConfig** | [**ApplicationAlertConfig**](ApplicationAlertConfig.md)|  | 

### Return type

[**ApplicationAlertConfigWithMetadata**](ApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationAlertConfig

> DeleteApplicationAlertConfig(ctx, id)

Delete Application Alert Config

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


## DisableApplicationAlertConfig

> DisableApplicationAlertConfig(ctx, id, optional)

Disable Application Alert Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***DisableApplicationAlertConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DisableApplicationAlertConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 

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


## EnableApplicationAlertConfig

> EnableApplicationAlertConfig(ctx, id, optional)

Enable Application Alert Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***EnableApplicationAlertConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnableApplicationAlertConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 

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


## FindActiveApplicationAlertConfigs

> []ApplicationAlertConfigWithMetadata FindActiveApplicationAlertConfigs(ctx, optional)

All Application Alert Configs

Configs are sorted descending by their created date.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindActiveApplicationAlertConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindActiveApplicationAlertConfigsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **optional.String**|  | 

### Return type

[**[]ApplicationAlertConfigWithMetadata**](ApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindApplicationAlertConfig

> ApplicationAlertConfigWithMetadata FindApplicationAlertConfig(ctx, id, optional)

Get Application Alert Config

Find a Application Alert Config by ID. This will deliver deleted configs too.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***FindApplicationAlertConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindApplicationAlertConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validOn** | **optional.Int64**|  | 

### Return type

[**ApplicationAlertConfigWithMetadata**](ApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindApplicationAlertConfigVersions

> []ConfigVersion FindApplicationAlertConfigVersions(ctx, id)

Get versions of Application Alert Config

Find all versions of a Application Alert Config by ID. This will deliver deleted configs too. Configs are sorted descending by their created date.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**[]ConfigVersion**](ConfigVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationAlertConfig

> ApplicationAlertConfigWithMetadata UpdateApplicationAlertConfig(ctx, id, applicationAlertConfig)

Update Application Alert Config

This API endpoint updates the Application Alert Configuration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**applicationAlertConfig** | [**ApplicationAlertConfig**](ApplicationAlertConfig.md)|  | 

### Return type

[**ApplicationAlertConfigWithMetadata**](ApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

