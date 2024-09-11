# DateRangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromDate** | Pointer to **NullableTime** |  | [optional] 
**ToDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewDateRangeRequest

`func NewDateRangeRequest() *DateRangeRequest`

NewDateRangeRequest instantiates a new DateRangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateRangeRequestWithDefaults

`func NewDateRangeRequestWithDefaults() *DateRangeRequest`

NewDateRangeRequestWithDefaults instantiates a new DateRangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromDate

`func (o *DateRangeRequest) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *DateRangeRequest) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *DateRangeRequest) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *DateRangeRequest) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### SetFromDateNil

`func (o *DateRangeRequest) SetFromDateNil(b bool)`

 SetFromDateNil sets the value for FromDate to be an explicit nil

### UnsetFromDate
`func (o *DateRangeRequest) UnsetFromDate()`

UnsetFromDate ensures that no value is present for FromDate, not even an explicit nil
### GetToDate

`func (o *DateRangeRequest) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *DateRangeRequest) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *DateRangeRequest) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *DateRangeRequest) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### SetToDateNil

`func (o *DateRangeRequest) SetToDateNil(b bool)`

 SetToDateNil sets the value for ToDate to be an explicit nil

### UnsetToDate
`func (o *DateRangeRequest) UnsetToDate()`

UnsetToDate ensures that no value is present for ToDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


