# DocumentMappingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**MappingType** | **string** |  | 
**MappingContent** | Pointer to **string** | required only when the mapppingType is not MAP_API | [optional] 
**MappingSourceRef** | **string** | For DWL_SCRIPT is Inline , DWL_FILE is the file name, and for MAP_API is TargetAPI endpoint Id. | 

## Methods

### NewDocumentMappingConfig

`func NewDocumentMappingConfig(mappingType string, mappingSourceRef string, ) *DocumentMappingConfig`

NewDocumentMappingConfig instantiates a new DocumentMappingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentMappingConfigWithDefaults

`func NewDocumentMappingConfigWithDefaults() *DocumentMappingConfig`

NewDocumentMappingConfigWithDefaults instantiates a new DocumentMappingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentMappingConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentMappingConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentMappingConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentMappingConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMappingType

`func (o *DocumentMappingConfig) GetMappingType() string`

GetMappingType returns the MappingType field if non-nil, zero value otherwise.

### GetMappingTypeOk

`func (o *DocumentMappingConfig) GetMappingTypeOk() (*string, bool)`

GetMappingTypeOk returns a tuple with the MappingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingType

`func (o *DocumentMappingConfig) SetMappingType(v string)`

SetMappingType sets MappingType field to given value.


### GetMappingContent

`func (o *DocumentMappingConfig) GetMappingContent() string`

GetMappingContent returns the MappingContent field if non-nil, zero value otherwise.

### GetMappingContentOk

`func (o *DocumentMappingConfig) GetMappingContentOk() (*string, bool)`

GetMappingContentOk returns a tuple with the MappingContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingContent

`func (o *DocumentMappingConfig) SetMappingContent(v string)`

SetMappingContent sets MappingContent field to given value.

### HasMappingContent

`func (o *DocumentMappingConfig) HasMappingContent() bool`

HasMappingContent returns a boolean if a field has been set.

### GetMappingSourceRef

`func (o *DocumentMappingConfig) GetMappingSourceRef() string`

GetMappingSourceRef returns the MappingSourceRef field if non-nil, zero value otherwise.

### GetMappingSourceRefOk

`func (o *DocumentMappingConfig) GetMappingSourceRefOk() (*string, bool)`

GetMappingSourceRefOk returns a tuple with the MappingSourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingSourceRef

`func (o *DocumentMappingConfig) SetMappingSourceRef(v string)`

SetMappingSourceRef sets MappingSourceRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


