# RideReceipt

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RideId** | **string** | The unique ID of this ride | [optional] [default to null]
**Price** | [**interface{}**](interface{}.md) | The total price for the current ride | [optional] [default to null]
**LineItems** | [**[]LineItem**](LineItem.md) | The break down of line items | [optional] [default to null]
**Charges** | [**[]Charge**](Charge.md) | The break down of charge method | [optional] [default to null]
**RequestedAt** | [**time.Time**](time.Time.md) | The ride requested timestamp in date and time | [optional] [default to null]
**RideProfile** | [**interface{}**](interface{}.md) | Indicates whether the ride was requested from the business profile or personal profile of the user.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


