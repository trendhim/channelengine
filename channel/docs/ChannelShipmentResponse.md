# ChannelShipmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelOrderNo** | **string** | The unique order reference used by the Channel. | 
**MerchantShipmentNo** | Pointer to **NullableString** | The unique shipment reference used by the Merchant. | [optional] 
**Lines** | [**[]ChannelShipmentLineResponse**](ChannelShipmentLineResponse.md) |  | 
**CreatedAt** | Pointer to **time.Time** | The date at which the shipment was created in ChannelEngine. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date at which the shipment was last modified in ChannelEngine. | [optional] 
**ExtraData** | Pointer to **map[string]string** | Extra data on the order. Each item must have an unqiue key | [optional] 
**TrackTraceNo** | Pointer to **NullableString** | The unique shipping reference used by the Shipping carrier (track&amp;trace number). | [optional] 
**TrackTraceUrl** | Pointer to **NullableString** | A link to a page of the carrier where the customer can track the shipping of her package. | [optional] 
**ReturnTrackTraceNo** | Pointer to **NullableString** | The unique return shipping reference that may be used by the Shipping carrier (track &amp; trace number) if the shipment is returned. | [optional] 
**Method** | Pointer to **NullableString** | Shipment method: the carrier used for shipping the package. E.g. DHL, postNL. | [optional] 
**ShippedFromCountryCode** | Pointer to **NullableString** | The code of the country from where the package is being shipped. | [optional] 
**ShippedFromStockLocationId** | Pointer to **NullableInt32** | The id of the stock location where you ship the package from | [optional] 
**ShipmentDate** | Pointer to **NullableTime** | The date at which the shipment was originally created in the source system (if available). | [optional] 

## Methods

### NewChannelShipmentResponse

`func NewChannelShipmentResponse(channelOrderNo string, lines []ChannelShipmentLineResponse, ) *ChannelShipmentResponse`

NewChannelShipmentResponse instantiates a new ChannelShipmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelShipmentResponseWithDefaults

`func NewChannelShipmentResponseWithDefaults() *ChannelShipmentResponse`

NewChannelShipmentResponseWithDefaults instantiates a new ChannelShipmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelOrderNo

`func (o *ChannelShipmentResponse) GetChannelOrderNo() string`

GetChannelOrderNo returns the ChannelOrderNo field if non-nil, zero value otherwise.

### GetChannelOrderNoOk

`func (o *ChannelShipmentResponse) GetChannelOrderNoOk() (*string, bool)`

GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderNo

`func (o *ChannelShipmentResponse) SetChannelOrderNo(v string)`

SetChannelOrderNo sets ChannelOrderNo field to given value.


### GetMerchantShipmentNo

`func (o *ChannelShipmentResponse) GetMerchantShipmentNo() string`

GetMerchantShipmentNo returns the MerchantShipmentNo field if non-nil, zero value otherwise.

### GetMerchantShipmentNoOk

`func (o *ChannelShipmentResponse) GetMerchantShipmentNoOk() (*string, bool)`

GetMerchantShipmentNoOk returns a tuple with the MerchantShipmentNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantShipmentNo

`func (o *ChannelShipmentResponse) SetMerchantShipmentNo(v string)`

SetMerchantShipmentNo sets MerchantShipmentNo field to given value.

### HasMerchantShipmentNo

`func (o *ChannelShipmentResponse) HasMerchantShipmentNo() bool`

HasMerchantShipmentNo returns a boolean if a field has been set.

### SetMerchantShipmentNoNil

`func (o *ChannelShipmentResponse) SetMerchantShipmentNoNil(b bool)`

 SetMerchantShipmentNoNil sets the value for MerchantShipmentNo to be an explicit nil

### UnsetMerchantShipmentNo
`func (o *ChannelShipmentResponse) UnsetMerchantShipmentNo()`

UnsetMerchantShipmentNo ensures that no value is present for MerchantShipmentNo, not even an explicit nil
### GetLines

`func (o *ChannelShipmentResponse) GetLines() []ChannelShipmentLineResponse`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *ChannelShipmentResponse) GetLinesOk() (*[]ChannelShipmentLineResponse, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *ChannelShipmentResponse) SetLines(v []ChannelShipmentLineResponse)`

SetLines sets Lines field to given value.


### GetCreatedAt

`func (o *ChannelShipmentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChannelShipmentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChannelShipmentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChannelShipmentResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ChannelShipmentResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChannelShipmentResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChannelShipmentResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ChannelShipmentResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExtraData

`func (o *ChannelShipmentResponse) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ChannelShipmentResponse) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ChannelShipmentResponse) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *ChannelShipmentResponse) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *ChannelShipmentResponse) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *ChannelShipmentResponse) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil
### GetTrackTraceNo

`func (o *ChannelShipmentResponse) GetTrackTraceNo() string`

GetTrackTraceNo returns the TrackTraceNo field if non-nil, zero value otherwise.

### GetTrackTraceNoOk

`func (o *ChannelShipmentResponse) GetTrackTraceNoOk() (*string, bool)`

GetTrackTraceNoOk returns a tuple with the TrackTraceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackTraceNo

`func (o *ChannelShipmentResponse) SetTrackTraceNo(v string)`

SetTrackTraceNo sets TrackTraceNo field to given value.

### HasTrackTraceNo

`func (o *ChannelShipmentResponse) HasTrackTraceNo() bool`

HasTrackTraceNo returns a boolean if a field has been set.

### SetTrackTraceNoNil

`func (o *ChannelShipmentResponse) SetTrackTraceNoNil(b bool)`

 SetTrackTraceNoNil sets the value for TrackTraceNo to be an explicit nil

### UnsetTrackTraceNo
`func (o *ChannelShipmentResponse) UnsetTrackTraceNo()`

UnsetTrackTraceNo ensures that no value is present for TrackTraceNo, not even an explicit nil
### GetTrackTraceUrl

`func (o *ChannelShipmentResponse) GetTrackTraceUrl() string`

GetTrackTraceUrl returns the TrackTraceUrl field if non-nil, zero value otherwise.

### GetTrackTraceUrlOk

`func (o *ChannelShipmentResponse) GetTrackTraceUrlOk() (*string, bool)`

GetTrackTraceUrlOk returns a tuple with the TrackTraceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackTraceUrl

`func (o *ChannelShipmentResponse) SetTrackTraceUrl(v string)`

SetTrackTraceUrl sets TrackTraceUrl field to given value.

### HasTrackTraceUrl

`func (o *ChannelShipmentResponse) HasTrackTraceUrl() bool`

HasTrackTraceUrl returns a boolean if a field has been set.

### SetTrackTraceUrlNil

`func (o *ChannelShipmentResponse) SetTrackTraceUrlNil(b bool)`

 SetTrackTraceUrlNil sets the value for TrackTraceUrl to be an explicit nil

### UnsetTrackTraceUrl
`func (o *ChannelShipmentResponse) UnsetTrackTraceUrl()`

UnsetTrackTraceUrl ensures that no value is present for TrackTraceUrl, not even an explicit nil
### GetReturnTrackTraceNo

`func (o *ChannelShipmentResponse) GetReturnTrackTraceNo() string`

GetReturnTrackTraceNo returns the ReturnTrackTraceNo field if non-nil, zero value otherwise.

### GetReturnTrackTraceNoOk

`func (o *ChannelShipmentResponse) GetReturnTrackTraceNoOk() (*string, bool)`

GetReturnTrackTraceNoOk returns a tuple with the ReturnTrackTraceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTrackTraceNo

`func (o *ChannelShipmentResponse) SetReturnTrackTraceNo(v string)`

SetReturnTrackTraceNo sets ReturnTrackTraceNo field to given value.

### HasReturnTrackTraceNo

`func (o *ChannelShipmentResponse) HasReturnTrackTraceNo() bool`

HasReturnTrackTraceNo returns a boolean if a field has been set.

### SetReturnTrackTraceNoNil

`func (o *ChannelShipmentResponse) SetReturnTrackTraceNoNil(b bool)`

 SetReturnTrackTraceNoNil sets the value for ReturnTrackTraceNo to be an explicit nil

### UnsetReturnTrackTraceNo
`func (o *ChannelShipmentResponse) UnsetReturnTrackTraceNo()`

UnsetReturnTrackTraceNo ensures that no value is present for ReturnTrackTraceNo, not even an explicit nil
### GetMethod

`func (o *ChannelShipmentResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ChannelShipmentResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ChannelShipmentResponse) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ChannelShipmentResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *ChannelShipmentResponse) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *ChannelShipmentResponse) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetShippedFromCountryCode

`func (o *ChannelShipmentResponse) GetShippedFromCountryCode() string`

GetShippedFromCountryCode returns the ShippedFromCountryCode field if non-nil, zero value otherwise.

### GetShippedFromCountryCodeOk

`func (o *ChannelShipmentResponse) GetShippedFromCountryCodeOk() (*string, bool)`

GetShippedFromCountryCodeOk returns a tuple with the ShippedFromCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedFromCountryCode

`func (o *ChannelShipmentResponse) SetShippedFromCountryCode(v string)`

SetShippedFromCountryCode sets ShippedFromCountryCode field to given value.

### HasShippedFromCountryCode

`func (o *ChannelShipmentResponse) HasShippedFromCountryCode() bool`

HasShippedFromCountryCode returns a boolean if a field has been set.

### SetShippedFromCountryCodeNil

`func (o *ChannelShipmentResponse) SetShippedFromCountryCodeNil(b bool)`

 SetShippedFromCountryCodeNil sets the value for ShippedFromCountryCode to be an explicit nil

### UnsetShippedFromCountryCode
`func (o *ChannelShipmentResponse) UnsetShippedFromCountryCode()`

UnsetShippedFromCountryCode ensures that no value is present for ShippedFromCountryCode, not even an explicit nil
### GetShippedFromStockLocationId

`func (o *ChannelShipmentResponse) GetShippedFromStockLocationId() int32`

GetShippedFromStockLocationId returns the ShippedFromStockLocationId field if non-nil, zero value otherwise.

### GetShippedFromStockLocationIdOk

`func (o *ChannelShipmentResponse) GetShippedFromStockLocationIdOk() (*int32, bool)`

GetShippedFromStockLocationIdOk returns a tuple with the ShippedFromStockLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedFromStockLocationId

`func (o *ChannelShipmentResponse) SetShippedFromStockLocationId(v int32)`

SetShippedFromStockLocationId sets ShippedFromStockLocationId field to given value.

### HasShippedFromStockLocationId

`func (o *ChannelShipmentResponse) HasShippedFromStockLocationId() bool`

HasShippedFromStockLocationId returns a boolean if a field has been set.

### SetShippedFromStockLocationIdNil

`func (o *ChannelShipmentResponse) SetShippedFromStockLocationIdNil(b bool)`

 SetShippedFromStockLocationIdNil sets the value for ShippedFromStockLocationId to be an explicit nil

### UnsetShippedFromStockLocationId
`func (o *ChannelShipmentResponse) UnsetShippedFromStockLocationId()`

UnsetShippedFromStockLocationId ensures that no value is present for ShippedFromStockLocationId, not even an explicit nil
### GetShipmentDate

`func (o *ChannelShipmentResponse) GetShipmentDate() time.Time`

GetShipmentDate returns the ShipmentDate field if non-nil, zero value otherwise.

### GetShipmentDateOk

`func (o *ChannelShipmentResponse) GetShipmentDateOk() (*time.Time, bool)`

GetShipmentDateOk returns a tuple with the ShipmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDate

`func (o *ChannelShipmentResponse) SetShipmentDate(v time.Time)`

SetShipmentDate sets ShipmentDate field to given value.

### HasShipmentDate

`func (o *ChannelShipmentResponse) HasShipmentDate() bool`

HasShipmentDate returns a boolean if a field has been set.

### SetShipmentDateNil

`func (o *ChannelShipmentResponse) SetShipmentDateNil(b bool)`

 SetShipmentDateNil sets the value for ShipmentDate to be an explicit nil

### UnsetShipmentDate
`func (o *ChannelShipmentResponse) UnsetShipmentDate()`

UnsetShipmentDate ensures that no value is present for ShipmentDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


