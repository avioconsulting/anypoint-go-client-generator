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

// DocumentMappingConfig Document Mapping configuration for document flow.
type DocumentMappingConfig struct {
	Id *string `json:"id,omitempty"`
	MappingType string `json:"mappingType"`
	// required only when the mapppingType is not MAP_API
	MappingContent *string `json:"mappingContent,omitempty"`
	// For DWL_SCRIPT is Inline , DWL_FILE is the file name, and for MAP_API is TargetAPI endpoint Id.
	MappingSourceRef string `json:"mappingSourceRef"`
}

// NewDocumentMappingConfig instantiates a new DocumentMappingConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentMappingConfig(mappingType string, mappingSourceRef string) *DocumentMappingConfig {
	this := DocumentMappingConfig{}
	this.MappingType = mappingType
	this.MappingSourceRef = mappingSourceRef
	return &this
}

// NewDocumentMappingConfigWithDefaults instantiates a new DocumentMappingConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentMappingConfigWithDefaults() *DocumentMappingConfig {
	this := DocumentMappingConfig{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentMappingConfig) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMappingConfig) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentMappingConfig) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocumentMappingConfig) SetId(v string) {
	o.Id = &v
}

// GetMappingType returns the MappingType field value
func (o *DocumentMappingConfig) GetMappingType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MappingType
}

// GetMappingTypeOk returns a tuple with the MappingType field value
// and a boolean to check if the value has been set.
func (o *DocumentMappingConfig) GetMappingTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MappingType, true
}

// SetMappingType sets field value
func (o *DocumentMappingConfig) SetMappingType(v string) {
	o.MappingType = v
}

// GetMappingContent returns the MappingContent field value if set, zero value otherwise.
func (o *DocumentMappingConfig) GetMappingContent() string {
	if o == nil || o.MappingContent == nil {
		var ret string
		return ret
	}
	return *o.MappingContent
}

// GetMappingContentOk returns a tuple with the MappingContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMappingConfig) GetMappingContentOk() (*string, bool) {
	if o == nil || o.MappingContent == nil {
		return nil, false
	}
	return o.MappingContent, true
}

// HasMappingContent returns a boolean if a field has been set.
func (o *DocumentMappingConfig) HasMappingContent() bool {
	if o != nil && o.MappingContent != nil {
		return true
	}

	return false
}

// SetMappingContent gets a reference to the given string and assigns it to the MappingContent field.
func (o *DocumentMappingConfig) SetMappingContent(v string) {
	o.MappingContent = &v
}

// GetMappingSourceRef returns the MappingSourceRef field value
func (o *DocumentMappingConfig) GetMappingSourceRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MappingSourceRef
}

// GetMappingSourceRefOk returns a tuple with the MappingSourceRef field value
// and a boolean to check if the value has been set.
func (o *DocumentMappingConfig) GetMappingSourceRefOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MappingSourceRef, true
}

// SetMappingSourceRef sets field value
func (o *DocumentMappingConfig) SetMappingSourceRef(v string) {
	o.MappingSourceRef = v
}

func (o DocumentMappingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["mappingType"] = o.MappingType
	}
	if o.MappingContent != nil {
		toSerialize["mappingContent"] = o.MappingContent
	}
	if true {
		toSerialize["mappingSourceRef"] = o.MappingSourceRef
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentMappingConfig struct {
	value *DocumentMappingConfig
	isSet bool
}

func (v NullableDocumentMappingConfig) Get() *DocumentMappingConfig {
	return v.value
}

func (v *NullableDocumentMappingConfig) Set(val *DocumentMappingConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentMappingConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentMappingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentMappingConfig(val *DocumentMappingConfig) *NullableDocumentMappingConfig {
	return &NullableDocumentMappingConfig{value: val, isSet: true}
}

func (v NullableDocumentMappingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentMappingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


