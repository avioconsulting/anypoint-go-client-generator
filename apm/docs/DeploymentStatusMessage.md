# DeploymentStatusMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**DeploymentStatusType**](DeploymentStatusType.md) |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentStatusMessage

`func NewDeploymentStatusMessage(status DeploymentStatusType, ) *DeploymentStatusMessage`

NewDeploymentStatusMessage instantiates a new DeploymentStatusMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStatusMessageWithDefaults

`func NewDeploymentStatusMessageWithDefaults() *DeploymentStatusMessage`

NewDeploymentStatusMessageWithDefaults instantiates a new DeploymentStatusMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DeploymentStatusMessage) GetStatus() DeploymentStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentStatusMessage) GetStatusOk() (*DeploymentStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentStatusMessage) SetStatus(v DeploymentStatusType)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *DeploymentStatusMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeploymentStatusMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeploymentStatusMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeploymentStatusMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


