# MerchantShipmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantShipmentNo** | **string** | The unique shipment reference used by the Merchant. | 
**MerchantOrderNo** | Pointer to **NullableString** | The unique order reference used by the Merchant. | [optional] 
**ChannelShipmentNo** | Pointer to **NullableString** | The unique shipment reference used by the Channel. | [optional] 
**ChannelOrderNo** | Pointer to **NullableString** | The unique order reference used by the Channel. | [optional] 
**ChannelId** | Pointer to **NullableInt32** | The unique ID of the channel for this specific environment/account. | [optional] 
**GlobalChannelId** | Pointer to **NullableInt32** | The unique ID of the channel across all of ChannelEngine. | [optional] 
**Lines** | Pointer to [**[]MerchantShipmentLineResponse**](MerchantShipmentLineResponse.md) |  | [optional] 
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

### NewMerchantShipmentResponse

`func NewMerchantShipmentResponse(merchantShipmentNo string, ) *MerchantShipmentResponse`

NewMerchantShipmentResponse instantiates a new MerchantShipmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantShipmentResponseWithDefaults

`func NewMerchantShipmentResponseWithDefaults() *MerchantShipmentResponse`

NewMerchantShipmentResponseWithDefaults instantiates a new MerchantShipmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantShipmentNo

`func (o *MerchantShipmentResponse) GetMerchantShipmentNo() string`

GetMerchantShipmentNo returns the MerchantShipmentNo field if non-nil, zero value otherwise.

### GetMerchantShipmentNoOk

`func (o *MerchantShipmentResponse) GetMerchantShipmentNoOk() (*string, bool)`

GetMerchantShipmentNoOk returns a tuple with the MerchantShipmentNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantShipmentNo

`func (o *MerchantShipmentResponse) SetMerchantShipmentNo(v string)`

SetMerchantShipmentNo sets MerchantShipmentNo field to given value.


### GetMerchantOrderNo

`func (o *MerchantShipmentResponse) GetMerchantOrderNo() string`

GetMerchantOrderNo returns the MerchantOrderNo field if non-nil, zero value otherwise.

### GetMerchantOrderNoOk

`func (o *MerchantShipmentResponse) GetMerchantOrderNoOk() (*string, bool)`

GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderNo

`func (o *MerchantShipmentResponse) SetMerchantOrderNo(v string)`

SetMerchantOrderNo sets MerchantOrderNo field to given value.

### HasMerchantOrderNo

`func (o *MerchantShipmentResponse) HasMerchantOrderNo() bool`

HasMerchantOrderNo returns a boolean if a field has been set.

### SetMerchantOrderNoNil

`func (o *MerchantShipmentResponse) SetMerchantOrderNoNil(b bool)`

 SetMerchantOrderNoNil sets the value for MerchantOrderNo to be an explicit nil

### UnsetMerchantOrderNo
`func (o *MerchantShipmentResponse) UnsetMerchantOrderNo()`

UnsetMerchantOrderNo ensures that no value is present for MerchantOrderNo, not even an explicit nil
### GetChannelShipmentNo

`func (o *MerchantShipmentResponse) GetChannelShipmentNo() string`

GetChannelShipmentNo returns the ChannelShipmentNo field if non-nil, zero value otherwise.

### GetChannelShipmentNoOk

`func (o *MerchantShipmentResponse) GetChannelShipmentNoOk() (*string, bool)`

GetChannelShipmentNoOk returns a tuple with the ChannelShipmentNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelShipmentNo

`func (o *MerchantShipmentResponse) SetChannelShipmentNo(v string)`

SetChannelShipmentNo sets ChannelShipmentNo field to given value.

### HasChannelShipmentNo

`func (o *MerchantShipmentResponse) HasChannelShipmentNo() bool`

HasChannelShipmentNo returns a boolean if a field has been set.

### SetChannelShipmentNoNil

`func (o *MerchantShipmentResponse) SetChannelShipmentNoNil(b bool)`

 SetChannelShipmentNoNil sets the value for ChannelShipmentNo to be an explicit nil

### UnsetChannelShipmentNo
`func (o *MerchantShipmentResponse) UnsetChannelShipmentNo()`

UnsetChannelShipmentNo ensures that no value is present for ChannelShipmentNo, not even an explicit nil
### GetChannelOrderNo

`func (o *MerchantShipmentResponse) GetChannelOrderNo() string`

GetChannelOrderNo returns the ChannelOrderNo field if non-nil, zero value otherwise.

### GetChannelOrderNoOk

`func (o *MerchantShipmentResponse) GetChannelOrderNoOk() (*string, bool)`

GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderNo

`func (o *MerchantShipmentResponse) SetChannelOrderNo(v string)`

SetChannelOrderNo sets ChannelOrderNo field to given value.

### HasChannelOrderNo

`func (o *MerchantShipmentResponse) HasChannelOrderNo() bool`

HasChannelOrderNo returns a boolean if a field has been set.

### SetChannelOrderNoNil

`func (o *MerchantShipmentResponse) SetChannelOrderNoNil(b bool)`

 SetChannelOrderNoNil sets the value for ChannelOrderNo to be an explicit nil

### UnsetChannelOrderNo
`func (o *MerchantShipmentResponse) UnsetChannelOrderNo()`

UnsetChannelOrderNo ensures that no value is present for ChannelOrderNo, not even an explicit nil
### GetChannelId

`func (o *MerchantShipmentResponse) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *MerchantShipmentResponse) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *MerchantShipmentResponse) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *MerchantShipmentResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *MerchantShipmentResponse) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *MerchantShipmentResponse) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetGlobalChannelId

`func (o *MerchantShipmentResponse) GetGlobalChannelId() int32`

GetGlobalChannelId returns the GlobalChannelId field if non-nil, zero value otherwise.

### GetGlobalChannelIdOk

`func (o *MerchantShipmentResponse) GetGlobalChannelIdOk() (*int32, bool)`

GetGlobalChannelIdOk returns a tuple with the GlobalChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalChannelId

`func (o *MerchantShipmentResponse) SetGlobalChannelId(v int32)`

SetGlobalChannelId sets GlobalChannelId field to given value.

### HasGlobalChannelId

`func (o *MerchantShipmentResponse) HasGlobalChannelId() bool`

HasGlobalChannelId returns a boolean if a field has been set.

### SetGlobalChannelIdNil

`func (o *MerchantShipmentResponse) SetGlobalChannelIdNil(b bool)`

 SetGlobalChannelIdNil sets the value for GlobalChannelId to be an explicit nil

### UnsetGlobalChannelId
`func (o *MerchantShipmentResponse) UnsetGlobalChannelId()`

UnsetGlobalChannelId ensures that no value is present for GlobalChannelId, not even an explicit nil
### GetLines

`func (o *MerchantShipmentResponse) GetLines() []MerchantShipmentLineResponse`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantShipmentResponse) GetLinesOk() (*[]MerchantShipmentLineResponse, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantShipmentResponse) SetLines(v []MerchantShipmentLineResponse)`

SetLines sets Lines field to given value.

### HasLines

`func (o *MerchantShipmentResponse) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *MerchantShipmentResponse) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *MerchantShipmentResponse) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil
### GetCreatedAt

`func (o *MerchantShipmentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MerchantShipmentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MerchantShipmentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MerchantShipmentResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MerchantShipmentResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MerchantShipmentResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MerchantShipmentResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MerchantShipmentResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExtraData

`func (o *MerchantShipmentResponse) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantShipmentResponse) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantShipmentResponse) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantShipmentResponse) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantShipmentResponse) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantShipmentResponse) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil
### GetTrackTraceNo

`func (o *MerchantShipmentResponse) GetTrackTraceNo() string`

GetTrackTraceNo returns the TrackTraceNo field if non-nil, zero value otherwise.

### GetTrackTraceNoOk

`func (o *MerchantShipmentResponse) GetTrackTraceNoOk() (*string, bool)`

GetTrackTraceNoOk returns a tuple with the TrackTraceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackTraceNo

`func (o *MerchantShipmentResponse) SetTrackTraceNo(v string)`

SetTrackTraceNo sets TrackTraceNo field to given value.

### HasTrackTraceNo

`func (o *MerchantShipmentResponse) HasTrackTraceNo() bool`

HasTrackTraceNo returns a boolean if a field has been set.

### SetTrackTraceNoNil

`func (o *MerchantShipmentResponse) SetTrackTraceNoNil(b bool)`

 SetTrackTraceNoNil sets the value for TrackTraceNo to be an explicit nil

### UnsetTrackTraceNo
`func (o *MerchantShipmentResponse) UnsetTrackTraceNo()`

UnsetTrackTraceNo ensures that no value is present for TrackTraceNo, not even an explicit nil
### GetTrackTraceUrl

`func (o *MerchantShipmentResponse) GetTrackTraceUrl() string`

GetTrackTraceUrl returns the TrackTraceUrl field if non-nil, zero value otherwise.

### GetTrackTraceUrlOk

`func (o *MerchantShipmentResponse) GetTrackTraceUrlOk() (*string, bool)`

GetTrackTraceUrlOk returns a tuple with the TrackTraceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackTraceUrl

`func (o *MerchantShipmentResponse) SetTrackTraceUrl(v string)`

SetTrackTraceUrl sets TrackTraceUrl field to given value.

### HasTrackTraceUrl

`func (o *MerchantShipmentResponse) HasTrackTraceUrl() bool`

HasTrackTraceUrl returns a boolean if a field has been set.

### SetTrackTraceUrlNil

`func (o *MerchantShipmentResponse) SetTrackTraceUrlNil(b bool)`

 SetTrackTraceUrlNil sets the value for TrackTraceUrl to be an explicit nil

### UnsetTrackTraceUrl
`func (o *MerchantShipmentResponse) UnsetTrackTraceUrl()`

UnsetTrackTraceUrl ensures that no value is present for TrackTraceUrl, not even an explicit nil
### GetReturnTrackTraceNo

`func (o *MerchantShipmentResponse) GetReturnTrackTraceNo() string`

GetReturnTrackTraceNo returns the ReturnTrackTraceNo field if non-nil, zero value otherwise.

### GetReturnTrackTraceNoOk

`func (o *MerchantShipmentResponse) GetReturnTrackTraceNoOk() (*string, bool)`

GetReturnTrackTraceNoOk returns a tuple with the ReturnTrackTraceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTrackTraceNo

`func (o *MerchantShipmentResponse) SetReturnTrackTraceNo(v string)`

SetReturnTrackTraceNo sets ReturnTrackTraceNo field to given value.

### HasReturnTrackTraceNo

`func (o *MerchantShipmentResponse) HasReturnTrackTraceNo() bool`

HasReturnTrackTraceNo returns a boolean if a field has been set.

### SetReturnTrackTraceNoNil

`func (o *MerchantShipmentResponse) SetReturnTrackTraceNoNil(b bool)`

 SetReturnTrackTraceNoNil sets the value for ReturnTrackTraceNo to be an explicit nil

### UnsetReturnTrackTraceNo
`func (o *MerchantShipmentResponse) UnsetReturnTrackTraceNo()`

UnsetReturnTrackTraceNo ensures that no value is present for ReturnTrackTraceNo, not even an explicit nil
### GetMethod

`func (o *MerchantShipmentResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MerchantShipmentResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MerchantShipmentResponse) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *MerchantShipmentResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *MerchantShipmentResponse) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *MerchantShipmentResponse) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetShippedFromCountryCode

`func (o *MerchantShipmentResponse) GetShippedFromCountryCode() string`

GetShippedFromCountryCode returns the ShippedFromCountryCode field if non-nil, zero value otherwise.

### GetShippedFromCountryCodeOk

`func (o *MerchantShipmentResponse) GetShippedFromCountryCodeOk() (*string, bool)`

GetShippedFromCountryCodeOk returns a tuple with the ShippedFromCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedFromCountryCode

`func (o *MerchantShipmentResponse) SetShippedFromCountryCode(v string)`

SetShippedFromCountryCode sets ShippedFromCountryCode field to given value.

### HasShippedFromCountryCode

`func (o *MerchantShipmentResponse) HasShippedFromCountryCode() bool`

HasShippedFromCountryCode returns a boolean if a field has been set.

### SetShippedFromCountryCodeNil

`func (o *MerchantShipmentResponse) SetShippedFromCountryCodeNil(b bool)`

 SetShippedFromCountryCodeNil sets the value for ShippedFromCountryCode to be an explicit nil

### UnsetShippedFromCountryCode
`func (o *MerchantShipmentResponse) UnsetShippedFromCountryCode()`

UnsetShippedFromCountryCode ensures that no value is present for ShippedFromCountryCode, not even an explicit nil
### GetShippedFromStockLocationId

`func (o *MerchantShipmentResponse) GetShippedFromStockLocationId() int32`

GetShippedFromStockLocationId returns the ShippedFromStockLocationId field if non-nil, zero value otherwise.

### GetShippedFromStockLocationIdOk

`func (o *MerchantShipmentResponse) GetShippedFromStockLocationIdOk() (*int32, bool)`

GetShippedFromStockLocationIdOk returns a tuple with the ShippedFromStockLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedFromStockLocationId

`func (o *MerchantShipmentResponse) SetShippedFromStockLocationId(v int32)`

SetShippedFromStockLocationId sets ShippedFromStockLocationId field to given value.

### HasShippedFromStockLocationId

`func (o *MerchantShipmentResponse) HasShippedFromStockLocationId() bool`

HasShippedFromStockLocationId returns a boolean if a field has been set.

### SetShippedFromStockLocationIdNil

`func (o *MerchantShipmentResponse) SetShippedFromStockLocationIdNil(b bool)`

 SetShippedFromStockLocationIdNil sets the value for ShippedFromStockLocationId to be an explicit nil

### UnsetShippedFromStockLocationId
`func (o *MerchantShipmentResponse) UnsetShippedFromStockLocationId()`

UnsetShippedFromStockLocationId ensures that no value is present for ShippedFromStockLocationId, not even an explicit nil
### GetShipmentDate

`func (o *MerchantShipmentResponse) GetShipmentDate() time.Time`

GetShipmentDate returns the ShipmentDate field if non-nil, zero value otherwise.

### GetShipmentDateOk

`func (o *MerchantShipmentResponse) GetShipmentDateOk() (*time.Time, bool)`

GetShipmentDateOk returns a tuple with the ShipmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDate

`func (o *MerchantShipmentResponse) SetShipmentDate(v time.Time)`

SetShipmentDate sets ShipmentDate field to given value.

### HasShipmentDate

`func (o *MerchantShipmentResponse) HasShipmentDate() bool`

HasShipmentDate returns a boolean if a field has been set.

### SetShipmentDateNil

`func (o *MerchantShipmentResponse) SetShipmentDateNil(b bool)`

 SetShipmentDateNil sets the value for ShipmentDate to be an explicit nil

### UnsetShipmentDate
`func (o *MerchantShipmentResponse) UnsetShipmentDate()`

UnsetShipmentDate ensures that no value is present for ShipmentDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


