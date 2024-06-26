# IdentifierTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Label** | **string** |  | 
**Description** | **string** |  | 
**Readonly** | Pointer to **bool** |  | [optional] 
**Qualifiers** | [**[]IdentifierTypeQualifier**](IdentifierTypeQualifier.md) |  | 
**EnvironmentId** | Pointer to **string** |  | [optional] 

## Methods

### NewIdentifierTypeAllOf

`func NewIdentifierTypeAllOf(id string, name string, label string, description string, qualifiers []IdentifierTypeQualifier, ) *IdentifierTypeAllOf`

NewIdentifierTypeAllOf instantiates a new IdentifierTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentifierTypeAllOfWithDefaults

`func NewIdentifierTypeAllOfWithDefaults() *IdentifierTypeAllOf`

NewIdentifierTypeAllOfWithDefaults instantiates a new IdentifierTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentifierTypeAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentifierTypeAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentifierTypeAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *IdentifierTypeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentifierTypeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentifierTypeAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *IdentifierTypeAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IdentifierTypeAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IdentifierTypeAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *IdentifierTypeAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentifierTypeAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentifierTypeAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetReadonly

`func (o *IdentifierTypeAllOf) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *IdentifierTypeAllOf) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *IdentifierTypeAllOf) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *IdentifierTypeAllOf) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetQualifiers

`func (o *IdentifierTypeAllOf) GetQualifiers() []IdentifierTypeQualifier`

GetQualifiers returns the Qualifiers field if non-nil, zero value otherwise.

### GetQualifiersOk

`func (o *IdentifierTypeAllOf) GetQualifiersOk() (*[]IdentifierTypeQualifier, bool)`

GetQualifiersOk returns a tuple with the Qualifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiers

`func (o *IdentifierTypeAllOf) SetQualifiers(v []IdentifierTypeQualifier)`

SetQualifiers sets Qualifiers field to given value.


### GetEnvironmentId

`func (o *IdentifierTypeAllOf) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *IdentifierTypeAllOf) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *IdentifierTypeAllOf) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *IdentifierTypeAllOf) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


