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

// EdiConfig struct for EdiConfig
type EdiConfig struct {
	FormatType string `json:"formatType"`
}

// NewEdiConfig instantiates a new EdiConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdiConfig(formatType string) *EdiConfig {
	this := EdiConfig{}
	this.FormatType = formatType
	return &this
}

// NewEdiConfigWithDefaults instantiates a new EdiConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdiConfigWithDefaults() *EdiConfig {
	this := EdiConfig{}
	return &this
}

// GetFormatType returns the FormatType field value
func (o *EdiConfig) GetFormatType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormatType
}

// GetFormatTypeOk returns a tuple with the FormatType field value
// and a boolean to check if the value has been set.
func (o *EdiConfig) GetFormatTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FormatType, true
}

// SetFormatType sets field value
func (o *EdiConfig) SetFormatType(v string) {
	o.FormatType = v
}

func (o EdiConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["formatType"] = o.FormatType
	}
	return json.Marshal(toSerialize)
}

type NullableEdiConfig struct {
	value *EdiConfig
	isSet bool
}

func (v NullableEdiConfig) Get() *EdiConfig {
	return v.value
}

func (v *NullableEdiConfig) Set(val *EdiConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEdiConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEdiConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdiConfig(val *EdiConfig) *NullableEdiConfig {
	return &NullableEdiConfig{value: val, isSet: true}
}

func (v NullableEdiConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdiConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


