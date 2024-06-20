# MerchantPurchaseOrderInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceNo** | Pointer to **NullableString** |  | [optional] 
**InvoiceType** | Pointer to [**ModulesPurchaseOrderInvoiceType**](ModulesPurchaseOrderInvoiceType.md) |  | [optional] 
**InvoiceTotalAmount** | Pointer to **float32** |  | [optional] 
**InvoiceTotalCurrencyCode** | Pointer to **NullableString** |  | [optional] 
**RemitToParty** | Pointer to [**MerchantVendorParty**](MerchantVendorParty.md) |  | [optional] 
**ShipToPartyId** | Pointer to **NullableInt32** |  | [optional] 
**BillToPartyId** | Pointer to **NullableInt32** |  | [optional] 
**AdditionalDetails** | Pointer to [**[]PurchaseOrderInvoiceAdditionalDetails**](PurchaseOrderInvoiceAdditionalDetails.md) |  | [optional] 
**Lines** | Pointer to [**[]MerchantPurchaseOrderInvoiceLine**](MerchantPurchaseOrderInvoiceLine.md) |  | [optional] 

## Methods

### NewMerchantPurchaseOrderInvoice

`func NewMerchantPurchaseOrderInvoice() *MerchantPurchaseOrderInvoice`

NewMerchantPurchaseOrderInvoice instantiates a new MerchantPurchaseOrderInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantPurchaseOrderInvoiceWithDefaults

`func NewMerchantPurchaseOrderInvoiceWithDefaults() *MerchantPurchaseOrderInvoice`

NewMerchantPurchaseOrderInvoiceWithDefaults instantiates a new MerchantPurchaseOrderInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceNo

`func (o *MerchantPurchaseOrderInvoice) GetInvoiceNo() string`

GetInvoiceNo returns the InvoiceNo field if non-nil, zero value otherwise.

### GetInvoiceNoOk

`func (o *MerchantPurchaseOrderInvoice) GetInvoiceNoOk() (*string, bool)`

GetInvoiceNoOk returns a tuple with the InvoiceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNo

`func (o *MerchantPurchaseOrderInvoice) SetInvoiceNo(v string)`

SetInvoiceNo sets InvoiceNo field to given value.

### HasInvoiceNo

`func (o *MerchantPurchaseOrderInvoice) HasInvoiceNo() bool`

HasInvoiceNo returns a boolean if a field has been set.

### SetInvoiceNoNil

`func (o *MerchantPurchaseOrderInvoice) SetInvoiceNoNil(b bool)`

 SetInvoiceNoNil sets the value for InvoiceNo to be an explicit nil

### UnsetInvoiceNo
`func (o *MerchantPurchaseOrderInvoice) UnsetInvoiceNo()`

UnsetInvoiceNo ensures that no value is present for InvoiceNo, not even an explicit nil
### GetInvoiceType

`func (o *MerchantPurchaseOrderInvoice) GetInvoiceType() ModulesPurchaseOrderInvoiceType`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *MerchantPurchaseOrderInvoice) GetInvoiceTypeOk() (*ModulesPurchaseOrderInvoiceType, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *MerchantPurchaseOrderInvoice) SetInvoiceType(v ModulesPurchaseOrderInvoiceType)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *MerchantPurchaseOrderInvoice) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetInvoiceTotalAmount

`func (o *MerchantPurchaseOrderInvoice) GetInvoiceTotalAmount() float32`

GetInvoiceTotalAmount returns the InvoiceTotalAmount field if non-nil, zero value otherwise.

### GetInvoiceTotalAmountOk

`func (o *MerchantPurchaseOrderInvoice) GetInvoiceTotalAmountOk() (*float32, bool)`

GetInvoiceTotalAmountOk returns a tuple with the InvoiceTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTotalAmount

`func (o *MerchantPurchaseOrderInvoice) SetInvoiceTotalAmount(v float32)`

SetInvoiceTotalAmount sets InvoiceTotalAmount field to given value.

### HasInvoiceTotalAmount

`func (o *MerchantPurchaseOrderInvoice) HasInvoiceTotalAmount() bool`

HasInvoiceTotalAmount returns a boolean if a field has been set.

### GetInvoiceTotalCurrencyCode

`func (o *MerchantPurchaseOrderInvoice) GetInvoiceTotalCurrencyCode() string`

GetInvoiceTotalCurrencyCode returns the InvoiceTotalCurrencyCode field if non-nil, zero value otherwise.

### GetInvoiceTotalCurrencyCodeOk

`func (o *MerchantPurchaseOrderInvoice) GetInvoiceTotalCurrencyCodeOk() (*string, bool)`

GetInvoiceTotalCurrencyCodeOk returns a tuple with the InvoiceTotalCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTotalCurrencyCode

`func (o *MerchantPurchaseOrderInvoice) SetInvoiceTotalCurrencyCode(v string)`

SetInvoiceTotalCurrencyCode sets InvoiceTotalCurrencyCode field to given value.

### HasInvoiceTotalCurrencyCode

`func (o *MerchantPurchaseOrderInvoice) HasInvoiceTotalCurrencyCode() bool`

HasInvoiceTotalCurrencyCode returns a boolean if a field has been set.

### SetInvoiceTotalCurrencyCodeNil

`func (o *MerchantPurchaseOrderInvoice) SetInvoiceTotalCurrencyCodeNil(b bool)`

 SetInvoiceTotalCurrencyCodeNil sets the value for InvoiceTotalCurrencyCode to be an explicit nil

### UnsetInvoiceTotalCurrencyCode
`func (o *MerchantPurchaseOrderInvoice) UnsetInvoiceTotalCurrencyCode()`

UnsetInvoiceTotalCurrencyCode ensures that no value is present for InvoiceTotalCurrencyCode, not even an explicit nil
### GetRemitToParty

`func (o *MerchantPurchaseOrderInvoice) GetRemitToParty() MerchantVendorParty`

GetRemitToParty returns the RemitToParty field if non-nil, zero value otherwise.

### GetRemitToPartyOk

`func (o *MerchantPurchaseOrderInvoice) GetRemitToPartyOk() (*MerchantVendorParty, bool)`

GetRemitToPartyOk returns a tuple with the RemitToParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemitToParty

`func (o *MerchantPurchaseOrderInvoice) SetRemitToParty(v MerchantVendorParty)`

SetRemitToParty sets RemitToParty field to given value.

### HasRemitToParty

`func (o *MerchantPurchaseOrderInvoice) HasRemitToParty() bool`

HasRemitToParty returns a boolean if a field has been set.

### GetShipToPartyId

`func (o *MerchantPurchaseOrderInvoice) GetShipToPartyId() int32`

GetShipToPartyId returns the ShipToPartyId field if non-nil, zero value otherwise.

### GetShipToPartyIdOk

`func (o *MerchantPurchaseOrderInvoice) GetShipToPartyIdOk() (*int32, bool)`

GetShipToPartyIdOk returns a tuple with the ShipToPartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToPartyId

`func (o *MerchantPurchaseOrderInvoice) SetShipToPartyId(v int32)`

SetShipToPartyId sets ShipToPartyId field to given value.

### HasShipToPartyId

`func (o *MerchantPurchaseOrderInvoice) HasShipToPartyId() bool`

HasShipToPartyId returns a boolean if a field has been set.

### SetShipToPartyIdNil

`func (o *MerchantPurchaseOrderInvoice) SetShipToPartyIdNil(b bool)`

 SetShipToPartyIdNil sets the value for ShipToPartyId to be an explicit nil

### UnsetShipToPartyId
`func (o *MerchantPurchaseOrderInvoice) UnsetShipToPartyId()`

UnsetShipToPartyId ensures that no value is present for ShipToPartyId, not even an explicit nil
### GetBillToPartyId

`func (o *MerchantPurchaseOrderInvoice) GetBillToPartyId() int32`

GetBillToPartyId returns the BillToPartyId field if non-nil, zero value otherwise.

### GetBillToPartyIdOk

`func (o *MerchantPurchaseOrderInvoice) GetBillToPartyIdOk() (*int32, bool)`

GetBillToPartyIdOk returns a tuple with the BillToPartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToPartyId

`func (o *MerchantPurchaseOrderInvoice) SetBillToPartyId(v int32)`

SetBillToPartyId sets BillToPartyId field to given value.

### HasBillToPartyId

`func (o *MerchantPurchaseOrderInvoice) HasBillToPartyId() bool`

HasBillToPartyId returns a boolean if a field has been set.

### SetBillToPartyIdNil

`func (o *MerchantPurchaseOrderInvoice) SetBillToPartyIdNil(b bool)`

 SetBillToPartyIdNil sets the value for BillToPartyId to be an explicit nil

### UnsetBillToPartyId
`func (o *MerchantPurchaseOrderInvoice) UnsetBillToPartyId()`

UnsetBillToPartyId ensures that no value is present for BillToPartyId, not even an explicit nil
### GetAdditionalDetails

`func (o *MerchantPurchaseOrderInvoice) GetAdditionalDetails() []PurchaseOrderInvoiceAdditionalDetails`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *MerchantPurchaseOrderInvoice) GetAdditionalDetailsOk() (*[]PurchaseOrderInvoiceAdditionalDetails, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *MerchantPurchaseOrderInvoice) SetAdditionalDetails(v []PurchaseOrderInvoiceAdditionalDetails)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *MerchantPurchaseOrderInvoice) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### SetAdditionalDetailsNil

`func (o *MerchantPurchaseOrderInvoice) SetAdditionalDetailsNil(b bool)`

 SetAdditionalDetailsNil sets the value for AdditionalDetails to be an explicit nil

### UnsetAdditionalDetails
`func (o *MerchantPurchaseOrderInvoice) UnsetAdditionalDetails()`

UnsetAdditionalDetails ensures that no value is present for AdditionalDetails, not even an explicit nil
### GetLines

`func (o *MerchantPurchaseOrderInvoice) GetLines() []MerchantPurchaseOrderInvoiceLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantPurchaseOrderInvoice) GetLinesOk() (*[]MerchantPurchaseOrderInvoiceLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantPurchaseOrderInvoice) SetLines(v []MerchantPurchaseOrderInvoiceLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *MerchantPurchaseOrderInvoice) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *MerchantPurchaseOrderInvoice) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *MerchantPurchaseOrderInvoice) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


