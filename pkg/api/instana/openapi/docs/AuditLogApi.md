# \AuditLogApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLogs1**](AuditLogApi.md#GetAuditLogs1) | **Get** /api/settings/auditlog | Audit log



## GetAuditLogs1

> AuditLogUiResponse GetAuditLogs1(ctx, optional)

Audit log

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAuditLogs1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAuditLogs1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**|  | 
 **query** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**AuditLogUiResponse**](AuditLogUiResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

