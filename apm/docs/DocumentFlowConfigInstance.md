# DocumentFlowConfigInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Status** | [**DeploymentStatusType**](DeploymentStatusType.md) |  | 
**ConfigId** | **string** |  | 
**Direction** | [**DirectionType**](DirectionType.md) |  | 
**Version** | **int32** |  | 
**PartnerToId** | **string** |  | 
**PartnerFromId** | **string** |  | 
**PreProcessingEndpoint** | Pointer to [**BaseEndpoint**](BaseEndpoint.md) |  | [optional] 
**ReceivingEndpoint** | [**BaseEndpoint**](BaseEndpoint.md) |  | 
**ReceivingAckEndpoint** | Pointer to [**BaseEndpoint**](BaseEndpoint.md) |  | [optional] 
**TargetEndpoint** | [**BaseEndpoint**](BaseEndpoint.md) |  | 
**SourceDocType** | [**BaseEdiDocument**](BaseEdiDocument.md) |  | 
**TargetDocType** | [**BaseEdiDocument**](BaseEdiDocument.md) |  | 
**DocumentMapping** | [**[]DocumentMappingConfig**](DocumentMappingConfig.md) |  | 
**MessageStatus** | Pointer to **string** |  | [optional] 
**ReceivingAckConfig** | Pointer to [**BaseAckConfig**](BaseAckConfig.md) |  | [optional] 

## Methods

### NewDocumentFlowConfigInstance

`func NewDocumentFlowConfigInstance(name string, status DeploymentStatusType, configId string, direction DirectionType, version int32, partnerToId string, partnerFromId string, receivingEndpoint BaseEndpoint, targetEndpoint BaseEndpoint, sourceDocType BaseEdiDocument, targetDocType BaseEdiDocument, documentMapping []DocumentMappingConfig, ) *DocumentFlowConfigInstance`

NewDocumentFlowConfigInstance instantiates a new DocumentFlowConfigInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentFlowConfigInstanceWithDefaults

`func NewDocumentFlowConfigInstanceWithDefaults() *DocumentFlowConfigInstance`

NewDocumentFlowConfigInstanceWithDefaults instantiates a new DocumentFlowConfigInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentFlowConfigInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentFlowConfigInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentFlowConfigInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentFlowConfigInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DocumentFlowConfigInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentFlowConfigInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentFlowConfigInstance) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DocumentFlowConfigInstance) GetStatus() DeploymentStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentFlowConfigInstance) GetStatusOk() (*DeploymentStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentFlowConfigInstance) SetStatus(v DeploymentStatusType)`

SetStatus sets Status field to given value.


### GetConfigId

`func (o *DocumentFlowConfigInstance) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *DocumentFlowConfigInstance) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *DocumentFlowConfigInstance) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.


### GetDirection

`func (o *DocumentFlowConfigInstance) GetDirection() DirectionType`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *DocumentFlowConfigInstance) GetDirectionOk() (*DirectionType, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *DocumentFlowConfigInstance) SetDirection(v DirectionType)`

SetDirection sets Direction field to given value.


### GetVersion

`func (o *DocumentFlowConfigInstance) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DocumentFlowConfigInstance) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DocumentFlowConfigInstance) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetPartnerToId

`func (o *DocumentFlowConfigInstance) GetPartnerToId() string`

GetPartnerToId returns the PartnerToId field if non-nil, zero value otherwise.

### GetPartnerToIdOk

`func (o *DocumentFlowConfigInstance) GetPartnerToIdOk() (*string, bool)`

GetPartnerToIdOk returns a tuple with the PartnerToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerToId

`func (o *DocumentFlowConfigInstance) SetPartnerToId(v string)`

SetPartnerToId sets PartnerToId field to given value.


### GetPartnerFromId

`func (o *DocumentFlowConfigInstance) GetPartnerFromId() string`

GetPartnerFromId returns the PartnerFromId field if non-nil, zero value otherwise.

### GetPartnerFromIdOk

`func (o *DocumentFlowConfigInstance) GetPartnerFromIdOk() (*string, bool)`

GetPartnerFromIdOk returns a tuple with the PartnerFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerFromId

`func (o *DocumentFlowConfigInstance) SetPartnerFromId(v string)`

SetPartnerFromId sets PartnerFromId field to given value.


### GetPreProcessingEndpoint

`func (o *DocumentFlowConfigInstance) GetPreProcessingEndpoint() BaseEndpoint`

GetPreProcessingEndpoint returns the PreProcessingEndpoint field if non-nil, zero value otherwise.

### GetPreProcessingEndpointOk

`func (o *DocumentFlowConfigInstance) GetPreProcessingEndpointOk() (*BaseEndpoint, bool)`

GetPreProcessingEndpointOk returns a tuple with the PreProcessingEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreProcessingEndpoint

`func (o *DocumentFlowConfigInstance) SetPreProcessingEndpoint(v BaseEndpoint)`

SetPreProcessingEndpoint sets PreProcessingEndpoint field to given value.

### HasPreProcessingEndpoint

`func (o *DocumentFlowConfigInstance) HasPreProcessingEndpoint() bool`

HasPreProcessingEndpoint returns a boolean if a field has been set.

### GetReceivingEndpoint

`func (o *DocumentFlowConfigInstance) GetReceivingEndpoint() BaseEndpoint`

GetReceivingEndpoint returns the ReceivingEndpoint field if non-nil, zero value otherwise.

### GetReceivingEndpointOk

`func (o *DocumentFlowConfigInstance) GetReceivingEndpointOk() (*BaseEndpoint, bool)`

GetReceivingEndpointOk returns a tuple with the ReceivingEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingEndpoint

`func (o *DocumentFlowConfigInstance) SetReceivingEndpoint(v BaseEndpoint)`

SetReceivingEndpoint sets ReceivingEndpoint field to given value.


### GetReceivingAckEndpoint

`func (o *DocumentFlowConfigInstance) GetReceivingAckEndpoint() BaseEndpoint`

GetReceivingAckEndpoint returns the ReceivingAckEndpoint field if non-nil, zero value otherwise.

### GetReceivingAckEndpointOk

`func (o *DocumentFlowConfigInstance) GetReceivingAckEndpointOk() (*BaseEndpoint, bool)`

GetReceivingAckEndpointOk returns a tuple with the ReceivingAckEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAckEndpoint

`func (o *DocumentFlowConfigInstance) SetReceivingAckEndpoint(v BaseEndpoint)`

SetReceivingAckEndpoint sets ReceivingAckEndpoint field to given value.

### HasReceivingAckEndpoint

`func (o *DocumentFlowConfigInstance) HasReceivingAckEndpoint() bool`

HasReceivingAckEndpoint returns a boolean if a field has been set.

### GetTargetEndpoint

`func (o *DocumentFlowConfigInstance) GetTargetEndpoint() BaseEndpoint`

GetTargetEndpoint returns the TargetEndpoint field if non-nil, zero value otherwise.

### GetTargetEndpointOk

`func (o *DocumentFlowConfigInstance) GetTargetEndpointOk() (*BaseEndpoint, bool)`

GetTargetEndpointOk returns a tuple with the TargetEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEndpoint

`func (o *DocumentFlowConfigInstance) SetTargetEndpoint(v BaseEndpoint)`

SetTargetEndpoint sets TargetEndpoint field to given value.


### GetSourceDocType

`func (o *DocumentFlowConfigInstance) GetSourceDocType() BaseEdiDocument`

GetSourceDocType returns the SourceDocType field if non-nil, zero value otherwise.

### GetSourceDocTypeOk

`func (o *DocumentFlowConfigInstance) GetSourceDocTypeOk() (*BaseEdiDocument, bool)`

GetSourceDocTypeOk returns a tuple with the SourceDocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocType

`func (o *DocumentFlowConfigInstance) SetSourceDocType(v BaseEdiDocument)`

SetSourceDocType sets SourceDocType field to given value.


### GetTargetDocType

`func (o *DocumentFlowConfigInstance) GetTargetDocType() BaseEdiDocument`

GetTargetDocType returns the TargetDocType field if non-nil, zero value otherwise.

### GetTargetDocTypeOk

`func (o *DocumentFlowConfigInstance) GetTargetDocTypeOk() (*BaseEdiDocument, bool)`

GetTargetDocTypeOk returns a tuple with the TargetDocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDocType

`func (o *DocumentFlowConfigInstance) SetTargetDocType(v BaseEdiDocument)`

SetTargetDocType sets TargetDocType field to given value.


### GetDocumentMapping

`func (o *DocumentFlowConfigInstance) GetDocumentMapping() []DocumentMappingConfig`

GetDocumentMapping returns the DocumentMapping field if non-nil, zero value otherwise.

### GetDocumentMappingOk

`func (o *DocumentFlowConfigInstance) GetDocumentMappingOk() (*[]DocumentMappingConfig, bool)`

GetDocumentMappingOk returns a tuple with the DocumentMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMapping

`func (o *DocumentFlowConfigInstance) SetDocumentMapping(v []DocumentMappingConfig)`

SetDocumentMapping sets DocumentMapping field to given value.


### GetMessageStatus

`func (o *DocumentFlowConfigInstance) GetMessageStatus() string`

GetMessageStatus returns the MessageStatus field if non-nil, zero value otherwise.

### GetMessageStatusOk

`func (o *DocumentFlowConfigInstance) GetMessageStatusOk() (*string, bool)`

GetMessageStatusOk returns a tuple with the MessageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageStatus

`func (o *DocumentFlowConfigInstance) SetMessageStatus(v string)`

SetMessageStatus sets MessageStatus field to given value.

### HasMessageStatus

`func (o *DocumentFlowConfigInstance) HasMessageStatus() bool`

HasMessageStatus returns a boolean if a field has been set.

### GetReceivingAckConfig

`func (o *DocumentFlowConfigInstance) GetReceivingAckConfig() BaseAckConfig`

GetReceivingAckConfig returns the ReceivingAckConfig field if non-nil, zero value otherwise.

### GetReceivingAckConfigOk

`func (o *DocumentFlowConfigInstance) GetReceivingAckConfigOk() (*BaseAckConfig, bool)`

GetReceivingAckConfigOk returns a tuple with the ReceivingAckConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAckConfig

`func (o *DocumentFlowConfigInstance) SetReceivingAckConfig(v BaseAckConfig)`

SetReceivingAckConfig sets ReceivingAckConfig field to given value.

### HasReceivingAckConfig

`func (o *DocumentFlowConfigInstance) HasReceivingAckConfig() bool`

HasReceivingAckConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


