# \CustomDashboardsApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomDashboard**](CustomDashboardsApi.md#AddCustomDashboard) | **Post** /api/custom-dashboard | Add custom dashboard
[**DeleteCustomDashboard**](CustomDashboardsApi.md#DeleteCustomDashboard) | **Delete** /api/custom-dashboard/{customDashboardId} | Remove custom dashboard
[**GetCustomDashboard**](CustomDashboardsApi.md#GetCustomDashboard) | **Get** /api/custom-dashboard/{customDashboardId} | Get custom dashboard
[**GetCustomDashboards**](CustomDashboardsApi.md#GetCustomDashboards) | **Get** /api/custom-dashboard | Get accessible custom dashboards
[**GetShareableApiTokens**](CustomDashboardsApi.md#GetShareableApiTokens) | **Get** /api/custom-dashboard/shareable-api-tokens | Get all API tokens.
[**GetShareableUsers**](CustomDashboardsApi.md#GetShareableUsers) | **Get** /api/custom-dashboard/shareable-users | Get all users (without invitations).
[**UpdateCustomDashboard**](CustomDashboardsApi.md#UpdateCustomDashboard) | **Put** /api/custom-dashboard/{customDashboardId} | Update custom dashboard



## AddCustomDashboard

> CustomDashboard AddCustomDashboard(ctx, optional)

Add custom dashboard

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddCustomDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddCustomDashboardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customDashboard** | [**optional.Interface of CustomDashboard**](CustomDashboard.md)|  | 

### Return type

[**CustomDashboard**](CustomDashboard.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomDashboard

> DeleteCustomDashboard(ctx, customDashboardId)

Remove custom dashboard

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customDashboardId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomDashboard

> CustomDashboard GetCustomDashboard(ctx, customDashboardId)

Get custom dashboard

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customDashboardId** | **string**|  | 

### Return type

[**CustomDashboard**](CustomDashboard.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomDashboards

> []CustomDashboardPreview GetCustomDashboards(ctx, )

Get accessible custom dashboards

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]CustomDashboardPreview**](CustomDashboardPreview.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShareableApiTokens

> []ApiToken GetShareableApiTokens(ctx, )

Get all API tokens.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShareableUsers

> []UserResult GetShareableUsers(ctx, )

Get all users (without invitations).

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]UserResult**](UserResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomDashboard

> CustomDashboard UpdateCustomDashboard(ctx, customDashboardId, optional)

Update custom dashboard

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customDashboardId** | **string**|  | 
 **optional** | ***UpdateCustomDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCustomDashboardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customDashboard** | [**optional.Interface of CustomDashboard**](CustomDashboard.md)|  | 

### Return type

[**CustomDashboard**](CustomDashboard.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

