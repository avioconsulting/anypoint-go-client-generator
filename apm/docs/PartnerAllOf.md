# PartnerAllOf

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

### NewPartnerAllOf

`func NewPartnerAllOf(id string, name string, environmentId string, hostFlag bool, ) *PartnerAllOf`

NewPartnerAllOf instantiates a new PartnerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerAllOfWithDefaults

`func NewPartnerAllOfWithDefaults() *PartnerAllOf`

NewPartnerAllOfWithDefaults instantiates a new PartnerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartnerAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnerAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnerAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PartnerAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PartnerAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PartnerAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PartnerAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PartnerAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *PartnerAllOf) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *PartnerAllOf) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *PartnerAllOf) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *PartnerAllOf) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetStatus

`func (o *PartnerAllOf) GetStatus() TimedStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PartnerAllOf) GetStatusOk() (*TimedStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PartnerAllOf) SetStatus(v TimedStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PartnerAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *PartnerAllOf) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *PartnerAllOf) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *PartnerAllOf) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetHostFlag

`func (o *PartnerAllOf) GetHostFlag() bool`

GetHostFlag returns the HostFlag field if non-nil, zero value otherwise.

### GetHostFlagOk

`func (o *PartnerAllOf) GetHostFlagOk() (*bool, bool)`

GetHostFlagOk returns a tuple with the HostFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostFlag

`func (o *PartnerAllOf) SetHostFlag(v bool)`

SetHostFlag sets HostFlag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


