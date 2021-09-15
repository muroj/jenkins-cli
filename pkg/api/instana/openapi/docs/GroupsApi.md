# \GroupsApi

All URIs are relative to *https://tron-ibmdataaiwai.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPermissionsOnGroup**](GroupsApi.md#AddPermissionsOnGroup) | **Put** /api/settings/rbac/groups/{groupId}/permissions | Add permissions to group
[**AddUsersToGroup**](GroupsApi.md#AddUsersToGroup) | **Put** /api/settings/rbac/groups/{groupId}/users | Add users to group
[**CreateGroup**](GroupsApi.md#CreateGroup) | **Post** /api/settings/rbac/groups | Create group
[**DeleteGroup**](GroupsApi.md#DeleteGroup) | **Delete** /api/settings/rbac/groups/{id} | Delete group
[**GetGroup**](GroupsApi.md#GetGroup) | **Get** /api/settings/rbac/groups/{id} | Get group
[**GetGroups**](GroupsApi.md#GetGroups) | **Get** /api/settings/rbac/groups | Get groups
[**GetGroupsByUser**](GroupsApi.md#GetGroupsByUser) | **Get** /api/settings/rbac/groups/user/{email} | Get groups of a single user
[**UpdateGroup**](GroupsApi.md#UpdateGroup) | **Put** /api/settings/rbac/groups/{id} | Update group



## AddPermissionsOnGroup

> ApiGroup AddPermissionsOnGroup(ctx, groupId, body)

Add permissions to group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**|  | 
**body** | **string**|  | 

### Return type

[**ApiGroup**](ApiGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUsersToGroup

> ApiGroup AddUsersToGroup(ctx, groupId, body)

Add users to group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**|  | 
**body** | **string**|  | 

### Return type

[**ApiGroup**](ApiGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> ApiGroup CreateGroup(ctx, apiGroup)

Create group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiGroup** | [**ApiGroup**](ApiGroup.md)|  | 

### Return type

[**ApiGroup**](ApiGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, id)

Delete group

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


## GetGroup

> ApiGroup GetGroup(ctx, id)

Get group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**ApiGroup**](ApiGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> []ApiGroup GetGroups(ctx, )

Get groups

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ApiGroup**](ApiGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupsByUser

> []ApiGroup GetGroupsByUser(ctx, email)

Get groups of a single user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**|  | 

### Return type

[**[]ApiGroup**](ApiGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> ApiGroup UpdateGroup(ctx, id, apiGroup)

Update group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**apiGroup** | [**ApiGroup**](ApiGroup.md)|  | 

### Return type

[**ApiGroup**](ApiGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

