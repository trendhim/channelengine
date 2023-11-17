# SettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipment** | Pointer to [**ShipmentSettingsResponse**](ShipmentSettingsResponse.md) |  | [optional] 
**Advanced** | Pointer to [**AdvanceSettingsResponse**](AdvanceSettingsResponse.md) |  | [optional] 

## Methods

### NewSettingsResponse

`func NewSettingsResponse() *SettingsResponse`

NewSettingsResponse instantiates a new SettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsResponseWithDefaults

`func NewSettingsResponseWithDefaults() *SettingsResponse`

NewSettingsResponseWithDefaults instantiates a new SettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipment

`func (o *SettingsResponse) GetShipment() ShipmentSettingsResponse`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *SettingsResponse) GetShipmentOk() (*ShipmentSettingsResponse, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *SettingsResponse) SetShipment(v ShipmentSettingsResponse)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *SettingsResponse) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetAdvanced

`func (o *SettingsResponse) GetAdvanced() AdvanceSettingsResponse`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *SettingsResponse) GetAdvancedOk() (*AdvanceSettingsResponse, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *SettingsResponse) SetAdvanced(v AdvanceSettingsResponse)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *SettingsResponse) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


