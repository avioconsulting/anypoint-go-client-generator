# GroupCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupCode** | **string** |  | 
**Codes** | [**[][]CodeKeyVal**]([]CodeKeyVal.md) |  | 

## Methods

### NewGroupCode

`func NewGroupCode(groupCode string, codes [][]CodeKeyVal, ) *GroupCode`

NewGroupCode instantiates a new GroupCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupCodeWithDefaults

`func NewGroupCodeWithDefaults() *GroupCode`

NewGroupCodeWithDefaults instantiates a new GroupCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupCode

`func (o *GroupCode) GetGroupCode() string`

GetGroupCode returns the GroupCode field if non-nil, zero value otherwise.

### GetGroupCodeOk

`func (o *GroupCode) GetGroupCodeOk() (*string, bool)`

GetGroupCodeOk returns a tuple with the GroupCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCode

`func (o *GroupCode) SetGroupCode(v string)`

SetGroupCode sets GroupCode field to given value.


### GetCodes

`func (o *GroupCode) GetCodes() [][]CodeKeyVal`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *GroupCode) GetCodesOk() (*[][]CodeKeyVal, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *GroupCode) SetCodes(v [][]CodeKeyVal)`

SetCodes sets Codes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


