# MerchantShipmentTrackingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | Shipment method (carrier). | 
**TrackTraceNo** | **string** | The unique shipping reference used by the Shipping carrier (track &amp; trace number). | 
**ReturnTrackTraceNo** | Pointer to **NullableString** | The unique return shipping reference that may be used by the Shipping carrier (track &amp; trace number) if the shipment is returned. | [optional] 
**TrackTraceUrl** | Pointer to **NullableString** | A link to a page of the carrier where the customer can track the shipping of her package. | [optional] 
**ShippedFromCountryCode** | Pointer to **NullableString** | The code of the country from where the package is being shipped. | [optional] 

## Methods

### NewMerchantShipmentTrackingRequest

`func NewMerchantShipmentTrackingRequest(method string, trackTraceNo string, ) *MerchantShipmentTrackingRequest`

NewMerchantShipmentTrackingRequest instantiates a new MerchantShipmentTrackingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantShipmentTrackingRequestWithDefaults

`func NewMerchantShipmentTrackingRequestWithDefaults() *MerchantShipmentTrackingRequest`

NewMerchantShipmentTrackingRequestWithDefaults instantiates a new MerchantShipmentTrackingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *MerchantShipmentTrackingRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MerchantShipmentTrackingRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MerchantShipmentTrackingRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetTrackTraceNo

`func (o *MerchantShipmentTrackingRequest) GetTrackTraceNo() string`

GetTrackTraceNo returns the TrackTraceNo field if non-nil, zero value otherwise.

### GetTrackTraceNoOk

`func (o *MerchantShipmentTrackingRequest) GetTrackTraceNoOk() (*string, bool)`

GetTrackTraceNoOk returns a tuple with the TrackTraceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackTraceNo

`func (o *MerchantShipmentTrackingRequest) SetTrackTraceNo(v string)`

SetTrackTraceNo sets TrackTraceNo field to given value.


### GetReturnTrackTraceNo

`func (o *MerchantShipmentTrackingRequest) GetReturnTrackTraceNo() string`

GetReturnTrackTraceNo returns the ReturnTrackTraceNo field if non-nil, zero value otherwise.

### GetReturnTrackTraceNoOk

`func (o *MerchantShipmentTrackingRequest) GetReturnTrackTraceNoOk() (*string, bool)`

GetReturnTrackTraceNoOk returns a tuple with the ReturnTrackTraceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTrackTraceNo

`func (o *MerchantShipmentTrackingRequest) SetReturnTrackTraceNo(v string)`

SetReturnTrackTraceNo sets ReturnTrackTraceNo field to given value.

### HasReturnTrackTraceNo

`func (o *MerchantShipmentTrackingRequest) HasReturnTrackTraceNo() bool`

HasReturnTrackTraceNo returns a boolean if a field has been set.

### SetReturnTrackTraceNoNil

`func (o *MerchantShipmentTrackingRequest) SetReturnTrackTraceNoNil(b bool)`

 SetReturnTrackTraceNoNil sets the value for ReturnTrackTraceNo to be an explicit nil

### UnsetReturnTrackTraceNo
`func (o *MerchantShipmentTrackingRequest) UnsetReturnTrackTraceNo()`

UnsetReturnTrackTraceNo ensures that no value is present for ReturnTrackTraceNo, not even an explicit nil
### GetTrackTraceUrl

`func (o *MerchantShipmentTrackingRequest) GetTrackTraceUrl() string`

GetTrackTraceUrl returns the TrackTraceUrl field if non-nil, zero value otherwise.

### GetTrackTraceUrlOk

`func (o *MerchantShipmentTrackingRequest) GetTrackTraceUrlOk() (*string, bool)`

GetTrackTraceUrlOk returns a tuple with the TrackTraceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackTraceUrl

`func (o *MerchantShipmentTrackingRequest) SetTrackTraceUrl(v string)`

SetTrackTraceUrl sets TrackTraceUrl field to given value.

### HasTrackTraceUrl

`func (o *MerchantShipmentTrackingRequest) HasTrackTraceUrl() bool`

HasTrackTraceUrl returns a boolean if a field has been set.

### SetTrackTraceUrlNil

`func (o *MerchantShipmentTrackingRequest) SetTrackTraceUrlNil(b bool)`

 SetTrackTraceUrlNil sets the value for TrackTraceUrl to be an explicit nil

### UnsetTrackTraceUrl
`func (o *MerchantShipmentTrackingRequest) UnsetTrackTraceUrl()`

UnsetTrackTraceUrl ensures that no value is present for TrackTraceUrl, not even an explicit nil
### GetShippedFromCountryCode

`func (o *MerchantShipmentTrackingRequest) GetShippedFromCountryCode() string`

GetShippedFromCountryCode returns the ShippedFromCountryCode field if non-nil, zero value otherwise.

### GetShippedFromCountryCodeOk

`func (o *MerchantShipmentTrackingRequest) GetShippedFromCountryCodeOk() (*string, bool)`

GetShippedFromCountryCodeOk returns a tuple with the ShippedFromCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedFromCountryCode

`func (o *MerchantShipmentTrackingRequest) SetShippedFromCountryCode(v string)`

SetShippedFromCountryCode sets ShippedFromCountryCode field to given value.

### HasShippedFromCountryCode

`func (o *MerchantShipmentTrackingRequest) HasShippedFromCountryCode() bool`

HasShippedFromCountryCode returns a boolean if a field has been set.

### SetShippedFromCountryCodeNil

`func (o *MerchantShipmentTrackingRequest) SetShippedFromCountryCodeNil(b bool)`

 SetShippedFromCountryCodeNil sets the value for ShippedFromCountryCode to be an explicit nil

### UnsetShippedFromCountryCode
`func (o *MerchantShipmentTrackingRequest) UnsetShippedFromCountryCode()`

UnsetShippedFromCountryCode ensures that no value is present for ShippedFromCountryCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


