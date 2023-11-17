# MerchantChannelLabelShipmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimensions** | [**MerchantShipmentPackageDimensionsRequest**](MerchantShipmentPackageDimensionsRequest.md) |  | 
**Weight** | [**MerchantShipmentPackageWeightRequest**](MerchantShipmentPackageWeightRequest.md) |  | 
**ChannelMethodCode** | **string** |  | 
**MerchantShipmentNo** | **string** | The unique shipment reference used by the Merchant. | 
**MerchantOrderNo** | **string** | The unique order reference used by the Merchant. | 
**ShippedFromCountryCode** | Pointer to **NullableString** | The code of the country from where the package is being shipped. | [optional] 
**ShippedFromStockLocationId** | Pointer to **NullableInt32** |  | [optional] 
**Lines** | [**[]MerchantShipmentLineRequest**](MerchantShipmentLineRequest.md) |  | 

## Methods

### NewMerchantChannelLabelShipmentRequest

`func NewMerchantChannelLabelShipmentRequest(dimensions MerchantShipmentPackageDimensionsRequest, weight MerchantShipmentPackageWeightRequest, channelMethodCode string, merchantShipmentNo string, merchantOrderNo string, lines []MerchantShipmentLineRequest, ) *MerchantChannelLabelShipmentRequest`

NewMerchantChannelLabelShipmentRequest instantiates a new MerchantChannelLabelShipmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantChannelLabelShipmentRequestWithDefaults

`func NewMerchantChannelLabelShipmentRequestWithDefaults() *MerchantChannelLabelShipmentRequest`

NewMerchantChannelLabelShipmentRequestWithDefaults instantiates a new MerchantChannelLabelShipmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensions

`func (o *MerchantChannelLabelShipmentRequest) GetDimensions() MerchantShipmentPackageDimensionsRequest`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *MerchantChannelLabelShipmentRequest) GetDimensionsOk() (*MerchantShipmentPackageDimensionsRequest, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *MerchantChannelLabelShipmentRequest) SetDimensions(v MerchantShipmentPackageDimensionsRequest)`

SetDimensions sets Dimensions field to given value.


### GetWeight

`func (o *MerchantChannelLabelShipmentRequest) GetWeight() MerchantShipmentPackageWeightRequest`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *MerchantChannelLabelShipmentRequest) GetWeightOk() (*MerchantShipmentPackageWeightRequest, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *MerchantChannelLabelShipmentRequest) SetWeight(v MerchantShipmentPackageWeightRequest)`

SetWeight sets Weight field to given value.


### GetChannelMethodCode

`func (o *MerchantChannelLabelShipmentRequest) GetChannelMethodCode() string`

GetChannelMethodCode returns the ChannelMethodCode field if non-nil, zero value otherwise.

### GetChannelMethodCodeOk

`func (o *MerchantChannelLabelShipmentRequest) GetChannelMethodCodeOk() (*string, bool)`

GetChannelMethodCodeOk returns a tuple with the ChannelMethodCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelMethodCode

`func (o *MerchantChannelLabelShipmentRequest) SetChannelMethodCode(v string)`

SetChannelMethodCode sets ChannelMethodCode field to given value.


### GetMerchantShipmentNo

`func (o *MerchantChannelLabelShipmentRequest) GetMerchantShipmentNo() string`

GetMerchantShipmentNo returns the MerchantShipmentNo field if non-nil, zero value otherwise.

### GetMerchantShipmentNoOk

`func (o *MerchantChannelLabelShipmentRequest) GetMerchantShipmentNoOk() (*string, bool)`

GetMerchantShipmentNoOk returns a tuple with the MerchantShipmentNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantShipmentNo

`func (o *MerchantChannelLabelShipmentRequest) SetMerchantShipmentNo(v string)`

SetMerchantShipmentNo sets MerchantShipmentNo field to given value.


### GetMerchantOrderNo

`func (o *MerchantChannelLabelShipmentRequest) GetMerchantOrderNo() string`

GetMerchantOrderNo returns the MerchantOrderNo field if non-nil, zero value otherwise.

### GetMerchantOrderNoOk

`func (o *MerchantChannelLabelShipmentRequest) GetMerchantOrderNoOk() (*string, bool)`

GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderNo

`func (o *MerchantChannelLabelShipmentRequest) SetMerchantOrderNo(v string)`

SetMerchantOrderNo sets MerchantOrderNo field to given value.


### GetShippedFromCountryCode

`func (o *MerchantChannelLabelShipmentRequest) GetShippedFromCountryCode() string`

GetShippedFromCountryCode returns the ShippedFromCountryCode field if non-nil, zero value otherwise.

### GetShippedFromCountryCodeOk

`func (o *MerchantChannelLabelShipmentRequest) GetShippedFromCountryCodeOk() (*string, bool)`

GetShippedFromCountryCodeOk returns a tuple with the ShippedFromCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedFromCountryCode

`func (o *MerchantChannelLabelShipmentRequest) SetShippedFromCountryCode(v string)`

SetShippedFromCountryCode sets ShippedFromCountryCode field to given value.

### HasShippedFromCountryCode

`func (o *MerchantChannelLabelShipmentRequest) HasShippedFromCountryCode() bool`

HasShippedFromCountryCode returns a boolean if a field has been set.

### SetShippedFromCountryCodeNil

`func (o *MerchantChannelLabelShipmentRequest) SetShippedFromCountryCodeNil(b bool)`

 SetShippedFromCountryCodeNil sets the value for ShippedFromCountryCode to be an explicit nil

### UnsetShippedFromCountryCode
`func (o *MerchantChannelLabelShipmentRequest) UnsetShippedFromCountryCode()`

UnsetShippedFromCountryCode ensures that no value is present for ShippedFromCountryCode, not even an explicit nil
### GetShippedFromStockLocationId

`func (o *MerchantChannelLabelShipmentRequest) GetShippedFromStockLocationId() int32`

GetShippedFromStockLocationId returns the ShippedFromStockLocationId field if non-nil, zero value otherwise.

### GetShippedFromStockLocationIdOk

`func (o *MerchantChannelLabelShipmentRequest) GetShippedFromStockLocationIdOk() (*int32, bool)`

GetShippedFromStockLocationIdOk returns a tuple with the ShippedFromStockLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedFromStockLocationId

`func (o *MerchantChannelLabelShipmentRequest) SetShippedFromStockLocationId(v int32)`

SetShippedFromStockLocationId sets ShippedFromStockLocationId field to given value.

### HasShippedFromStockLocationId

`func (o *MerchantChannelLabelShipmentRequest) HasShippedFromStockLocationId() bool`

HasShippedFromStockLocationId returns a boolean if a field has been set.

### SetShippedFromStockLocationIdNil

`func (o *MerchantChannelLabelShipmentRequest) SetShippedFromStockLocationIdNil(b bool)`

 SetShippedFromStockLocationIdNil sets the value for ShippedFromStockLocationId to be an explicit nil

### UnsetShippedFromStockLocationId
`func (o *MerchantChannelLabelShipmentRequest) UnsetShippedFromStockLocationId()`

UnsetShippedFromStockLocationId ensures that no value is present for ShippedFromStockLocationId, not even an explicit nil
### GetLines

`func (o *MerchantChannelLabelShipmentRequest) GetLines() []MerchantShipmentLineRequest`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantChannelLabelShipmentRequest) GetLinesOk() (*[]MerchantShipmentLineRequest, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantChannelLabelShipmentRequest) SetLines(v []MerchantShipmentLineRequest)`

SetLines sets Lines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


