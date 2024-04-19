# DocumentFlowConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DocumentFlowId** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | [default to "DRAFT"]
**Version** | Pointer to **int32** |  | [optional] 
**PreProcessingEndpointId** | Pointer to **string** |  | [optional] 
**ReceivingEndpointId** | Pointer to **string** |  | [optional] 
**ReceivingAckEndpointId** | Pointer to **string** |  | [optional] 
**TargetEndpointId** | Pointer to **string** |  | [optional] 
**SourceDocTypeId** | Pointer to **string** |  | [optional] 
**TargetDocTypeId** | Pointer to **string** |  | [optional] 
**DocumentMapping** | Pointer to [**[]DocumentMappingConfig**](DocumentMappingConfig.md) |  | [optional] 
**ReceivingAckConfig** | Pointer to [**BaseAckConfig**](BaseAckConfig.md) |  | [optional] 
**PartnerIdentifiersConfig** | Pointer to [**DocumentFlowConfigAllOfPartnerIdentifiersConfig**](DocumentFlowConfigAllOfPartnerIdentifiersConfig.md) |  | [optional] 

## Methods

### NewDocumentFlowConfigAllOf

`func NewDocumentFlowConfigAllOf(status string, ) *DocumentFlowConfigAllOf`

NewDocumentFlowConfigAllOf instantiates a new DocumentFlowConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentFlowConfigAllOfWithDefaults

`func NewDocumentFlowConfigAllOfWithDefaults() *DocumentFlowConfigAllOf`

NewDocumentFlowConfigAllOfWithDefaults instantiates a new DocumentFlowConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentFlowConfigAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentFlowConfigAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentFlowConfigAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentFlowConfigAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocumentFlowId

`func (o *DocumentFlowConfigAllOf) GetDocumentFlowId() string`

GetDocumentFlowId returns the DocumentFlowId field if non-nil, zero value otherwise.

### GetDocumentFlowIdOk

`func (o *DocumentFlowConfigAllOf) GetDocumentFlowIdOk() (*string, bool)`

GetDocumentFlowIdOk returns a tuple with the DocumentFlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentFlowId

`func (o *DocumentFlowConfigAllOf) SetDocumentFlowId(v string)`

SetDocumentFlowId sets DocumentFlowId field to given value.

### HasDocumentFlowId

`func (o *DocumentFlowConfigAllOf) HasDocumentFlowId() bool`

HasDocumentFlowId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DocumentFlowConfigAllOf) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DocumentFlowConfigAllOf) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DocumentFlowConfigAllOf) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DocumentFlowConfigAllOf) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetStatus

`func (o *DocumentFlowConfigAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentFlowConfigAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentFlowConfigAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *DocumentFlowConfigAllOf) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DocumentFlowConfigAllOf) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DocumentFlowConfigAllOf) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DocumentFlowConfigAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPreProcessingEndpointId

`func (o *DocumentFlowConfigAllOf) GetPreProcessingEndpointId() string`

GetPreProcessingEndpointId returns the PreProcessingEndpointId field if non-nil, zero value otherwise.

### GetPreProcessingEndpointIdOk

`func (o *DocumentFlowConfigAllOf) GetPreProcessingEndpointIdOk() (*string, bool)`

GetPreProcessingEndpointIdOk returns a tuple with the PreProcessingEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreProcessingEndpointId

`func (o *DocumentFlowConfigAllOf) SetPreProcessingEndpointId(v string)`

SetPreProcessingEndpointId sets PreProcessingEndpointId field to given value.

### HasPreProcessingEndpointId

`func (o *DocumentFlowConfigAllOf) HasPreProcessingEndpointId() bool`

HasPreProcessingEndpointId returns a boolean if a field has been set.

### GetReceivingEndpointId

`func (o *DocumentFlowConfigAllOf) GetReceivingEndpointId() string`

GetReceivingEndpointId returns the ReceivingEndpointId field if non-nil, zero value otherwise.

### GetReceivingEndpointIdOk

`func (o *DocumentFlowConfigAllOf) GetReceivingEndpointIdOk() (*string, bool)`

GetReceivingEndpointIdOk returns a tuple with the ReceivingEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingEndpointId

`func (o *DocumentFlowConfigAllOf) SetReceivingEndpointId(v string)`

SetReceivingEndpointId sets ReceivingEndpointId field to given value.

### HasReceivingEndpointId

`func (o *DocumentFlowConfigAllOf) HasReceivingEndpointId() bool`

HasReceivingEndpointId returns a boolean if a field has been set.

### GetReceivingAckEndpointId

`func (o *DocumentFlowConfigAllOf) GetReceivingAckEndpointId() string`

GetReceivingAckEndpointId returns the ReceivingAckEndpointId field if non-nil, zero value otherwise.

### GetReceivingAckEndpointIdOk

`func (o *DocumentFlowConfigAllOf) GetReceivingAckEndpointIdOk() (*string, bool)`

GetReceivingAckEndpointIdOk returns a tuple with the ReceivingAckEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAckEndpointId

`func (o *DocumentFlowConfigAllOf) SetReceivingAckEndpointId(v string)`

SetReceivingAckEndpointId sets ReceivingAckEndpointId field to given value.

### HasReceivingAckEndpointId

`func (o *DocumentFlowConfigAllOf) HasReceivingAckEndpointId() bool`

HasReceivingAckEndpointId returns a boolean if a field has been set.

### GetTargetEndpointId

`func (o *DocumentFlowConfigAllOf) GetTargetEndpointId() string`

GetTargetEndpointId returns the TargetEndpointId field if non-nil, zero value otherwise.

### GetTargetEndpointIdOk

`func (o *DocumentFlowConfigAllOf) GetTargetEndpointIdOk() (*string, bool)`

GetTargetEndpointIdOk returns a tuple with the TargetEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEndpointId

`func (o *DocumentFlowConfigAllOf) SetTargetEndpointId(v string)`

SetTargetEndpointId sets TargetEndpointId field to given value.

### HasTargetEndpointId

`func (o *DocumentFlowConfigAllOf) HasTargetEndpointId() bool`

HasTargetEndpointId returns a boolean if a field has been set.

### GetSourceDocTypeId

`func (o *DocumentFlowConfigAllOf) GetSourceDocTypeId() string`

GetSourceDocTypeId returns the SourceDocTypeId field if non-nil, zero value otherwise.

### GetSourceDocTypeIdOk

`func (o *DocumentFlowConfigAllOf) GetSourceDocTypeIdOk() (*string, bool)`

GetSourceDocTypeIdOk returns a tuple with the SourceDocTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocTypeId

`func (o *DocumentFlowConfigAllOf) SetSourceDocTypeId(v string)`

SetSourceDocTypeId sets SourceDocTypeId field to given value.

### HasSourceDocTypeId

`func (o *DocumentFlowConfigAllOf) HasSourceDocTypeId() bool`

HasSourceDocTypeId returns a boolean if a field has been set.

### GetTargetDocTypeId

`func (o *DocumentFlowConfigAllOf) GetTargetDocTypeId() string`

GetTargetDocTypeId returns the TargetDocTypeId field if non-nil, zero value otherwise.

### GetTargetDocTypeIdOk

`func (o *DocumentFlowConfigAllOf) GetTargetDocTypeIdOk() (*string, bool)`

GetTargetDocTypeIdOk returns a tuple with the TargetDocTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDocTypeId

`func (o *DocumentFlowConfigAllOf) SetTargetDocTypeId(v string)`

SetTargetDocTypeId sets TargetDocTypeId field to given value.

### HasTargetDocTypeId

`func (o *DocumentFlowConfigAllOf) HasTargetDocTypeId() bool`

HasTargetDocTypeId returns a boolean if a field has been set.

### GetDocumentMapping

`func (o *DocumentFlowConfigAllOf) GetDocumentMapping() []DocumentMappingConfig`

GetDocumentMapping returns the DocumentMapping field if non-nil, zero value otherwise.

### GetDocumentMappingOk

`func (o *DocumentFlowConfigAllOf) GetDocumentMappingOk() (*[]DocumentMappingConfig, bool)`

GetDocumentMappingOk returns a tuple with the DocumentMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMapping

`func (o *DocumentFlowConfigAllOf) SetDocumentMapping(v []DocumentMappingConfig)`

SetDocumentMapping sets DocumentMapping field to given value.

### HasDocumentMapping

`func (o *DocumentFlowConfigAllOf) HasDocumentMapping() bool`

HasDocumentMapping returns a boolean if a field has been set.

### GetReceivingAckConfig

`func (o *DocumentFlowConfigAllOf) GetReceivingAckConfig() BaseAckConfig`

GetReceivingAckConfig returns the ReceivingAckConfig field if non-nil, zero value otherwise.

### GetReceivingAckConfigOk

`func (o *DocumentFlowConfigAllOf) GetReceivingAckConfigOk() (*BaseAckConfig, bool)`

GetReceivingAckConfigOk returns a tuple with the ReceivingAckConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAckConfig

`func (o *DocumentFlowConfigAllOf) SetReceivingAckConfig(v BaseAckConfig)`

SetReceivingAckConfig sets ReceivingAckConfig field to given value.

### HasReceivingAckConfig

`func (o *DocumentFlowConfigAllOf) HasReceivingAckConfig() bool`

HasReceivingAckConfig returns a boolean if a field has been set.

### GetPartnerIdentifiersConfig

`func (o *DocumentFlowConfigAllOf) GetPartnerIdentifiersConfig() DocumentFlowConfigAllOfPartnerIdentifiersConfig`

GetPartnerIdentifiersConfig returns the PartnerIdentifiersConfig field if non-nil, zero value otherwise.

### GetPartnerIdentifiersConfigOk

`func (o *DocumentFlowConfigAllOf) GetPartnerIdentifiersConfigOk() (*DocumentFlowConfigAllOfPartnerIdentifiersConfig, bool)`

GetPartnerIdentifiersConfigOk returns a tuple with the PartnerIdentifiersConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerIdentifiersConfig

`func (o *DocumentFlowConfigAllOf) SetPartnerIdentifiersConfig(v DocumentFlowConfigAllOfPartnerIdentifiersConfig)`

SetPartnerIdentifiersConfig sets PartnerIdentifiersConfig field to given value.

### HasPartnerIdentifiersConfig

`func (o *DocumentFlowConfigAllOf) HasPartnerIdentifiersConfig() bool`

HasPartnerIdentifiersConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


