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

// Identifier Identifier Type
type Identifier struct {
	Id *string `json:"id,omitempty"`
	IdentifierTypeQualifierId string `json:"identifierTypeQualifierId"`
	Status string `json:"status"`
	QualifierLabel *string `json:"qualifierLabel,omitempty"`
	TypeLabel *string `json:"typeLabel,omitempty"`
	Code *string `json:"code,omitempty"`
	Value string `json:"value"`
}

// NewIdentifier instantiates a new Identifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentifier(identifierTypeQualifierId string, status string, value string) *Identifier {
	this := Identifier{}
	this.IdentifierTypeQualifierId = identifierTypeQualifierId
	this.Status = status
	this.Value = value
	return &this
}

// NewIdentifierWithDefaults instantiates a new Identifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentifierWithDefaults() *Identifier {
	this := Identifier{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Identifier) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identifier) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Identifier) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Identifier) SetId(v string) {
	o.Id = &v
}

// GetIdentifierTypeQualifierId returns the IdentifierTypeQualifierId field value
func (o *Identifier) GetIdentifierTypeQualifierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentifierTypeQualifierId
}

// GetIdentifierTypeQualifierIdOk returns a tuple with the IdentifierTypeQualifierId field value
// and a boolean to check if the value has been set.
func (o *Identifier) GetIdentifierTypeQualifierIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdentifierTypeQualifierId, true
}

// SetIdentifierTypeQualifierId sets field value
func (o *Identifier) SetIdentifierTypeQualifierId(v string) {
	o.IdentifierTypeQualifierId = v
}

// GetStatus returns the Status field value
func (o *Identifier) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Identifier) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Identifier) SetStatus(v string) {
	o.Status = v
}

// GetQualifierLabel returns the QualifierLabel field value if set, zero value otherwise.
func (o *Identifier) GetQualifierLabel() string {
	if o == nil || o.QualifierLabel == nil {
		var ret string
		return ret
	}
	return *o.QualifierLabel
}

// GetQualifierLabelOk returns a tuple with the QualifierLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identifier) GetQualifierLabelOk() (*string, bool) {
	if o == nil || o.QualifierLabel == nil {
		return nil, false
	}
	return o.QualifierLabel, true
}

// HasQualifierLabel returns a boolean if a field has been set.
func (o *Identifier) HasQualifierLabel() bool {
	if o != nil && o.QualifierLabel != nil {
		return true
	}

	return false
}

// SetQualifierLabel gets a reference to the given string and assigns it to the QualifierLabel field.
func (o *Identifier) SetQualifierLabel(v string) {
	o.QualifierLabel = &v
}

// GetTypeLabel returns the TypeLabel field value if set, zero value otherwise.
func (o *Identifier) GetTypeLabel() string {
	if o == nil || o.TypeLabel == nil {
		var ret string
		return ret
	}
	return *o.TypeLabel
}

// GetTypeLabelOk returns a tuple with the TypeLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identifier) GetTypeLabelOk() (*string, bool) {
	if o == nil || o.TypeLabel == nil {
		return nil, false
	}
	return o.TypeLabel, true
}

// HasTypeLabel returns a boolean if a field has been set.
func (o *Identifier) HasTypeLabel() bool {
	if o != nil && o.TypeLabel != nil {
		return true
	}

	return false
}

// SetTypeLabel gets a reference to the given string and assigns it to the TypeLabel field.
func (o *Identifier) SetTypeLabel(v string) {
	o.TypeLabel = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Identifier) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identifier) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Identifier) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Identifier) SetCode(v string) {
	o.Code = &v
}

// GetValue returns the Value field value
func (o *Identifier) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Identifier) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Identifier) SetValue(v string) {
	o.Value = v
}

func (o Identifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["identifierTypeQualifierId"] = o.IdentifierTypeQualifierId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.QualifierLabel != nil {
		toSerialize["qualifierLabel"] = o.QualifierLabel
	}
	if o.TypeLabel != nil {
		toSerialize["typeLabel"] = o.TypeLabel
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableIdentifier struct {
	value *Identifier
	isSet bool
}

func (v NullableIdentifier) Get() *Identifier {
	return v.value
}

func (v *NullableIdentifier) Set(val *Identifier) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentifier(val *Identifier) *NullableIdentifier {
	return &NullableIdentifier{value: val, isSet: true}
}

func (v NullableIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


