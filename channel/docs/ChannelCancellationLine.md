# ChannelCancellationLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderLineIdentifier** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 

## Methods

### NewChannelCancellationLine

`func NewChannelCancellationLine() *ChannelCancellationLine`

NewChannelCancellationLine instantiates a new ChannelCancellationLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelCancellationLineWithDefaults

`func NewChannelCancellationLineWithDefaults() *ChannelCancellationLine`

NewChannelCancellationLineWithDefaults instantiates a new ChannelCancellationLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderLineIdentifier

`func (o *ChannelCancellationLine) GetOrderLineIdentifier() string`

GetOrderLineIdentifier returns the OrderLineIdentifier field if non-nil, zero value otherwise.

### GetOrderLineIdentifierOk

`func (o *ChannelCancellationLine) GetOrderLineIdentifierOk() (*string, bool)`

GetOrderLineIdentifierOk returns a tuple with the OrderLineIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLineIdentifier

`func (o *ChannelCancellationLine) SetOrderLineIdentifier(v string)`

SetOrderLineIdentifier sets OrderLineIdentifier field to given value.

### HasOrderLineIdentifier

`func (o *ChannelCancellationLine) HasOrderLineIdentifier() bool`

HasOrderLineIdentifier returns a boolean if a field has been set.

### SetOrderLineIdentifierNil

`func (o *ChannelCancellationLine) SetOrderLineIdentifierNil(b bool)`

 SetOrderLineIdentifierNil sets the value for OrderLineIdentifier to be an explicit nil

### UnsetOrderLineIdentifier
`func (o *ChannelCancellationLine) UnsetOrderLineIdentifier()`

UnsetOrderLineIdentifier ensures that no value is present for OrderLineIdentifier, not even an explicit nil
### GetQuantity

`func (o *ChannelCancellationLine) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ChannelCancellationLine) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ChannelCancellationLine) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ChannelCancellationLine) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


