# PartnerProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logo** | Pointer to [**Logo**](Logo.md) |  | [optional] 
**Contacts** | Pointer to [**[]Contact**](Contact.md) |  | [optional] 
**Identifiers** | [**[]Identifier**](Identifier.md) |  | 
**Addresses** | Pointer to [**[]Address**](Address.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | [**UserDetail**](UserDetail.md) |  | 
**UpdatedAt** | **time.Time** |  | 
**UpdatedBy** | [**UserDetail**](UserDetail.md) |  | 
**Protocols** | **[]string** |  | 
**Standars** | **[]string** |  | 
**UsedInDeployments** | [**UsedInDeployments**](UsedInDeployments.md) |  | 

## Methods

### NewPartnerProfileAllOf

`func NewPartnerProfileAllOf(identifiers []Identifier, createdAt time.Time, createdBy UserDetail, updatedAt time.Time, updatedBy UserDetail, protocols []string, standars []string, usedInDeployments UsedInDeployments, ) *PartnerProfileAllOf`

NewPartnerProfileAllOf instantiates a new PartnerProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerProfileAllOfWithDefaults

`func NewPartnerProfileAllOfWithDefaults() *PartnerProfileAllOf`

NewPartnerProfileAllOfWithDefaults instantiates a new PartnerProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogo

`func (o *PartnerProfileAllOf) GetLogo() Logo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *PartnerProfileAllOf) GetLogoOk() (*Logo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *PartnerProfileAllOf) SetLogo(v Logo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *PartnerProfileAllOf) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetContacts

`func (o *PartnerProfileAllOf) GetContacts() []Contact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *PartnerProfileAllOf) GetContactsOk() (*[]Contact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *PartnerProfileAllOf) SetContacts(v []Contact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *PartnerProfileAllOf) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetIdentifiers

`func (o *PartnerProfileAllOf) GetIdentifiers() []Identifier`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *PartnerProfileAllOf) GetIdentifiersOk() (*[]Identifier, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *PartnerProfileAllOf) SetIdentifiers(v []Identifier)`

SetIdentifiers sets Identifiers field to given value.


### GetAddresses

`func (o *PartnerProfileAllOf) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PartnerProfileAllOf) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PartnerProfileAllOf) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *PartnerProfileAllOf) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PartnerProfileAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PartnerProfileAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PartnerProfileAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *PartnerProfileAllOf) GetCreatedBy() UserDetail`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PartnerProfileAllOf) GetCreatedByOk() (*UserDetail, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PartnerProfileAllOf) SetCreatedBy(v UserDetail)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedAt

`func (o *PartnerProfileAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PartnerProfileAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PartnerProfileAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedBy

`func (o *PartnerProfileAllOf) GetUpdatedBy() UserDetail`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *PartnerProfileAllOf) GetUpdatedByOk() (*UserDetail, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *PartnerProfileAllOf) SetUpdatedBy(v UserDetail)`

SetUpdatedBy sets UpdatedBy field to given value.


### GetProtocols

`func (o *PartnerProfileAllOf) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *PartnerProfileAllOf) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *PartnerProfileAllOf) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.


### GetStandars

`func (o *PartnerProfileAllOf) GetStandars() []string`

GetStandars returns the Standars field if non-nil, zero value otherwise.

### GetStandarsOk

`func (o *PartnerProfileAllOf) GetStandarsOk() (*[]string, bool)`

GetStandarsOk returns a tuple with the Standars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandars

`func (o *PartnerProfileAllOf) SetStandars(v []string)`

SetStandars sets Standars field to given value.


### GetUsedInDeployments

`func (o *PartnerProfileAllOf) GetUsedInDeployments() UsedInDeployments`

GetUsedInDeployments returns the UsedInDeployments field if non-nil, zero value otherwise.

### GetUsedInDeploymentsOk

`func (o *PartnerProfileAllOf) GetUsedInDeploymentsOk() (*UsedInDeployments, bool)`

GetUsedInDeploymentsOk returns a tuple with the UsedInDeployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedInDeployments

`func (o *PartnerProfileAllOf) SetUsedInDeployments(v UsedInDeployments)`

SetUsedInDeployments sets UsedInDeployments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


