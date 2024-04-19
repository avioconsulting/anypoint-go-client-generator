# HttpConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** |  | [optional] [default to "HTTP"]
**ServerAddress** | Pointer to **string** |  | [optional] 
**ServerPort** | Pointer to **int32** |  | [optional] [default to 80]
**Path** | Pointer to **string** |  | [optional] 
**AllowedMethods** | Pointer to **string** |  | [optional] [default to "POST"]
**PersistentConnections** | Pointer to **bool** |  | [optional] [default to false]
**ConnectionIdleTimeout** | Pointer to **int32** |  | [optional] [default to 30000]
**ResponseTimeout** | Pointer to **int32** |  | [optional] [default to 15000]
**AuthMode** | Pointer to [**AuthModeBaseConfig**](AuthModeBaseConfig.md) |  | [optional] 
**TlsContext** | Pointer to [**TLSContext**](TLSContext.md) |  | [optional] 

## Methods

### NewHttpConfigAllOf

`func NewHttpConfigAllOf() *HttpConfigAllOf`

NewHttpConfigAllOf instantiates a new HttpConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpConfigAllOfWithDefaults

`func NewHttpConfigAllOfWithDefaults() *HttpConfigAllOf`

NewHttpConfigAllOfWithDefaults instantiates a new HttpConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *HttpConfigAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HttpConfigAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HttpConfigAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *HttpConfigAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetServerAddress

`func (o *HttpConfigAllOf) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *HttpConfigAllOf) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *HttpConfigAllOf) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *HttpConfigAllOf) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.

### GetServerPort

`func (o *HttpConfigAllOf) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *HttpConfigAllOf) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *HttpConfigAllOf) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *HttpConfigAllOf) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetPath

`func (o *HttpConfigAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HttpConfigAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HttpConfigAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HttpConfigAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetAllowedMethods

`func (o *HttpConfigAllOf) GetAllowedMethods() string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *HttpConfigAllOf) GetAllowedMethodsOk() (*string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *HttpConfigAllOf) SetAllowedMethods(v string)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *HttpConfigAllOf) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.

### GetPersistentConnections

`func (o *HttpConfigAllOf) GetPersistentConnections() bool`

GetPersistentConnections returns the PersistentConnections field if non-nil, zero value otherwise.

### GetPersistentConnectionsOk

`func (o *HttpConfigAllOf) GetPersistentConnectionsOk() (*bool, bool)`

GetPersistentConnectionsOk returns a tuple with the PersistentConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentConnections

`func (o *HttpConfigAllOf) SetPersistentConnections(v bool)`

SetPersistentConnections sets PersistentConnections field to given value.

### HasPersistentConnections

`func (o *HttpConfigAllOf) HasPersistentConnections() bool`

HasPersistentConnections returns a boolean if a field has been set.

### GetConnectionIdleTimeout

`func (o *HttpConfigAllOf) GetConnectionIdleTimeout() int32`

GetConnectionIdleTimeout returns the ConnectionIdleTimeout field if non-nil, zero value otherwise.

### GetConnectionIdleTimeoutOk

`func (o *HttpConfigAllOf) GetConnectionIdleTimeoutOk() (*int32, bool)`

GetConnectionIdleTimeoutOk returns a tuple with the ConnectionIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionIdleTimeout

`func (o *HttpConfigAllOf) SetConnectionIdleTimeout(v int32)`

SetConnectionIdleTimeout sets ConnectionIdleTimeout field to given value.

### HasConnectionIdleTimeout

`func (o *HttpConfigAllOf) HasConnectionIdleTimeout() bool`

HasConnectionIdleTimeout returns a boolean if a field has been set.

### GetResponseTimeout

`func (o *HttpConfigAllOf) GetResponseTimeout() int32`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *HttpConfigAllOf) GetResponseTimeoutOk() (*int32, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *HttpConfigAllOf) SetResponseTimeout(v int32)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *HttpConfigAllOf) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### GetAuthMode

`func (o *HttpConfigAllOf) GetAuthMode() AuthModeBaseConfig`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *HttpConfigAllOf) GetAuthModeOk() (*AuthModeBaseConfig, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *HttpConfigAllOf) SetAuthMode(v AuthModeBaseConfig)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *HttpConfigAllOf) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetTlsContext

`func (o *HttpConfigAllOf) GetTlsContext() TLSContext`

GetTlsContext returns the TlsContext field if non-nil, zero value otherwise.

### GetTlsContextOk

`func (o *HttpConfigAllOf) GetTlsContextOk() (*TLSContext, bool)`

GetTlsContextOk returns a tuple with the TlsContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsContext

`func (o *HttpConfigAllOf) SetTlsContext(v TLSContext)`

SetTlsContext sets TlsContext field to given value.

### HasTlsContext

`func (o *HttpConfigAllOf) HasTlsContext() bool`

HasTlsContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


