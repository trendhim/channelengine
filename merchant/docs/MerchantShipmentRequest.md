# MerchantShipmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantShipmentNo** | **string** | The unique shipment reference used by the Merchant. | 
**MerchantOrderNo** | **string** | The unique order reference used by the Merchant. | 
**Lines** | [**[]MerchantShipmentLineRequest**](MerchantShipmentLineRequest.md) |  | 
**ExtraData** | Pointer to **map[string]string** | Extra data on the order. Each item must have an unqiue key | [optional] 
**TrackTraceNo** | Pointer to **NullableString** | The unique shipping reference used by the Shipping carrier (track&amp;trace number). | [optional] 
**TrackTraceUrl** | Pointer to **NullableString** | A link to a page of the carrier where the customer can track the shipping of her package. | [optional] 
**ReturnTrackTraceNo** | Pointer to **NullableString** | The unique return shipping reference that may be used by the Shipping carrier (track &amp; trace number) if the shipment is returned. | [optional] 
**Method** | Pointer to **NullableString** | Shipment method: the carrier used for shipping the package. E.g. DHL, postNL. | [optional] 
**ShippedFromCountryCode** | Pointer to **NullableString** | The code of the country from where the package is being shipped. | [optional] 
**ShippedFromStockLocationId** | Pointer to **NullableInt32** | The id of the stock location where you ship the package from | [optional] 
**ShipmentDate** | Pointer to **NullableTime** | The date at which the shipment was originally created in the source system (if available). | [optional] 

## Methods

### NewMerchantShipmentRequest

`func NewMerchantShipmentRequest(merchantShipmentNo string, merchantOrderNo string, lines []MerchantShipmentLineRequest, ) *MerchantShipmentRequest`

NewMerchantShipmentRequest instantiates a new MerchantShipmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantShipmentRequestWithDefaults

`func NewMerchantShipmentRequestWithDefaults() *MerchantShipmentRequest`

NewMerchantShipmentRequestWithDefaults instantiates a new MerchantShipmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantShipmentNo

`func (o *MerchantShipmentRequest) GetMerchantShipmentNo() string`

GetMerchantShipmentNo returns the MerchantShipmentNo field if non-nil, zero value otherwise.

### GetMerchantShipmentNoOk

`func (o *MerchantShipmentRequest) GetMerchantShipmentNoOk() (*string, bool)`

GetMerchantShipmentNoOk returns a tuple with the MerchantShipmentNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantShipmentNo

`func (o *MerchantShipmentRequest) SetMerchantShipmentNo(v string)`

SetMerchantShipmentNo sets MerchantShipmentNo field to given value.


### GetMerchantOrderNo

`func (o *MerchantShipmentRequest) GetMerchantOrderNo() string`

GetMerchantOrderNo returns the MerchantOrderNo field if non-nil, zero value otherwise.

### GetMerchantOrderNoOk

`func (o *MerchantShipmentRequest) GetMerchantOrderNoOk() (*string, bool)`

GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderNo

`func (o *MerchantShipmentRequest) SetMerchantOrderNo(v string)`

SetMerchantOrderNo sets MerchantOrderNo field to given value.


### GetLines

`func (o *MerchantShipmentRequest) GetLines() []MerchantShipmentLineRequest`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantShipmentRequest) GetLinesOk() (*[]MerchantShipmentLineRequest, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantShipmentRequest) SetLines(v []MerchantShipmentLineRequest)`

SetLines sets Lines field to given value.


### GetExtraData

`func (o *MerchantShipmentRequest) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantShipmentRequest) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantShipmentRequest) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantShipmentRequest) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantShipmentRequest) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantShipmentRequest) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil
### GetTrackTraceNo

`func (o *MerchantShipmentRequest) GetTrackTraceNo() string`

GetTrackTraceNo returns the TrackTraceNo field if non-nil, zero value otherwise.

### GetTrackTraceNoOk

`func (o *MerchantShipmentRequest) GetTrackTraceNoOk() (*string, bool)`

GetTrackTraceNoOk returns a tuple with the TrackTraceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackTraceNo

`func (o *MerchantShipmentRequest) SetTrackTraceNo(v string)`

SetTrackTraceNo sets TrackTraceNo field to given value.

### HasTrackTraceNo

`func (o *MerchantShipmentRequest) HasTrackTraceNo() bool`

HasTrackTraceNo returns a boolean if a field has been set.

### SetTrackTraceNoNil

`func (o *MerchantShipmentRequest) SetTrackTraceNoNil(b bool)`

 SetTrackTraceNoNil sets the value for TrackTraceNo to be an explicit nil

### UnsetTrackTraceNo
`func (o *MerchantShipmentRequest) UnsetTrackTraceNo()`

UnsetTrackTraceNo ensures that no value is present for TrackTraceNo, not even an explicit nil
### GetTrackTraceUrl

`func (o *MerchantShipmentRequest) GetTrackTraceUrl() string`

GetTrackTraceUrl returns the TrackTraceUrl field if non-nil, zero value otherwise.

### GetTrackTraceUrlOk

`func (o *MerchantShipmentRequest) GetTrackTraceUrlOk() (*string, bool)`

GetTrackTraceUrlOk returns a tuple with the TrackTraceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackTraceUrl

`func (o *MerchantShipmentRequest) SetTrackTraceUrl(v string)`

SetTrackTraceUrl sets TrackTraceUrl field to given value.

### HasTrackTraceUrl

`func (o *MerchantShipmentRequest) HasTrackTraceUrl() bool`

HasTrackTraceUrl returns a boolean if a field has been set.

### SetTrackTraceUrlNil

`func (o *MerchantShipmentRequest) SetTrackTraceUrlNil(b bool)`

 SetTrackTraceUrlNil sets the value for TrackTraceUrl to be an explicit nil

### UnsetTrackTraceUrl
`func (o *MerchantShipmentRequest) UnsetTrackTraceUrl()`

UnsetTrackTraceUrl ensures that no value is present for TrackTraceUrl, not even an explicit nil
### GetReturnTrackTraceNo

`func (o *MerchantShipmentRequest) GetReturnTrackTraceNo() string`

GetReturnTrackTraceNo returns the ReturnTrackTraceNo field if non-nil, zero value otherwise.

### GetReturnTrackTraceNoOk

`func (o *MerchantShipmentRequest) GetReturnTrackTraceNoOk() (*string, bool)`

GetReturnTrackTraceNoOk returns a tuple with the ReturnTrackTraceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTrackTraceNo

`func (o *MerchantShipmentRequest) SetReturnTrackTraceNo(v string)`

SetReturnTrackTraceNo sets ReturnTrackTraceNo field to given value.

### HasReturnTrackTraceNo

`func (o *MerchantShipmentRequest) HasReturnTrackTraceNo() bool`

HasReturnTrackTraceNo returns a boolean if a field has been set.

### SetReturnTrackTraceNoNil

`func (o *MerchantShipmentRequest) SetReturnTrackTraceNoNil(b bool)`

 SetReturnTrackTraceNoNil sets the value for ReturnTrackTraceNo to be an explicit nil

### UnsetReturnTrackTraceNo
`func (o *MerchantShipmentRequest) UnsetReturnTrackTraceNo()`

UnsetReturnTrackTraceNo ensures that no value is present for ReturnTrackTraceNo, not even an explicit nil
### GetMethod

`func (o *MerchantShipmentRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MerchantShipmentRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MerchantShipmentRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *MerchantShipmentRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *MerchantShipmentRequest) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *MerchantShipmentRequest) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetShippedFromCountryCode

`func (o *MerchantShipmentRequest) GetShippedFromCountryCode() string`

GetShippedFromCountryCode returns the ShippedFromCountryCode field if non-nil, zero value otherwise.

### GetShippedFromCountryCodeOk

`func (o *MerchantShipmentRequest) GetShippedFromCountryCodeOk() (*string, bool)`

GetShippedFromCountryCodeOk returns a tuple with the ShippedFromCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedFromCountryCode

`func (o *MerchantShipmentRequest) SetShippedFromCountryCode(v string)`

SetShippedFromCountryCode sets ShippedFromCountryCode field to given value.

### HasShippedFromCountryCode

`func (o *MerchantShipmentRequest) HasShippedFromCountryCode() bool`

HasShippedFromCountryCode returns a boolean if a field has been set.

### SetShippedFromCountryCodeNil

`func (o *MerchantShipmentRequest) SetShippedFromCountryCodeNil(b bool)`

 SetShippedFromCountryCodeNil sets the value for ShippedFromCountryCode to be an explicit nil

### UnsetShippedFromCountryCode
`func (o *MerchantShipmentRequest) UnsetShippedFromCountryCode()`

UnsetShippedFromCountryCode ensures that no value is present for ShippedFromCountryCode, not even an explicit nil
### GetShippedFromStockLocationId

`func (o *MerchantShipmentRequest) GetShippedFromStockLocationId() int32`

GetShippedFromStockLocationId returns the ShippedFromStockLocationId field if non-nil, zero value otherwise.

### GetShippedFromStockLocationIdOk

`func (o *MerchantShipmentRequest) GetShippedFromStockLocationIdOk() (*int32, bool)`

GetShippedFromStockLocationIdOk returns a tuple with the ShippedFromStockLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedFromStockLocationId

`func (o *MerchantShipmentRequest) SetShippedFromStockLocationId(v int32)`

SetShippedFromStockLocationId sets ShippedFromStockLocationId field to given value.

### HasShippedFromStockLocationId

`func (o *MerchantShipmentRequest) HasShippedFromStockLocationId() bool`

HasShippedFromStockLocationId returns a boolean if a field has been set.

### SetShippedFromStockLocationIdNil

`func (o *MerchantShipmentRequest) SetShippedFromStockLocationIdNil(b bool)`

 SetShippedFromStockLocationIdNil sets the value for ShippedFromStockLocationId to be an explicit nil

### UnsetShippedFromStockLocationId
`func (o *MerchantShipmentRequest) UnsetShippedFromStockLocationId()`

UnsetShippedFromStockLocationId ensures that no value is present for ShippedFromStockLocationId, not even an explicit nil
### GetShipmentDate

`func (o *MerchantShipmentRequest) GetShipmentDate() time.Time`

GetShipmentDate returns the ShipmentDate field if non-nil, zero value otherwise.

### GetShipmentDateOk

`func (o *MerchantShipmentRequest) GetShipmentDateOk() (*time.Time, bool)`

GetShipmentDateOk returns a tuple with the ShipmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDate

`func (o *MerchantShipmentRequest) SetShipmentDate(v time.Time)`

SetShipmentDate sets ShipmentDate field to given value.

### HasShipmentDate

`func (o *MerchantShipmentRequest) HasShipmentDate() bool`

HasShipmentDate returns a boolean if a field has been set.

### SetShipmentDateNil

`func (o *MerchantShipmentRequest) SetShipmentDateNil(b bool)`

 SetShipmentDateNil sets the value for ShipmentDate to be an explicit nil

### UnsetShipmentDate
`func (o *MerchantShipmentRequest) UnsetShipmentDate()`

UnsetShipmentDate ensures that no value is present for ShipmentDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


