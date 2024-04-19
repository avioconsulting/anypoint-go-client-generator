# BaseEndpointConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigName** | **string** |  | 
**ServerAddress** | Pointer to **string** | Defines Server/Host Address | [optional] 
**ServerPort** | Pointer to **int32** |  | [optional] 
**AuthMode** | Pointer to [**AuthModeBaseConfig**](AuthModeBaseConfig.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**MovedPath** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBaseEndpointConfig

`func NewBaseEndpointConfig(configName string, ) *BaseEndpointConfig`

NewBaseEndpointConfig instantiates a new BaseEndpointConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEndpointConfigWithDefaults

`func NewBaseEndpointConfigWithDefaults() *BaseEndpointConfig`

NewBaseEndpointConfigWithDefaults instantiates a new BaseEndpointConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigName

`func (o *BaseEndpointConfig) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *BaseEndpointConfig) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *BaseEndpointConfig) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.


### GetServerAddress

`func (o *BaseEndpointConfig) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *BaseEndpointConfig) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *BaseEndpointConfig) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *BaseEndpointConfig) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.

### GetServerPort

`func (o *BaseEndpointConfig) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *BaseEndpointConfig) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *BaseEndpointConfig) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *BaseEndpointConfig) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetAuthMode

`func (o *BaseEndpointConfig) GetAuthMode() AuthModeBaseConfig`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *BaseEndpointConfig) GetAuthModeOk() (*AuthModeBaseConfig, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *BaseEndpointConfig) SetAuthMode(v AuthModeBaseConfig)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *BaseEndpointConfig) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetPath

`func (o *BaseEndpointConfig) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BaseEndpointConfig) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BaseEndpointConfig) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BaseEndpointConfig) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMovedPath

`func (o *BaseEndpointConfig) GetMovedPath() string`

GetMovedPath returns the MovedPath field if non-nil, zero value otherwise.

### GetMovedPathOk

`func (o *BaseEndpointConfig) GetMovedPathOk() (*string, bool)`

GetMovedPathOk returns a tuple with the MovedPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovedPath

`func (o *BaseEndpointConfig) SetMovedPath(v string)`

SetMovedPath sets MovedPath field to given value.

### HasMovedPath

`func (o *BaseEndpointConfig) HasMovedPath() bool`

HasMovedPath returns a boolean if a field has been set.

### SetMovedPathNil

`func (o *BaseEndpointConfig) SetMovedPathNil(b bool)`

 SetMovedPathNil sets the value for MovedPath to be an explicit nil

### UnsetMovedPath
`func (o *BaseEndpointConfig) UnsetMovedPath()`

UnsetMovedPath ensures that no value is present for MovedPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


