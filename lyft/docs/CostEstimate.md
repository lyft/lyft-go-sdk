# CostEstimate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RideType** | [**RideTypeEnum**](RideTypeEnum.md) |  | [optional] [default to null]
**DisplayName** | **string** | A human readable description of the ride type | [optional] [default to null]
**Currency** | **string** | The ISO 4217 currency code for the amount (e.g. &#39;USD&#39;) | [optional] [default to null]
**EstimatedCostCentsMin** | **int32** | Estimated lower bound for trip cost, in minor units (cents). Estimates are not guaranteed, and only provide a reasonable range based on current conditions.  | [optional] [default to null]
**EstimatedCostCentsMax** | **int32** | Estimated upper bound for trip cost, in minor units (cents). Estimates are not guaranteed, and only provide a reasonable range based on current conditions.  | [optional] [default to null]
**EstimatedDistanceMiles** | **float64** | Estimated distance for this trip  | [optional] [default to null]
**EstimatedDurationSeconds** | **int32** | Estimated time to get from the start location to the end.  | [optional] [default to null]
**IsValidEstimate** | **bool** | The validity of the cost estimate returned | [optional] [default to null]
**PrimetimePercentage** | **string** | Current Prime Time Percentage. Prime Time adds a percentage to ride costs, prior to other applicable fees. When ride requests greatly outnumber available drivers, our system will automatically turn on Prime Time. If Prime Time is inactive, the value returned will be &#39;0%&#39;. Note: The returned estimate already has Prime Time factored in. The value is returned here for reference and to allow users to confirm/accept Prime Time prior to initiating a ride.  | [optional] [default to null]
**PrimetimeConfirmationToken** | **string** | This token is needed when requesting rides. (Deprecated) | [optional] [default to null]
**CostToken** | **string** | A token that confirms the user has accepted current Prime Time and/or fixed price charges. See &#39;Request a Lyft&#39; for more details | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


