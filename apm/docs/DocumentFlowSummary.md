# DocumentFlowSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to [**DirectionType**](DirectionType.md) |  | [optional] 
**PartnerFrom** | Pointer to [**Partner**](Partner.md) |  | [optional] 
**PartnerTo** | Pointer to [**Partner**](Partner.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**DeploymentSummary** | Pointer to [**[]DocumentFlowDeploymentSummary**](DocumentFlowDeploymentSummary.md) |  | [optional] 

## Methods

### NewDocumentFlowSummary

`func NewDocumentFlowSummary() *DocumentFlowSummary`

NewDocumentFlowSummary instantiates a new DocumentFlowSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentFlowSummaryWithDefaults

`func NewDocumentFlowSummaryWithDefaults() *DocumentFlowSummary`

NewDocumentFlowSummaryWithDefaults instantiates a new DocumentFlowSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentFlowSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentFlowSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentFlowSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentFlowSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DocumentFlowSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentFlowSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentFlowSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentFlowSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDirection

`func (o *DocumentFlowSummary) GetDirection() DirectionType`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *DocumentFlowSummary) GetDirectionOk() (*DirectionType, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *DocumentFlowSummary) SetDirection(v DirectionType)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *DocumentFlowSummary) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetPartnerFrom

`func (o *DocumentFlowSummary) GetPartnerFrom() Partner`

GetPartnerFrom returns the PartnerFrom field if non-nil, zero value otherwise.

### GetPartnerFromOk

`func (o *DocumentFlowSummary) GetPartnerFromOk() (*Partner, bool)`

GetPartnerFromOk returns a tuple with the PartnerFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerFrom

`func (o *DocumentFlowSummary) SetPartnerFrom(v Partner)`

SetPartnerFrom sets PartnerFrom field to given value.

### HasPartnerFrom

`func (o *DocumentFlowSummary) HasPartnerFrom() bool`

HasPartnerFrom returns a boolean if a field has been set.

### GetPartnerTo

`func (o *DocumentFlowSummary) GetPartnerTo() Partner`

GetPartnerTo returns the PartnerTo field if non-nil, zero value otherwise.

### GetPartnerToOk

`func (o *DocumentFlowSummary) GetPartnerToOk() (*Partner, bool)`

GetPartnerToOk returns a tuple with the PartnerTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerTo

`func (o *DocumentFlowSummary) SetPartnerTo(v Partner)`

SetPartnerTo sets PartnerTo field to given value.

### HasPartnerTo

`func (o *DocumentFlowSummary) HasPartnerTo() bool`

HasPartnerTo returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocumentFlowSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentFlowSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentFlowSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocumentFlowSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DocumentFlowSummary) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DocumentFlowSummary) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DocumentFlowSummary) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DocumentFlowSummary) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DocumentFlowSummary) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DocumentFlowSummary) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DocumentFlowSummary) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DocumentFlowSummary) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DocumentFlowSummary) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DocumentFlowSummary) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DocumentFlowSummary) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DocumentFlowSummary) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetDeploymentSummary

`func (o *DocumentFlowSummary) GetDeploymentSummary() []DocumentFlowDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *DocumentFlowSummary) GetDeploymentSummaryOk() (*[]DocumentFlowDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSummary

`func (o *DocumentFlowSummary) SetDeploymentSummary(v []DocumentFlowDeploymentSummary)`

SetDeploymentSummary sets DeploymentSummary field to given value.

### HasDeploymentSummary

`func (o *DocumentFlowSummary) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


