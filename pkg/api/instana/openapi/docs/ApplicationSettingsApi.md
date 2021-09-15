# \ApplicationSettingsApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplicationConfig**](ApplicationSettingsApi.md#AddApplicationConfig) | **Post** /api/application-monitoring/settings/application | Add application configuration
[**AddManualServiceConfig**](ApplicationSettingsApi.md#AddManualServiceConfig) | **Post** /api/application-monitoring/settings/manual-service | Add manual service configuration
[**AddServiceConfig**](ApplicationSettingsApi.md#AddServiceConfig) | **Post** /api/application-monitoring/settings/service | Add service configuration
[**CreateEndpointConfig**](ApplicationSettingsApi.md#CreateEndpointConfig) | **Post** /api/application-monitoring/settings/http-endpoint | Create endpoint configuration
[**DeleteApplicationConfig**](ApplicationSettingsApi.md#DeleteApplicationConfig) | **Delete** /api/application-monitoring/settings/application/{id} | Delete application configuration
[**DeleteEndpointConfig**](ApplicationSettingsApi.md#DeleteEndpointConfig) | **Delete** /api/application-monitoring/settings/http-endpoint/{id} | Delete endpoint configuration
[**DeleteManualServiceConfig**](ApplicationSettingsApi.md#DeleteManualServiceConfig) | **Delete** /api/application-monitoring/settings/manual-service/{id} | Delete manual service configuration
[**DeleteServiceConfig**](ApplicationSettingsApi.md#DeleteServiceConfig) | **Delete** /api/application-monitoring/settings/service/{id} | Delete service configuration
[**GetAllManualServiceConfigs**](ApplicationSettingsApi.md#GetAllManualServiceConfigs) | **Get** /api/application-monitoring/settings/manual-service | All manual service configurations
[**GetApplicationConfig**](ApplicationSettingsApi.md#GetApplicationConfig) | **Get** /api/application-monitoring/settings/application/{id} | Application configuration
[**GetApplicationConfigs**](ApplicationSettingsApi.md#GetApplicationConfigs) | **Get** /api/application-monitoring/settings/application | All Application configurations
[**GetEndpointConfig**](ApplicationSettingsApi.md#GetEndpointConfig) | **Get** /api/application-monitoring/settings/http-endpoint/{id} | Endpoint configuration
[**GetEndpointConfigs**](ApplicationSettingsApi.md#GetEndpointConfigs) | **Get** /api/application-monitoring/settings/http-endpoint | All Endpoint configurations
[**GetServiceConfig**](ApplicationSettingsApi.md#GetServiceConfig) | **Get** /api/application-monitoring/settings/service/{id} | Service configuration
[**GetServiceConfigs**](ApplicationSettingsApi.md#GetServiceConfigs) | **Get** /api/application-monitoring/settings/service | All service configurations
[**OrderServiceConfig**](ApplicationSettingsApi.md#OrderServiceConfig) | **Put** /api/application-monitoring/settings/service/order | Order of service configuration
[**PutApplicationConfig**](ApplicationSettingsApi.md#PutApplicationConfig) | **Put** /api/application-monitoring/settings/application/{id} | Update application configuration
[**PutServiceConfig**](ApplicationSettingsApi.md#PutServiceConfig) | **Put** /api/application-monitoring/settings/service/{id} | Update service configuration
[**ReplaceAll**](ApplicationSettingsApi.md#ReplaceAll) | **Put** /api/application-monitoring/settings/service | Replace all service configurations
[**ReplaceAllManualServiceConfigs**](ApplicationSettingsApi.md#ReplaceAllManualServiceConfigs) | **Put** /api/application-monitoring/settings/manual-service | Replace all manual service configurations
[**UpdateEndpointConfig**](ApplicationSettingsApi.md#UpdateEndpointConfig) | **Put** /api/application-monitoring/settings/http-endpoint/{id} | Update endpoint configuration
[**UpdateManualServiceConfig**](ApplicationSettingsApi.md#UpdateManualServiceConfig) | **Put** /api/application-monitoring/settings/manual-service/{id} | Update manual service configuration



## AddApplicationConfig

> ApplicationConfig AddApplicationConfig(ctx, newApplicationConfig)

Add application configuration

Create a new Application Perspective.  ## Deprecated Parameters **matchSpecification:** A binary tree sturcture of match expression connected with binary operator AND or OR. It is replaced by **tagFilterExpression** which is also used in Application Analyze API endpoints.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newApplicationConfig** | [**NewApplicationConfig**](NewApplicationConfig.md)|  | 

### Return type

[**ApplicationConfig**](ApplicationConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddManualServiceConfig

> ManualServiceConfig AddManualServiceConfig(ctx, newManualServiceConfig)

Add manual service configuration

Add a manual service configuration.  **This is an experimental endpoint to workaround service mapping issues.**  ### Mandatory Parameters:  **tagFilterExpression** A tag filter expression to match calls on which the manual service configuration will be applied. Only call tags are allowed in the expression.  ### Optional Parameters: **unmonitoredServiceName** Specify a service name if you want to map calls to an unmonitored service.  **existingServiceId** Specify a service id if you want to map calls to an existing service.  **description** A description of the configuration.  **enabled** Enable or disable the configuration.  Note: Either unmonitoredServiceName or existingServiceId should be specified in a configuration.  ### Defaults **enabled** `true`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newManualServiceConfig** | [**NewManualServiceConfig**](NewManualServiceConfig.md)|  | 

### Return type

[**ManualServiceConfig**](ManualServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddServiceConfig

> ServiceConfig AddServiceConfig(ctx, serviceConfig)

Add service configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceConfig** | [**ServiceConfig**](ServiceConfig.md)|  | 

### Return type

[**ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEndpointConfig

> HttpEndpointConfig CreateEndpointConfig(ctx, httpEndpointConfig)

Create endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**httpEndpointConfig** | [**HttpEndpointConfig**](HttpEndpointConfig.md)|  | 

### Return type

[**HttpEndpointConfig**](HttpEndpointConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationConfig

> DeleteApplicationConfig(ctx, id)

Delete application configuration

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


## DeleteEndpointConfig

> DeleteEndpointConfig(ctx, id)

Delete endpoint configuration

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


## DeleteManualServiceConfig

> DeleteManualServiceConfig(ctx, id)

Delete manual service configuration

Delete a manual service configuration.  **This is an experimental endpoint to workaround service mapping issues.**

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


## DeleteServiceConfig

> DeleteServiceConfig(ctx, id)

Delete service configuration

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


## GetAllManualServiceConfigs

> []ManualServiceConfig GetAllManualServiceConfigs(ctx, )

All manual service configurations

Get all manual service configurations.  **This is an experimental endpoint to workaround service mapping issues.**

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ManualServiceConfig**](ManualServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationConfig

> ApplicationConfig GetApplicationConfig(ctx, id)

Application configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**ApplicationConfig**](ApplicationConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationConfigs

> []ApplicationConfig GetApplicationConfigs(ctx, )

All Application configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ApplicationConfig**](ApplicationConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointConfig

> HttpEndpointConfig GetEndpointConfig(ctx, id)

Endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**HttpEndpointConfig**](HttpEndpointConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointConfigs

> []HttpEndpointConfig GetEndpointConfigs(ctx, )

All Endpoint configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]HttpEndpointConfig**](HttpEndpointConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceConfig

> ServiceConfig GetServiceConfig(ctx, id)

Service configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceConfigs

> []ServiceConfig GetServiceConfigs(ctx, )

All service configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderServiceConfig

> OrderServiceConfig(ctx, requestBody)

Order of service configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestBody** | [**[]string**](string.md)|  | 

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


## PutApplicationConfig

> ApplicationConfig PutApplicationConfig(ctx, id, applicationConfig)

Update application configuration

Update an existing Application Perspective.  ## Deprecated Parameters **matchSpecification:** A binary tree sturcture of match expression connected with binary operator AND or OR. It is replaced by **tagFilterExpression** which is also used in Application Analyze API endpoints.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**applicationConfig** | [**ApplicationConfig**](ApplicationConfig.md)|  | 

### Return type

[**ApplicationConfig**](ApplicationConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceConfig

> ServiceConfig PutServiceConfig(ctx, id, serviceConfig)

Update service configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**serviceConfig** | [**ServiceConfig**](ServiceConfig.md)|  | 

### Return type

[**ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceAll

> []ServiceConfig ReplaceAll(ctx, serviceConfig)

Replace all service configurations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceConfig** | [**[]ServiceConfig**](ServiceConfig.md)|  | 

### Return type

[**[]ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceAllManualServiceConfigs

> []ManualServiceConfig ReplaceAllManualServiceConfigs(ctx, newManualServiceConfig)

Replace all manual service configurations

Replace all manual service configurations.  **This is an experimental endpoint to workaround service mapping issues.**  ### Mandatory Parameters:  **tagFilterExpression** A tag filter expression to match calls on which the manual service configuration will be applied. Only call tags are allowed in the expression.  ### Optional Parameters: **unmonitoredServiceName** Specify a service name if you want to map calls to an unmonitored service.  **existingServiceId** Specify a service id if you want to map calls to an existing service.  **description** A description of the configuration.  **enabled** Enable or disable the configuration.  Note: Either unmonitoredServiceName or existingServiceId should be specified in a configuration.  ### Defaults **enabled** `true`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newManualServiceConfig** | [**[]NewManualServiceConfig**](NewManualServiceConfig.md)|  | 

### Return type

[**[]ManualServiceConfig**](ManualServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpointConfig

> HttpEndpointConfig UpdateEndpointConfig(ctx, id, httpEndpointConfig)

Update endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**httpEndpointConfig** | [**HttpEndpointConfig**](HttpEndpointConfig.md)|  | 

### Return type

[**HttpEndpointConfig**](HttpEndpointConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManualServiceConfig

> ManualServiceConfig UpdateManualServiceConfig(ctx, id, manualServiceConfig)

Update manual service configuration

Update a manual service configuration.  **This is an experimental endpoint to workaround service mapping issues.**  ### Mandatory Parameters:  **tagFilterExpression** A tag filter expression to match calls on which the manual service configuration will be applied. Only call tags are allowed in the expression.  ### Optional Parameters: **unmonitoredServiceName** Specify a service name if you want to map calls to an unmonitored service.  **existingServiceId** Specify a service id if you want to map calls to an existing service.  **description** A description of the configuration.  **enabled** Enable or disable the configuration.  Note: Either unmonitoredServiceName or existingServiceId should be specified in a configuration.  ### Defaults **enabled** `true`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**manualServiceConfig** | [**ManualServiceConfig**](ManualServiceConfig.md)|  | 

### Return type

[**ManualServiceConfig**](ManualServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

