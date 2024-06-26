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

// IdentifierTypeQualifierAllOf struct for IdentifierTypeQualifierAllOf
type IdentifierTypeQualifierAllOf struct {
	Id *string `json:"id,omitempty"`
	IdentifierTypeId *string `json:"identifierTypeId,omitempty"`
	Code *string `json:"code,omitempty"`
	Label *string `json:"label,omitempty"`
	SegmentIdentifier *string `json:"segmentIdentifier,omitempty"`
	Descritpion *string `json:"descritpion,omitempty"`
	EnvironmentId *string `json:"environmentId,omitempty"`
}

// NewIdentifierTypeQualifierAllOf instantiates a new IdentifierTypeQualifierAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentifierTypeQualifierAllOf() *IdentifierTypeQualifierAllOf {
	this := IdentifierTypeQualifierAllOf{}
	return &this
}

// NewIdentifierTypeQualifierAllOfWithDefaults instantiates a new IdentifierTypeQualifierAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentifierTypeQualifierAllOfWithDefaults() *IdentifierTypeQualifierAllOf {
	this := IdentifierTypeQualifierAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentifierTypeQualifierAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifierTypeQualifierAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentifierTypeQualifierAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentifierTypeQualifierAllOf) SetId(v string) {
	o.Id = &v
}

// GetIdentifierTypeId returns the IdentifierTypeId field value if set, zero value otherwise.
func (o *IdentifierTypeQualifierAllOf) GetIdentifierTypeId() string {
	if o == nil || o.IdentifierTypeId == nil {
		var ret string
		return ret
	}
	return *o.IdentifierTypeId
}

// GetIdentifierTypeIdOk returns a tuple with the IdentifierTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifierTypeQualifierAllOf) GetIdentifierTypeIdOk() (*string, bool) {
	if o == nil || o.IdentifierTypeId == nil {
		return nil, false
	}
	return o.IdentifierTypeId, true
}

// HasIdentifierTypeId returns a boolean if a field has been set.
func (o *IdentifierTypeQualifierAllOf) HasIdentifierTypeId() bool {
	if o != nil && o.IdentifierTypeId != nil {
		return true
	}

	return false
}

// SetIdentifierTypeId gets a reference to the given string and assigns it to the IdentifierTypeId field.
func (o *IdentifierTypeQualifierAllOf) SetIdentifierTypeId(v string) {
	o.IdentifierTypeId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *IdentifierTypeQualifierAllOf) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifierTypeQualifierAllOf) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *IdentifierTypeQualifierAllOf) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *IdentifierTypeQualifierAllOf) SetCode(v string) {
	o.Code = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *IdentifierTypeQualifierAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifierTypeQualifierAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *IdentifierTypeQualifierAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *IdentifierTypeQualifierAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetSegmentIdentifier returns the SegmentIdentifier field value if set, zero value otherwise.
func (o *IdentifierTypeQualifierAllOf) GetSegmentIdentifier() string {
	if o == nil || o.SegmentIdentifier == nil {
		var ret string
		return ret
	}
	return *o.SegmentIdentifier
}

// GetSegmentIdentifierOk returns a tuple with the SegmentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifierTypeQualifierAllOf) GetSegmentIdentifierOk() (*string, bool) {
	if o == nil || o.SegmentIdentifier == nil {
		return nil, false
	}
	return o.SegmentIdentifier, true
}

// HasSegmentIdentifier returns a boolean if a field has been set.
func (o *IdentifierTypeQualifierAllOf) HasSegmentIdentifier() bool {
	if o != nil && o.SegmentIdentifier != nil {
		return true
	}

	return false
}

// SetSegmentIdentifier gets a reference to the given string and assigns it to the SegmentIdentifier field.
func (o *IdentifierTypeQualifierAllOf) SetSegmentIdentifier(v string) {
	o.SegmentIdentifier = &v
}

// GetDescritpion returns the Descritpion field value if set, zero value otherwise.
func (o *IdentifierTypeQualifierAllOf) GetDescritpion() string {
	if o == nil || o.Descritpion == nil {
		var ret string
		return ret
	}
	return *o.Descritpion
}

// GetDescritpionOk returns a tuple with the Descritpion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifierTypeQualifierAllOf) GetDescritpionOk() (*string, bool) {
	if o == nil || o.Descritpion == nil {
		return nil, false
	}
	return o.Descritpion, true
}

// HasDescritpion returns a boolean if a field has been set.
func (o *IdentifierTypeQualifierAllOf) HasDescritpion() bool {
	if o != nil && o.Descritpion != nil {
		return true
	}

	return false
}

// SetDescritpion gets a reference to the given string and assigns it to the Descritpion field.
func (o *IdentifierTypeQualifierAllOf) SetDescritpion(v string) {
	o.Descritpion = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *IdentifierTypeQualifierAllOf) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifierTypeQualifierAllOf) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *IdentifierTypeQualifierAllOf) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId != nil {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *IdentifierTypeQualifierAllOf) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

func (o IdentifierTypeQualifierAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdentifierTypeId != nil {
		toSerialize["identifierTypeId"] = o.IdentifierTypeId
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.SegmentIdentifier != nil {
		toSerialize["segmentIdentifier"] = o.SegmentIdentifier
	}
	if o.Descritpion != nil {
		toSerialize["descritpion"] = o.Descritpion
	}
	if o.EnvironmentId != nil {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	return json.Marshal(toSerialize)
}

type NullableIdentifierTypeQualifierAllOf struct {
	value *IdentifierTypeQualifierAllOf
	isSet bool
}

func (v NullableIdentifierTypeQualifierAllOf) Get() *IdentifierTypeQualifierAllOf {
	return v.value
}

func (v *NullableIdentifierTypeQualifierAllOf) Set(val *IdentifierTypeQualifierAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentifierTypeQualifierAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentifierTypeQualifierAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentifierTypeQualifierAllOf(val *IdentifierTypeQualifierAllOf) *NullableIdentifierTypeQualifierAllOf {
	return &NullableIdentifierTypeQualifierAllOf{value: val, isSet: true}
}

func (v NullableIdentifierTypeQualifierAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentifierTypeQualifierAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


