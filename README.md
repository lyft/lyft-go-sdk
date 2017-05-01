# Lyft Go SDK

lyft-go-sdk is a Go client library for accessing the [Lyft API](https://developer.lyft.com/docs).

## Registration
- You must first create a Lyft Developer account [here](https://www.lyft.com/developers).
- Once registered, you will be assigned a Client ID and a Client Secret.

## Installation

```console
go get github.com/lyft/lyft-go-sdk/lyft
```

## Usage

The `lyft` package provides a `Client` for accessing the Lyft API. Authentication is handled by the `http.Client` passed to `NewAPIClient` to handle OAuth2. The Lyft API uses the [Client Credentials](https://developer.lyft.com/docs/authentication#section-client-credentials-2-legged-flow-for-public-endpoints) grant for public endpoints and the [Authorization Code](https://developer.lyft.com/docs/authentication#section-3-legged-flow-for-accessing-user-specific-endpoints) grant for endpoints that require a user context; see example for the [former](https://github.com/lyft/lyft-go-sdk/blob/master/examples/public/main.go) and the [latter](https://github.com/lyft/lyft-go-sdk/blob/master/examples/user/main.go). 

```go
client := lyft.NewAPIClient(httpClient, "my-sample-app")

// Get a list of rides taken by the authenticated user within the past 30 days.
result, resp, err := client.UserApi.GetRides(time.Now().AddDate(0, 0, -30), nil)

// Get driver ETA for the specified location.
result, _, err := client.PublicApi.GetETA(37.7763, -122.3918, nil)

// Preset Prime Time to 25% in the specified location.
resp, err := client.SandboxApi.SetPrimeTime(lyft.SandboxPrimetime{
	Lat:                 37.7884,
	Lng:                 -122.4076,
	PrimetimePercentage: "25%",
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.lyft.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*PublicApi* | [**GetCost**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/PublicApi.md#getcost) | **Get** /cost | Cost estimates
*PublicApi* | [**GetDrivers**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/PublicApi.md#getdrivers) | **Get** /drivers | Available drivers nearby
*PublicApi* | [**GetETA**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/PublicApi.md#geteta) | **Get** /eta | Pickup ETAs
*PublicApi* | [**GetRideTypes**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/PublicApi.md#getridetypes) | **Get** /ridetypes | Types of rides
*SandboxApi* | [**SetPrimeTime**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/SandboxApi.md#setprimetime) | **Put** /sandbox/primetime | Preset Prime Time percentage
*SandboxApi* | [**SetRideStatus**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/SandboxApi.md#setridestatus) | **Put** /sandbox/rides/{id} | Propagate ride through ride status
*SandboxApi* | [**SetRideTypeAvailability**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/SandboxApi.md#setridetypeavailability) | **Put** /sandbox/ridetypes/{ride_type} | Driver availability for processing ride request
*SandboxApi* | [**SetRideTypes**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/SandboxApi.md#setridetypes) | **Put** /sandbox/ridetypes | Preset types of rides for sandbox
*UserApi* | [**CancelRide**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/UserApi.md#cancelride) | **Post** /rides/{id}/cancel | Cancel a ongoing requested ride
*UserApi* | [**GetProfile**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/UserApi.md#getprofile) | **Get** /profile | The user&#39;s general info
*UserApi* | [**GetRide**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/UserApi.md#getride) | **Get** /rides/{id} | Get the ride detail of a given ride ID
*UserApi* | [**GetRideReceipt**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/UserApi.md#getridereceipt) | **Get** /rides/{id}/receipt | Get the receipt of the rides.
*UserApi* | [**GetRides**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/UserApi.md#getrides) | **Get** /rides | List rides
*UserApi* | [**NewRide**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/UserApi.md#newride) | **Post** /rides | Request a Lyft
*UserApi* | [**SetRideDestination**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/UserApi.md#setridedestination) | **Put** /rides/{id}/destination | Update the destination of the ride
*UserApi* | [**SetRideRating**](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/UserApi.md#setriderating) | **Put** /rides/{id}/rating | Add the passenger&#39;s rating, feedback, and tip


## Documentation For Models

 - [ApiError](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/ApiError.md)
 - [CancellationCost](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/CancellationCost.md)
 - [CancellationCostError](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/CancellationCostError.md)
 - [CancellationRequest](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/CancellationRequest.md)
 - [Charge](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/Charge.md)
 - [Cost](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/Cost.md)
 - [CostEstimate](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/CostEstimate.md)
 - [CostEstimateResponse](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/CostEstimateResponse.md)
 - [CurrentRideLocation](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/CurrentRideLocation.md)
 - [DriverDetail](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/DriverDetail.md)
 - [ErrorDetail](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/ErrorDetail.md)
 - [Eta](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/Eta.md)
 - [EtaEstimateResponse](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/EtaEstimateResponse.md)
 - [LatLng](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/LatLng.md)
 - [LineItem](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/LineItem.md)
 - [Location](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/Location.md)
 - [NearbyDriver](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/NearbyDriver.md)
 - [NearbyDriversByRideType](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/NearbyDriversByRideType.md)
 - [NearbyDriversResponse](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/NearbyDriversResponse.md)
 - [PassengerDetail](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/PassengerDetail.md)
 - [PickupDropoffLocation](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/PickupDropoffLocation.md)
 - [PricingDetails](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/PricingDetails.md)
 - [Profile](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/Profile.md)
 - [RatingRequest](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/RatingRequest.md)
 - [Ride](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/Ride.md)
 - [RideDetail](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/RideDetail.md)
 - [RideLocation](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/RideLocation.md)
 - [RideProfileEnum](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/RideProfileEnum.md)
 - [RideReceipt](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/RideReceipt.md)
 - [RideRequest](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/RideRequest.md)
 - [RideRequestError](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/RideRequestError.md)
 - [RideStatusEnum](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/RideStatusEnum.md)
 - [RideType](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/RideType.md)
 - [RideTypeEnum](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/RideTypeEnum.md)
 - [RideTypeEnumWithOther](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/RideTypeEnumWithOther.md)
 - [RideTypesResponse](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/RideTypesResponse.md)
 - [RidesResponse](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/RidesResponse.md)
 - [SandboxDriverAvailability](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/SandboxDriverAvailability.md)
 - [SandboxPrimetime](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/SandboxPrimetime.md)
 - [SandboxRideStatus](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/SandboxRideStatus.md)
 - [SandboxRideType](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/SandboxRideType.md)
 - [SandboxRideUpdate](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/SandboxRideUpdate.md)
 - [Tip](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/Tip.md)
 - [TipParams](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/TipParams.md)
 - [UserDetail](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/UserDetail.md)
 - [VehicleDetail](https://github.com/lyft/lyft-go-sdk/tree/master/lyft/docs/VehicleDetail.md)


## Support

If you're looking for help configuring or using the SDK, or if you have general questions related to our APIs, the Lyft Developer Platform team provides support through our [forum](https://developer.lyft.com/discuss) as well as on Stack Overflow (using the `lyft-api` tag).

## Reporting security vulnerabilities

If you've found a vulnerability or a potential vulnerability in the Lyft Go SDK,
please let us know at security@lyft.com. We'll send a confirmation email to
acknowledge your report, and we'll send an additional email when we've
identified the issue positively or negatively.