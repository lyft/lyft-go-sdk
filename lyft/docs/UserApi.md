# UserApi

All URIs are relative to *https://api.lyft.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRide**](UserApi.md#CancelRide) | **Post** /rides/{id}/cancel | Cancel a ongoing requested ride
[**GetProfile**](UserApi.md#GetProfile) | **Get** /profile | The user&#39;s general info
[**GetRide**](UserApi.md#GetRide) | **Get** /rides/{id} | Get the ride detail of a given ride ID
[**GetRideReceipt**](UserApi.md#GetRideReceipt) | **Get** /rides/{id}/receipt | Get the receipt of the rides.
[**GetRides**](UserApi.md#GetRides) | **Get** /rides | List rides
[**NewRide**](UserApi.md#NewRide) | **Post** /rides | Request a Lyft
[**SetRideDestination**](UserApi.md#SetRideDestination) | **Put** /rides/{id}/destination | Update the destination of the ride
[**SetRideRating**](UserApi.md#SetRideRating) | **Put** /rides/{id}/rating | Add the passenger&#39;s rating, feedback, and tip


# **CancelRide**
> CancelRide(id, optional)  (empty response body)

Cancel a ongoing ride which was requested earlier by providing the ride id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string**| The ID of the ride | 
**optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**CancellationRequest**](CancellationRequest.md)|  | 

# **GetProfile**
> GetProfile() [**Profile**](Profile.md)

The v1 of this endpoint returns the user's ID, v2 will return more general info about the user. We require authentication for this endpoint, so we extract the user ID from the access token. 

### Required Parameters
This endpoint does not need any parameter.

# **GetRide**
> GetRide(id) [**RideDetail**](RideDetail.md)

Get the status of a ride along with information about the driver, vehicle and price of a given ride ID 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string**| The ID of the ride | 

# **GetRideReceipt**
> GetRideReceipt(id) [**RideReceipt**](RideReceipt.md)

Get the receipt information of a processed ride by providing the ride id. Receipts will only be available to view once the payment has been processed. In the case of canceled ride, cancellation penalty is included if applicable. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string**| The ID of the ride | 

# **GetRides**
> GetRides(startTime, optional) [**RidesResponse**](RidesResponse.md)

Get a list of past & current rides for this passenger. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**startTime** | **time.Time**| Restrict to rides starting after this point in time. The earliest supported date is 2015-01-01T00:00:00+00:00  | 
**optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endTime** | **time.Time**| Restrict to rides starting before this point in time. The earliest supported date is 2015-01-01T00:00:00+00:00  | 
 **limit** | **int32**| The maximum number of rides to return. The default limit is 10 if not specified. The maximum allowed value is 50, an integer greater that 50 will return at most 50 results.  | [default to 10]

# **NewRide**
> NewRide(request) [**RideRequest**](RideRequest.md)

Request a Lyft come pick you up at the given location. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**Ride**](Ride.md)| Ride request information | 

# **SetRideDestination**
> SetRideDestination(id, request) [**Location**](Location.md)

Add or update the ride's destination. Note that the ride must still be active (not droppedOff or canceled), and that destinations on Lyft Line rides can not be changed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string**| The ID of the ride | 
**request** | [**Location**](Location.md)| The coordinates and optional address of the destination | 

# **SetRideRating**
> SetRideRating(id, request)  (empty response body)

Add the passenger's 1 to 5 star rating of the ride, optional written feedback, and optional tip amount in minor units and currency. The ride must already be dropped off, and ratings must be given within 24 hours of drop off. For purposes of display, 5 is considered the default rating. When this endpoint is successfully called, payment processing will begin. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string**| The ID of the ride | 
**request** | [**RatingRequest**](RatingRequest.md)| The rating and optional feedback | 

