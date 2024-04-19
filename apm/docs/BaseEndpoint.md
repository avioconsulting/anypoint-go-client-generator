# BaseEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Description** | **string** |  | 
**EnvironmentId** | **string** |  | 
**EndpointTypeId** | **string** |  | 
**Partner** | Pointer to [**Partner**](Partner.md) |  | [optional] 
**PartnerId** | Pointer to **string** |  | [optional] 
**Deployment** | Pointer to [**EndpointDeployment**](EndpointDeployment.md) |  | [optional] 
**EndpointRole** | [**EndpointRoleType**](EndpointRoleType.md) |  | 
**EndpointType** | Pointer to **string** |  | [optional] 
**Config** | [**BaseEndpointConfig**](BaseEndpointConfig.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**UserDetail**](UserDetail.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to [**UserDetail**](UserDetail.md) |  | [optional] 

## Methods

### NewBaseEndpoint

`func NewBaseEndpoint(name string, description string, environmentId string, endpointTypeId string, endpointRole EndpointRoleType, config BaseEndpointConfig, ) *BaseEndpoint`

NewBaseEndpoint instantiates a new BaseEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEndpointWithDefaults

`func NewBaseEndpointWithDefaults() *BaseEndpoint`

NewBaseEndpointWithDefaults instantiates a new BaseEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseEndpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BaseEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseEndpoint) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BaseEndpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseEndpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseEndpoint) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnvironmentId

`func (o *BaseEndpoint) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *BaseEndpoint) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *BaseEndpoint) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetEndpointTypeId

`func (o *BaseEndpoint) GetEndpointTypeId() string`

GetEndpointTypeId returns the EndpointTypeId field if non-nil, zero value otherwise.

### GetEndpointTypeIdOk

`func (o *BaseEndpoint) GetEndpointTypeIdOk() (*string, bool)`

GetEndpointTypeIdOk returns a tuple with the EndpointTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointTypeId

`func (o *BaseEndpoint) SetEndpointTypeId(v string)`

SetEndpointTypeId sets EndpointTypeId field to given value.


### GetPartner

`func (o *BaseEndpoint) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *BaseEndpoint) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *BaseEndpoint) SetPartner(v Partner)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *BaseEndpoint) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetPartnerId

`func (o *BaseEndpoint) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *BaseEndpoint) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *BaseEndpoint) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *BaseEndpoint) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetDeployment

`func (o *BaseEndpoint) GetDeployment() EndpointDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *BaseEndpoint) GetDeploymentOk() (*EndpointDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *BaseEndpoint) SetDeployment(v EndpointDeployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *BaseEndpoint) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetEndpointRole

`func (o *BaseEndpoint) GetEndpointRole() EndpointRoleType`

GetEndpointRole returns the EndpointRole field if non-nil, zero value otherwise.

### GetEndpointRoleOk

`func (o *BaseEndpoint) GetEndpointRoleOk() (*EndpointRoleType, bool)`

GetEndpointRoleOk returns a tuple with the EndpointRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointRole

`func (o *BaseEndpoint) SetEndpointRole(v EndpointRoleType)`

SetEndpointRole sets EndpointRole field to given value.


### GetEndpointType

`func (o *BaseEndpoint) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *BaseEndpoint) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *BaseEndpoint) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *BaseEndpoint) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### GetConfig

`func (o *BaseEndpoint) GetConfig() BaseEndpointConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *BaseEndpoint) GetConfigOk() (*BaseEndpointConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *BaseEndpoint) SetConfig(v BaseEndpointConfig)`

SetConfig sets Config field to given value.


### GetCreatedAt

`func (o *BaseEndpoint) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseEndpoint) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseEndpoint) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BaseEndpoint) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BaseEndpoint) GetCreatedBy() UserDetail`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BaseEndpoint) GetCreatedByOk() (*UserDetail, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BaseEndpoint) SetCreatedBy(v UserDetail)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BaseEndpoint) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BaseEndpoint) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseEndpoint) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseEndpoint) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BaseEndpoint) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *BaseEndpoint) GetUpdatedBy() UserDetail`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *BaseEndpoint) GetUpdatedByOk() (*UserDetail, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *BaseEndpoint) SetUpdatedBy(v UserDetail)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *BaseEndpoint) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


