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

// UserDetail User details from core services
type UserDetail struct {
	UserId *string `json:"userId,omitempty"`
	Firstname *string `json:"firstname,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

// NewUserDetail instantiates a new UserDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDetail() *UserDetail {
	this := UserDetail{}
	return &this
}

// NewUserDetailWithDefaults instantiates a new UserDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDetailWithDefaults() *UserDetail {
	this := UserDetail{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UserDetail) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetail) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UserDetail) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *UserDetail) SetUserId(v string) {
	o.UserId = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *UserDetail) GetFirstname() string {
	if o == nil || o.Firstname == nil {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetail) GetFirstnameOk() (*string, bool) {
	if o == nil || o.Firstname == nil {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *UserDetail) HasFirstname() bool {
	if o != nil && o.Firstname != nil {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *UserDetail) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *UserDetail) GetLastname() string {
	if o == nil || o.Lastname == nil {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetail) GetLastnameOk() (*string, bool) {
	if o == nil || o.Lastname == nil {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *UserDetail) HasLastname() bool {
	if o != nil && o.Lastname != nil {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *UserDetail) SetLastname(v string) {
	o.Lastname = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *UserDetail) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetail) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *UserDetail) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *UserDetail) SetUserName(v string) {
	o.UserName = &v
}

func (o UserDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.Firstname != nil {
		toSerialize["firstname"] = o.Firstname
	}
	if o.Lastname != nil {
		toSerialize["lastname"] = o.Lastname
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	return json.Marshal(toSerialize)
}

type NullableUserDetail struct {
	value *UserDetail
	isSet bool
}

func (v NullableUserDetail) Get() *UserDetail {
	return v.value
}

func (v *NullableUserDetail) Set(val *UserDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDetail(val *UserDetail) *NullableUserDetail {
	return &NullableUserDetail{value: val, isSet: true}
}

func (v NullableUserDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

