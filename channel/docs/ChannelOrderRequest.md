# ChannelOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress** | [**ChannelAddressRequest**](ChannelAddressRequest.md) |  | 
**ShippingAddress** | [**ChannelAddressRequest**](ChannelAddressRequest.md) |  | 
**ChannelOrderNo** | **string** | The unique order reference used by the Channel. | 
**CommercialOrderNo** | Pointer to **NullableString** | Optional (if identical to the ChannelOrderNo). The order reference used by the Channel for commercial purposes (e.g. on the invoice).  Can be different from the ChannelOrderNo. For example when the internal unique order reference is a unique id or guid, while the commercial order reference is a human readable reference that can be reused. | [optional] 
**IsBusinessOrder** | Pointer to **NullableBool** | Optional. Is a business order (default value is false).  If not provided the VAT Number will be checked. If a VAT Number is found, IsBusinessOrder will be set to true.  No VAT will be calculated when set to true. | [optional] 
**KeyIsMerchantProductNo** | Pointer to **bool** | Optional. Is the MPN used as key for the product (default value is false). | [optional] 
**Lines** | [**[]ChannelOrderLineRequest**](ChannelOrderLineRequest.md) | The order lines. | 
**ShippingCostsInclVat** | **float32** | The shipping fee including VAT  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | 
**OrderFee** | Pointer to **float32** | The fee on order itself (besides the fees on the order lines)  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
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

### NewChannelOrderRequest

`func NewChannelOrderRequest(billingAddress ChannelAddressRequest, shippingAddress ChannelAddressRequest, channelOrderNo string, lines []ChannelOrderLineRequest, shippingCostsInclVat float32, email string, currencyCode string, orderDate time.Time, ) *ChannelOrderRequest`

NewChannelOrderRequest instantiates a new ChannelOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelOrderRequestWithDefaults

`func NewChannelOrderRequestWithDefaults() *ChannelOrderRequest`

NewChannelOrderRequestWithDefaults instantiates a new ChannelOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress

`func (o *ChannelOrderRequest) GetBillingAddress() ChannelAddressRequest`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *ChannelOrderRequest) GetBillingAddressOk() (*ChannelAddressRequest, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *ChannelOrderRequest) SetBillingAddress(v ChannelAddressRequest)`

SetBillingAddress sets BillingAddress field to given value.


### GetShippingAddress

`func (o *ChannelOrderRequest) GetShippingAddress() ChannelAddressRequest`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *ChannelOrderRequest) GetShippingAddressOk() (*ChannelAddressRequest, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *ChannelOrderRequest) SetShippingAddress(v ChannelAddressRequest)`

SetShippingAddress sets ShippingAddress field to given value.


### GetChannelOrderNo

`func (o *ChannelOrderRequest) GetChannelOrderNo() string`

GetChannelOrderNo returns the ChannelOrderNo field if non-nil, zero value otherwise.

### GetChannelOrderNoOk

`func (o *ChannelOrderRequest) GetChannelOrderNoOk() (*string, bool)`

GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderNo

`func (o *ChannelOrderRequest) SetChannelOrderNo(v string)`

SetChannelOrderNo sets ChannelOrderNo field to given value.


### GetCommercialOrderNo

`func (o *ChannelOrderRequest) GetCommercialOrderNo() string`

GetCommercialOrderNo returns the CommercialOrderNo field if non-nil, zero value otherwise.

### GetCommercialOrderNoOk

`func (o *ChannelOrderRequest) GetCommercialOrderNoOk() (*string, bool)`

GetCommercialOrderNoOk returns a tuple with the CommercialOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialOrderNo

`func (o *ChannelOrderRequest) SetCommercialOrderNo(v string)`

SetCommercialOrderNo sets CommercialOrderNo field to given value.

### HasCommercialOrderNo

`func (o *ChannelOrderRequest) HasCommercialOrderNo() bool`

HasCommercialOrderNo returns a boolean if a field has been set.

### SetCommercialOrderNoNil

`func (o *ChannelOrderRequest) SetCommercialOrderNoNil(b bool)`

 SetCommercialOrderNoNil sets the value for CommercialOrderNo to be an explicit nil

### UnsetCommercialOrderNo
`func (o *ChannelOrderRequest) UnsetCommercialOrderNo()`

UnsetCommercialOrderNo ensures that no value is present for CommercialOrderNo, not even an explicit nil
### GetIsBusinessOrder

`func (o *ChannelOrderRequest) GetIsBusinessOrder() bool`

GetIsBusinessOrder returns the IsBusinessOrder field if non-nil, zero value otherwise.

### GetIsBusinessOrderOk

`func (o *ChannelOrderRequest) GetIsBusinessOrderOk() (*bool, bool)`

GetIsBusinessOrderOk returns a tuple with the IsBusinessOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBusinessOrder

`func (o *ChannelOrderRequest) SetIsBusinessOrder(v bool)`

SetIsBusinessOrder sets IsBusinessOrder field to given value.

### HasIsBusinessOrder

`func (o *ChannelOrderRequest) HasIsBusinessOrder() bool`

HasIsBusinessOrder returns a boolean if a field has been set.

### SetIsBusinessOrderNil

`func (o *ChannelOrderRequest) SetIsBusinessOrderNil(b bool)`

 SetIsBusinessOrderNil sets the value for IsBusinessOrder to be an explicit nil

### UnsetIsBusinessOrder
`func (o *ChannelOrderRequest) UnsetIsBusinessOrder()`

UnsetIsBusinessOrder ensures that no value is present for IsBusinessOrder, not even an explicit nil
### GetKeyIsMerchantProductNo

`func (o *ChannelOrderRequest) GetKeyIsMerchantProductNo() bool`

GetKeyIsMerchantProductNo returns the KeyIsMerchantProductNo field if non-nil, zero value otherwise.

### GetKeyIsMerchantProductNoOk

`func (o *ChannelOrderRequest) GetKeyIsMerchantProductNoOk() (*bool, bool)`

GetKeyIsMerchantProductNoOk returns a tuple with the KeyIsMerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyIsMerchantProductNo

`func (o *ChannelOrderRequest) SetKeyIsMerchantProductNo(v bool)`

SetKeyIsMerchantProductNo sets KeyIsMerchantProductNo field to given value.

### HasKeyIsMerchantProductNo

`func (o *ChannelOrderRequest) HasKeyIsMerchantProductNo() bool`

HasKeyIsMerchantProductNo returns a boolean if a field has been set.

### GetLines

`func (o *ChannelOrderRequest) GetLines() []ChannelOrderLineRequest`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *ChannelOrderRequest) GetLinesOk() (*[]ChannelOrderLineRequest, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *ChannelOrderRequest) SetLines(v []ChannelOrderLineRequest)`

SetLines sets Lines field to given value.


### GetShippingCostsInclVat

`func (o *ChannelOrderRequest) GetShippingCostsInclVat() float32`

GetShippingCostsInclVat returns the ShippingCostsInclVat field if non-nil, zero value otherwise.

### GetShippingCostsInclVatOk

`func (o *ChannelOrderRequest) GetShippingCostsInclVatOk() (*float32, bool)`

GetShippingCostsInclVatOk returns a tuple with the ShippingCostsInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCostsInclVat

`func (o *ChannelOrderRequest) SetShippingCostsInclVat(v float32)`

SetShippingCostsInclVat sets ShippingCostsInclVat field to given value.


### GetOrderFee

`func (o *ChannelOrderRequest) GetOrderFee() float32`

GetOrderFee returns the OrderFee field if non-nil, zero value otherwise.

### GetOrderFeeOk

`func (o *ChannelOrderRequest) GetOrderFeeOk() (*float32, bool)`

GetOrderFeeOk returns a tuple with the OrderFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderFee

`func (o *ChannelOrderRequest) SetOrderFee(v float32)`

SetOrderFee sets OrderFee field to given value.

### HasOrderFee

`func (o *ChannelOrderRequest) HasOrderFee() bool`

HasOrderFee returns a boolean if a field has been set.

### GetPhone

`func (o *ChannelOrderRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ChannelOrderRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ChannelOrderRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ChannelOrderRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *ChannelOrderRequest) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *ChannelOrderRequest) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *ChannelOrderRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ChannelOrderRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ChannelOrderRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLanguageCode

`func (o *ChannelOrderRequest) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *ChannelOrderRequest) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *ChannelOrderRequest) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *ChannelOrderRequest) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### SetLanguageCodeNil

`func (o *ChannelOrderRequest) SetLanguageCodeNil(b bool)`

 SetLanguageCodeNil sets the value for LanguageCode to be an explicit nil

### UnsetLanguageCode
`func (o *ChannelOrderRequest) UnsetLanguageCode()`

UnsetLanguageCode ensures that no value is present for LanguageCode, not even an explicit nil
### GetCompanyRegistrationNo

`func (o *ChannelOrderRequest) GetCompanyRegistrationNo() string`

GetCompanyRegistrationNo returns the CompanyRegistrationNo field if non-nil, zero value otherwise.

### GetCompanyRegistrationNoOk

`func (o *ChannelOrderRequest) GetCompanyRegistrationNoOk() (*string, bool)`

GetCompanyRegistrationNoOk returns a tuple with the CompanyRegistrationNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyRegistrationNo

`func (o *ChannelOrderRequest) SetCompanyRegistrationNo(v string)`

SetCompanyRegistrationNo sets CompanyRegistrationNo field to given value.

### HasCompanyRegistrationNo

`func (o *ChannelOrderRequest) HasCompanyRegistrationNo() bool`

HasCompanyRegistrationNo returns a boolean if a field has been set.

### SetCompanyRegistrationNoNil

`func (o *ChannelOrderRequest) SetCompanyRegistrationNoNil(b bool)`

 SetCompanyRegistrationNoNil sets the value for CompanyRegistrationNo to be an explicit nil

### UnsetCompanyRegistrationNo
`func (o *ChannelOrderRequest) UnsetCompanyRegistrationNo()`

UnsetCompanyRegistrationNo ensures that no value is present for CompanyRegistrationNo, not even an explicit nil
### GetVatNo

`func (o *ChannelOrderRequest) GetVatNo() string`

GetVatNo returns the VatNo field if non-nil, zero value otherwise.

### GetVatNoOk

`func (o *ChannelOrderRequest) GetVatNoOk() (*string, bool)`

GetVatNoOk returns a tuple with the VatNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNo

`func (o *ChannelOrderRequest) SetVatNo(v string)`

SetVatNo sets VatNo field to given value.

### HasVatNo

`func (o *ChannelOrderRequest) HasVatNo() bool`

HasVatNo returns a boolean if a field has been set.

### SetVatNoNil

`func (o *ChannelOrderRequest) SetVatNoNil(b bool)`

 SetVatNoNil sets the value for VatNo to be an explicit nil

### UnsetVatNo
`func (o *ChannelOrderRequest) UnsetVatNo()`

UnsetVatNo ensures that no value is present for VatNo, not even an explicit nil
### GetPaymentMethod

`func (o *ChannelOrderRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *ChannelOrderRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *ChannelOrderRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *ChannelOrderRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### SetPaymentMethodNil

`func (o *ChannelOrderRequest) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *ChannelOrderRequest) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
### GetPaymentReferenceNo

`func (o *ChannelOrderRequest) GetPaymentReferenceNo() string`

GetPaymentReferenceNo returns the PaymentReferenceNo field if non-nil, zero value otherwise.

### GetPaymentReferenceNoOk

`func (o *ChannelOrderRequest) GetPaymentReferenceNoOk() (*string, bool)`

GetPaymentReferenceNoOk returns a tuple with the PaymentReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentReferenceNo

`func (o *ChannelOrderRequest) SetPaymentReferenceNo(v string)`

SetPaymentReferenceNo sets PaymentReferenceNo field to given value.

### HasPaymentReferenceNo

`func (o *ChannelOrderRequest) HasPaymentReferenceNo() bool`

HasPaymentReferenceNo returns a boolean if a field has been set.

### SetPaymentReferenceNoNil

`func (o *ChannelOrderRequest) SetPaymentReferenceNoNil(b bool)`

 SetPaymentReferenceNoNil sets the value for PaymentReferenceNo to be an explicit nil

### UnsetPaymentReferenceNo
`func (o *ChannelOrderRequest) UnsetPaymentReferenceNo()`

UnsetPaymentReferenceNo ensures that no value is present for PaymentReferenceNo, not even an explicit nil
### GetCurrencyCode

`func (o *ChannelOrderRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ChannelOrderRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ChannelOrderRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetOrderDate

`func (o *ChannelOrderRequest) GetOrderDate() time.Time`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *ChannelOrderRequest) GetOrderDateOk() (*time.Time, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *ChannelOrderRequest) SetOrderDate(v time.Time)`

SetOrderDate sets OrderDate field to given value.


### GetChannelCustomerNo

`func (o *ChannelOrderRequest) GetChannelCustomerNo() string`

GetChannelCustomerNo returns the ChannelCustomerNo field if non-nil, zero value otherwise.

### GetChannelCustomerNoOk

`func (o *ChannelOrderRequest) GetChannelCustomerNoOk() (*string, bool)`

GetChannelCustomerNoOk returns a tuple with the ChannelCustomerNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCustomerNo

`func (o *ChannelOrderRequest) SetChannelCustomerNo(v string)`

SetChannelCustomerNo sets ChannelCustomerNo field to given value.

### HasChannelCustomerNo

`func (o *ChannelOrderRequest) HasChannelCustomerNo() bool`

HasChannelCustomerNo returns a boolean if a field has been set.

### SetChannelCustomerNoNil

`func (o *ChannelOrderRequest) SetChannelCustomerNoNil(b bool)`

 SetChannelCustomerNoNil sets the value for ChannelCustomerNo to be an explicit nil

### UnsetChannelCustomerNo
`func (o *ChannelOrderRequest) UnsetChannelCustomerNo()`

UnsetChannelCustomerNo ensures that no value is present for ChannelCustomerNo, not even an explicit nil
### GetExtraData

`func (o *ChannelOrderRequest) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ChannelOrderRequest) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ChannelOrderRequest) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *ChannelOrderRequest) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *ChannelOrderRequest) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *ChannelOrderRequest) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


