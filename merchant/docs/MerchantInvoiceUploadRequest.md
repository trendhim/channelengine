# MerchantInvoiceUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceContent** | **string** | Data needed to upload an invoice. | 
**InvoiceNumber** | **string** | The invoice number used in the invoice. | 

## Methods

### NewMerchantInvoiceUploadRequest

`func NewMerchantInvoiceUploadRequest(invoiceContent string, invoiceNumber string, ) *MerchantInvoiceUploadRequest`

NewMerchantInvoiceUploadRequest instantiates a new MerchantInvoiceUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantInvoiceUploadRequestWithDefaults

`func NewMerchantInvoiceUploadRequestWithDefaults() *MerchantInvoiceUploadRequest`

NewMerchantInvoiceUploadRequestWithDefaults instantiates a new MerchantInvoiceUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceContent

`func (o *MerchantInvoiceUploadRequest) GetInvoiceContent() string`

GetInvoiceContent returns the InvoiceContent field if non-nil, zero value otherwise.

### GetInvoiceContentOk

`func (o *MerchantInvoiceUploadRequest) GetInvoiceContentOk() (*string, bool)`

GetInvoiceContentOk returns a tuple with the InvoiceContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceContent

`func (o *MerchantInvoiceUploadRequest) SetInvoiceContent(v string)`

SetInvoiceContent sets InvoiceContent field to given value.


### GetInvoiceNumber

`func (o *MerchantInvoiceUploadRequest) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *MerchantInvoiceUploadRequest) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *MerchantInvoiceUploadRequest) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


