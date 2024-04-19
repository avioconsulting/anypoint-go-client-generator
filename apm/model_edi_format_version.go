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

// EdiFormatVersion struct for EdiFormatVersion
type EdiFormatVersion struct {
	Id *string `json:"id,omitempty"`
	FormatType string `json:"formatType"`
	Version string `json:"version"`
	Description string `json:"description"`
	Label string `json:"label"`
}

// NewEdiFormatVersion instantiates a new EdiFormatVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdiFormatVersion(formatType string, version string, description string, label string) *EdiFormatVersion {
	this := EdiFormatVersion{}
	this.FormatType = formatType
	this.Version = version
	this.Description = description
	this.Label = label
	return &this
}

// NewEdiFormatVersionWithDefaults instantiates a new EdiFormatVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdiFormatVersionWithDefaults() *EdiFormatVersion {
	this := EdiFormatVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EdiFormatVersion) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatVersion) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EdiFormatVersion) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EdiFormatVersion) SetId(v string) {
	o.Id = &v
}

// GetFormatType returns the FormatType field value
func (o *EdiFormatVersion) GetFormatType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormatType
}

// GetFormatTypeOk returns a tuple with the FormatType field value
// and a boolean to check if the value has been set.
func (o *EdiFormatVersion) GetFormatTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FormatType, true
}

// SetFormatType sets field value
func (o *EdiFormatVersion) SetFormatType(v string) {
	o.FormatType = v
}

// GetVersion returns the Version field value
func (o *EdiFormatVersion) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EdiFormatVersion) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *EdiFormatVersion) SetVersion(v string) {
	o.Version = v
}

// GetDescription returns the Description field value
func (o *EdiFormatVersion) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EdiFormatVersion) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EdiFormatVersion) SetDescription(v string) {
	o.Description = v
}

// GetLabel returns the Label field value
func (o *EdiFormatVersion) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *EdiFormatVersion) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *EdiFormatVersion) SetLabel(v string) {
	o.Label = v
}

func (o EdiFormatVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["formatType"] = o.FormatType
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableEdiFormatVersion struct {
	value *EdiFormatVersion
	isSet bool
}

func (v NullableEdiFormatVersion) Get() *EdiFormatVersion {
	return v.value
}

func (v *NullableEdiFormatVersion) Set(val *EdiFormatVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableEdiFormatVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableEdiFormatVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdiFormatVersion(val *EdiFormatVersion) *NullableEdiFormatVersion {
	return &NullableEdiFormatVersion{value: val, isSet: true}
}

func (v NullableEdiFormatVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdiFormatVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


