# Http

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**HttpConfig**](HttpConfig.md) |  | 

## Methods

### NewHttp

`func NewHttp(config HttpConfig, ) *Http`

NewHttp instantiates a new Http object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpWithDefaults

`func NewHttpWithDefaults() *Http`

NewHttpWithDefaults instantiates a new Http object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *Http) GetConfig() HttpConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Http) GetConfigOk() (*HttpConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Http) SetConfig(v HttpConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


