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

// EdiFormatConfigurationTerminatorDelimiter struct for EdiFormatConfigurationTerminatorDelimiter
type EdiFormatConfigurationTerminatorDelimiter struct {
	SegmentTerminatorCharacter *string `json:"segmentTerminatorCharacter,omitempty"`
	DataElementDelimiter *string `json:"dataElementDelimiter,omitempty"`
	StringSubstituionCharacter *string `json:"stringSubstituionCharacter,omitempty"`
}

// NewEdiFormatConfigurationTerminatorDelimiter instantiates a new EdiFormatConfigurationTerminatorDelimiter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdiFormatConfigurationTerminatorDelimiter() *EdiFormatConfigurationTerminatorDelimiter {
	this := EdiFormatConfigurationTerminatorDelimiter{}
	return &this
}

// NewEdiFormatConfigurationTerminatorDelimiterWithDefaults instantiates a new EdiFormatConfigurationTerminatorDelimiter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdiFormatConfigurationTerminatorDelimiterWithDefaults() *EdiFormatConfigurationTerminatorDelimiter {
	this := EdiFormatConfigurationTerminatorDelimiter{}
	return &this
}

// GetSegmentTerminatorCharacter returns the SegmentTerminatorCharacter field value if set, zero value otherwise.
func (o *EdiFormatConfigurationTerminatorDelimiter) GetSegmentTerminatorCharacter() string {
	if o == nil || o.SegmentTerminatorCharacter == nil {
		var ret string
		return ret
	}
	return *o.SegmentTerminatorCharacter
}

// GetSegmentTerminatorCharacterOk returns a tuple with the SegmentTerminatorCharacter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationTerminatorDelimiter) GetSegmentTerminatorCharacterOk() (*string, bool) {
	if o == nil || o.SegmentTerminatorCharacter == nil {
		return nil, false
	}
	return o.SegmentTerminatorCharacter, true
}

// HasSegmentTerminatorCharacter returns a boolean if a field has been set.
func (o *EdiFormatConfigurationTerminatorDelimiter) HasSegmentTerminatorCharacter() bool {
	if o != nil && o.SegmentTerminatorCharacter != nil {
		return true
	}

	return false
}

// SetSegmentTerminatorCharacter gets a reference to the given string and assigns it to the SegmentTerminatorCharacter field.
func (o *EdiFormatConfigurationTerminatorDelimiter) SetSegmentTerminatorCharacter(v string) {
	o.SegmentTerminatorCharacter = &v
}

// GetDataElementDelimiter returns the DataElementDelimiter field value if set, zero value otherwise.
func (o *EdiFormatConfigurationTerminatorDelimiter) GetDataElementDelimiter() string {
	if o == nil || o.DataElementDelimiter == nil {
		var ret string
		return ret
	}
	return *o.DataElementDelimiter
}

// GetDataElementDelimiterOk returns a tuple with the DataElementDelimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationTerminatorDelimiter) GetDataElementDelimiterOk() (*string, bool) {
	if o == nil || o.DataElementDelimiter == nil {
		return nil, false
	}
	return o.DataElementDelimiter, true
}

// HasDataElementDelimiter returns a boolean if a field has been set.
func (o *EdiFormatConfigurationTerminatorDelimiter) HasDataElementDelimiter() bool {
	if o != nil && o.DataElementDelimiter != nil {
		return true
	}

	return false
}

// SetDataElementDelimiter gets a reference to the given string and assigns it to the DataElementDelimiter field.
func (o *EdiFormatConfigurationTerminatorDelimiter) SetDataElementDelimiter(v string) {
	o.DataElementDelimiter = &v
}

// GetStringSubstituionCharacter returns the StringSubstituionCharacter field value if set, zero value otherwise.
func (o *EdiFormatConfigurationTerminatorDelimiter) GetStringSubstituionCharacter() string {
	if o == nil || o.StringSubstituionCharacter == nil {
		var ret string
		return ret
	}
	return *o.StringSubstituionCharacter
}

// GetStringSubstituionCharacterOk returns a tuple with the StringSubstituionCharacter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationTerminatorDelimiter) GetStringSubstituionCharacterOk() (*string, bool) {
	if o == nil || o.StringSubstituionCharacter == nil {
		return nil, false
	}
	return o.StringSubstituionCharacter, true
}

// HasStringSubstituionCharacter returns a boolean if a field has been set.
func (o *EdiFormatConfigurationTerminatorDelimiter) HasStringSubstituionCharacter() bool {
	if o != nil && o.StringSubstituionCharacter != nil {
		return true
	}

	return false
}

// SetStringSubstituionCharacter gets a reference to the given string and assigns it to the StringSubstituionCharacter field.
func (o *EdiFormatConfigurationTerminatorDelimiter) SetStringSubstituionCharacter(v string) {
	o.StringSubstituionCharacter = &v
}

func (o EdiFormatConfigurationTerminatorDelimiter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SegmentTerminatorCharacter != nil {
		toSerialize["segmentTerminatorCharacter"] = o.SegmentTerminatorCharacter
	}
	if o.DataElementDelimiter != nil {
		toSerialize["dataElementDelimiter"] = o.DataElementDelimiter
	}
	if o.StringSubstituionCharacter != nil {
		toSerialize["stringSubstituionCharacter"] = o.StringSubstituionCharacter
	}
	return json.Marshal(toSerialize)
}

type NullableEdiFormatConfigurationTerminatorDelimiter struct {
	value *EdiFormatConfigurationTerminatorDelimiter
	isSet bool
}

func (v NullableEdiFormatConfigurationTerminatorDelimiter) Get() *EdiFormatConfigurationTerminatorDelimiter {
	return v.value
}

func (v *NullableEdiFormatConfigurationTerminatorDelimiter) Set(val *EdiFormatConfigurationTerminatorDelimiter) {
	v.value = val
	v.isSet = true
}

func (v NullableEdiFormatConfigurationTerminatorDelimiter) IsSet() bool {
	return v.isSet
}

func (v *NullableEdiFormatConfigurationTerminatorDelimiter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdiFormatConfigurationTerminatorDelimiter(val *EdiFormatConfigurationTerminatorDelimiter) *NullableEdiFormatConfigurationTerminatorDelimiter {
	return &NullableEdiFormatConfigurationTerminatorDelimiter{value: val, isSet: true}
}

func (v NullableEdiFormatConfigurationTerminatorDelimiter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdiFormatConfigurationTerminatorDelimiter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

