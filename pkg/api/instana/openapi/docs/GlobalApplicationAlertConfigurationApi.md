# \GlobalApplicationAlertConfigurationApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlobalApplicationAlertConfig**](GlobalApplicationAlertConfigurationApi.md#CreateGlobalApplicationAlertConfig) | **Post** /api/events/settings/global-alert-configs/applications | Create Global Application Alert Config
[**DeleteGlobalApplicationAlertConfig**](GlobalApplicationAlertConfigurationApi.md#DeleteGlobalApplicationAlertConfig) | **Delete** /api/events/settings/global-alert-configs/applications/{id} | Delete Global Application Alert Config
[**DisableGlobalApplicationAlertConfig**](GlobalApplicationAlertConfigurationApi.md#DisableGlobalApplicationAlertConfig) | **Put** /api/events/settings/global-alert-configs/applications/{id}/disable | Disable Global Application Alert Config
[**EnableGlobalApplicationAlertConfig**](GlobalApplicationAlertConfigurationApi.md#EnableGlobalApplicationAlertConfig) | **Put** /api/events/settings/global-alert-configs/applications/{id}/enable | Enable Global Application Alert Config
[**FindActiveGlobalApplicationAlertConfigs**](GlobalApplicationAlertConfigurationApi.md#FindActiveGlobalApplicationAlertConfigs) | **Get** /api/events/settings/global-alert-configs/applications | All Global Application Alert Configs
[**FindGlobalApplicationAlertConfig**](GlobalApplicationAlertConfigurationApi.md#FindGlobalApplicationAlertConfig) | **Get** /api/events/settings/global-alert-configs/applications/{id} | Get Global Application Alert Config
[**FindGlobalApplicationAlertConfigVersions**](GlobalApplicationAlertConfigurationApi.md#FindGlobalApplicationAlertConfigVersions) | **Get** /api/events/settings/global-alert-configs/applications/{id}/versions | Get versions of Global Application Alert Config
[**UpdateGlobalApplicationAlertConfig**](GlobalApplicationAlertConfigurationApi.md#UpdateGlobalApplicationAlertConfig) | **Post** /api/events/settings/global-alert-configs/applications/{id} | Update Global Application Alert Config



## CreateGlobalApplicationAlertConfig

> GlobalApplicationAlertConfigWithMetadata CreateGlobalApplicationAlertConfig(ctx, globalApplicationsAlertConfig)

Create Global Application Alert Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalApplicationsAlertConfig** | [**GlobalApplicationsAlertConfig**](GlobalApplicationsAlertConfig.md)|  | 

### Return type

[**GlobalApplicationAlertConfigWithMetadata**](GlobalApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlobalApplicationAlertConfig

> DeleteGlobalApplicationAlertConfig(ctx, id)

Delete Global Application Alert Config

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


## DisableGlobalApplicationAlertConfig

> DisableGlobalApplicationAlertConfig(ctx, id, optional)

Disable Global Application Alert Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***DisableGlobalApplicationAlertConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DisableGlobalApplicationAlertConfigOpts struct


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


## EnableGlobalApplicationAlertConfig

> EnableGlobalApplicationAlertConfig(ctx, id, optional)

Enable Global Application Alert Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***EnableGlobalApplicationAlertConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnableGlobalApplicationAlertConfigOpts struct


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


## FindActiveGlobalApplicationAlertConfigs

> []GlobalApplicationAlertConfigWithMetadata FindActiveGlobalApplicationAlertConfigs(ctx, optional)

All Global Application Alert Configs

Configs are sorted descending by their created date.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindActiveGlobalApplicationAlertConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindActiveGlobalApplicationAlertConfigsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **optional.String**|  | 

### Return type

[**[]GlobalApplicationAlertConfigWithMetadata**](GlobalApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindGlobalApplicationAlertConfig

> GlobalApplicationAlertConfigWithMetadata FindGlobalApplicationAlertConfig(ctx, id, optional)

Get Global Application Alert Config

Find a Global Application Alert Config by ID. This will deliver deleted configs too.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***FindGlobalApplicationAlertConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindGlobalApplicationAlertConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validOn** | **optional.Int64**|  | 

### Return type

[**GlobalApplicationAlertConfigWithMetadata**](GlobalApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindGlobalApplicationAlertConfigVersions

> []ConfigVersion FindGlobalApplicationAlertConfigVersions(ctx, id)

Get versions of Global Application Alert Config

Find all versions of a Global Application Alert Config by ID. This will deliver deleted configs too. Configs are sorted descending by their created date.

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


## UpdateGlobalApplicationAlertConfig

> GlobalApplicationAlertConfigWithMetadata UpdateGlobalApplicationAlertConfig(ctx, id, globalApplicationsAlertConfig)

Update Global Application Alert Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**globalApplicationsAlertConfig** | [**GlobalApplicationsAlertConfig**](GlobalApplicationsAlertConfig.md)|  | 

### Return type

[**GlobalApplicationAlertConfigWithMetadata**](GlobalApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

