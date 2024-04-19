# Identifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**IdentifierTypeQualifierId** | **string** |  | 
**Status** | **string** |  | 
**QualifierLabel** | Pointer to **string** |  | [optional] 
**TypeLabel** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Value** | **string** |  | 

## Methods

### NewIdentifier

`func NewIdentifier(identifierTypeQualifierId string, status string, value string, ) *Identifier`

NewIdentifier instantiates a new Identifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentifierWithDefaults

`func NewIdentifierWithDefaults() *Identifier`

NewIdentifierWithDefaults instantiates a new Identifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Identifier) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Identifier) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Identifier) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Identifier) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifierTypeQualifierId

`func (o *Identifier) GetIdentifierTypeQualifierId() string`

GetIdentifierTypeQualifierId returns the IdentifierTypeQualifierId field if non-nil, zero value otherwise.

### GetIdentifierTypeQualifierIdOk

`func (o *Identifier) GetIdentifierTypeQualifierIdOk() (*string, bool)`

GetIdentifierTypeQualifierIdOk returns a tuple with the IdentifierTypeQualifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierTypeQualifierId

`func (o *Identifier) SetIdentifierTypeQualifierId(v string)`

SetIdentifierTypeQualifierId sets IdentifierTypeQualifierId field to given value.


### GetStatus

`func (o *Identifier) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Identifier) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Identifier) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetQualifierLabel

`func (o *Identifier) GetQualifierLabel() string`

GetQualifierLabel returns the QualifierLabel field if non-nil, zero value otherwise.

### GetQualifierLabelOk

`func (o *Identifier) GetQualifierLabelOk() (*string, bool)`

GetQualifierLabelOk returns a tuple with the QualifierLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifierLabel

`func (o *Identifier) SetQualifierLabel(v string)`

SetQualifierLabel sets QualifierLabel field to given value.

### HasQualifierLabel

`func (o *Identifier) HasQualifierLabel() bool`

HasQualifierLabel returns a boolean if a field has been set.

### GetTypeLabel

`func (o *Identifier) GetTypeLabel() string`

GetTypeLabel returns the TypeLabel field if non-nil, zero value otherwise.

### GetTypeLabelOk

`func (o *Identifier) GetTypeLabelOk() (*string, bool)`

GetTypeLabelOk returns a tuple with the TypeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeLabel

`func (o *Identifier) SetTypeLabel(v string)`

SetTypeLabel sets TypeLabel field to given value.

### HasTypeLabel

`func (o *Identifier) HasTypeLabel() bool`

HasTypeLabel returns a boolean if a field has been set.

### GetCode

`func (o *Identifier) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Identifier) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Identifier) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Identifier) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetValue

`func (o *Identifier) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Identifier) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Identifier) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


