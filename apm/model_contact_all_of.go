/*
 * Partner Manager Partners API
 *
 * Anypoint Partner Manager - Partners API
 *
 * API version: v1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apm

import (
	"encoding/json"
)

// ContactAllOf struct for ContactAllOf
type ContactAllOf struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// Field referring the status of the contact (for crud operations)
	Status *string `json:"status,omitempty"`
	ContactType *ContactType `json:"contactType,omitempty"`
}

// NewContactAllOf instantiates a new ContactAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactAllOf() *ContactAllOf {
	this := ContactAllOf{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// NewContactAllOfWithDefaults instantiates a new ContactAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactAllOfWithDefaults() *ContactAllOf {
	this := ContactAllOf{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContactAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContactAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContactAllOf) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContactAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContactAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContactAllOf) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ContactAllOf) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactAllOf) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ContactAllOf) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ContactAllOf) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *ContactAllOf) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactAllOf) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ContactAllOf) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *ContactAllOf) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ContactAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ContactAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ContactAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetContactType returns the ContactType field value if set, zero value otherwise.
func (o *ContactAllOf) GetContactType() ContactType {
	if o == nil || o.ContactType == nil {
		var ret ContactType
		return ret
	}
	return *o.ContactType
}

// GetContactTypeOk returns a tuple with the ContactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactAllOf) GetContactTypeOk() (*ContactType, bool) {
	if o == nil || o.ContactType == nil {
		return nil, false
	}
	return o.ContactType, true
}

// HasContactType returns a boolean if a field has been set.
func (o *ContactAllOf) HasContactType() bool {
	if o != nil && o.ContactType != nil {
		return true
	}

	return false
}

// SetContactType gets a reference to the given ContactType and assigns it to the ContactType field.
func (o *ContactAllOf) SetContactType(v ContactType) {
	o.ContactType = &v
}

func (o ContactAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ContactType != nil {
		toSerialize["contactType"] = o.ContactType
	}
	return json.Marshal(toSerialize)
}

type NullableContactAllOf struct {
	value *ContactAllOf
	isSet bool
}

func (v NullableContactAllOf) Get() *ContactAllOf {
	return v.value
}

func (v *NullableContactAllOf) Set(val *ContactAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContactAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContactAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactAllOf(val *ContactAllOf) *NullableContactAllOf {
	return &NullableContactAllOf{value: val, isSet: true}
}

func (v NullableContactAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


