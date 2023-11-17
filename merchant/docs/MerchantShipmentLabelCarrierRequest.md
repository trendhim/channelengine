# MerchantShipmentLabelCarrierRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lines** | [**[]MerchantShipmentLineRequest**](MerchantShipmentLineRequest.md) |  | 
**Dimensions** | [**MerchantShipmentPackageDimensionsRequest**](MerchantShipmentPackageDimensionsRequest.md) |  | 
**Weight** | [**MerchantShipmentPackageWeightRequest**](MerchantShipmentPackageWeightRequest.md) |  | 

## Methods

### NewMerchantShipmentLabelCarrierRequest

`func NewMerchantShipmentLabelCarrierRequest(lines []MerchantShipmentLineRequest, dimensions MerchantShipmentPackageDimensionsRequest, weight MerchantShipmentPackageWeightRequest, ) *MerchantShipmentLabelCarrierRequest`

NewMerchantShipmentLabelCarrierRequest instantiates a new MerchantShipmentLabelCarrierRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantShipmentLabelCarrierRequestWithDefaults

`func NewMerchantShipmentLabelCarrierRequestWithDefaults() *MerchantShipmentLabelCarrierRequest`

NewMerchantShipmentLabelCarrierRequestWithDefaults instantiates a new MerchantShipmentLabelCarrierRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLines

`func (o *MerchantShipmentLabelCarrierRequest) GetLines() []MerchantShipmentLineRequest`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantShipmentLabelCarrierRequest) GetLinesOk() (*[]MerchantShipmentLineRequest, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantShipmentLabelCarrierRequest) SetLines(v []MerchantShipmentLineRequest)`

SetLines sets Lines field to given value.


### GetDimensions

`func (o *MerchantShipmentLabelCarrierRequest) GetDimensions() MerchantShipmentPackageDimensionsRequest`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *MerchantShipmentLabelCarrierRequest) GetDimensionsOk() (*MerchantShipmentPackageDimensionsRequest, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *MerchantShipmentLabelCarrierRequest) SetDimensions(v MerchantShipmentPackageDimensionsRequest)`

SetDimensions sets Dimensions field to given value.


### GetWeight

`func (o *MerchantShipmentLabelCarrierRequest) GetWeight() MerchantShipmentPackageWeightRequest`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *MerchantShipmentLabelCarrierRequest) GetWeightOk() (*MerchantShipmentPackageWeightRequest, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *MerchantShipmentLabelCarrierRequest) SetWeight(v MerchantShipmentPackageWeightRequest)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


