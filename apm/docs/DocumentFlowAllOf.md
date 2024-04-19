# DocumentFlowAllOf

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

### NewDocumentFlowAllOf

`func NewDocumentFlowAllOf() *DocumentFlowAllOf`

NewDocumentFlowAllOf instantiates a new DocumentFlowAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentFlowAllOfWithDefaults

`func NewDocumentFlowAllOfWithDefaults() *DocumentFlowAllOf`

NewDocumentFlowAllOfWithDefaults instantiates a new DocumentFlowAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentFlowAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentFlowAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentFlowAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentFlowAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DocumentFlowAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentFlowAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentFlowAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentFlowAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDirection

`func (o *DocumentFlowAllOf) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *DocumentFlowAllOf) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *DocumentFlowAllOf) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *DocumentFlowAllOf) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetPartnerFromId

`func (o *DocumentFlowAllOf) GetPartnerFromId() string`

GetPartnerFromId returns the PartnerFromId field if non-nil, zero value otherwise.

### GetPartnerFromIdOk

`func (o *DocumentFlowAllOf) GetPartnerFromIdOk() (*string, bool)`

GetPartnerFromIdOk returns a tuple with the PartnerFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerFromId

`func (o *DocumentFlowAllOf) SetPartnerFromId(v string)`

SetPartnerFromId sets PartnerFromId field to given value.

### HasPartnerFromId

`func (o *DocumentFlowAllOf) HasPartnerFromId() bool`

HasPartnerFromId returns a boolean if a field has been set.

### GetPartnerToId

`func (o *DocumentFlowAllOf) GetPartnerToId() string`

GetPartnerToId returns the PartnerToId field if non-nil, zero value otherwise.

### GetPartnerToIdOk

`func (o *DocumentFlowAllOf) GetPartnerToIdOk() (*string, bool)`

GetPartnerToIdOk returns a tuple with the PartnerToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerToId

`func (o *DocumentFlowAllOf) SetPartnerToId(v string)`

SetPartnerToId sets PartnerToId field to given value.

### HasPartnerToId

`func (o *DocumentFlowAllOf) HasPartnerToId() bool`

HasPartnerToId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocumentFlowAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentFlowAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentFlowAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocumentFlowAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DocumentFlowAllOf) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DocumentFlowAllOf) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DocumentFlowAllOf) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DocumentFlowAllOf) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifierAt

`func (o *DocumentFlowAllOf) GetModifierAt() time.Time`

GetModifierAt returns the ModifierAt field if non-nil, zero value otherwise.

### GetModifierAtOk

`func (o *DocumentFlowAllOf) GetModifierAtOk() (*time.Time, bool)`

GetModifierAtOk returns a tuple with the ModifierAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierAt

`func (o *DocumentFlowAllOf) SetModifierAt(v time.Time)`

SetModifierAt sets ModifierAt field to given value.

### HasModifierAt

`func (o *DocumentFlowAllOf) HasModifierAt() bool`

HasModifierAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *DocumentFlowAllOf) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *DocumentFlowAllOf) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *DocumentFlowAllOf) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *DocumentFlowAllOf) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetConfigurations

`func (o *DocumentFlowAllOf) GetConfigurations() []DocumentFlowConfig`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *DocumentFlowAllOf) GetConfigurationsOk() (*[]DocumentFlowConfig, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *DocumentFlowAllOf) SetConfigurations(v []DocumentFlowConfig)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *DocumentFlowAllOf) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


