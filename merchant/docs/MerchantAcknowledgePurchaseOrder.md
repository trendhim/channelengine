# MerchantAcknowledgePurchaseOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **NullableString** |  | [optional] 
**MerchantOrderNo** | Pointer to **NullableString** |  | [optional] 
**Lines** | Pointer to [**[]MerchantAcknowledgePurchaseOrderLine**](MerchantAcknowledgePurchaseOrderLine.md) |  | [optional] 

## Methods

### NewMerchantAcknowledgePurchaseOrder

`func NewMerchantAcknowledgePurchaseOrder() *MerchantAcknowledgePurchaseOrder`

NewMerchantAcknowledgePurchaseOrder instantiates a new MerchantAcknowledgePurchaseOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantAcknowledgePurchaseOrderWithDefaults

`func NewMerchantAcknowledgePurchaseOrderWithDefaults() *MerchantAcknowledgePurchaseOrder`

NewMerchantAcknowledgePurchaseOrderWithDefaults instantiates a new MerchantAcknowledgePurchaseOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *MerchantAcknowledgePurchaseOrder) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MerchantAcknowledgePurchaseOrder) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MerchantAcknowledgePurchaseOrder) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *MerchantAcknowledgePurchaseOrder) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *MerchantAcknowledgePurchaseOrder) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *MerchantAcknowledgePurchaseOrder) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetMerchantOrderNo

`func (o *MerchantAcknowledgePurchaseOrder) GetMerchantOrderNo() string`

GetMerchantOrderNo returns the MerchantOrderNo field if non-nil, zero value otherwise.

### GetMerchantOrderNoOk

`func (o *MerchantAcknowledgePurchaseOrder) GetMerchantOrderNoOk() (*string, bool)`

GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderNo

`func (o *MerchantAcknowledgePurchaseOrder) SetMerchantOrderNo(v string)`

SetMerchantOrderNo sets MerchantOrderNo field to given value.

### HasMerchantOrderNo

`func (o *MerchantAcknowledgePurchaseOrder) HasMerchantOrderNo() bool`

HasMerchantOrderNo returns a boolean if a field has been set.

### SetMerchantOrderNoNil

`func (o *MerchantAcknowledgePurchaseOrder) SetMerchantOrderNoNil(b bool)`

 SetMerchantOrderNoNil sets the value for MerchantOrderNo to be an explicit nil

### UnsetMerchantOrderNo
`func (o *MerchantAcknowledgePurchaseOrder) UnsetMerchantOrderNo()`

UnsetMerchantOrderNo ensures that no value is present for MerchantOrderNo, not even an explicit nil
### GetLines

`func (o *MerchantAcknowledgePurchaseOrder) GetLines() []MerchantAcknowledgePurchaseOrderLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantAcknowledgePurchaseOrder) GetLinesOk() (*[]MerchantAcknowledgePurchaseOrderLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantAcknowledgePurchaseOrder) SetLines(v []MerchantAcknowledgePurchaseOrderLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *MerchantAcknowledgePurchaseOrder) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *MerchantAcknowledgePurchaseOrder) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *MerchantAcknowledgePurchaseOrder) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


