# IPurchaseOrderByFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**ChannelId** | Pointer to **int32** |  | [optional] 
**ChannelName** | Pointer to **NullableString** |  | [optional] 
**GlobalChannelName** | Pointer to **NullableString** |  | [optional] 
**GlobalChannelId** | Pointer to **NullableInt32** |  | [optional] 
**ChannelPurchaseOrderNo** | Pointer to **NullableString** |  | [optional] 
**MerchantPurchaseOrderNo** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**ModulesPurchaseOrderStatus**](ModulesPurchaseOrderStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**Type** | Pointer to [**ModulesPurchaseOrderType**](ModulesPurchaseOrderType.md) |  | [optional] 
**FromShipDate** | Pointer to **NullableTime** |  | [optional] 
**ToShipDate** | Pointer to **NullableTime** |  | [optional] 
**FromDeliveryDate** | Pointer to **NullableTime** |  | [optional] 
**ToDeliveryDate** | Pointer to **NullableTime** |  | [optional] 
**ImportInformation** | Pointer to [**IImportInformation**](IImportInformation.md) |  | [optional] 
**SellingParty** | Pointer to [**IVendorParty**](IVendorParty.md) |  | [optional] 
**ShipToParty** | Pointer to [**IVendorParty**](IVendorParty.md) |  | [optional] 
**BuyingParty** | Pointer to [**IVendorParty**](IVendorParty.md) |  | [optional] 
**BillingParty** | Pointer to [**IVendorParty**](IVendorParty.md) |  | [optional] 
**Lines** | Pointer to [**[]IPurchaseOrderLineByFilter**](IPurchaseOrderLineByFilter.md) |  | [optional] 

## Methods

### NewIPurchaseOrderByFilter

`func NewIPurchaseOrderByFilter() *IPurchaseOrderByFilter`

NewIPurchaseOrderByFilter instantiates a new IPurchaseOrderByFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPurchaseOrderByFilterWithDefaults

`func NewIPurchaseOrderByFilterWithDefaults() *IPurchaseOrderByFilter`

NewIPurchaseOrderByFilterWithDefaults instantiates a new IPurchaseOrderByFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPurchaseOrderByFilter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPurchaseOrderByFilter) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPurchaseOrderByFilter) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IPurchaseOrderByFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IPurchaseOrderByFilter) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IPurchaseOrderByFilter) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetChannelId

`func (o *IPurchaseOrderByFilter) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *IPurchaseOrderByFilter) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *IPurchaseOrderByFilter) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *IPurchaseOrderByFilter) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetChannelName

`func (o *IPurchaseOrderByFilter) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *IPurchaseOrderByFilter) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *IPurchaseOrderByFilter) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *IPurchaseOrderByFilter) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### SetChannelNameNil

`func (o *IPurchaseOrderByFilter) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *IPurchaseOrderByFilter) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetGlobalChannelName

`func (o *IPurchaseOrderByFilter) GetGlobalChannelName() string`

GetGlobalChannelName returns the GlobalChannelName field if non-nil, zero value otherwise.

### GetGlobalChannelNameOk

`func (o *IPurchaseOrderByFilter) GetGlobalChannelNameOk() (*string, bool)`

GetGlobalChannelNameOk returns a tuple with the GlobalChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalChannelName

`func (o *IPurchaseOrderByFilter) SetGlobalChannelName(v string)`

SetGlobalChannelName sets GlobalChannelName field to given value.

### HasGlobalChannelName

`func (o *IPurchaseOrderByFilter) HasGlobalChannelName() bool`

HasGlobalChannelName returns a boolean if a field has been set.

### SetGlobalChannelNameNil

`func (o *IPurchaseOrderByFilter) SetGlobalChannelNameNil(b bool)`

 SetGlobalChannelNameNil sets the value for GlobalChannelName to be an explicit nil

### UnsetGlobalChannelName
`func (o *IPurchaseOrderByFilter) UnsetGlobalChannelName()`

UnsetGlobalChannelName ensures that no value is present for GlobalChannelName, not even an explicit nil
### GetGlobalChannelId

`func (o *IPurchaseOrderByFilter) GetGlobalChannelId() int32`

GetGlobalChannelId returns the GlobalChannelId field if non-nil, zero value otherwise.

### GetGlobalChannelIdOk

`func (o *IPurchaseOrderByFilter) GetGlobalChannelIdOk() (*int32, bool)`

GetGlobalChannelIdOk returns a tuple with the GlobalChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalChannelId

`func (o *IPurchaseOrderByFilter) SetGlobalChannelId(v int32)`

SetGlobalChannelId sets GlobalChannelId field to given value.

### HasGlobalChannelId

`func (o *IPurchaseOrderByFilter) HasGlobalChannelId() bool`

HasGlobalChannelId returns a boolean if a field has been set.

### SetGlobalChannelIdNil

`func (o *IPurchaseOrderByFilter) SetGlobalChannelIdNil(b bool)`

 SetGlobalChannelIdNil sets the value for GlobalChannelId to be an explicit nil

### UnsetGlobalChannelId
`func (o *IPurchaseOrderByFilter) UnsetGlobalChannelId()`

UnsetGlobalChannelId ensures that no value is present for GlobalChannelId, not even an explicit nil
### GetChannelPurchaseOrderNo

`func (o *IPurchaseOrderByFilter) GetChannelPurchaseOrderNo() string`

GetChannelPurchaseOrderNo returns the ChannelPurchaseOrderNo field if non-nil, zero value otherwise.

### GetChannelPurchaseOrderNoOk

`func (o *IPurchaseOrderByFilter) GetChannelPurchaseOrderNoOk() (*string, bool)`

GetChannelPurchaseOrderNoOk returns a tuple with the ChannelPurchaseOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPurchaseOrderNo

`func (o *IPurchaseOrderByFilter) SetChannelPurchaseOrderNo(v string)`

SetChannelPurchaseOrderNo sets ChannelPurchaseOrderNo field to given value.

### HasChannelPurchaseOrderNo

`func (o *IPurchaseOrderByFilter) HasChannelPurchaseOrderNo() bool`

HasChannelPurchaseOrderNo returns a boolean if a field has been set.

### SetChannelPurchaseOrderNoNil

`func (o *IPurchaseOrderByFilter) SetChannelPurchaseOrderNoNil(b bool)`

 SetChannelPurchaseOrderNoNil sets the value for ChannelPurchaseOrderNo to be an explicit nil

### UnsetChannelPurchaseOrderNo
`func (o *IPurchaseOrderByFilter) UnsetChannelPurchaseOrderNo()`

UnsetChannelPurchaseOrderNo ensures that no value is present for ChannelPurchaseOrderNo, not even an explicit nil
### GetMerchantPurchaseOrderNo

`func (o *IPurchaseOrderByFilter) GetMerchantPurchaseOrderNo() string`

GetMerchantPurchaseOrderNo returns the MerchantPurchaseOrderNo field if non-nil, zero value otherwise.

### GetMerchantPurchaseOrderNoOk

`func (o *IPurchaseOrderByFilter) GetMerchantPurchaseOrderNoOk() (*string, bool)`

GetMerchantPurchaseOrderNoOk returns a tuple with the MerchantPurchaseOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantPurchaseOrderNo

`func (o *IPurchaseOrderByFilter) SetMerchantPurchaseOrderNo(v string)`

SetMerchantPurchaseOrderNo sets MerchantPurchaseOrderNo field to given value.

### HasMerchantPurchaseOrderNo

`func (o *IPurchaseOrderByFilter) HasMerchantPurchaseOrderNo() bool`

HasMerchantPurchaseOrderNo returns a boolean if a field has been set.

### SetMerchantPurchaseOrderNoNil

`func (o *IPurchaseOrderByFilter) SetMerchantPurchaseOrderNoNil(b bool)`

 SetMerchantPurchaseOrderNoNil sets the value for MerchantPurchaseOrderNo to be an explicit nil

### UnsetMerchantPurchaseOrderNo
`func (o *IPurchaseOrderByFilter) UnsetMerchantPurchaseOrderNo()`

UnsetMerchantPurchaseOrderNo ensures that no value is present for MerchantPurchaseOrderNo, not even an explicit nil
### GetStatus

`func (o *IPurchaseOrderByFilter) GetStatus() ModulesPurchaseOrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IPurchaseOrderByFilter) GetStatusOk() (*ModulesPurchaseOrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IPurchaseOrderByFilter) SetStatus(v ModulesPurchaseOrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IPurchaseOrderByFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IPurchaseOrderByFilter) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IPurchaseOrderByFilter) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IPurchaseOrderByFilter) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IPurchaseOrderByFilter) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *IPurchaseOrderByFilter) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *IPurchaseOrderByFilter) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *IPurchaseOrderByFilter) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IPurchaseOrderByFilter) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IPurchaseOrderByFilter) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IPurchaseOrderByFilter) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *IPurchaseOrderByFilter) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *IPurchaseOrderByFilter) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetType

`func (o *IPurchaseOrderByFilter) GetType() ModulesPurchaseOrderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPurchaseOrderByFilter) GetTypeOk() (*ModulesPurchaseOrderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPurchaseOrderByFilter) SetType(v ModulesPurchaseOrderType)`

SetType sets Type field to given value.

### HasType

`func (o *IPurchaseOrderByFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFromShipDate

`func (o *IPurchaseOrderByFilter) GetFromShipDate() time.Time`

GetFromShipDate returns the FromShipDate field if non-nil, zero value otherwise.

### GetFromShipDateOk

`func (o *IPurchaseOrderByFilter) GetFromShipDateOk() (*time.Time, bool)`

GetFromShipDateOk returns a tuple with the FromShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromShipDate

`func (o *IPurchaseOrderByFilter) SetFromShipDate(v time.Time)`

SetFromShipDate sets FromShipDate field to given value.

### HasFromShipDate

`func (o *IPurchaseOrderByFilter) HasFromShipDate() bool`

HasFromShipDate returns a boolean if a field has been set.

### SetFromShipDateNil

`func (o *IPurchaseOrderByFilter) SetFromShipDateNil(b bool)`

 SetFromShipDateNil sets the value for FromShipDate to be an explicit nil

### UnsetFromShipDate
`func (o *IPurchaseOrderByFilter) UnsetFromShipDate()`

UnsetFromShipDate ensures that no value is present for FromShipDate, not even an explicit nil
### GetToShipDate

`func (o *IPurchaseOrderByFilter) GetToShipDate() time.Time`

GetToShipDate returns the ToShipDate field if non-nil, zero value otherwise.

### GetToShipDateOk

`func (o *IPurchaseOrderByFilter) GetToShipDateOk() (*time.Time, bool)`

GetToShipDateOk returns a tuple with the ToShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToShipDate

`func (o *IPurchaseOrderByFilter) SetToShipDate(v time.Time)`

SetToShipDate sets ToShipDate field to given value.

### HasToShipDate

`func (o *IPurchaseOrderByFilter) HasToShipDate() bool`

HasToShipDate returns a boolean if a field has been set.

### SetToShipDateNil

`func (o *IPurchaseOrderByFilter) SetToShipDateNil(b bool)`

 SetToShipDateNil sets the value for ToShipDate to be an explicit nil

### UnsetToShipDate
`func (o *IPurchaseOrderByFilter) UnsetToShipDate()`

UnsetToShipDate ensures that no value is present for ToShipDate, not even an explicit nil
### GetFromDeliveryDate

`func (o *IPurchaseOrderByFilter) GetFromDeliveryDate() time.Time`

GetFromDeliveryDate returns the FromDeliveryDate field if non-nil, zero value otherwise.

### GetFromDeliveryDateOk

`func (o *IPurchaseOrderByFilter) GetFromDeliveryDateOk() (*time.Time, bool)`

GetFromDeliveryDateOk returns a tuple with the FromDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDeliveryDate

`func (o *IPurchaseOrderByFilter) SetFromDeliveryDate(v time.Time)`

SetFromDeliveryDate sets FromDeliveryDate field to given value.

### HasFromDeliveryDate

`func (o *IPurchaseOrderByFilter) HasFromDeliveryDate() bool`

HasFromDeliveryDate returns a boolean if a field has been set.

### SetFromDeliveryDateNil

`func (o *IPurchaseOrderByFilter) SetFromDeliveryDateNil(b bool)`

 SetFromDeliveryDateNil sets the value for FromDeliveryDate to be an explicit nil

### UnsetFromDeliveryDate
`func (o *IPurchaseOrderByFilter) UnsetFromDeliveryDate()`

UnsetFromDeliveryDate ensures that no value is present for FromDeliveryDate, not even an explicit nil
### GetToDeliveryDate

`func (o *IPurchaseOrderByFilter) GetToDeliveryDate() time.Time`

GetToDeliveryDate returns the ToDeliveryDate field if non-nil, zero value otherwise.

### GetToDeliveryDateOk

`func (o *IPurchaseOrderByFilter) GetToDeliveryDateOk() (*time.Time, bool)`

GetToDeliveryDateOk returns a tuple with the ToDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDeliveryDate

`func (o *IPurchaseOrderByFilter) SetToDeliveryDate(v time.Time)`

SetToDeliveryDate sets ToDeliveryDate field to given value.

### HasToDeliveryDate

`func (o *IPurchaseOrderByFilter) HasToDeliveryDate() bool`

HasToDeliveryDate returns a boolean if a field has been set.

### SetToDeliveryDateNil

`func (o *IPurchaseOrderByFilter) SetToDeliveryDateNil(b bool)`

 SetToDeliveryDateNil sets the value for ToDeliveryDate to be an explicit nil

### UnsetToDeliveryDate
`func (o *IPurchaseOrderByFilter) UnsetToDeliveryDate()`

UnsetToDeliveryDate ensures that no value is present for ToDeliveryDate, not even an explicit nil
### GetImportInformation

`func (o *IPurchaseOrderByFilter) GetImportInformation() IImportInformation`

GetImportInformation returns the ImportInformation field if non-nil, zero value otherwise.

### GetImportInformationOk

`func (o *IPurchaseOrderByFilter) GetImportInformationOk() (*IImportInformation, bool)`

GetImportInformationOk returns a tuple with the ImportInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportInformation

`func (o *IPurchaseOrderByFilter) SetImportInformation(v IImportInformation)`

SetImportInformation sets ImportInformation field to given value.

### HasImportInformation

`func (o *IPurchaseOrderByFilter) HasImportInformation() bool`

HasImportInformation returns a boolean if a field has been set.

### GetSellingParty

`func (o *IPurchaseOrderByFilter) GetSellingParty() IVendorParty`

GetSellingParty returns the SellingParty field if non-nil, zero value otherwise.

### GetSellingPartyOk

`func (o *IPurchaseOrderByFilter) GetSellingPartyOk() (*IVendorParty, bool)`

GetSellingPartyOk returns a tuple with the SellingParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingParty

`func (o *IPurchaseOrderByFilter) SetSellingParty(v IVendorParty)`

SetSellingParty sets SellingParty field to given value.

### HasSellingParty

`func (o *IPurchaseOrderByFilter) HasSellingParty() bool`

HasSellingParty returns a boolean if a field has been set.

### GetShipToParty

`func (o *IPurchaseOrderByFilter) GetShipToParty() IVendorParty`

GetShipToParty returns the ShipToParty field if non-nil, zero value otherwise.

### GetShipToPartyOk

`func (o *IPurchaseOrderByFilter) GetShipToPartyOk() (*IVendorParty, bool)`

GetShipToPartyOk returns a tuple with the ShipToParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToParty

`func (o *IPurchaseOrderByFilter) SetShipToParty(v IVendorParty)`

SetShipToParty sets ShipToParty field to given value.

### HasShipToParty

`func (o *IPurchaseOrderByFilter) HasShipToParty() bool`

HasShipToParty returns a boolean if a field has been set.

### GetBuyingParty

`func (o *IPurchaseOrderByFilter) GetBuyingParty() IVendorParty`

GetBuyingParty returns the BuyingParty field if non-nil, zero value otherwise.

### GetBuyingPartyOk

`func (o *IPurchaseOrderByFilter) GetBuyingPartyOk() (*IVendorParty, bool)`

GetBuyingPartyOk returns a tuple with the BuyingParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyingParty

`func (o *IPurchaseOrderByFilter) SetBuyingParty(v IVendorParty)`

SetBuyingParty sets BuyingParty field to given value.

### HasBuyingParty

`func (o *IPurchaseOrderByFilter) HasBuyingParty() bool`

HasBuyingParty returns a boolean if a field has been set.

### GetBillingParty

`func (o *IPurchaseOrderByFilter) GetBillingParty() IVendorParty`

GetBillingParty returns the BillingParty field if non-nil, zero value otherwise.

### GetBillingPartyOk

`func (o *IPurchaseOrderByFilter) GetBillingPartyOk() (*IVendorParty, bool)`

GetBillingPartyOk returns a tuple with the BillingParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingParty

`func (o *IPurchaseOrderByFilter) SetBillingParty(v IVendorParty)`

SetBillingParty sets BillingParty field to given value.

### HasBillingParty

`func (o *IPurchaseOrderByFilter) HasBillingParty() bool`

HasBillingParty returns a boolean if a field has been set.

### GetLines

`func (o *IPurchaseOrderByFilter) GetLines() []IPurchaseOrderLineByFilter`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *IPurchaseOrderByFilter) GetLinesOk() (*[]IPurchaseOrderLineByFilter, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *IPurchaseOrderByFilter) SetLines(v []IPurchaseOrderLineByFilter)`

SetLines sets Lines field to given value.

### HasLines

`func (o *IPurchaseOrderByFilter) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *IPurchaseOrderByFilter) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *IPurchaseOrderByFilter) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


