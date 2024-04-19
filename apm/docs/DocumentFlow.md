# DocumentFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**PartnerFromId** | Pointer to **string** |  | [optional] 
**PartnerToId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**ModifierAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**Configurations** | Pointer to [**[]DocumentFlowConfig**](DocumentFlowConfig.md) |  | [optional] 

## Methods

### NewDocumentFlow

`func NewDocumentFlow() *DocumentFlow`

NewDocumentFlow instantiates a new DocumentFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentFlowWithDefaults

`func NewDocumentFlowWithDefaults() *DocumentFlow`

NewDocumentFlowWithDefaults instantiates a new DocumentFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentFlow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentFlow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DocumentFlow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentFlow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentFlow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentFlow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDirection

`func (o *DocumentFlow) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *DocumentFlow) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *DocumentFlow) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *DocumentFlow) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetPartnerFromId

`func (o *DocumentFlow) GetPartnerFromId() string`

GetPartnerFromId returns the PartnerFromId field if non-nil, zero value otherwise.

### GetPartnerFromIdOk

`func (o *DocumentFlow) GetPartnerFromIdOk() (*string, bool)`

GetPartnerFromIdOk returns a tuple with the PartnerFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerFromId

`func (o *DocumentFlow) SetPartnerFromId(v string)`

SetPartnerFromId sets PartnerFromId field to given value.

### HasPartnerFromId

`func (o *DocumentFlow) HasPartnerFromId() bool`

HasPartnerFromId returns a boolean if a field has been set.

### GetPartnerToId

`func (o *DocumentFlow) GetPartnerToId() string`

GetPartnerToId returns the PartnerToId field if non-nil, zero value otherwise.

### GetPartnerToIdOk

`func (o *DocumentFlow) GetPartnerToIdOk() (*string, bool)`

GetPartnerToIdOk returns a tuple with the PartnerToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerToId

`func (o *DocumentFlow) SetPartnerToId(v string)`

SetPartnerToId sets PartnerToId field to given value.

### HasPartnerToId

`func (o *DocumentFlow) HasPartnerToId() bool`

HasPartnerToId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocumentFlow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentFlow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentFlow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocumentFlow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DocumentFlow) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DocumentFlow) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DocumentFlow) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DocumentFlow) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifierAt

`func (o *DocumentFlow) GetModifierAt() time.Time`

GetModifierAt returns the ModifierAt field if non-nil, zero value otherwise.

### GetModifierAtOk

`func (o *DocumentFlow) GetModifierAtOk() (*time.Time, bool)`

GetModifierAtOk returns a tuple with the ModifierAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierAt

`func (o *DocumentFlow) SetModifierAt(v time.Time)`

SetModifierAt sets ModifierAt field to given value.

### HasModifierAt

`func (o *DocumentFlow) HasModifierAt() bool`

HasModifierAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *DocumentFlow) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *DocumentFlow) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *DocumentFlow) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *DocumentFlow) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetConfigurations

`func (o *DocumentFlow) GetConfigurations() []DocumentFlowConfig`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *DocumentFlow) GetConfigurationsOk() (*[]DocumentFlowConfig, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *DocumentFlow) SetConfigurations(v []DocumentFlowConfig)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *DocumentFlow) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


