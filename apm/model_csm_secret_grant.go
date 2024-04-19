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

// CSMSecretGrant struct for CSMSecretGrant
type CSMSecretGrant struct {
	Path string `json:"path"`
	AccessGrant string `json:"accessGrant"`
}

// NewCSMSecretGrant instantiates a new CSMSecretGrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSMSecretGrant(path string, accessGrant string) *CSMSecretGrant {
	this := CSMSecretGrant{}
	this.Path = path
	this.AccessGrant = accessGrant
	return &this
}

// NewCSMSecretGrantWithDefaults instantiates a new CSMSecretGrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSMSecretGrantWithDefaults() *CSMSecretGrant {
	this := CSMSecretGrant{}
	return &this
}

// GetPath returns the Path field value
func (o *CSMSecretGrant) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *CSMSecretGrant) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *CSMSecretGrant) SetPath(v string) {
	o.Path = v
}

// GetAccessGrant returns the AccessGrant field value
func (o *CSMSecretGrant) GetAccessGrant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessGrant
}

// GetAccessGrantOk returns a tuple with the AccessGrant field value
// and a boolean to check if the value has been set.
func (o *CSMSecretGrant) GetAccessGrantOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessGrant, true
}

// SetAccessGrant sets field value
func (o *CSMSecretGrant) SetAccessGrant(v string) {
	o.AccessGrant = v
}

func (o CSMSecretGrant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["accessGrant"] = o.AccessGrant
	}
	return json.Marshal(toSerialize)
}

type NullableCSMSecretGrant struct {
	value *CSMSecretGrant
	isSet bool
}

func (v NullableCSMSecretGrant) Get() *CSMSecretGrant {
	return v.value
}

func (v *NullableCSMSecretGrant) Set(val *CSMSecretGrant) {
	v.value = val
	v.isSet = true
}

func (v NullableCSMSecretGrant) IsSet() bool {
	return v.isSet
}

func (v *NullableCSMSecretGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSMSecretGrant(val *CSMSecretGrant) *NullableCSMSecretGrant {
	return &NullableCSMSecretGrant{value: val, isSet: true}
}

func (v NullableCSMSecretGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSMSecretGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

