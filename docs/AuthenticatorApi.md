# \AuthenticatorApi

All URIs are relative to *https://iot.worthcloud.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenToken**](AuthenticatorApi.md#GenToken) | **Get** /v1/authenticator/{access_key}/{secret_key} | 生成token



## GenToken

> GenTokenResponse GenToken(ctx, accessKey, secretKey)

生成token

生成token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessKey** | **string**| access_key | 
**secretKey** | **string**| secret_key | 

### Return type

[**GenTokenResponse**](GenTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

