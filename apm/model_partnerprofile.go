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

// Partnerprofile Partner Profile Type
type Partnerprofile struct {
	PartnerType string `json:"partnerType"`
	EnvironmentId string `json:"environmentId"`
	Name string `json:"name"`
	Identifiers []Partnerprofileidentifier `json:"identifiers"`
}

// NewPartnerprofile instantiates a new Partnerprofile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerprofile(partnerType string, environmentId string, name string, identifiers []Partnerprofileidentifier) *Partnerprofile {
	this := Partnerprofile{}
	this.PartnerType = partnerType
	this.EnvironmentId = environmentId
	this.Name = name
	this.Identifiers = identifiers
	return &this
}

// NewPartnerprofileWithDefaults instantiates a new Partnerprofile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerprofileWithDefaults() *Partnerprofile {
	this := Partnerprofile{}
	return &this
}

// GetPartnerType returns the PartnerType field value
func (o *Partnerprofile) GetPartnerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartnerType
}

// GetPartnerTypeOk returns a tuple with the PartnerType field value
// and a boolean to check if the value has been set.
func (o *Partnerprofile) GetPartnerTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PartnerType, true
}

// SetPartnerType sets field value
func (o *Partnerprofile) SetPartnerType(v string) {
	o.PartnerType = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *Partnerprofile) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *Partnerprofile) GetEnvironmentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *Partnerprofile) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetName returns the Name field value
func (o *Partnerprofile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Partnerprofile) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Partnerprofile) SetName(v string) {
	o.Name = v
}

// GetIdentifiers returns the Identifiers field value
func (o *Partnerprofile) GetIdentifiers() []Partnerprofileidentifier {
	if o == nil {
		var ret []Partnerprofileidentifier
		return ret
	}

	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value
// and a boolean to check if the value has been set.
func (o *Partnerprofile) GetIdentifiersOk() (*[]Partnerprofileidentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Identifiers, true
}

// SetIdentifiers sets field value
func (o *Partnerprofile) SetIdentifiers(v []Partnerprofileidentifier) {
	o.Identifiers = v
}

func (o Partnerprofile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["partnerType"] = o.PartnerType
	}
	if true {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["identifiers"] = o.Identifiers
	}
	return json.Marshal(toSerialize)
}

type NullablePartnerprofile struct {
	value *Partnerprofile
	isSet bool
}

func (v NullablePartnerprofile) Get() *Partnerprofile {
	return v.value
}

func (v *NullablePartnerprofile) Set(val *Partnerprofile) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerprofile) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerprofile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerprofile(val *Partnerprofile) *NullablePartnerprofile {
	return &NullablePartnerprofile{value: val, isSet: true}
}

func (v NullablePartnerprofile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerprofile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

