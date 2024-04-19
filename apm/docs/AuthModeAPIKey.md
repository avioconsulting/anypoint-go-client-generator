# AuthModeAPIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** |  | [optional] 
**HttpHeaderName** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthModeAPIKey

`func NewAuthModeAPIKey() *AuthModeAPIKey`

NewAuthModeAPIKey instantiates a new AuthModeAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthModeAPIKeyWithDefaults

`func NewAuthModeAPIKeyWithDefaults() *AuthModeAPIKey`

NewAuthModeAPIKeyWithDefaults instantiates a new AuthModeAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *AuthModeAPIKey) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AuthModeAPIKey) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AuthModeAPIKey) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AuthModeAPIKey) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetHttpHeaderName

`func (o *AuthModeAPIKey) GetHttpHeaderName() string`

GetHttpHeaderName returns the HttpHeaderName field if non-nil, zero value otherwise.

### GetHttpHeaderNameOk

`func (o *AuthModeAPIKey) GetHttpHeaderNameOk() (*string, bool)`

GetHttpHeaderNameOk returns a tuple with the HttpHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaderName

`func (o *AuthModeAPIKey) SetHttpHeaderName(v string)`

SetHttpHeaderName sets HttpHeaderName field to given value.

### HasHttpHeaderName

`func (o *AuthModeAPIKey) HasHttpHeaderName() bool`

HasHttpHeaderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


