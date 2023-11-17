# MerchantCancellationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantCancellationNo** | **string** | The unique cancellation reference used by the Merchant (sku). | 
**MerchantOrderNo** | **string** | The unique order reference used by the Merchant (sku). | 
**Lines** | [**[]MerchantCancellationLineRequest**](MerchantCancellationLineRequest.md) |  | 
**Reason** | Pointer to **NullableString** | Reason for cancellation (text). | [optional] 
**ReasonCode** | Pointer to [**MancoReason**](MancoReason.md) |  | [optional] 

## Methods

### NewMerchantCancellationRequest

`func NewMerchantCancellationRequest(merchantCancellationNo string, merchantOrderNo string, lines []MerchantCancellationLineRequest, ) *MerchantCancellationRequest`

NewMerchantCancellationRequest instantiates a new MerchantCancellationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantCancellationRequestWithDefaults

`func NewMerchantCancellationRequestWithDefaults() *MerchantCancellationRequest`

NewMerchantCancellationRequestWithDefaults instantiates a new MerchantCancellationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantCancellationNo

`func (o *MerchantCancellationRequest) GetMerchantCancellationNo() string`

GetMerchantCancellationNo returns the MerchantCancellationNo field if non-nil, zero value otherwise.

### GetMerchantCancellationNoOk

`func (o *MerchantCancellationRequest) GetMerchantCancellationNoOk() (*string, bool)`

GetMerchantCancellationNoOk returns a tuple with the MerchantCancellationNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCancellationNo

`func (o *MerchantCancellationRequest) SetMerchantCancellationNo(v string)`

SetMerchantCancellationNo sets MerchantCancellationNo field to given value.


### GetMerchantOrderNo

`func (o *MerchantCancellationRequest) GetMerchantOrderNo() string`

GetMerchantOrderNo returns the MerchantOrderNo field if non-nil, zero value otherwise.

### GetMerchantOrderNoOk

`func (o *MerchantCancellationRequest) GetMerchantOrderNoOk() (*string, bool)`

GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderNo

`func (o *MerchantCancellationRequest) SetMerchantOrderNo(v string)`

SetMerchantOrderNo sets MerchantOrderNo field to given value.


### GetLines

`func (o *MerchantCancellationRequest) GetLines() []MerchantCancellationLineRequest`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantCancellationRequest) GetLinesOk() (*[]MerchantCancellationLineRequest, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantCancellationRequest) SetLines(v []MerchantCancellationLineRequest)`

SetLines sets Lines field to given value.


### GetReason

`func (o *MerchantCancellationRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MerchantCancellationRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MerchantCancellationRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MerchantCancellationRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *MerchantCancellationRequest) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *MerchantCancellationRequest) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetReasonCode

`func (o *MerchantCancellationRequest) GetReasonCode() MancoReason`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *MerchantCancellationRequest) GetReasonCodeOk() (*MancoReason, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *MerchantCancellationRequest) SetReasonCode(v MancoReason)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *MerchantCancellationRequest) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


