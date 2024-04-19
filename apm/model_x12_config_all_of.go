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

// X12ConfigAllOf struct for X12ConfigAllOf
type X12ConfigAllOf struct {
	ConfigType *string `json:"configType,omitempty"`
}

// NewX12ConfigAllOf instantiates a new X12ConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX12ConfigAllOf() *X12ConfigAllOf {
	this := X12ConfigAllOf{}
	return &this
}

// NewX12ConfigAllOfWithDefaults instantiates a new X12ConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX12ConfigAllOfWithDefaults() *X12ConfigAllOf {
	this := X12ConfigAllOf{}
	return &this
}

// GetConfigType returns the ConfigType field value if set, zero value otherwise.
func (o *X12ConfigAllOf) GetConfigType() string {
	if o == nil || o.ConfigType == nil {
		var ret string
		return ret
	}
	return *o.ConfigType
}

// GetConfigTypeOk returns a tuple with the ConfigType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X12ConfigAllOf) GetConfigTypeOk() (*string, bool) {
	if o == nil || o.ConfigType == nil {
		return nil, false
	}
	return o.ConfigType, true
}

// HasConfigType returns a boolean if a field has been set.
func (o *X12ConfigAllOf) HasConfigType() bool {
	if o != nil && o.ConfigType != nil {
		return true
	}

	return false
}

// SetConfigType gets a reference to the given string and assigns it to the ConfigType field.
func (o *X12ConfigAllOf) SetConfigType(v string) {
	o.ConfigType = &v
}

func (o X12ConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfigType != nil {
		toSerialize["configType"] = o.ConfigType
	}
	return json.Marshal(toSerialize)
}

type NullableX12ConfigAllOf struct {
	value *X12ConfigAllOf
	isSet bool
}

func (v NullableX12ConfigAllOf) Get() *X12ConfigAllOf {
	return v.value
}

func (v *NullableX12ConfigAllOf) Set(val *X12ConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableX12ConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableX12ConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableX12ConfigAllOf(val *X12ConfigAllOf) *NullableX12ConfigAllOf {
	return &NullableX12ConfigAllOf{value: val, isSet: true}
}

func (v NullableX12ConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableX12ConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


