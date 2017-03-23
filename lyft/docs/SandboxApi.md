# \SandboxApi

All URIs are relative to *https://api.lyft.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetPrimeTime**](SandboxApi.md#SetPrimeTime) | **Put** /sandbox/primetime | Preset Prime Time percentage
[**SetRideStatus**](SandboxApi.md#SetRideStatus) | **Put** /sandbox/rides/{id} | Propagate ride through ride status
[**SetRideTypeAvailability**](SandboxApi.md#SetRideTypeAvailability) | **Put** /sandbox/ridetypes/{ride_type} | Driver availability for processing ride request
[**SetRideTypes**](SandboxApi.md#SetRideTypes) | **Put** /sandbox/ridetypes | Preset types of rides for sandbox


# **SetPrimeTime**
> SetPrimeTime(ctx, request)
Preset Prime Time percentage

Preset a Prime Time percentage in the region surrounding the specified location. This Prime Time percentage will be applied when requesting cost, or when requesting a ride in sandbox mode. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **request** | [**SandboxPrimetime**](SandboxPrimetime.md)| Prime Time to be preset in the region surrounding the lat, lng | 

### Return type

 (empty response body)

### Authorization

[Client Authentication](../README.md#Client Authentication), [User Authentication](../README.md#User Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetRideStatus**
> SandboxRideUpdate SetRideStatus(ctx, id, request)
Propagate ride through ride status

Propagate a sandbox-ride through various ride status 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The ID of the ride | 
  **request** | [**SandboxRideStatus**](SandboxRideStatus.md)| status to propagate the ride into | 

### Return type

[**SandboxRideUpdate**](SandboxRideUpdate.md)

### Authorization

[User Authentication](../README.md#User Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetRideTypeAvailability**
> SetRideTypeAvailability(ctx, rideType, request)
Driver availability for processing ride request

Set driver availability for the provided ride_type in the city/region surrounding the specified location 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **rideType** | **string**|  | 
  **request** | [**SandboxDriverAvailability**](SandboxDriverAvailability.md)| Driver availability to be preset in the region surrounding the lat, lng | 

### Return type

 (empty response body)

### Authorization

[Client Authentication](../README.md#Client Authentication), [User Authentication](../README.md#User Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetRideTypes**
> SandboxRideType SetRideTypes(ctx, request)
Preset types of rides for sandbox

The sandbox-ridetypes endpoint allows you to preset the ridetypes in the region surrounding the specified latitude and longitude to allow testing different scenarios 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **request** | [**SandboxRideType**](SandboxRideType.md)| Ridetypes to be preset in the region surrounding the lat, lng | 

### Return type

[**SandboxRideType**](SandboxRideType.md)

### Authorization

[Client Authentication](../README.md#Client Authentication), [User Authentication](../README.md#User Authentication)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

