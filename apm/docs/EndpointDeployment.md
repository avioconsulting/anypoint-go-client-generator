# EndpointDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Name** | **string** |  | 
**RuntimeDeploymentId** | **int32** |  | 
**Port** | **int32** |  | 

## Methods

### NewEndpointDeployment

`func NewEndpointDeployment(status string, name string, runtimeDeploymentId int32, port int32, ) *EndpointDeployment`

NewEndpointDeployment instantiates a new EndpointDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointDeploymentWithDefaults

`func NewEndpointDeploymentWithDefaults() *EndpointDeployment`

NewEndpointDeploymentWithDefaults instantiates a new EndpointDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndpointDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointDeployment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EndpointDeployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *EndpointDeployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EndpointDeployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EndpointDeployment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetName

`func (o *EndpointDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointDeployment) SetName(v string)`

SetName sets Name field to given value.


### GetRuntimeDeploymentId

`func (o *EndpointDeployment) GetRuntimeDeploymentId() int32`

GetRuntimeDeploymentId returns the RuntimeDeploymentId field if non-nil, zero value otherwise.

### GetRuntimeDeploymentIdOk

`func (o *EndpointDeployment) GetRuntimeDeploymentIdOk() (*int32, bool)`

GetRuntimeDeploymentIdOk returns a tuple with the RuntimeDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeDeploymentId

`func (o *EndpointDeployment) SetRuntimeDeploymentId(v int32)`

SetRuntimeDeploymentId sets RuntimeDeploymentId field to given value.


### GetPort

`func (o *EndpointDeployment) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EndpointDeployment) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EndpointDeployment) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


