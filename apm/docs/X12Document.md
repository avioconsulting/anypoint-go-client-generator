# X12Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X12Version** | Pointer to **string** |  | [optional] 
**TransactionSet** | Pointer to **string** |  | [optional] 

## Methods

### NewX12Document

`func NewX12Document() *X12Document`

NewX12Document instantiates a new X12Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX12DocumentWithDefaults

`func NewX12DocumentWithDefaults() *X12Document`

NewX12DocumentWithDefaults instantiates a new X12Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX12Version

`func (o *X12Document) GetX12Version() string`

GetX12Version returns the X12Version field if non-nil, zero value otherwise.

### GetX12VersionOk

`func (o *X12Document) GetX12VersionOk() (*string, bool)`

GetX12VersionOk returns a tuple with the X12Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX12Version

`func (o *X12Document) SetX12Version(v string)`

SetX12Version sets X12Version field to given value.

### HasX12Version

`func (o *X12Document) HasX12Version() bool`

HasX12Version returns a boolean if a field has been set.

### GetTransactionSet

`func (o *X12Document) GetTransactionSet() string`

GetTransactionSet returns the TransactionSet field if non-nil, zero value otherwise.

### GetTransactionSetOk

`func (o *X12Document) GetTransactionSetOk() (*string, bool)`

GetTransactionSetOk returns a tuple with the TransactionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionSet

`func (o *X12Document) SetTransactionSet(v string)`

SetTransactionSet sets TransactionSet field to given value.

### HasTransactionSet

`func (o *X12Document) HasTransactionSet() bool`

HasTransactionSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


