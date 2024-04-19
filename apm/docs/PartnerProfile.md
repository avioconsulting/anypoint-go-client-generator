# Partnerprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerType** | **string** |  | 
**EnvironmentId** | **string** |  | 
**Name** | **string** |  | 
**Identifiers** | [**[]Partnerprofileidentifier**](Partnerprofileidentifier.md) |  | 

## Methods

### NewPartnerprofile

`func NewPartnerprofile(partnerType string, environmentId string, name string, identifiers []Partnerprofileidentifier, ) *Partnerprofile`

NewPartnerprofile instantiates a new Partnerprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerprofileWithDefaults

`func NewPartnerprofileWithDefaults() *Partnerprofile`

NewPartnerprofileWithDefaults instantiates a new Partnerprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerType

`func (o *Partnerprofile) GetPartnerType() string`

GetPartnerType returns the PartnerType field if non-nil, zero value otherwise.

### GetPartnerTypeOk

`func (o *Partnerprofile) GetPartnerTypeOk() (*string, bool)`

GetPartnerTypeOk returns a tuple with the PartnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerType

`func (o *Partnerprofile) SetPartnerType(v string)`

SetPartnerType sets PartnerType field to given value.


### GetEnvironmentId

`func (o *Partnerprofile) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *Partnerprofile) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *Partnerprofile) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetName

`func (o *Partnerprofile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Partnerprofile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Partnerprofile) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifiers

`func (o *Partnerprofile) GetIdentifiers() []Partnerprofileidentifier`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *Partnerprofile) GetIdentifiersOk() (*[]Partnerprofileidentifier, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *Partnerprofile) SetIdentifiers(v []Partnerprofileidentifier)`

SetIdentifiers sets Identifiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


