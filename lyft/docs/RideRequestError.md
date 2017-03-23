# RideRequestError

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error_** | **string** | A \&quot;slug\&quot; that serves as the error code (eg. \&quot;bad_parameter\&quot;) | [default to null]
**ErrorDetail** | [**[]ErrorDetail**](ErrorDetail.md) |  | [optional] [default to null]
**ErrorDescription** | **string** | A user-friendly description of the error (appropriate to show to an end-user) | [optional] [default to null]
**PrimetimePercentage** | **string** | Current Prime Time percentage | [optional] [default to null]
**PrimetimeMultiplier** | **float32** | Current Prime Time multiplier (eg. if primetime_percentage is 100%, primetime_multiplier will be 2.0) | [optional] [default to null]
**PrimetimeConfirmationToken** | **string** | A token that confirms the user has accepted current Prime Time charges | [optional] [default to null]
**CostToken** | **string** | Reserved for future use | [optional] [default to null]
**TokenDuration** | **string** | Validity of the token in seconds | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


