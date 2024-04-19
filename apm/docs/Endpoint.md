# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**EndpointTypeId** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **string** |  | [optional] 
**EndpointRole** | Pointer to **string** |  | [optional] 
**EndpointType** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**SupportedFormatTypes** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**EndpointConfig**](EndpointConfig.md) |  | [optional] 

## Methods

### NewEndpoint

`func NewEndpoint() *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Endpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Endpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Endpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Endpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Endpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Endpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Endpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Endpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Endpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Endpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Endpoint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Endpoint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *Endpoint) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *Endpoint) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *Endpoint) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *Endpoint) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEndpointTypeId

`func (o *Endpoint) GetEndpointTypeId() string`

GetEndpointTypeId returns the EndpointTypeId field if non-nil, zero value otherwise.

### GetEndpointTypeIdOk

`func (o *Endpoint) GetEndpointTypeIdOk() (*string, bool)`

GetEndpointTypeIdOk returns a tuple with the EndpointTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointTypeId

`func (o *Endpoint) SetEndpointTypeId(v string)`

SetEndpointTypeId sets EndpointTypeId field to given value.

### HasEndpointTypeId

`func (o *Endpoint) HasEndpointTypeId() bool`

HasEndpointTypeId returns a boolean if a field has been set.

### GetPartnerId

`func (o *Endpoint) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *Endpoint) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *Endpoint) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *Endpoint) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetEndpointRole

`func (o *Endpoint) GetEndpointRole() string`

GetEndpointRole returns the EndpointRole field if non-nil, zero value otherwise.

### GetEndpointRoleOk

`func (o *Endpoint) GetEndpointRoleOk() (*string, bool)`

GetEndpointRoleOk returns a tuple with the EndpointRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointRole

`func (o *Endpoint) SetEndpointRole(v string)`

SetEndpointRole sets EndpointRole field to given value.

### HasEndpointRole

`func (o *Endpoint) HasEndpointRole() bool`

HasEndpointRole returns a boolean if a field has been set.

### GetEndpointType

`func (o *Endpoint) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *Endpoint) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *Endpoint) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *Endpoint) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### GetVisibility

`func (o *Endpoint) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Endpoint) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Endpoint) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Endpoint) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetSupportedFormatTypes

`func (o *Endpoint) GetSupportedFormatTypes() string`

GetSupportedFormatTypes returns the SupportedFormatTypes field if non-nil, zero value otherwise.

### GetSupportedFormatTypesOk

`func (o *Endpoint) GetSupportedFormatTypesOk() (*string, bool)`

GetSupportedFormatTypesOk returns a tuple with the SupportedFormatTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFormatTypes

`func (o *Endpoint) SetSupportedFormatTypes(v string)`

SetSupportedFormatTypes sets SupportedFormatTypes field to given value.

### HasSupportedFormatTypes

`func (o *Endpoint) HasSupportedFormatTypes() bool`

HasSupportedFormatTypes returns a boolean if a field has been set.

### GetConfig

`func (o *Endpoint) GetConfig() EndpointConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Endpoint) GetConfigOk() (*EndpointConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Endpoint) SetConfig(v EndpointConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Endpoint) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


