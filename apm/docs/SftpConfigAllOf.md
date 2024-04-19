# SftpConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerAddress** | Pointer to **string** |  | [optional] 
**ServerPort** | Pointer to **int32** |  | [optional] 
**AuthMode** | Pointer to [**AuthModeBasic**](AuthModeBasic.md) |  | [optional] 
**FileAge** | Pointer to **float32** |  | [optional] 
**PollingFrequency** | Pointer to **float32** |  | [optional] 
**SizeCheckWaitTime** | Pointer to **int32** |  | [optional] 

## Methods

### NewSftpConfigAllOf

`func NewSftpConfigAllOf() *SftpConfigAllOf`

NewSftpConfigAllOf instantiates a new SftpConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSftpConfigAllOfWithDefaults

`func NewSftpConfigAllOfWithDefaults() *SftpConfigAllOf`

NewSftpConfigAllOfWithDefaults instantiates a new SftpConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerAddress

`func (o *SftpConfigAllOf) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *SftpConfigAllOf) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *SftpConfigAllOf) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *SftpConfigAllOf) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.

### GetServerPort

`func (o *SftpConfigAllOf) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *SftpConfigAllOf) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *SftpConfigAllOf) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *SftpConfigAllOf) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetAuthMode

`func (o *SftpConfigAllOf) GetAuthMode() AuthModeBasic`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SftpConfigAllOf) GetAuthModeOk() (*AuthModeBasic, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SftpConfigAllOf) SetAuthMode(v AuthModeBasic)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *SftpConfigAllOf) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetFileAge

`func (o *SftpConfigAllOf) GetFileAge() float32`

GetFileAge returns the FileAge field if non-nil, zero value otherwise.

### GetFileAgeOk

`func (o *SftpConfigAllOf) GetFileAgeOk() (*float32, bool)`

GetFileAgeOk returns a tuple with the FileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAge

`func (o *SftpConfigAllOf) SetFileAge(v float32)`

SetFileAge sets FileAge field to given value.

### HasFileAge

`func (o *SftpConfigAllOf) HasFileAge() bool`

HasFileAge returns a boolean if a field has been set.

### GetPollingFrequency

`func (o *SftpConfigAllOf) GetPollingFrequency() float32`

GetPollingFrequency returns the PollingFrequency field if non-nil, zero value otherwise.

### GetPollingFrequencyOk

`func (o *SftpConfigAllOf) GetPollingFrequencyOk() (*float32, bool)`

GetPollingFrequencyOk returns a tuple with the PollingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingFrequency

`func (o *SftpConfigAllOf) SetPollingFrequency(v float32)`

SetPollingFrequency sets PollingFrequency field to given value.

### HasPollingFrequency

`func (o *SftpConfigAllOf) HasPollingFrequency() bool`

HasPollingFrequency returns a boolean if a field has been set.

### GetSizeCheckWaitTime

`func (o *SftpConfigAllOf) GetSizeCheckWaitTime() int32`

GetSizeCheckWaitTime returns the SizeCheckWaitTime field if non-nil, zero value otherwise.

### GetSizeCheckWaitTimeOk

`func (o *SftpConfigAllOf) GetSizeCheckWaitTimeOk() (*int32, bool)`

GetSizeCheckWaitTimeOk returns a tuple with the SizeCheckWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeCheckWaitTime

`func (o *SftpConfigAllOf) SetSizeCheckWaitTime(v int32)`

SetSizeCheckWaitTime sets SizeCheckWaitTime field to given value.

### HasSizeCheckWaitTime

`func (o *SftpConfigAllOf) HasSizeCheckWaitTime() bool`

HasSizeCheckWaitTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


