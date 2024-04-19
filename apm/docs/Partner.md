# Partner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**WebsiteUrl** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TimedStatus**](TimedStatus.md) |  | [optional] 
**EnvironmentId** | **string** |  | 
**HostFlag** | **bool** |  | [default to false]

## Methods

### NewPartner

`func NewPartner(id string, name string, environmentId string, hostFlag bool, ) *Partner`

NewPartner instantiates a new Partner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerWithDefaults

`func NewPartnerWithDefaults() *Partner`

NewPartnerWithDefaults instantiates a new Partner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Partner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Partner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Partner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Partner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Partner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Partner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Partner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Partner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Partner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Partner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *Partner) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *Partner) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *Partner) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *Partner) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetStatus

`func (o *Partner) GetStatus() TimedStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Partner) GetStatusOk() (*TimedStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Partner) SetStatus(v TimedStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Partner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *Partner) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *Partner) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *Partner) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetHostFlag

`func (o *Partner) GetHostFlag() bool`

GetHostFlag returns the HostFlag field if non-nil, zero value otherwise.

### GetHostFlagOk

`func (o *Partner) GetHostFlagOk() (*bool, bool)`

GetHostFlagOk returns a tuple with the HostFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostFlag

`func (o *Partner) SetHostFlag(v bool)`

SetHostFlag sets HostFlag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


