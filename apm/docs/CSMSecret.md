# CSMSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretId** | **string** |  | 
**SecretGroupId** | **string** |  | 
**GrantResponse** | [**CSMSecretGrant**](CSMSecretGrant.md) |  | 
**JsonDetails** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCSMSecret

`func NewCSMSecret(secretId string, secretGroupId string, grantResponse CSMSecretGrant, ) *CSMSecret`

NewCSMSecret instantiates a new CSMSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSMSecretWithDefaults

`func NewCSMSecretWithDefaults() *CSMSecret`

NewCSMSecretWithDefaults instantiates a new CSMSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretId

`func (o *CSMSecret) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *CSMSecret) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *CSMSecret) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.


### GetSecretGroupId

`func (o *CSMSecret) GetSecretGroupId() string`

GetSecretGroupId returns the SecretGroupId field if non-nil, zero value otherwise.

### GetSecretGroupIdOk

`func (o *CSMSecret) GetSecretGroupIdOk() (*string, bool)`

GetSecretGroupIdOk returns a tuple with the SecretGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretGroupId

`func (o *CSMSecret) SetSecretGroupId(v string)`

SetSecretGroupId sets SecretGroupId field to given value.


### GetGrantResponse

`func (o *CSMSecret) GetGrantResponse() CSMSecretGrant`

GetGrantResponse returns the GrantResponse field if non-nil, zero value otherwise.

### GetGrantResponseOk

`func (o *CSMSecret) GetGrantResponseOk() (*CSMSecretGrant, bool)`

GetGrantResponseOk returns a tuple with the GrantResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantResponse

`func (o *CSMSecret) SetGrantResponse(v CSMSecretGrant)`

SetGrantResponse sets GrantResponse field to given value.


### GetJsonDetails

`func (o *CSMSecret) GetJsonDetails() string`

GetJsonDetails returns the JsonDetails field if non-nil, zero value otherwise.

### GetJsonDetailsOk

`func (o *CSMSecret) GetJsonDetailsOk() (*string, bool)`

GetJsonDetailsOk returns a tuple with the JsonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonDetails

`func (o *CSMSecret) SetJsonDetails(v string)`

SetJsonDetails sets JsonDetails field to given value.

### HasJsonDetails

`func (o *CSMSecret) HasJsonDetails() bool`

HasJsonDetails returns a boolean if a field has been set.

### SetJsonDetailsNil

`func (o *CSMSecret) SetJsonDetailsNil(b bool)`

 SetJsonDetailsNil sets the value for JsonDetails to be an explicit nil

### UnsetJsonDetails
`func (o *CSMSecret) UnsetJsonDetails()`

UnsetJsonDetails ensures that no value is present for JsonDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


