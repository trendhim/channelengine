# MerchantShipmentLabelCarrierResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The channel&#39;s name for the shipping label carrier | [optional] 
**Code** | Pointer to **NullableString** | The channel&#39;s code for the shipping label carrier | [optional] 
**Restrictions** | Pointer to **NullableString** | Optional. Any restrictions on this carriers, e.g. weight and/or dimensions | [optional] 
**Price** | Pointer to **NullableFloat32** | Optional. Price for this shipping label | [optional] 
**Recommendation** | Pointer to [**ChannelCarrierRecommendationApi**](ChannelCarrierRecommendationApi.md) |  | [optional] 
**CollectionMethod** | Pointer to [**ChannelCarrierCollectionMethodApi**](ChannelCarrierCollectionMethodApi.md) |  | [optional] 
**HandoverDateTime** | Pointer to **NullableTime** | Optional. When to handover the package to the carrier  It is in the ISO format including the timezone offset.  E.g. 2020-10-03T18:00:00+02:00  which is 3rd Oct 2020, at 18:00 PM in Amsterdam (Summer Time aka CEST: UTC +2:00 ) | [optional] 

## Methods

### NewMerchantShipmentLabelCarrierResponse

`func NewMerchantShipmentLabelCarrierResponse() *MerchantShipmentLabelCarrierResponse`

NewMerchantShipmentLabelCarrierResponse instantiates a new MerchantShipmentLabelCarrierResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantShipmentLabelCarrierResponseWithDefaults

`func NewMerchantShipmentLabelCarrierResponseWithDefaults() *MerchantShipmentLabelCarrierResponse`

NewMerchantShipmentLabelCarrierResponseWithDefaults instantiates a new MerchantShipmentLabelCarrierResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MerchantShipmentLabelCarrierResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantShipmentLabelCarrierResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantShipmentLabelCarrierResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MerchantShipmentLabelCarrierResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MerchantShipmentLabelCarrierResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MerchantShipmentLabelCarrierResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *MerchantShipmentLabelCarrierResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MerchantShipmentLabelCarrierResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MerchantShipmentLabelCarrierResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MerchantShipmentLabelCarrierResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *MerchantShipmentLabelCarrierResponse) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *MerchantShipmentLabelCarrierResponse) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetRestrictions

`func (o *MerchantShipmentLabelCarrierResponse) GetRestrictions() string`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *MerchantShipmentLabelCarrierResponse) GetRestrictionsOk() (*string, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *MerchantShipmentLabelCarrierResponse) SetRestrictions(v string)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *MerchantShipmentLabelCarrierResponse) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### SetRestrictionsNil

`func (o *MerchantShipmentLabelCarrierResponse) SetRestrictionsNil(b bool)`

 SetRestrictionsNil sets the value for Restrictions to be an explicit nil

### UnsetRestrictions
`func (o *MerchantShipmentLabelCarrierResponse) UnsetRestrictions()`

UnsetRestrictions ensures that no value is present for Restrictions, not even an explicit nil
### GetPrice

`func (o *MerchantShipmentLabelCarrierResponse) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MerchantShipmentLabelCarrierResponse) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MerchantShipmentLabelCarrierResponse) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MerchantShipmentLabelCarrierResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *MerchantShipmentLabelCarrierResponse) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *MerchantShipmentLabelCarrierResponse) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetRecommendation

`func (o *MerchantShipmentLabelCarrierResponse) GetRecommendation() ChannelCarrierRecommendationApi`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *MerchantShipmentLabelCarrierResponse) GetRecommendationOk() (*ChannelCarrierRecommendationApi, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *MerchantShipmentLabelCarrierResponse) SetRecommendation(v ChannelCarrierRecommendationApi)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *MerchantShipmentLabelCarrierResponse) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### GetCollectionMethod

`func (o *MerchantShipmentLabelCarrierResponse) GetCollectionMethod() ChannelCarrierCollectionMethodApi`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *MerchantShipmentLabelCarrierResponse) GetCollectionMethodOk() (*ChannelCarrierCollectionMethodApi, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *MerchantShipmentLabelCarrierResponse) SetCollectionMethod(v ChannelCarrierCollectionMethodApi)`

SetCollectionMethod sets CollectionMethod field to given value.

### HasCollectionMethod

`func (o *MerchantShipmentLabelCarrierResponse) HasCollectionMethod() bool`

HasCollectionMethod returns a boolean if a field has been set.

### GetHandoverDateTime

`func (o *MerchantShipmentLabelCarrierResponse) GetHandoverDateTime() time.Time`

GetHandoverDateTime returns the HandoverDateTime field if non-nil, zero value otherwise.

### GetHandoverDateTimeOk

`func (o *MerchantShipmentLabelCarrierResponse) GetHandoverDateTimeOk() (*time.Time, bool)`

GetHandoverDateTimeOk returns a tuple with the HandoverDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoverDateTime

`func (o *MerchantShipmentLabelCarrierResponse) SetHandoverDateTime(v time.Time)`

SetHandoverDateTime sets HandoverDateTime field to given value.

### HasHandoverDateTime

`func (o *MerchantShipmentLabelCarrierResponse) HasHandoverDateTime() bool`

HasHandoverDateTime returns a boolean if a field has been set.

### SetHandoverDateTimeNil

`func (o *MerchantShipmentLabelCarrierResponse) SetHandoverDateTimeNil(b bool)`

 SetHandoverDateTimeNil sets the value for HandoverDateTime to be an explicit nil

### UnsetHandoverDateTime
`func (o *MerchantShipmentLabelCarrierResponse) UnsetHandoverDateTime()`

UnsetHandoverDateTime ensures that no value is present for HandoverDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


