# PublicApi

All URIs are relative to *https://api.lyft.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCost**](PublicApi.md#GetCost) | **Get** /cost | Cost estimates
[**GetDrivers**](PublicApi.md#GetDrivers) | **Get** /drivers | Available drivers nearby
[**GetETA**](PublicApi.md#GetETA) | **Get** /eta | Pickup ETAs
[**GetRideTypes**](PublicApi.md#GetRideTypes) | **Get** /ridetypes | Types of rides


# **GetCost**
> GetCost(startLat, startLng, optional) [**CostEstimateResponse**](CostEstimateResponse.md)

Cost estimates

Estimate the cost of taking a Lyft between two points. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**startLat** | **float64**| Latitude of the starting location | 
**startLng** | **float64**| Longitude of the starting location | 
**optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rideType** | **string**| ID of a ride type | 
 **endLat** | **float64**| Latitude of the ending location | 
 **endLng** | **float64**| Longitude of the ending location | 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDrivers**
> GetDrivers(lat, lng) [**NearbyDriversResponse**](NearbyDriversResponse.md)

Available drivers nearby

The drivers endpoint returns a list of nearby drivers' lat and lng at a given location. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**lat** | **float64**| Latitude of a location | 
**lng** | **float64**| Longitude of a location | 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetETA**
> GetETA(lat, lng, optional) [**EtaEstimateResponse**](EtaEstimateResponse.md)

Pickup ETAs

The ETA endpoint lets you know how quickly a Lyft driver can come get you 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**lat** | **float64**| Latitude of a location | 
**lng** | **float64**| Longitude of a location | 
**optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationLat** | **float64**| Latitude of destination location | 
 **destinationLng** | **float64**| Longitude of destination location | 
 **rideType** | **string**| ID of a ride type | 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRideTypes**
> GetRideTypes(lat, lng, optional) [**RideTypesResponse**](RideTypesResponse.md)

Types of rides

The ride types endpoint returns information about what kinds of Lyft rides you can request at a given location. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**lat** | **float64**| Latitude of a location | 
**lng** | **float64**| Longitude of a location | 
**optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rideType** | **string**| ID of a ride type | 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

