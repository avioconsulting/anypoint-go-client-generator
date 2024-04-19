# BaseEdiDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**IsStandard** | **bool** |  | 
**EdiDocumentTypeId** | **string** |  | 
**BaseType** | Pointer to [**EdiDocumentType**](EdiDocumentType.md) |  | [optional] 

## Methods

### NewBaseEdiDocument

`func NewBaseEdiDocument(name string, isStandard bool, ediDocumentTypeId string, ) *BaseEdiDocument`

NewBaseEdiDocument instantiates a new BaseEdiDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEdiDocumentWithDefaults

`func NewBaseEdiDocumentWithDefaults() *BaseEdiDocument`

NewBaseEdiDocumentWithDefaults instantiates a new BaseEdiDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseEdiDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseEdiDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseEdiDocument) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseEdiDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BaseEdiDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseEdiDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseEdiDocument) SetName(v string)`

SetName sets Name field to given value.


### GetIsStandard

`func (o *BaseEdiDocument) GetIsStandard() bool`

GetIsStandard returns the IsStandard field if non-nil, zero value otherwise.

### GetIsStandardOk

`func (o *BaseEdiDocument) GetIsStandardOk() (*bool, bool)`

GetIsStandardOk returns a tuple with the IsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandard

`func (o *BaseEdiDocument) SetIsStandard(v bool)`

SetIsStandard sets IsStandard field to given value.


### GetEdiDocumentTypeId

`func (o *BaseEdiDocument) GetEdiDocumentTypeId() string`

GetEdiDocumentTypeId returns the EdiDocumentTypeId field if non-nil, zero value otherwise.

### GetEdiDocumentTypeIdOk

`func (o *BaseEdiDocument) GetEdiDocumentTypeIdOk() (*string, bool)`

GetEdiDocumentTypeIdOk returns a tuple with the EdiDocumentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdiDocumentTypeId

`func (o *BaseEdiDocument) SetEdiDocumentTypeId(v string)`

SetEdiDocumentTypeId sets EdiDocumentTypeId field to given value.


### GetBaseType

`func (o *BaseEdiDocument) GetBaseType() EdiDocumentType`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *BaseEdiDocument) GetBaseTypeOk() (*EdiDocumentType, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *BaseEdiDocument) SetBaseType(v EdiDocumentType)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *BaseEdiDocument) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


