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

// DocumentFlowConfigAllOfPartnerIdentifiersConfig struct for DocumentFlowConfigAllOfPartnerIdentifiersConfig
type DocumentFlowConfigAllOfPartnerIdentifiersConfig struct {
	UseAllHostIdentifiers *bool `json:"useAllHostIdentifiers,omitempty"`
	UseAllPartnerIdentifiers *bool `json:"useAllPartnerIdentifiers,omitempty"`
	Identifiers *[]PartnerIdentifiersConfigIdentifiers `json:"identifiers,omitempty"`
}

// NewDocumentFlowConfigAllOfPartnerIdentifiersConfig instantiates a new DocumentFlowConfigAllOfPartnerIdentifiersConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentFlowConfigAllOfPartnerIdentifiersConfig() *DocumentFlowConfigAllOfPartnerIdentifiersConfig {
	this := DocumentFlowConfigAllOfPartnerIdentifiersConfig{}
	return &this
}

// NewDocumentFlowConfigAllOfPartnerIdentifiersConfigWithDefaults instantiates a new DocumentFlowConfigAllOfPartnerIdentifiersConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentFlowConfigAllOfPartnerIdentifiersConfigWithDefaults() *DocumentFlowConfigAllOfPartnerIdentifiersConfig {
	this := DocumentFlowConfigAllOfPartnerIdentifiersConfig{}
	return &this
}

// GetUseAllHostIdentifiers returns the UseAllHostIdentifiers field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOfPartnerIdentifiersConfig) GetUseAllHostIdentifiers() bool {
	if o == nil || o.UseAllHostIdentifiers == nil {
		var ret bool
		return ret
	}
	return *o.UseAllHostIdentifiers
}

// GetUseAllHostIdentifiersOk returns a tuple with the UseAllHostIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOfPartnerIdentifiersConfig) GetUseAllHostIdentifiersOk() (*bool, bool) {
	if o == nil || o.UseAllHostIdentifiers == nil {
		return nil, false
	}
	return o.UseAllHostIdentifiers, true
}

// HasUseAllHostIdentifiers returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOfPartnerIdentifiersConfig) HasUseAllHostIdentifiers() bool {
	if o != nil && o.UseAllHostIdentifiers != nil {
		return true
	}

	return false
}

// SetUseAllHostIdentifiers gets a reference to the given bool and assigns it to the UseAllHostIdentifiers field.
func (o *DocumentFlowConfigAllOfPartnerIdentifiersConfig) SetUseAllHostIdentifiers(v bool) {
	o.UseAllHostIdentifiers = &v
}

// GetUseAllPartnerIdentifiers returns the UseAllPartnerIdentifiers field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOfPartnerIdentifiersConfig) GetUseAllPartnerIdentifiers() bool {
	if o == nil || o.UseAllPartnerIdentifiers == nil {
		var ret bool
		return ret
	}
	return *o.UseAllPartnerIdentifiers
}

// GetUseAllPartnerIdentifiersOk returns a tuple with the UseAllPartnerIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOfPartnerIdentifiersConfig) GetUseAllPartnerIdentifiersOk() (*bool, bool) {
	if o == nil || o.UseAllPartnerIdentifiers == nil {
		return nil, false
	}
	return o.UseAllPartnerIdentifiers, true
}

// HasUseAllPartnerIdentifiers returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOfPartnerIdentifiersConfig) HasUseAllPartnerIdentifiers() bool {
	if o != nil && o.UseAllPartnerIdentifiers != nil {
		return true
	}

	return false
}

// SetUseAllPartnerIdentifiers gets a reference to the given bool and assigns it to the UseAllPartnerIdentifiers field.
func (o *DocumentFlowConfigAllOfPartnerIdentifiersConfig) SetUseAllPartnerIdentifiers(v bool) {
	o.UseAllPartnerIdentifiers = &v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOfPartnerIdentifiersConfig) GetIdentifiers() []PartnerIdentifiersConfigIdentifiers {
	if o == nil || o.Identifiers == nil {
		var ret []PartnerIdentifiersConfigIdentifiers
		return ret
	}
	return *o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOfPartnerIdentifiersConfig) GetIdentifiersOk() (*[]PartnerIdentifiersConfigIdentifiers, bool) {
	if o == nil || o.Identifiers == nil {
		return nil, false
	}
	return o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOfPartnerIdentifiersConfig) HasIdentifiers() bool {
	if o != nil && o.Identifiers != nil {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []PartnerIdentifiersConfigIdentifiers and assigns it to the Identifiers field.
func (o *DocumentFlowConfigAllOfPartnerIdentifiersConfig) SetIdentifiers(v []PartnerIdentifiersConfigIdentifiers) {
	o.Identifiers = &v
}

func (o DocumentFlowConfigAllOfPartnerIdentifiersConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UseAllHostIdentifiers != nil {
		toSerialize["useAllHostIdentifiers"] = o.UseAllHostIdentifiers
	}
	if o.UseAllPartnerIdentifiers != nil {
		toSerialize["useAllPartnerIdentifiers"] = o.UseAllPartnerIdentifiers
	}
	if o.Identifiers != nil {
		toSerialize["identifiers"] = o.Identifiers
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentFlowConfigAllOfPartnerIdentifiersConfig struct {
	value *DocumentFlowConfigAllOfPartnerIdentifiersConfig
	isSet bool
}

func (v NullableDocumentFlowConfigAllOfPartnerIdentifiersConfig) Get() *DocumentFlowConfigAllOfPartnerIdentifiersConfig {
	return v.value
}

func (v *NullableDocumentFlowConfigAllOfPartnerIdentifiersConfig) Set(val *DocumentFlowConfigAllOfPartnerIdentifiersConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentFlowConfigAllOfPartnerIdentifiersConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentFlowConfigAllOfPartnerIdentifiersConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentFlowConfigAllOfPartnerIdentifiersConfig(val *DocumentFlowConfigAllOfPartnerIdentifiersConfig) *NullableDocumentFlowConfigAllOfPartnerIdentifiersConfig {
	return &NullableDocumentFlowConfigAllOfPartnerIdentifiersConfig{value: val, isSet: true}
}

func (v NullableDocumentFlowConfigAllOfPartnerIdentifiersConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentFlowConfigAllOfPartnerIdentifiersConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


