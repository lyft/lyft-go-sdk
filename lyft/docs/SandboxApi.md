# SandboxApi

All URIs are relative to *https://api.lyft.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetPrimeTime**](SandboxApi.md#SetPrimeTime) | **Put** /sandbox/primetime | Preset Prime Time percentage
[**SetRideStatus**](SandboxApi.md#SetRideStatus) | **Put** /sandbox/rides/{id} | Propagate ride through ride status
[**SetRideTypeAvailability**](SandboxApi.md#SetRideTypeAvailability) | **Put** /sandbox/ridetypes/{ride_type} | Driver availability for processing ride request
[**SetRideTypes**](SandboxApi.md#SetRideTypes) | **Put** /sandbox/ridetypes | Preset types of rides for sandbox


# **SetPrimeTime**
> SetPrimeTime(request)  (empty response body)

Preset a Prime Time percentage in the region surrounding the specified location. This Prime Time percentage will be applied when requesting cost, or when requesting a ride in sandbox mode. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**SandboxPrimetime**](SandboxPrimetime.md)| Prime Time to be preset in the region surrounding the lat, lng | 

# **SetRideStatus**
> SetRideStatus(id, request) [**SandboxRideUpdate**](SandboxRideUpdate.md)

Propagate a sandbox-ride through various ride status 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string**| The ID of the ride | 
**request** | [**SandboxRideStatus**](SandboxRideStatus.md)| status to propagate the ride into | 

# **SetRideTypeAvailability**
> SetRideTypeAvailability(rideType, request)  (empty response body)

Set driver availability for the provided ride_type in the city/region surrounding the specified location 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**rideType** | **string**|  | 
**request** | [**SandboxDriverAvailability**](SandboxDriverAvailability.md)| Driver availability to be preset in the region surrounding the lat, lng | 

# **SetRideTypes**
> SetRideTypes(request) [**SandboxRideType**](SandboxRideType.md)

The sandbox-ridetypes endpoint allows you to preset the ridetypes in the region surrounding the specified latitude and longitude to allow testing different scenarios 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**SandboxRideType**](SandboxRideType.md)| Ridetypes to be preset in the region surrounding the lat, lng | 

