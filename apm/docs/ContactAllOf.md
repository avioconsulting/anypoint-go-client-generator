# ContactAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Field referring the status of the contact (for crud operations) | [optional] [default to "ACTIVE"]
**ContactType** | Pointer to [**ContactType**](ContactType.md) |  | [optional] 

## Methods

### NewContactAllOf

`func NewContactAllOf() *ContactAllOf`

NewContactAllOf instantiates a new ContactAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactAllOfWithDefaults

`func NewContactAllOfWithDefaults() *ContactAllOf`

NewContactAllOfWithDefaults instantiates a new ContactAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContactAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ContactAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContactAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *ContactAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactAllOf) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ContactAllOf) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ContactAllOf) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ContactAllOf) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ContactAllOf) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStatus

`func (o *ContactAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContactAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContactAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContactAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContactType

`func (o *ContactAllOf) GetContactType() ContactType`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *ContactAllOf) GetContactTypeOk() (*ContactType, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *ContactAllOf) SetContactType(v ContactType)`

SetContactType sets ContactType field to given value.

### HasContactType

`func (o *ContactAllOf) HasContactType() bool`

HasContactType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


