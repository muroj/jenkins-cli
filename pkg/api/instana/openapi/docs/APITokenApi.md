# \APITokenApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiToken**](APITokenApi.md#DeleteApiToken) | **Delete** /api/settings/api-tokens/{internalId} | Delete API token
[**GetApiToken**](APITokenApi.md#GetApiToken) | **Get** /api/settings/api-tokens/{internalId} | API token
[**GetApiTokens**](APITokenApi.md#GetApiTokens) | **Get** /api/settings/api-tokens | All API tokens
[**PostApiToken**](APITokenApi.md#PostApiToken) | **Post** /api/settings/api-tokens | Create an API token
[**PutApiToken**](APITokenApi.md#PutApiToken) | **Put** /api/settings/api-tokens/{internalId} | Create or update an API token



## DeleteApiToken

> DeleteApiToken(ctx, internalId)

Delete API token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internalId** | **string**|  | 

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


## GetApiToken

> ApiToken GetApiToken(ctx, internalId)

API token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internalId** | **string**|  | 

### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiTokens

> []ApiToken GetApiTokens(ctx, )

All API tokens

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


## PostApiToken

> ApiToken PostApiToken(ctx, apiToken)

Create an API token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiToken** | [**ApiToken**](ApiToken.md)|  | 

### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApiToken

> ApiToken PutApiToken(ctx, internalId, apiToken)

Create or update an API token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internalId** | **string**|  | 
**apiToken** | [**ApiToken**](ApiToken.md)|  | 

### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

