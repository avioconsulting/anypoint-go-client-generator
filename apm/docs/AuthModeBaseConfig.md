# AuthModeBaseConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **string** |  | 
**CsmSecret** | Pointer to [**CSMSecret**](CSMSecret.md) |  | [optional] 

## Methods

### NewAuthModeBaseConfig

`func NewAuthModeBaseConfig(authType string, ) *AuthModeBaseConfig`

NewAuthModeBaseConfig instantiates a new AuthModeBaseConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthModeBaseConfigWithDefaults

`func NewAuthModeBaseConfigWithDefaults() *AuthModeBaseConfig`

NewAuthModeBaseConfigWithDefaults instantiates a new AuthModeBaseConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *AuthModeBaseConfig) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AuthModeBaseConfig) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AuthModeBaseConfig) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetCsmSecret

`func (o *AuthModeBaseConfig) GetCsmSecret() CSMSecret`

GetCsmSecret returns the CsmSecret field if non-nil, zero value otherwise.

### GetCsmSecretOk

`func (o *AuthModeBaseConfig) GetCsmSecretOk() (*CSMSecret, bool)`

GetCsmSecretOk returns a tuple with the CsmSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsmSecret

`func (o *AuthModeBaseConfig) SetCsmSecret(v CSMSecret)`

SetCsmSecret sets CsmSecret field to given value.

### HasCsmSecret

`func (o *AuthModeBaseConfig) HasCsmSecret() bool`

HasCsmSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


