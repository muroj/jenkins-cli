# \ApplicationCatalogApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationCatalogMetrics**](ApplicationCatalogApi.md#GetApplicationCatalogMetrics) | **Get** /api/application-monitoring/catalog/metrics | Get Metric catalog
[**GetApplicationTagCatalog**](ApplicationCatalogApi.md#GetApplicationTagCatalog) | **Get** /api/application-monitoring/catalog | Get application tag catalog
[**GetApplicationTags**](ApplicationCatalogApi.md#GetApplicationTags) | **Get** /api/application-monitoring/catalog/tags | Get application tags



## GetApplicationCatalogMetrics

> []MetricDescription GetApplicationCatalogMetrics(ctx, )

Get Metric catalog

This endpoint retrieves all available metric definitions for application monitoring. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]MetricDescription**](MetricDescription.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationTagCatalog

> TagCatalog GetApplicationTagCatalog(ctx, optional)

Get application tag catalog

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetApplicationTagCatalogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationTagCatalogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **optional.Int64**|  | 
 **dataSource** | **optional.String**|  | 
 **useCase** | **optional.String**|  | 

### Return type

[**TagCatalog**](TagCatalog.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationTags

> []Tag GetApplicationTags(ctx, optional)

Get application tags

This endpoint retrieves all available tags for your monitored system.  These tags can be used to group metric results. ``` \"group\": {   \"groupbyTag\": \"service.name\" } ```  These tags can be used to filter metric results. ``` \"tagFilters\": [{  \"name\": \"application.name\",  \"operator\": \"EQUALS\",  \"value\": \"example\" }] ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetApplicationTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **optional.Int64**|  | 
 **dataSource** | **optional.String**|  | 
 **useCase** | **optional.String**|  | 

### Return type

[**[]Tag**](Tag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

