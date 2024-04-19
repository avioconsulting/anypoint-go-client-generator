# CustomAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Searchable** | **bool** |  | 
**Label** | **string** |  | 
**Alias** | **string** |  | 

## Methods

### NewCustomAttribute

`func NewCustomAttribute(searchable bool, label string, alias string, ) *CustomAttribute`

NewCustomAttribute instantiates a new CustomAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeWithDefaults

`func NewCustomAttributeWithDefaults() *CustomAttribute`

NewCustomAttributeWithDefaults instantiates a new CustomAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomAttribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSearchable

`func (o *CustomAttribute) GetSearchable() bool`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *CustomAttribute) GetSearchableOk() (*bool, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *CustomAttribute) SetSearchable(v bool)`

SetSearchable sets Searchable field to given value.


### GetLabel

`func (o *CustomAttribute) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomAttribute) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomAttribute) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetAlias

`func (o *CustomAttribute) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CustomAttribute) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CustomAttribute) SetAlias(v string)`

SetAlias sets Alias field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


