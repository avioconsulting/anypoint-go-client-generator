# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DocumentFlowDirection** | **string** |  | 
**DocumentTypeRole** | **string** |  | 
**EdiDocumentTypeId** | **string** |  | 
**IsStandard** | **bool** |  | 
**Name** | **string** |  | 
**MessageDataScript** | Pointer to **string** |  | [optional] 
**DocumentTypeParameters** | Pointer to **string** |  | [optional] 

## Methods

### NewDocument

`func NewDocument(documentFlowDirection string, documentTypeRole string, ediDocumentTypeId string, isStandard bool, name string, ) *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Document) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Document) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Document) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Document) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocumentFlowDirection

`func (o *Document) GetDocumentFlowDirection() string`

GetDocumentFlowDirection returns the DocumentFlowDirection field if non-nil, zero value otherwise.

### GetDocumentFlowDirectionOk

`func (o *Document) GetDocumentFlowDirectionOk() (*string, bool)`

GetDocumentFlowDirectionOk returns a tuple with the DocumentFlowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentFlowDirection

`func (o *Document) SetDocumentFlowDirection(v string)`

SetDocumentFlowDirection sets DocumentFlowDirection field to given value.


### GetDocumentTypeRole

`func (o *Document) GetDocumentTypeRole() string`

GetDocumentTypeRole returns the DocumentTypeRole field if non-nil, zero value otherwise.

### GetDocumentTypeRoleOk

`func (o *Document) GetDocumentTypeRoleOk() (*string, bool)`

GetDocumentTypeRoleOk returns a tuple with the DocumentTypeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeRole

`func (o *Document) SetDocumentTypeRole(v string)`

SetDocumentTypeRole sets DocumentTypeRole field to given value.


### GetEdiDocumentTypeId

`func (o *Document) GetEdiDocumentTypeId() string`

GetEdiDocumentTypeId returns the EdiDocumentTypeId field if non-nil, zero value otherwise.

### GetEdiDocumentTypeIdOk

`func (o *Document) GetEdiDocumentTypeIdOk() (*string, bool)`

GetEdiDocumentTypeIdOk returns a tuple with the EdiDocumentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdiDocumentTypeId

`func (o *Document) SetEdiDocumentTypeId(v string)`

SetEdiDocumentTypeId sets EdiDocumentTypeId field to given value.


### GetIsStandard

`func (o *Document) GetIsStandard() bool`

GetIsStandard returns the IsStandard field if non-nil, zero value otherwise.

### GetIsStandardOk

`func (o *Document) GetIsStandardOk() (*bool, bool)`

GetIsStandardOk returns a tuple with the IsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandard

`func (o *Document) SetIsStandard(v bool)`

SetIsStandard sets IsStandard field to given value.


### GetName

`func (o *Document) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Document) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Document) SetName(v string)`

SetName sets Name field to given value.


### GetMessageDataScript

`func (o *Document) GetMessageDataScript() string`

GetMessageDataScript returns the MessageDataScript field if non-nil, zero value otherwise.

### GetMessageDataScriptOk

`func (o *Document) GetMessageDataScriptOk() (*string, bool)`

GetMessageDataScriptOk returns a tuple with the MessageDataScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDataScript

`func (o *Document) SetMessageDataScript(v string)`

SetMessageDataScript sets MessageDataScript field to given value.

### HasMessageDataScript

`func (o *Document) HasMessageDataScript() bool`

HasMessageDataScript returns a boolean if a field has been set.

### GetDocumentTypeParameters

`func (o *Document) GetDocumentTypeParameters() string`

GetDocumentTypeParameters returns the DocumentTypeParameters field if non-nil, zero value otherwise.

### GetDocumentTypeParametersOk

`func (o *Document) GetDocumentTypeParametersOk() (*string, bool)`

GetDocumentTypeParametersOk returns a tuple with the DocumentTypeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeParameters

`func (o *Document) SetDocumentTypeParameters(v string)`

SetDocumentTypeParameters sets DocumentTypeParameters field to given value.

### HasDocumentTypeParameters

`func (o *Document) HasDocumentTypeParameters() bool`

HasDocumentTypeParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


