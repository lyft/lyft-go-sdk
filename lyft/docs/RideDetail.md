# RideDetail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RideId** | **string** | The unique ID of this ride | [optional] [default to null]
**Status** | [**RideStatusEnum**](RideStatusEnum.md) |  | [optional] [default to null]
**RideType** | [**RideTypeEnumWithOther**](RideTypeEnumWithOther.md) |  | [optional] [default to null]
**Passenger** | [**PassengerDetail**](PassengerDetail.md) |  | [optional] [default to null]
**Driver** | [**DriverDetail**](DriverDetail.md) |  | [optional] [default to null]
**Vehicle** | [**VehicleDetail**](VehicleDetail.md) |  | [optional] [default to null]
**Origin** | [**interface{}**](interface{}.md) | The *requested* location for passenger pickup | [optional] [default to null]
**Destination** | [**interface{}**](interface{}.md) | The *requested* location for passenger drop off | [optional] [default to null]
**Pickup** | [**interface{}**](interface{}.md) | The *actual* location of passenger pickup | [optional] [default to null]
**Dropoff** | [**interface{}**](interface{}.md) | The *actual* location of passenger drop off | [optional] [default to null]
**Location** | [**interface{}**](interface{}.md) | The *current* location info of the ride | [optional] [default to null]
**PrimetimePercentage** | **string** | The Prime Time percentage applied to the base price | [optional] [default to null]
**DistanceMiles** | **float32** | The distance, in miles, that this ride traveled. This field is only present after drop-off | [optional] [default to null]
**DurationSeconds** | **int32** | Duration of the ride in seconds from pickup to drop-off. This field is only present after drop-off. | [optional] [default to null]
**Price** | [**interface{}**](interface{}.md) | The total price for the current ride | [optional] [default to null]
**LineItems** | [**[]LineItem**](LineItem.md) | The break down of cost | [optional] [default to null]
**CanCancel** | **[]string** |  | [optional] [default to null]
**CanceledBy** | **string** | The role of user who canceled the ride (if applicable) | [optional] [default to null]
**CancellationPrice** | [**interface{}**](interface{}.md) | The cost of cancellation if there would be a penalty | [optional] [default to null]
**Rating** | **int32** | The rating the user left for this ride, from 1 to 5 | [optional] [default to null]
**Feedback** | **string** | The written feedback the user left for this ride | [optional] [default to null]
**PricingDetailsUrl** | **string** | The web view showing the pricing structure for the geographic area where the ride was taken           | [optional] [default to null]
**RouteUrl** | **string** | The web view showing the passenger, driver, and route for this ride. This field will only be present for rides created through this API, or that have been shared through the \&quot;Share my Route\&quot; feature  | [optional] [default to null]
**RequestedAt** | [**time.Time**](time.Time.md) | The ride requested timestamp in date and time | [optional] [default to null]
**GeneratedAt** | [**time.Time**](time.Time.md) | The request timestamp in date and time | [optional] [default to null]
**RideProfile** | [**interface{}**](interface{}.md) | Indicates whether the ride was requested from the business profile or personal profile of the user.  | [optional] [default to null]
**BeaconColor** | **string** | Hex color code of the driver AMP device. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


