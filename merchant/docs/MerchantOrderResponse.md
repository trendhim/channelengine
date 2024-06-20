# MerchantOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique identifier used by ChannelEngine. This identifier does  not have to be saved. It should only be used in a call to acknowledge the order. | [optional] 
**ChannelName** | Pointer to **NullableString** | The name of the channel for this specific environment/account. | [optional] 
**ChannelId** | Pointer to **NullableInt32** | The unique ID of the channel for this specific environment/account. | [optional] 
**GlobalChannelName** | Pointer to **NullableString** | The name of the channel across all of ChannelEngine. | [optional] 
**GlobalChannelId** | Pointer to **NullableInt32** | The unique ID of the channel across all of ChannelEngine. | [optional] 
**ChannelOrderSupport** | Pointer to [**OrderSupport**](OrderSupport.md) |  | [optional] 
**ChannelOrderNo** | Pointer to **NullableString** | The order reference used by the channel.  This number is not guaranteed to be unique accross all orders,  because different channels can use the same order number format. | [optional] 
**CommercialOrderNo** | Pointer to **NullableString** | The order reference used by the channel for commercial purposes (e.g. on the invoice). Can be different from the ChannelOrderNo.  For example when the internal unique order reference is a unique id or guid, while the commercial order reference is (usually) a human readable reference that can be reused or used for multiple sellers by the channel. | [optional] 
**MerchantOrderNo** | Pointer to **NullableString** | The unique order reference used by the Merchant | [optional] 
**Status** | Pointer to [**OrderStatusView**](OrderStatusView.md) |  | [optional] 
**IsBusinessOrder** | Pointer to **bool** | Indicating whether the order is a business order. | [optional] 
**AcknowledgedDate** | Pointer to **NullableTime** | The date the order was acknowledged in ChannelEngine. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | The date the order was created in ChannelEngine. | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | The date the order was last updated in ChannelEngine. | [optional] 
**ClosedAt** | Pointer to **NullableTime** | The date the order was last updated in ChannelEngine. | [optional] 
**MerchantComment** | Pointer to **NullableString** | The optional comment a merchant can add to an order. | [optional] 
**BillingAddress** | Pointer to [**MerchantAddressResponse**](MerchantAddressResponse.md) |  | [optional] 
**ShippingAddress** | Pointer to [**MerchantAddressResponse**](MerchantAddressResponse.md) |  | [optional] 
**SubTotalInclVat** | Pointer to **NullableFloat32** | The total value of the order lines including VAT  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | [optional] 
**SubTotalVat** | Pointer to **NullableFloat32** | The total amount of VAT charged over the order lines  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | [optional] 
**ShippingCostsInclVat** | Pointer to **float32** |  | [optional] 
**ShippingCostsVat** | Pointer to **NullableFloat32** | The total amount of VAT charged over the shipping fee  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | [optional] 
**TotalInclVat** | Pointer to **float32** | The total value of the order including VAT  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | [optional] 
**TotalVat** | Pointer to **NullableFloat32** | The total amount of VAT charged over the total value of te order  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | [optional] 
**OriginalSubTotalInclVat** | Pointer to **NullableFloat32** | The total value of the order lines including VAT  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**OriginalSubTotalVat** | Pointer to **NullableFloat32** | The total amount of VAT charged over the order lines  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**OriginalShippingCostsInclVat** | Pointer to **NullableFloat32** | The shipping fee including VAT  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**OriginalShippingCostsVat** | Pointer to **NullableFloat32** | The total amount of VAT charged over the shipping fee  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**OriginalTotalInclVat** | Pointer to **NullableFloat32** | The total value of the order including VAT  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**OriginalTotalVat** | Pointer to **NullableFloat32** | The total amount of VAT charged over the total value of te order  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**SubTotalExclVat** | Pointer to **NullableFloat32** |  | [optional] 
**TotalExclVat** | Pointer to **NullableFloat32** |  | [optional] 
**ShippingCostsExclVat** | Pointer to **NullableFloat32** |  | [optional] 
**OriginalSubTotalExclVat** | Pointer to **NullableFloat32** |  | [optional] 
**OriginalShippingCostsExclVat** | Pointer to **NullableFloat32** |  | [optional] 
**OriginalTotalExclVat** | Pointer to **NullableFloat32** |  | [optional] 
**OriginalSubTotalFee** | Pointer to **float32** | The sum of the fees on the order lines  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**SubTotalFee** | Pointer to **float32** | The sum of the fees on the order lines  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | [optional] 
**OriginalOrderFee** | Pointer to **float32** | The fee on order itself (besides the fees on the order lines)  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**OrderFee** | Pointer to **float32** | The fee on order itself (besides the fees on the order lines)  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | [optional] 
**OriginalTotalFee** | Pointer to **float32** | The total fee: the fees on the order lines + the fee on the order itself  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**TotalFee** | Pointer to **float32** | The total fee: the fees on the order lines + the fee on the order itself  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | [optional] 
**Lines** | Pointer to [**[]MerchantOrderLineResponse**](MerchantOrderLineResponse.md) |  | [optional] 
**Phone** | Pointer to **NullableString** | The customer&#39;s telephone number. | [optional] 
**Email** | **string** | The customer&#39;s email. | 
**LanguageCode** | Pointer to **NullableString** | The language of the order. Has to be a 2-letter ISO language code. | [optional] 
**CompanyRegistrationNo** | Pointer to **NullableString** | Optional. A company&#39;s chamber of commerce number. | [optional] 
**VatNo** | Pointer to **NullableString** | Optional. A company&#39;s VAT number. | [optional] 
**PaymentMethod** | Pointer to **NullableString** | The payment method used on the order. | [optional] 
**PaymentReferenceNo** | Pointer to **NullableString** | Reference or transaction id for the payment | [optional] 
**CurrencyCode** | **string** | The currency code for the amounts of the order. | 
**OrderDate** | **time.Time** | The date the order was created at the channel. | 
**ChannelCustomerNo** | Pointer to **NullableString** | The unique customer reference used by the channel. | [optional] 
**ExtraData** | Pointer to **map[string]string** | Extra data on the order. | [optional] 

## Methods

### NewMerchantOrderResponse

`func NewMerchantOrderResponse(email string, currencyCode string, orderDate time.Time, ) *MerchantOrderResponse`

NewMerchantOrderResponse instantiates a new MerchantOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantOrderResponseWithDefaults

`func NewMerchantOrderResponseWithDefaults() *MerchantOrderResponse`

NewMerchantOrderResponseWithDefaults instantiates a new MerchantOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MerchantOrderResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantOrderResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantOrderResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MerchantOrderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChannelName

`func (o *MerchantOrderResponse) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *MerchantOrderResponse) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *MerchantOrderResponse) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *MerchantOrderResponse) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### SetChannelNameNil

`func (o *MerchantOrderResponse) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *MerchantOrderResponse) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetChannelId

`func (o *MerchantOrderResponse) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *MerchantOrderResponse) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *MerchantOrderResponse) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *MerchantOrderResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *MerchantOrderResponse) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *MerchantOrderResponse) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetGlobalChannelName

`func (o *MerchantOrderResponse) GetGlobalChannelName() string`

GetGlobalChannelName returns the GlobalChannelName field if non-nil, zero value otherwise.

### GetGlobalChannelNameOk

`func (o *MerchantOrderResponse) GetGlobalChannelNameOk() (*string, bool)`

GetGlobalChannelNameOk returns a tuple with the GlobalChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalChannelName

`func (o *MerchantOrderResponse) SetGlobalChannelName(v string)`

SetGlobalChannelName sets GlobalChannelName field to given value.

### HasGlobalChannelName

`func (o *MerchantOrderResponse) HasGlobalChannelName() bool`

HasGlobalChannelName returns a boolean if a field has been set.

### SetGlobalChannelNameNil

`func (o *MerchantOrderResponse) SetGlobalChannelNameNil(b bool)`

 SetGlobalChannelNameNil sets the value for GlobalChannelName to be an explicit nil

### UnsetGlobalChannelName
`func (o *MerchantOrderResponse) UnsetGlobalChannelName()`

UnsetGlobalChannelName ensures that no value is present for GlobalChannelName, not even an explicit nil
### GetGlobalChannelId

`func (o *MerchantOrderResponse) GetGlobalChannelId() int32`

GetGlobalChannelId returns the GlobalChannelId field if non-nil, zero value otherwise.

### GetGlobalChannelIdOk

`func (o *MerchantOrderResponse) GetGlobalChannelIdOk() (*int32, bool)`

GetGlobalChannelIdOk returns a tuple with the GlobalChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalChannelId

`func (o *MerchantOrderResponse) SetGlobalChannelId(v int32)`

SetGlobalChannelId sets GlobalChannelId field to given value.

### HasGlobalChannelId

`func (o *MerchantOrderResponse) HasGlobalChannelId() bool`

HasGlobalChannelId returns a boolean if a field has been set.

### SetGlobalChannelIdNil

`func (o *MerchantOrderResponse) SetGlobalChannelIdNil(b bool)`

 SetGlobalChannelIdNil sets the value for GlobalChannelId to be an explicit nil

### UnsetGlobalChannelId
`func (o *MerchantOrderResponse) UnsetGlobalChannelId()`

UnsetGlobalChannelId ensures that no value is present for GlobalChannelId, not even an explicit nil
### GetChannelOrderSupport

`func (o *MerchantOrderResponse) GetChannelOrderSupport() OrderSupport`

GetChannelOrderSupport returns the ChannelOrderSupport field if non-nil, zero value otherwise.

### GetChannelOrderSupportOk

`func (o *MerchantOrderResponse) GetChannelOrderSupportOk() (*OrderSupport, bool)`

GetChannelOrderSupportOk returns a tuple with the ChannelOrderSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderSupport

`func (o *MerchantOrderResponse) SetChannelOrderSupport(v OrderSupport)`

SetChannelOrderSupport sets ChannelOrderSupport field to given value.

### HasChannelOrderSupport

`func (o *MerchantOrderResponse) HasChannelOrderSupport() bool`

HasChannelOrderSupport returns a boolean if a field has been set.

### GetChannelOrderNo

`func (o *MerchantOrderResponse) GetChannelOrderNo() string`

GetChannelOrderNo returns the ChannelOrderNo field if non-nil, zero value otherwise.

### GetChannelOrderNoOk

`func (o *MerchantOrderResponse) GetChannelOrderNoOk() (*string, bool)`

GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderNo

`func (o *MerchantOrderResponse) SetChannelOrderNo(v string)`

SetChannelOrderNo sets ChannelOrderNo field to given value.

### HasChannelOrderNo

`func (o *MerchantOrderResponse) HasChannelOrderNo() bool`

HasChannelOrderNo returns a boolean if a field has been set.

### SetChannelOrderNoNil

`func (o *MerchantOrderResponse) SetChannelOrderNoNil(b bool)`

 SetChannelOrderNoNil sets the value for ChannelOrderNo to be an explicit nil

### UnsetChannelOrderNo
`func (o *MerchantOrderResponse) UnsetChannelOrderNo()`

UnsetChannelOrderNo ensures that no value is present for ChannelOrderNo, not even an explicit nil
### GetCommercialOrderNo

`func (o *MerchantOrderResponse) GetCommercialOrderNo() string`

GetCommercialOrderNo returns the CommercialOrderNo field if non-nil, zero value otherwise.

### GetCommercialOrderNoOk

`func (o *MerchantOrderResponse) GetCommercialOrderNoOk() (*string, bool)`

GetCommercialOrderNoOk returns a tuple with the CommercialOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialOrderNo

`func (o *MerchantOrderResponse) SetCommercialOrderNo(v string)`

SetCommercialOrderNo sets CommercialOrderNo field to given value.

### HasCommercialOrderNo

`func (o *MerchantOrderResponse) HasCommercialOrderNo() bool`

HasCommercialOrderNo returns a boolean if a field has been set.

### SetCommercialOrderNoNil

`func (o *MerchantOrderResponse) SetCommercialOrderNoNil(b bool)`

 SetCommercialOrderNoNil sets the value for CommercialOrderNo to be an explicit nil

### UnsetCommercialOrderNo
`func (o *MerchantOrderResponse) UnsetCommercialOrderNo()`

UnsetCommercialOrderNo ensures that no value is present for CommercialOrderNo, not even an explicit nil
### GetMerchantOrderNo

`func (o *MerchantOrderResponse) GetMerchantOrderNo() string`

GetMerchantOrderNo returns the MerchantOrderNo field if non-nil, zero value otherwise.

### GetMerchantOrderNoOk

`func (o *MerchantOrderResponse) GetMerchantOrderNoOk() (*string, bool)`

GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderNo

`func (o *MerchantOrderResponse) SetMerchantOrderNo(v string)`

SetMerchantOrderNo sets MerchantOrderNo field to given value.

### HasMerchantOrderNo

`func (o *MerchantOrderResponse) HasMerchantOrderNo() bool`

HasMerchantOrderNo returns a boolean if a field has been set.

### SetMerchantOrderNoNil

`func (o *MerchantOrderResponse) SetMerchantOrderNoNil(b bool)`

 SetMerchantOrderNoNil sets the value for MerchantOrderNo to be an explicit nil

### UnsetMerchantOrderNo
`func (o *MerchantOrderResponse) UnsetMerchantOrderNo()`

UnsetMerchantOrderNo ensures that no value is present for MerchantOrderNo, not even an explicit nil
### GetStatus

`func (o *MerchantOrderResponse) GetStatus() OrderStatusView`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MerchantOrderResponse) GetStatusOk() (*OrderStatusView, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MerchantOrderResponse) SetStatus(v OrderStatusView)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MerchantOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsBusinessOrder

`func (o *MerchantOrderResponse) GetIsBusinessOrder() bool`

GetIsBusinessOrder returns the IsBusinessOrder field if non-nil, zero value otherwise.

### GetIsBusinessOrderOk

`func (o *MerchantOrderResponse) GetIsBusinessOrderOk() (*bool, bool)`

GetIsBusinessOrderOk returns a tuple with the IsBusinessOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBusinessOrder

`func (o *MerchantOrderResponse) SetIsBusinessOrder(v bool)`

SetIsBusinessOrder sets IsBusinessOrder field to given value.

### HasIsBusinessOrder

`func (o *MerchantOrderResponse) HasIsBusinessOrder() bool`

HasIsBusinessOrder returns a boolean if a field has been set.

### GetAcknowledgedDate

`func (o *MerchantOrderResponse) GetAcknowledgedDate() time.Time`

GetAcknowledgedDate returns the AcknowledgedDate field if non-nil, zero value otherwise.

### GetAcknowledgedDateOk

`func (o *MerchantOrderResponse) GetAcknowledgedDateOk() (*time.Time, bool)`

GetAcknowledgedDateOk returns a tuple with the AcknowledgedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedDate

`func (o *MerchantOrderResponse) SetAcknowledgedDate(v time.Time)`

SetAcknowledgedDate sets AcknowledgedDate field to given value.

### HasAcknowledgedDate

`func (o *MerchantOrderResponse) HasAcknowledgedDate() bool`

HasAcknowledgedDate returns a boolean if a field has been set.

### SetAcknowledgedDateNil

`func (o *MerchantOrderResponse) SetAcknowledgedDateNil(b bool)`

 SetAcknowledgedDateNil sets the value for AcknowledgedDate to be an explicit nil

### UnsetAcknowledgedDate
`func (o *MerchantOrderResponse) UnsetAcknowledgedDate()`

UnsetAcknowledgedDate ensures that no value is present for AcknowledgedDate, not even an explicit nil
### GetCreatedAt

`func (o *MerchantOrderResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MerchantOrderResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MerchantOrderResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MerchantOrderResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *MerchantOrderResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *MerchantOrderResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *MerchantOrderResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MerchantOrderResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MerchantOrderResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MerchantOrderResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *MerchantOrderResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *MerchantOrderResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetClosedAt

`func (o *MerchantOrderResponse) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *MerchantOrderResponse) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *MerchantOrderResponse) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *MerchantOrderResponse) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### SetClosedAtNil

`func (o *MerchantOrderResponse) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *MerchantOrderResponse) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetMerchantComment

`func (o *MerchantOrderResponse) GetMerchantComment() string`

GetMerchantComment returns the MerchantComment field if non-nil, zero value otherwise.

### GetMerchantCommentOk

`func (o *MerchantOrderResponse) GetMerchantCommentOk() (*string, bool)`

GetMerchantCommentOk returns a tuple with the MerchantComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantComment

`func (o *MerchantOrderResponse) SetMerchantComment(v string)`

SetMerchantComment sets MerchantComment field to given value.

### HasMerchantComment

`func (o *MerchantOrderResponse) HasMerchantComment() bool`

HasMerchantComment returns a boolean if a field has been set.

### SetMerchantCommentNil

`func (o *MerchantOrderResponse) SetMerchantCommentNil(b bool)`

 SetMerchantCommentNil sets the value for MerchantComment to be an explicit nil

### UnsetMerchantComment
`func (o *MerchantOrderResponse) UnsetMerchantComment()`

UnsetMerchantComment ensures that no value is present for MerchantComment, not even an explicit nil
### GetBillingAddress

`func (o *MerchantOrderResponse) GetBillingAddress() MerchantAddressResponse`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *MerchantOrderResponse) GetBillingAddressOk() (*MerchantAddressResponse, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *MerchantOrderResponse) SetBillingAddress(v MerchantAddressResponse)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *MerchantOrderResponse) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *MerchantOrderResponse) GetShippingAddress() MerchantAddressResponse`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *MerchantOrderResponse) GetShippingAddressOk() (*MerchantAddressResponse, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *MerchantOrderResponse) SetShippingAddress(v MerchantAddressResponse)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *MerchantOrderResponse) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetSubTotalInclVat

`func (o *MerchantOrderResponse) GetSubTotalInclVat() float32`

GetSubTotalInclVat returns the SubTotalInclVat field if non-nil, zero value otherwise.

### GetSubTotalInclVatOk

`func (o *MerchantOrderResponse) GetSubTotalInclVatOk() (*float32, bool)`

GetSubTotalInclVatOk returns a tuple with the SubTotalInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotalInclVat

`func (o *MerchantOrderResponse) SetSubTotalInclVat(v float32)`

SetSubTotalInclVat sets SubTotalInclVat field to given value.

### HasSubTotalInclVat

`func (o *MerchantOrderResponse) HasSubTotalInclVat() bool`

HasSubTotalInclVat returns a boolean if a field has been set.

### SetSubTotalInclVatNil

`func (o *MerchantOrderResponse) SetSubTotalInclVatNil(b bool)`

 SetSubTotalInclVatNil sets the value for SubTotalInclVat to be an explicit nil

### UnsetSubTotalInclVat
`func (o *MerchantOrderResponse) UnsetSubTotalInclVat()`

UnsetSubTotalInclVat ensures that no value is present for SubTotalInclVat, not even an explicit nil
### GetSubTotalVat

`func (o *MerchantOrderResponse) GetSubTotalVat() float32`

GetSubTotalVat returns the SubTotalVat field if non-nil, zero value otherwise.

### GetSubTotalVatOk

`func (o *MerchantOrderResponse) GetSubTotalVatOk() (*float32, bool)`

GetSubTotalVatOk returns a tuple with the SubTotalVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotalVat

`func (o *MerchantOrderResponse) SetSubTotalVat(v float32)`

SetSubTotalVat sets SubTotalVat field to given value.

### HasSubTotalVat

`func (o *MerchantOrderResponse) HasSubTotalVat() bool`

HasSubTotalVat returns a boolean if a field has been set.

### SetSubTotalVatNil

`func (o *MerchantOrderResponse) SetSubTotalVatNil(b bool)`

 SetSubTotalVatNil sets the value for SubTotalVat to be an explicit nil

### UnsetSubTotalVat
`func (o *MerchantOrderResponse) UnsetSubTotalVat()`

UnsetSubTotalVat ensures that no value is present for SubTotalVat, not even an explicit nil
### GetShippingCostsInclVat

`func (o *MerchantOrderResponse) GetShippingCostsInclVat() float32`

GetShippingCostsInclVat returns the ShippingCostsInclVat field if non-nil, zero value otherwise.

### GetShippingCostsInclVatOk

`func (o *MerchantOrderResponse) GetShippingCostsInclVatOk() (*float32, bool)`

GetShippingCostsInclVatOk returns a tuple with the ShippingCostsInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCostsInclVat

`func (o *MerchantOrderResponse) SetShippingCostsInclVat(v float32)`

SetShippingCostsInclVat sets ShippingCostsInclVat field to given value.

### HasShippingCostsInclVat

`func (o *MerchantOrderResponse) HasShippingCostsInclVat() bool`

HasShippingCostsInclVat returns a boolean if a field has been set.

### GetShippingCostsVat

`func (o *MerchantOrderResponse) GetShippingCostsVat() float32`

GetShippingCostsVat returns the ShippingCostsVat field if non-nil, zero value otherwise.

### GetShippingCostsVatOk

`func (o *MerchantOrderResponse) GetShippingCostsVatOk() (*float32, bool)`

GetShippingCostsVatOk returns a tuple with the ShippingCostsVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCostsVat

`func (o *MerchantOrderResponse) SetShippingCostsVat(v float32)`

SetShippingCostsVat sets ShippingCostsVat field to given value.

### HasShippingCostsVat

`func (o *MerchantOrderResponse) HasShippingCostsVat() bool`

HasShippingCostsVat returns a boolean if a field has been set.

### SetShippingCostsVatNil

`func (o *MerchantOrderResponse) SetShippingCostsVatNil(b bool)`

 SetShippingCostsVatNil sets the value for ShippingCostsVat to be an explicit nil

### UnsetShippingCostsVat
`func (o *MerchantOrderResponse) UnsetShippingCostsVat()`

UnsetShippingCostsVat ensures that no value is present for ShippingCostsVat, not even an explicit nil
### GetTotalInclVat

`func (o *MerchantOrderResponse) GetTotalInclVat() float32`

GetTotalInclVat returns the TotalInclVat field if non-nil, zero value otherwise.

### GetTotalInclVatOk

`func (o *MerchantOrderResponse) GetTotalInclVatOk() (*float32, bool)`

GetTotalInclVatOk returns a tuple with the TotalInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInclVat

`func (o *MerchantOrderResponse) SetTotalInclVat(v float32)`

SetTotalInclVat sets TotalInclVat field to given value.

### HasTotalInclVat

`func (o *MerchantOrderResponse) HasTotalInclVat() bool`

HasTotalInclVat returns a boolean if a field has been set.

### GetTotalVat

`func (o *MerchantOrderResponse) GetTotalVat() float32`

GetTotalVat returns the TotalVat field if non-nil, zero value otherwise.

### GetTotalVatOk

`func (o *MerchantOrderResponse) GetTotalVatOk() (*float32, bool)`

GetTotalVatOk returns a tuple with the TotalVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVat

`func (o *MerchantOrderResponse) SetTotalVat(v float32)`

SetTotalVat sets TotalVat field to given value.

### HasTotalVat

`func (o *MerchantOrderResponse) HasTotalVat() bool`

HasTotalVat returns a boolean if a field has been set.

### SetTotalVatNil

`func (o *MerchantOrderResponse) SetTotalVatNil(b bool)`

 SetTotalVatNil sets the value for TotalVat to be an explicit nil

### UnsetTotalVat
`func (o *MerchantOrderResponse) UnsetTotalVat()`

UnsetTotalVat ensures that no value is present for TotalVat, not even an explicit nil
### GetOriginalSubTotalInclVat

`func (o *MerchantOrderResponse) GetOriginalSubTotalInclVat() float32`

GetOriginalSubTotalInclVat returns the OriginalSubTotalInclVat field if non-nil, zero value otherwise.

### GetOriginalSubTotalInclVatOk

`func (o *MerchantOrderResponse) GetOriginalSubTotalInclVatOk() (*float32, bool)`

GetOriginalSubTotalInclVatOk returns a tuple with the OriginalSubTotalInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSubTotalInclVat

`func (o *MerchantOrderResponse) SetOriginalSubTotalInclVat(v float32)`

SetOriginalSubTotalInclVat sets OriginalSubTotalInclVat field to given value.

### HasOriginalSubTotalInclVat

`func (o *MerchantOrderResponse) HasOriginalSubTotalInclVat() bool`

HasOriginalSubTotalInclVat returns a boolean if a field has been set.

### SetOriginalSubTotalInclVatNil

`func (o *MerchantOrderResponse) SetOriginalSubTotalInclVatNil(b bool)`

 SetOriginalSubTotalInclVatNil sets the value for OriginalSubTotalInclVat to be an explicit nil

### UnsetOriginalSubTotalInclVat
`func (o *MerchantOrderResponse) UnsetOriginalSubTotalInclVat()`

UnsetOriginalSubTotalInclVat ensures that no value is present for OriginalSubTotalInclVat, not even an explicit nil
### GetOriginalSubTotalVat

`func (o *MerchantOrderResponse) GetOriginalSubTotalVat() float32`

GetOriginalSubTotalVat returns the OriginalSubTotalVat field if non-nil, zero value otherwise.

### GetOriginalSubTotalVatOk

`func (o *MerchantOrderResponse) GetOriginalSubTotalVatOk() (*float32, bool)`

GetOriginalSubTotalVatOk returns a tuple with the OriginalSubTotalVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSubTotalVat

`func (o *MerchantOrderResponse) SetOriginalSubTotalVat(v float32)`

SetOriginalSubTotalVat sets OriginalSubTotalVat field to given value.

### HasOriginalSubTotalVat

`func (o *MerchantOrderResponse) HasOriginalSubTotalVat() bool`

HasOriginalSubTotalVat returns a boolean if a field has been set.

### SetOriginalSubTotalVatNil

`func (o *MerchantOrderResponse) SetOriginalSubTotalVatNil(b bool)`

 SetOriginalSubTotalVatNil sets the value for OriginalSubTotalVat to be an explicit nil

### UnsetOriginalSubTotalVat
`func (o *MerchantOrderResponse) UnsetOriginalSubTotalVat()`

UnsetOriginalSubTotalVat ensures that no value is present for OriginalSubTotalVat, not even an explicit nil
### GetOriginalShippingCostsInclVat

`func (o *MerchantOrderResponse) GetOriginalShippingCostsInclVat() float32`

GetOriginalShippingCostsInclVat returns the OriginalShippingCostsInclVat field if non-nil, zero value otherwise.

### GetOriginalShippingCostsInclVatOk

`func (o *MerchantOrderResponse) GetOriginalShippingCostsInclVatOk() (*float32, bool)`

GetOriginalShippingCostsInclVatOk returns a tuple with the OriginalShippingCostsInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalShippingCostsInclVat

`func (o *MerchantOrderResponse) SetOriginalShippingCostsInclVat(v float32)`

SetOriginalShippingCostsInclVat sets OriginalShippingCostsInclVat field to given value.

### HasOriginalShippingCostsInclVat

`func (o *MerchantOrderResponse) HasOriginalShippingCostsInclVat() bool`

HasOriginalShippingCostsInclVat returns a boolean if a field has been set.

### SetOriginalShippingCostsInclVatNil

`func (o *MerchantOrderResponse) SetOriginalShippingCostsInclVatNil(b bool)`

 SetOriginalShippingCostsInclVatNil sets the value for OriginalShippingCostsInclVat to be an explicit nil

### UnsetOriginalShippingCostsInclVat
`func (o *MerchantOrderResponse) UnsetOriginalShippingCostsInclVat()`

UnsetOriginalShippingCostsInclVat ensures that no value is present for OriginalShippingCostsInclVat, not even an explicit nil
### GetOriginalShippingCostsVat

`func (o *MerchantOrderResponse) GetOriginalShippingCostsVat() float32`

GetOriginalShippingCostsVat returns the OriginalShippingCostsVat field if non-nil, zero value otherwise.

### GetOriginalShippingCostsVatOk

`func (o *MerchantOrderResponse) GetOriginalShippingCostsVatOk() (*float32, bool)`

GetOriginalShippingCostsVatOk returns a tuple with the OriginalShippingCostsVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalShippingCostsVat

`func (o *MerchantOrderResponse) SetOriginalShippingCostsVat(v float32)`

SetOriginalShippingCostsVat sets OriginalShippingCostsVat field to given value.

### HasOriginalShippingCostsVat

`func (o *MerchantOrderResponse) HasOriginalShippingCostsVat() bool`

HasOriginalShippingCostsVat returns a boolean if a field has been set.

### SetOriginalShippingCostsVatNil

`func (o *MerchantOrderResponse) SetOriginalShippingCostsVatNil(b bool)`

 SetOriginalShippingCostsVatNil sets the value for OriginalShippingCostsVat to be an explicit nil

### UnsetOriginalShippingCostsVat
`func (o *MerchantOrderResponse) UnsetOriginalShippingCostsVat()`

UnsetOriginalShippingCostsVat ensures that no value is present for OriginalShippingCostsVat, not even an explicit nil
### GetOriginalTotalInclVat

`func (o *MerchantOrderResponse) GetOriginalTotalInclVat() float32`

GetOriginalTotalInclVat returns the OriginalTotalInclVat field if non-nil, zero value otherwise.

### GetOriginalTotalInclVatOk

`func (o *MerchantOrderResponse) GetOriginalTotalInclVatOk() (*float32, bool)`

GetOriginalTotalInclVatOk returns a tuple with the OriginalTotalInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTotalInclVat

`func (o *MerchantOrderResponse) SetOriginalTotalInclVat(v float32)`

SetOriginalTotalInclVat sets OriginalTotalInclVat field to given value.

### HasOriginalTotalInclVat

`func (o *MerchantOrderResponse) HasOriginalTotalInclVat() bool`

HasOriginalTotalInclVat returns a boolean if a field has been set.

### SetOriginalTotalInclVatNil

`func (o *MerchantOrderResponse) SetOriginalTotalInclVatNil(b bool)`

 SetOriginalTotalInclVatNil sets the value for OriginalTotalInclVat to be an explicit nil

### UnsetOriginalTotalInclVat
`func (o *MerchantOrderResponse) UnsetOriginalTotalInclVat()`

UnsetOriginalTotalInclVat ensures that no value is present for OriginalTotalInclVat, not even an explicit nil
### GetOriginalTotalVat

`func (o *MerchantOrderResponse) GetOriginalTotalVat() float32`

GetOriginalTotalVat returns the OriginalTotalVat field if non-nil, zero value otherwise.

### GetOriginalTotalVatOk

`func (o *MerchantOrderResponse) GetOriginalTotalVatOk() (*float32, bool)`

GetOriginalTotalVatOk returns a tuple with the OriginalTotalVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTotalVat

`func (o *MerchantOrderResponse) SetOriginalTotalVat(v float32)`

SetOriginalTotalVat sets OriginalTotalVat field to given value.

### HasOriginalTotalVat

`func (o *MerchantOrderResponse) HasOriginalTotalVat() bool`

HasOriginalTotalVat returns a boolean if a field has been set.

### SetOriginalTotalVatNil

`func (o *MerchantOrderResponse) SetOriginalTotalVatNil(b bool)`

 SetOriginalTotalVatNil sets the value for OriginalTotalVat to be an explicit nil

### UnsetOriginalTotalVat
`func (o *MerchantOrderResponse) UnsetOriginalTotalVat()`

UnsetOriginalTotalVat ensures that no value is present for OriginalTotalVat, not even an explicit nil
### GetSubTotalExclVat

`func (o *MerchantOrderResponse) GetSubTotalExclVat() float32`

GetSubTotalExclVat returns the SubTotalExclVat field if non-nil, zero value otherwise.

### GetSubTotalExclVatOk

`func (o *MerchantOrderResponse) GetSubTotalExclVatOk() (*float32, bool)`

GetSubTotalExclVatOk returns a tuple with the SubTotalExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotalExclVat

`func (o *MerchantOrderResponse) SetSubTotalExclVat(v float32)`

SetSubTotalExclVat sets SubTotalExclVat field to given value.

### HasSubTotalExclVat

`func (o *MerchantOrderResponse) HasSubTotalExclVat() bool`

HasSubTotalExclVat returns a boolean if a field has been set.

### SetSubTotalExclVatNil

`func (o *MerchantOrderResponse) SetSubTotalExclVatNil(b bool)`

 SetSubTotalExclVatNil sets the value for SubTotalExclVat to be an explicit nil

### UnsetSubTotalExclVat
`func (o *MerchantOrderResponse) UnsetSubTotalExclVat()`

UnsetSubTotalExclVat ensures that no value is present for SubTotalExclVat, not even an explicit nil
### GetTotalExclVat

`func (o *MerchantOrderResponse) GetTotalExclVat() float32`

GetTotalExclVat returns the TotalExclVat field if non-nil, zero value otherwise.

### GetTotalExclVatOk

`func (o *MerchantOrderResponse) GetTotalExclVatOk() (*float32, bool)`

GetTotalExclVatOk returns a tuple with the TotalExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExclVat

`func (o *MerchantOrderResponse) SetTotalExclVat(v float32)`

SetTotalExclVat sets TotalExclVat field to given value.

### HasTotalExclVat

`func (o *MerchantOrderResponse) HasTotalExclVat() bool`

HasTotalExclVat returns a boolean if a field has been set.

### SetTotalExclVatNil

`func (o *MerchantOrderResponse) SetTotalExclVatNil(b bool)`

 SetTotalExclVatNil sets the value for TotalExclVat to be an explicit nil

### UnsetTotalExclVat
`func (o *MerchantOrderResponse) UnsetTotalExclVat()`

UnsetTotalExclVat ensures that no value is present for TotalExclVat, not even an explicit nil
### GetShippingCostsExclVat

`func (o *MerchantOrderResponse) GetShippingCostsExclVat() float32`

GetShippingCostsExclVat returns the ShippingCostsExclVat field if non-nil, zero value otherwise.

### GetShippingCostsExclVatOk

`func (o *MerchantOrderResponse) GetShippingCostsExclVatOk() (*float32, bool)`

GetShippingCostsExclVatOk returns a tuple with the ShippingCostsExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCostsExclVat

`func (o *MerchantOrderResponse) SetShippingCostsExclVat(v float32)`

SetShippingCostsExclVat sets ShippingCostsExclVat field to given value.

### HasShippingCostsExclVat

`func (o *MerchantOrderResponse) HasShippingCostsExclVat() bool`

HasShippingCostsExclVat returns a boolean if a field has been set.

### SetShippingCostsExclVatNil

`func (o *MerchantOrderResponse) SetShippingCostsExclVatNil(b bool)`

 SetShippingCostsExclVatNil sets the value for ShippingCostsExclVat to be an explicit nil

### UnsetShippingCostsExclVat
`func (o *MerchantOrderResponse) UnsetShippingCostsExclVat()`

UnsetShippingCostsExclVat ensures that no value is present for ShippingCostsExclVat, not even an explicit nil
### GetOriginalSubTotalExclVat

`func (o *MerchantOrderResponse) GetOriginalSubTotalExclVat() float32`

GetOriginalSubTotalExclVat returns the OriginalSubTotalExclVat field if non-nil, zero value otherwise.

### GetOriginalSubTotalExclVatOk

`func (o *MerchantOrderResponse) GetOriginalSubTotalExclVatOk() (*float32, bool)`

GetOriginalSubTotalExclVatOk returns a tuple with the OriginalSubTotalExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSubTotalExclVat

`func (o *MerchantOrderResponse) SetOriginalSubTotalExclVat(v float32)`

SetOriginalSubTotalExclVat sets OriginalSubTotalExclVat field to given value.

### HasOriginalSubTotalExclVat

`func (o *MerchantOrderResponse) HasOriginalSubTotalExclVat() bool`

HasOriginalSubTotalExclVat returns a boolean if a field has been set.

### SetOriginalSubTotalExclVatNil

`func (o *MerchantOrderResponse) SetOriginalSubTotalExclVatNil(b bool)`

 SetOriginalSubTotalExclVatNil sets the value for OriginalSubTotalExclVat to be an explicit nil

### UnsetOriginalSubTotalExclVat
`func (o *MerchantOrderResponse) UnsetOriginalSubTotalExclVat()`

UnsetOriginalSubTotalExclVat ensures that no value is present for OriginalSubTotalExclVat, not even an explicit nil
### GetOriginalShippingCostsExclVat

`func (o *MerchantOrderResponse) GetOriginalShippingCostsExclVat() float32`

GetOriginalShippingCostsExclVat returns the OriginalShippingCostsExclVat field if non-nil, zero value otherwise.

### GetOriginalShippingCostsExclVatOk

`func (o *MerchantOrderResponse) GetOriginalShippingCostsExclVatOk() (*float32, bool)`

GetOriginalShippingCostsExclVatOk returns a tuple with the OriginalShippingCostsExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalShippingCostsExclVat

`func (o *MerchantOrderResponse) SetOriginalShippingCostsExclVat(v float32)`

SetOriginalShippingCostsExclVat sets OriginalShippingCostsExclVat field to given value.

### HasOriginalShippingCostsExclVat

`func (o *MerchantOrderResponse) HasOriginalShippingCostsExclVat() bool`

HasOriginalShippingCostsExclVat returns a boolean if a field has been set.

### SetOriginalShippingCostsExclVatNil

`func (o *MerchantOrderResponse) SetOriginalShippingCostsExclVatNil(b bool)`

 SetOriginalShippingCostsExclVatNil sets the value for OriginalShippingCostsExclVat to be an explicit nil

### UnsetOriginalShippingCostsExclVat
`func (o *MerchantOrderResponse) UnsetOriginalShippingCostsExclVat()`

UnsetOriginalShippingCostsExclVat ensures that no value is present for OriginalShippingCostsExclVat, not even an explicit nil
### GetOriginalTotalExclVat

`func (o *MerchantOrderResponse) GetOriginalTotalExclVat() float32`

GetOriginalTotalExclVat returns the OriginalTotalExclVat field if non-nil, zero value otherwise.

### GetOriginalTotalExclVatOk

`func (o *MerchantOrderResponse) GetOriginalTotalExclVatOk() (*float32, bool)`

GetOriginalTotalExclVatOk returns a tuple with the OriginalTotalExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTotalExclVat

`func (o *MerchantOrderResponse) SetOriginalTotalExclVat(v float32)`

SetOriginalTotalExclVat sets OriginalTotalExclVat field to given value.

### HasOriginalTotalExclVat

`func (o *MerchantOrderResponse) HasOriginalTotalExclVat() bool`

HasOriginalTotalExclVat returns a boolean if a field has been set.

### SetOriginalTotalExclVatNil

`func (o *MerchantOrderResponse) SetOriginalTotalExclVatNil(b bool)`

 SetOriginalTotalExclVatNil sets the value for OriginalTotalExclVat to be an explicit nil

### UnsetOriginalTotalExclVat
`func (o *MerchantOrderResponse) UnsetOriginalTotalExclVat()`

UnsetOriginalTotalExclVat ensures that no value is present for OriginalTotalExclVat, not even an explicit nil
### GetOriginalSubTotalFee

`func (o *MerchantOrderResponse) GetOriginalSubTotalFee() float32`

GetOriginalSubTotalFee returns the OriginalSubTotalFee field if non-nil, zero value otherwise.

### GetOriginalSubTotalFeeOk

`func (o *MerchantOrderResponse) GetOriginalSubTotalFeeOk() (*float32, bool)`

GetOriginalSubTotalFeeOk returns a tuple with the OriginalSubTotalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSubTotalFee

`func (o *MerchantOrderResponse) SetOriginalSubTotalFee(v float32)`

SetOriginalSubTotalFee sets OriginalSubTotalFee field to given value.

### HasOriginalSubTotalFee

`func (o *MerchantOrderResponse) HasOriginalSubTotalFee() bool`

HasOriginalSubTotalFee returns a boolean if a field has been set.

### GetSubTotalFee

`func (o *MerchantOrderResponse) GetSubTotalFee() float32`

GetSubTotalFee returns the SubTotalFee field if non-nil, zero value otherwise.

### GetSubTotalFeeOk

`func (o *MerchantOrderResponse) GetSubTotalFeeOk() (*float32, bool)`

GetSubTotalFeeOk returns a tuple with the SubTotalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotalFee

`func (o *MerchantOrderResponse) SetSubTotalFee(v float32)`

SetSubTotalFee sets SubTotalFee field to given value.

### HasSubTotalFee

`func (o *MerchantOrderResponse) HasSubTotalFee() bool`

HasSubTotalFee returns a boolean if a field has been set.

### GetOriginalOrderFee

`func (o *MerchantOrderResponse) GetOriginalOrderFee() float32`

GetOriginalOrderFee returns the OriginalOrderFee field if non-nil, zero value otherwise.

### GetOriginalOrderFeeOk

`func (o *MerchantOrderResponse) GetOriginalOrderFeeOk() (*float32, bool)`

GetOriginalOrderFeeOk returns a tuple with the OriginalOrderFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalOrderFee

`func (o *MerchantOrderResponse) SetOriginalOrderFee(v float32)`

SetOriginalOrderFee sets OriginalOrderFee field to given value.

### HasOriginalOrderFee

`func (o *MerchantOrderResponse) HasOriginalOrderFee() bool`

HasOriginalOrderFee returns a boolean if a field has been set.

### GetOrderFee

`func (o *MerchantOrderResponse) GetOrderFee() float32`

GetOrderFee returns the OrderFee field if non-nil, zero value otherwise.

### GetOrderFeeOk

`func (o *MerchantOrderResponse) GetOrderFeeOk() (*float32, bool)`

GetOrderFeeOk returns a tuple with the OrderFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderFee

`func (o *MerchantOrderResponse) SetOrderFee(v float32)`

SetOrderFee sets OrderFee field to given value.

### HasOrderFee

`func (o *MerchantOrderResponse) HasOrderFee() bool`

HasOrderFee returns a boolean if a field has been set.

### GetOriginalTotalFee

`func (o *MerchantOrderResponse) GetOriginalTotalFee() float32`

GetOriginalTotalFee returns the OriginalTotalFee field if non-nil, zero value otherwise.

### GetOriginalTotalFeeOk

`func (o *MerchantOrderResponse) GetOriginalTotalFeeOk() (*float32, bool)`

GetOriginalTotalFeeOk returns a tuple with the OriginalTotalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTotalFee

`func (o *MerchantOrderResponse) SetOriginalTotalFee(v float32)`

SetOriginalTotalFee sets OriginalTotalFee field to given value.

### HasOriginalTotalFee

`func (o *MerchantOrderResponse) HasOriginalTotalFee() bool`

HasOriginalTotalFee returns a boolean if a field has been set.

### GetTotalFee

`func (o *MerchantOrderResponse) GetTotalFee() float32`

GetTotalFee returns the TotalFee field if non-nil, zero value otherwise.

### GetTotalFeeOk

`func (o *MerchantOrderResponse) GetTotalFeeOk() (*float32, bool)`

GetTotalFeeOk returns a tuple with the TotalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFee

`func (o *MerchantOrderResponse) SetTotalFee(v float32)`

SetTotalFee sets TotalFee field to given value.

### HasTotalFee

`func (o *MerchantOrderResponse) HasTotalFee() bool`

HasTotalFee returns a boolean if a field has been set.

### GetLines

`func (o *MerchantOrderResponse) GetLines() []MerchantOrderLineResponse`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantOrderResponse) GetLinesOk() (*[]MerchantOrderLineResponse, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantOrderResponse) SetLines(v []MerchantOrderLineResponse)`

SetLines sets Lines field to given value.

### HasLines

`func (o *MerchantOrderResponse) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *MerchantOrderResponse) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *MerchantOrderResponse) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil
### GetPhone

`func (o *MerchantOrderResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MerchantOrderResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MerchantOrderResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MerchantOrderResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *MerchantOrderResponse) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *MerchantOrderResponse) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *MerchantOrderResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MerchantOrderResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MerchantOrderResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLanguageCode

`func (o *MerchantOrderResponse) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *MerchantOrderResponse) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *MerchantOrderResponse) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *MerchantOrderResponse) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### SetLanguageCodeNil

`func (o *MerchantOrderResponse) SetLanguageCodeNil(b bool)`

 SetLanguageCodeNil sets the value for LanguageCode to be an explicit nil

### UnsetLanguageCode
`func (o *MerchantOrderResponse) UnsetLanguageCode()`

UnsetLanguageCode ensures that no value is present for LanguageCode, not even an explicit nil
### GetCompanyRegistrationNo

`func (o *MerchantOrderResponse) GetCompanyRegistrationNo() string`

GetCompanyRegistrationNo returns the CompanyRegistrationNo field if non-nil, zero value otherwise.

### GetCompanyRegistrationNoOk

`func (o *MerchantOrderResponse) GetCompanyRegistrationNoOk() (*string, bool)`

GetCompanyRegistrationNoOk returns a tuple with the CompanyRegistrationNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyRegistrationNo

`func (o *MerchantOrderResponse) SetCompanyRegistrationNo(v string)`

SetCompanyRegistrationNo sets CompanyRegistrationNo field to given value.

### HasCompanyRegistrationNo

`func (o *MerchantOrderResponse) HasCompanyRegistrationNo() bool`

HasCompanyRegistrationNo returns a boolean if a field has been set.

### SetCompanyRegistrationNoNil

`func (o *MerchantOrderResponse) SetCompanyRegistrationNoNil(b bool)`

 SetCompanyRegistrationNoNil sets the value for CompanyRegistrationNo to be an explicit nil

### UnsetCompanyRegistrationNo
`func (o *MerchantOrderResponse) UnsetCompanyRegistrationNo()`

UnsetCompanyRegistrationNo ensures that no value is present for CompanyRegistrationNo, not even an explicit nil
### GetVatNo

`func (o *MerchantOrderResponse) GetVatNo() string`

GetVatNo returns the VatNo field if non-nil, zero value otherwise.

### GetVatNoOk

`func (o *MerchantOrderResponse) GetVatNoOk() (*string, bool)`

GetVatNoOk returns a tuple with the VatNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNo

`func (o *MerchantOrderResponse) SetVatNo(v string)`

SetVatNo sets VatNo field to given value.

### HasVatNo

`func (o *MerchantOrderResponse) HasVatNo() bool`

HasVatNo returns a boolean if a field has been set.

### SetVatNoNil

`func (o *MerchantOrderResponse) SetVatNoNil(b bool)`

 SetVatNoNil sets the value for VatNo to be an explicit nil

### UnsetVatNo
`func (o *MerchantOrderResponse) UnsetVatNo()`

UnsetVatNo ensures that no value is present for VatNo, not even an explicit nil
### GetPaymentMethod

`func (o *MerchantOrderResponse) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *MerchantOrderResponse) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *MerchantOrderResponse) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *MerchantOrderResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### SetPaymentMethodNil

`func (o *MerchantOrderResponse) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *MerchantOrderResponse) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
### GetPaymentReferenceNo

`func (o *MerchantOrderResponse) GetPaymentReferenceNo() string`

GetPaymentReferenceNo returns the PaymentReferenceNo field if non-nil, zero value otherwise.

### GetPaymentReferenceNoOk

`func (o *MerchantOrderResponse) GetPaymentReferenceNoOk() (*string, bool)`

GetPaymentReferenceNoOk returns a tuple with the PaymentReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentReferenceNo

`func (o *MerchantOrderResponse) SetPaymentReferenceNo(v string)`

SetPaymentReferenceNo sets PaymentReferenceNo field to given value.

### HasPaymentReferenceNo

`func (o *MerchantOrderResponse) HasPaymentReferenceNo() bool`

HasPaymentReferenceNo returns a boolean if a field has been set.

### SetPaymentReferenceNoNil

`func (o *MerchantOrderResponse) SetPaymentReferenceNoNil(b bool)`

 SetPaymentReferenceNoNil sets the value for PaymentReferenceNo to be an explicit nil

### UnsetPaymentReferenceNo
`func (o *MerchantOrderResponse) UnsetPaymentReferenceNo()`

UnsetPaymentReferenceNo ensures that no value is present for PaymentReferenceNo, not even an explicit nil
### GetCurrencyCode

`func (o *MerchantOrderResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *MerchantOrderResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *MerchantOrderResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetOrderDate

`func (o *MerchantOrderResponse) GetOrderDate() time.Time`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *MerchantOrderResponse) GetOrderDateOk() (*time.Time, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *MerchantOrderResponse) SetOrderDate(v time.Time)`

SetOrderDate sets OrderDate field to given value.


### GetChannelCustomerNo

`func (o *MerchantOrderResponse) GetChannelCustomerNo() string`

GetChannelCustomerNo returns the ChannelCustomerNo field if non-nil, zero value otherwise.

### GetChannelCustomerNoOk

`func (o *MerchantOrderResponse) GetChannelCustomerNoOk() (*string, bool)`

GetChannelCustomerNoOk returns a tuple with the ChannelCustomerNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCustomerNo

`func (o *MerchantOrderResponse) SetChannelCustomerNo(v string)`

SetChannelCustomerNo sets ChannelCustomerNo field to given value.

### HasChannelCustomerNo

`func (o *MerchantOrderResponse) HasChannelCustomerNo() bool`

HasChannelCustomerNo returns a boolean if a field has been set.

### SetChannelCustomerNoNil

`func (o *MerchantOrderResponse) SetChannelCustomerNoNil(b bool)`

 SetChannelCustomerNoNil sets the value for ChannelCustomerNo to be an explicit nil

### UnsetChannelCustomerNo
`func (o *MerchantOrderResponse) UnsetChannelCustomerNo()`

UnsetChannelCustomerNo ensures that no value is present for ChannelCustomerNo, not even an explicit nil
### GetExtraData

`func (o *MerchantOrderResponse) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantOrderResponse) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantOrderResponse) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantOrderResponse) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantOrderResponse) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantOrderResponse) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


