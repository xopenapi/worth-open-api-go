# \DeviceApi

All URIs are relative to *https://iot.worthcloud.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendAction**](DeviceApi.md#SendAction) | **Post** /v1/devices/send_action | 发送指令



## SendAction

> ApiResponse SendAction(ctx, body)

发送指令

发送指令

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ActionReq**](ActionReq.md)|  | 

### Return type

[**ApiResponse**](APIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

