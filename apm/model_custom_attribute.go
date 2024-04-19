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

// CustomAttribute Custom Attribute
type CustomAttribute struct {
	Id *string `json:"id,omitempty"`
	Searchable bool `json:"searchable"`
	Label string `json:"label"`
	Alias string `json:"alias"`
}

// NewCustomAttribute instantiates a new CustomAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAttribute(searchable bool, label string, alias string) *CustomAttribute {
	this := CustomAttribute{}
	this.Searchable = searchable
	this.Label = label
	this.Alias = alias
	return &this
}

// NewCustomAttributeWithDefaults instantiates a new CustomAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAttributeWithDefaults() *CustomAttribute {
	this := CustomAttribute{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomAttribute) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttribute) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomAttribute) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomAttribute) SetId(v string) {
	o.Id = &v
}

// GetSearchable returns the Searchable field value
func (o *CustomAttribute) GetSearchable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Searchable
}

// GetSearchableOk returns a tuple with the Searchable field value
// and a boolean to check if the value has been set.
func (o *CustomAttribute) GetSearchableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Searchable, true
}

// SetSearchable sets field value
func (o *CustomAttribute) SetSearchable(v bool) {
	o.Searchable = v
}

// GetLabel returns the Label field value
func (o *CustomAttribute) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *CustomAttribute) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *CustomAttribute) SetLabel(v string) {
	o.Label = v
}

// GetAlias returns the Alias field value
func (o *CustomAttribute) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *CustomAttribute) GetAliasOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *CustomAttribute) SetAlias(v string) {
	o.Alias = v
}

func (o CustomAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["searchable"] = o.Searchable
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["alias"] = o.Alias
	}
	return json.Marshal(toSerialize)
}

type NullableCustomAttribute struct {
	value *CustomAttribute
	isSet bool
}

func (v NullableCustomAttribute) Get() *CustomAttribute {
	return v.value
}

func (v *NullableCustomAttribute) Set(val *CustomAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAttribute(val *CustomAttribute) *NullableCustomAttribute {
	return &NullableCustomAttribute{value: val, isSet: true}
}

func (v NullableCustomAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


