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

// X12Config struct for X12Config
type X12Config struct {
	EdiConfig
	ConfigType *string `json:"configType,omitempty"`
}

// NewX12Config instantiates a new X12Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX12Config(formatType string) *X12Config {
	this := X12Config{}
	this.FormatType = formatType
	return &this
}

// NewX12ConfigWithDefaults instantiates a new X12Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX12ConfigWithDefaults() *X12Config {
	this := X12Config{}
	return &this
}

// GetConfigType returns the ConfigType field value if set, zero value otherwise.
func (o *X12Config) GetConfigType() string {
	if o == nil || o.ConfigType == nil {
		var ret string
		return ret
	}
	return *o.ConfigType
}

// GetConfigTypeOk returns a tuple with the ConfigType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X12Config) GetConfigTypeOk() (*string, bool) {
	if o == nil || o.ConfigType == nil {
		return nil, false
	}
	return o.ConfigType, true
}

// HasConfigType returns a boolean if a field has been set.
func (o *X12Config) HasConfigType() bool {
	if o != nil && o.ConfigType != nil {
		return true
	}

	return false
}

// SetConfigType gets a reference to the given string and assigns it to the ConfigType field.
func (o *X12Config) SetConfigType(v string) {
	o.ConfigType = &v
}

func (o X12Config) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEdiConfig, errEdiConfig := json.Marshal(o.EdiConfig)
	if errEdiConfig != nil {
		return []byte{}, errEdiConfig
	}
	errEdiConfig = json.Unmarshal([]byte(serializedEdiConfig), &toSerialize)
	if errEdiConfig != nil {
		return []byte{}, errEdiConfig
	}
	if o.ConfigType != nil {
		toSerialize["configType"] = o.ConfigType
	}
	return json.Marshal(toSerialize)
}

type NullableX12Config struct {
	value *X12Config
	isSet bool
}

func (v NullableX12Config) Get() *X12Config {
	return v.value
}

func (v *NullableX12Config) Set(val *X12Config) {
	v.value = val
	v.isSet = true
}

func (v NullableX12Config) IsSet() bool {
	return v.isSet
}

func (v *NullableX12Config) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableX12Config(val *X12Config) *NullableX12Config {
	return &NullableX12Config{value: val, isSet: true}
}

func (v NullableX12Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableX12Config) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

