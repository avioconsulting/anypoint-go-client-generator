# HttpResponseMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **float32** |  | 
**Message** | **string** |  | 

## Methods

### NewHttpResponseMessage

`func NewHttpResponseMessage(errorCode float32, message string, ) *HttpResponseMessage`

NewHttpResponseMessage instantiates a new HttpResponseMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpResponseMessageWithDefaults

`func NewHttpResponseMessageWithDefaults() *HttpResponseMessage`

NewHttpResponseMessageWithDefaults instantiates a new HttpResponseMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *HttpResponseMessage) GetErrorCode() float32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *HttpResponseMessage) GetErrorCodeOk() (*float32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *HttpResponseMessage) SetErrorCode(v float32)`

SetErrorCode sets ErrorCode field to given value.


### GetMessage

`func (o *HttpResponseMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HttpResponseMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HttpResponseMessage) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


