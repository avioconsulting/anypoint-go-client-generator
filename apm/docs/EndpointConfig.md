# EndpointConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigName** | Pointer to **string** |  | [optional] 
**ServerAddress** | Pointer to **string** |  | [optional] 
**ServerPort** | Pointer to **int32** |  | [optional] 
**AuthMode** | Pointer to [**AuthModeConfig**](AuthModeConfig.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**MovedPath** | Pointer to **string** |  | [optional] 
**PollingFrequency** | Pointer to **int32** |  | [optional] 
**SizeCheckWaitTime** | Pointer to **int32** |  | [optional] 
**ResponseTimeout** | Pointer to **int32** |  | [optional] 
**ConnectionIdleTimeout** | Pointer to **int32** |  | [optional] 
**ManageWithApiManager** | Pointer to **bool** |  | [optional] 
**AutoDiscoveryId** | Pointer to **string** |  | [optional] 
**CredentialType** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to **string** |  | [optional] 
**AssociatedMessageTypeId** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**ValidWithoutPartner** | Pointer to **bool** |  | [optional] 
**AssociatedReferenceIdentifierId** | Pointer to **string** |  | [optional] 
**FileNamePattern** | Pointer to **string** |  | [optional] 
**TlsContext** | Pointer to [**TLSContext**](TLSContext.md) |  | [optional] 
**Vpc** | Pointer to [**EndpointConfigVpc**](EndpointConfigVpc.md) |  | [optional] 

## Methods

### NewEndpointConfig

`func NewEndpointConfig() *EndpointConfig`

NewEndpointConfig instantiates a new EndpointConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointConfigWithDefaults

`func NewEndpointConfigWithDefaults() *EndpointConfig`

NewEndpointConfigWithDefaults instantiates a new EndpointConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigName

`func (o *EndpointConfig) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *EndpointConfig) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *EndpointConfig) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *EndpointConfig) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### GetServerAddress

`func (o *EndpointConfig) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *EndpointConfig) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *EndpointConfig) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *EndpointConfig) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.

### GetServerPort

`func (o *EndpointConfig) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *EndpointConfig) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *EndpointConfig) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *EndpointConfig) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetAuthMode

`func (o *EndpointConfig) GetAuthMode() AuthModeConfig`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *EndpointConfig) GetAuthModeOk() (*AuthModeConfig, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *EndpointConfig) SetAuthMode(v AuthModeConfig)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *EndpointConfig) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetPath

`func (o *EndpointConfig) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EndpointConfig) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EndpointConfig) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *EndpointConfig) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMovedPath

`func (o *EndpointConfig) GetMovedPath() string`

GetMovedPath returns the MovedPath field if non-nil, zero value otherwise.

### GetMovedPathOk

`func (o *EndpointConfig) GetMovedPathOk() (*string, bool)`

GetMovedPathOk returns a tuple with the MovedPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovedPath

`func (o *EndpointConfig) SetMovedPath(v string)`

SetMovedPath sets MovedPath field to given value.

### HasMovedPath

`func (o *EndpointConfig) HasMovedPath() bool`

HasMovedPath returns a boolean if a field has been set.

### GetPollingFrequency

`func (o *EndpointConfig) GetPollingFrequency() int32`

GetPollingFrequency returns the PollingFrequency field if non-nil, zero value otherwise.

### GetPollingFrequencyOk

`func (o *EndpointConfig) GetPollingFrequencyOk() (*int32, bool)`

GetPollingFrequencyOk returns a tuple with the PollingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingFrequency

`func (o *EndpointConfig) SetPollingFrequency(v int32)`

SetPollingFrequency sets PollingFrequency field to given value.

### HasPollingFrequency

`func (o *EndpointConfig) HasPollingFrequency() bool`

HasPollingFrequency returns a boolean if a field has been set.

### GetSizeCheckWaitTime

`func (o *EndpointConfig) GetSizeCheckWaitTime() int32`

GetSizeCheckWaitTime returns the SizeCheckWaitTime field if non-nil, zero value otherwise.

### GetSizeCheckWaitTimeOk

`func (o *EndpointConfig) GetSizeCheckWaitTimeOk() (*int32, bool)`

GetSizeCheckWaitTimeOk returns a tuple with the SizeCheckWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeCheckWaitTime

`func (o *EndpointConfig) SetSizeCheckWaitTime(v int32)`

SetSizeCheckWaitTime sets SizeCheckWaitTime field to given value.

### HasSizeCheckWaitTime

`func (o *EndpointConfig) HasSizeCheckWaitTime() bool`

HasSizeCheckWaitTime returns a boolean if a field has been set.

### GetResponseTimeout

`func (o *EndpointConfig) GetResponseTimeout() int32`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *EndpointConfig) GetResponseTimeoutOk() (*int32, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *EndpointConfig) SetResponseTimeout(v int32)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *EndpointConfig) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### GetConnectionIdleTimeout

`func (o *EndpointConfig) GetConnectionIdleTimeout() int32`

GetConnectionIdleTimeout returns the ConnectionIdleTimeout field if non-nil, zero value otherwise.

### GetConnectionIdleTimeoutOk

`func (o *EndpointConfig) GetConnectionIdleTimeoutOk() (*int32, bool)`

GetConnectionIdleTimeoutOk returns a tuple with the ConnectionIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionIdleTimeout

`func (o *EndpointConfig) SetConnectionIdleTimeout(v int32)`

SetConnectionIdleTimeout sets ConnectionIdleTimeout field to given value.

### HasConnectionIdleTimeout

`func (o *EndpointConfig) HasConnectionIdleTimeout() bool`

HasConnectionIdleTimeout returns a boolean if a field has been set.

### GetManageWithApiManager

`func (o *EndpointConfig) GetManageWithApiManager() bool`

GetManageWithApiManager returns the ManageWithApiManager field if non-nil, zero value otherwise.

### GetManageWithApiManagerOk

`func (o *EndpointConfig) GetManageWithApiManagerOk() (*bool, bool)`

GetManageWithApiManagerOk returns a tuple with the ManageWithApiManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageWithApiManager

`func (o *EndpointConfig) SetManageWithApiManager(v bool)`

SetManageWithApiManager sets ManageWithApiManager field to given value.

### HasManageWithApiManager

`func (o *EndpointConfig) HasManageWithApiManager() bool`

HasManageWithApiManager returns a boolean if a field has been set.

### GetAutoDiscoveryId

`func (o *EndpointConfig) GetAutoDiscoveryId() string`

GetAutoDiscoveryId returns the AutoDiscoveryId field if non-nil, zero value otherwise.

### GetAutoDiscoveryIdOk

`func (o *EndpointConfig) GetAutoDiscoveryIdOk() (*string, bool)`

GetAutoDiscoveryIdOk returns a tuple with the AutoDiscoveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDiscoveryId

`func (o *EndpointConfig) SetAutoDiscoveryId(v string)`

SetAutoDiscoveryId sets AutoDiscoveryId field to given value.

### HasAutoDiscoveryId

`func (o *EndpointConfig) HasAutoDiscoveryId() bool`

HasAutoDiscoveryId returns a boolean if a field has been set.

### GetCredentialType

`func (o *EndpointConfig) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *EndpointConfig) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *EndpointConfig) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *EndpointConfig) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetUsage

`func (o *EndpointConfig) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *EndpointConfig) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *EndpointConfig) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *EndpointConfig) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetAssociatedMessageTypeId

`func (o *EndpointConfig) GetAssociatedMessageTypeId() string`

GetAssociatedMessageTypeId returns the AssociatedMessageTypeId field if non-nil, zero value otherwise.

### GetAssociatedMessageTypeIdOk

`func (o *EndpointConfig) GetAssociatedMessageTypeIdOk() (*string, bool)`

GetAssociatedMessageTypeIdOk returns a tuple with the AssociatedMessageTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedMessageTypeId

`func (o *EndpointConfig) SetAssociatedMessageTypeId(v string)`

SetAssociatedMessageTypeId sets AssociatedMessageTypeId field to given value.

### HasAssociatedMessageTypeId

`func (o *EndpointConfig) HasAssociatedMessageTypeId() bool`

HasAssociatedMessageTypeId returns a boolean if a field has been set.

### GetProtocol

`func (o *EndpointConfig) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *EndpointConfig) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *EndpointConfig) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *EndpointConfig) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetValidWithoutPartner

`func (o *EndpointConfig) GetValidWithoutPartner() bool`

GetValidWithoutPartner returns the ValidWithoutPartner field if non-nil, zero value otherwise.

### GetValidWithoutPartnerOk

`func (o *EndpointConfig) GetValidWithoutPartnerOk() (*bool, bool)`

GetValidWithoutPartnerOk returns a tuple with the ValidWithoutPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidWithoutPartner

`func (o *EndpointConfig) SetValidWithoutPartner(v bool)`

SetValidWithoutPartner sets ValidWithoutPartner field to given value.

### HasValidWithoutPartner

`func (o *EndpointConfig) HasValidWithoutPartner() bool`

HasValidWithoutPartner returns a boolean if a field has been set.

### GetAssociatedReferenceIdentifierId

`func (o *EndpointConfig) GetAssociatedReferenceIdentifierId() string`

GetAssociatedReferenceIdentifierId returns the AssociatedReferenceIdentifierId field if non-nil, zero value otherwise.

### GetAssociatedReferenceIdentifierIdOk

`func (o *EndpointConfig) GetAssociatedReferenceIdentifierIdOk() (*string, bool)`

GetAssociatedReferenceIdentifierIdOk returns a tuple with the AssociatedReferenceIdentifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedReferenceIdentifierId

`func (o *EndpointConfig) SetAssociatedReferenceIdentifierId(v string)`

SetAssociatedReferenceIdentifierId sets AssociatedReferenceIdentifierId field to given value.

### HasAssociatedReferenceIdentifierId

`func (o *EndpointConfig) HasAssociatedReferenceIdentifierId() bool`

HasAssociatedReferenceIdentifierId returns a boolean if a field has been set.

### GetFileNamePattern

`func (o *EndpointConfig) GetFileNamePattern() string`

GetFileNamePattern returns the FileNamePattern field if non-nil, zero value otherwise.

### GetFileNamePatternOk

`func (o *EndpointConfig) GetFileNamePatternOk() (*string, bool)`

GetFileNamePatternOk returns a tuple with the FileNamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNamePattern

`func (o *EndpointConfig) SetFileNamePattern(v string)`

SetFileNamePattern sets FileNamePattern field to given value.

### HasFileNamePattern

`func (o *EndpointConfig) HasFileNamePattern() bool`

HasFileNamePattern returns a boolean if a field has been set.

### GetTlsContext

`func (o *EndpointConfig) GetTlsContext() TLSContext`

GetTlsContext returns the TlsContext field if non-nil, zero value otherwise.

### GetTlsContextOk

`func (o *EndpointConfig) GetTlsContextOk() (*TLSContext, bool)`

GetTlsContextOk returns a tuple with the TlsContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsContext

`func (o *EndpointConfig) SetTlsContext(v TLSContext)`

SetTlsContext sets TlsContext field to given value.

### HasTlsContext

`func (o *EndpointConfig) HasTlsContext() bool`

HasTlsContext returns a boolean if a field has been set.

### GetVpc

`func (o *EndpointConfig) GetVpc() EndpointConfigVpc`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *EndpointConfig) GetVpcOk() (*EndpointConfigVpc, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *EndpointConfig) SetVpc(v EndpointConfigVpc)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *EndpointConfig) HasVpc() bool`

HasVpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


