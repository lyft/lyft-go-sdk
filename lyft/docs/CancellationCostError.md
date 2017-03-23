# CancellationCostError

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Total price of the ride | [default to null]
**Currency** | **string** | The ISO 4217 currency code for the amount (e.g. USD) | [default to null]
**Description** | **string** | The description for the cost | [default to null]
**Token** | **string** | Token used to confirm the fee when cancelling a request | [optional] [default to null]
**TokenDuration** | **int32** | How long, in seconds, before the token expires | [optional] [default to null]
**Error_** | **string** | A \&quot;slug\&quot; that serves as the error code (eg. \&quot;bad_parameter\&quot;) | [optional] [default to null]
**ErrorDetail** | [**[]ErrorDetail**](ErrorDetail.md) |  | [optional] [default to null]
**ErrorDescription** | **string** | A user-friendly description of the error (appropriate to show to an end-user) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


