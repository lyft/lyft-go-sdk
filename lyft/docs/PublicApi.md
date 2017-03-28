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

# **GetDrivers**
> GetDrivers(lat, lng) [**NearbyDriversResponse**](NearbyDriversResponse.md)

The drivers endpoint returns a list of nearby drivers' lat and lng at a given location. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**lat** | **float64**| Latitude of a location | 
**lng** | **float64**| Longitude of a location | 

# **GetETA**
> GetETA(lat, lng, optional) [**EtaEstimateResponse**](EtaEstimateResponse.md)

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

# **GetRideTypes**
> GetRideTypes(lat, lng, optional) [**RideTypesResponse**](RideTypesResponse.md)

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

