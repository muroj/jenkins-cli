# \WebsiteCatalogApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebsiteCatalogMetrics**](WebsiteCatalogApi.md#GetWebsiteCatalogMetrics) | **Get** /api/website-monitoring/catalog/metrics | Metric catalog
[**GetWebsiteCatalogTags**](WebsiteCatalogApi.md#GetWebsiteCatalogTags) | **Get** /api/website-monitoring/catalog/tags | Get all existing website tags
[**GetWebsiteTagCatalog**](WebsiteCatalogApi.md#GetWebsiteTagCatalog) | **Get** /api/website-monitoring/catalog | Get website tag catalog



## GetWebsiteCatalogMetrics

> []WebsiteMonitoringMetricDescription GetWebsiteCatalogMetrics(ctx, )

Metric catalog

This endpoint retrieves all available metric definitions for website monitoring. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]WebsiteMonitoringMetricDescription**](WebsiteMonitoringMetricDescription.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebsiteCatalogTags

> []Tag GetWebsiteCatalogTags(ctx, )

Get all existing website tags

This endpoint retrieves all available tags for your monitored system.  These tags can be used to group metric results. ``` \"group\": {   \"groupbyTag\": \"beacon.page.name\" } ```  These tags can be used to filter metric results. ``` \"tagFilters\": [{  \"name\": \"beacon.website.name\",  \"operator\": \"EQUALS\",  \"value\": \"example\" }] ``` 

### Required Parameters

This endpoint does not need any parameter.

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


## GetWebsiteTagCatalog

> TagCatalog GetWebsiteTagCatalog(ctx, beaconType, useCase)

Get website tag catalog

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**beaconType** | **string**|  | 
**useCase** | **string**|  | 

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

