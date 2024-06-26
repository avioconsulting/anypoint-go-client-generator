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

// EdiFormatConfigurationCharacterSetEncoding struct for EdiFormatConfigurationCharacterSetEncoding
type EdiFormatConfigurationCharacterSetEncoding struct {
	CharacterSet *string `json:"characterSet,omitempty"`
	CharacterEncoding *string `json:"characterEncoding,omitempty"`
	LineEndingBetweenSegments *string `json:"lineEndingBetweenSegments,omitempty"`
	ComponentElementSeparatorISA16 *string `json:"componentElementSeparatorISA16,omitempty"`
	SubstitutionCharacter *string `json:"substitutionCharacter,omitempty"`
	WriteUseCRLFLastLine *bool `json:"writeUseCRLFLastLine,omitempty"`
}

// NewEdiFormatConfigurationCharacterSetEncoding instantiates a new EdiFormatConfigurationCharacterSetEncoding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdiFormatConfigurationCharacterSetEncoding() *EdiFormatConfigurationCharacterSetEncoding {
	this := EdiFormatConfigurationCharacterSetEncoding{}
	return &this
}

// NewEdiFormatConfigurationCharacterSetEncodingWithDefaults instantiates a new EdiFormatConfigurationCharacterSetEncoding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdiFormatConfigurationCharacterSetEncodingWithDefaults() *EdiFormatConfigurationCharacterSetEncoding {
	this := EdiFormatConfigurationCharacterSetEncoding{}
	return &this
}

// GetCharacterSet returns the CharacterSet field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetEncoding) GetCharacterSet() string {
	if o == nil || o.CharacterSet == nil {
		var ret string
		return ret
	}
	return *o.CharacterSet
}

// GetCharacterSetOk returns a tuple with the CharacterSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetEncoding) GetCharacterSetOk() (*string, bool) {
	if o == nil || o.CharacterSet == nil {
		return nil, false
	}
	return o.CharacterSet, true
}

// HasCharacterSet returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetEncoding) HasCharacterSet() bool {
	if o != nil && o.CharacterSet != nil {
		return true
	}

	return false
}

// SetCharacterSet gets a reference to the given string and assigns it to the CharacterSet field.
func (o *EdiFormatConfigurationCharacterSetEncoding) SetCharacterSet(v string) {
	o.CharacterSet = &v
}

// GetCharacterEncoding returns the CharacterEncoding field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetEncoding) GetCharacterEncoding() string {
	if o == nil || o.CharacterEncoding == nil {
		var ret string
		return ret
	}
	return *o.CharacterEncoding
}

// GetCharacterEncodingOk returns a tuple with the CharacterEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetEncoding) GetCharacterEncodingOk() (*string, bool) {
	if o == nil || o.CharacterEncoding == nil {
		return nil, false
	}
	return o.CharacterEncoding, true
}

// HasCharacterEncoding returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetEncoding) HasCharacterEncoding() bool {
	if o != nil && o.CharacterEncoding != nil {
		return true
	}

	return false
}

// SetCharacterEncoding gets a reference to the given string and assigns it to the CharacterEncoding field.
func (o *EdiFormatConfigurationCharacterSetEncoding) SetCharacterEncoding(v string) {
	o.CharacterEncoding = &v
}

// GetLineEndingBetweenSegments returns the LineEndingBetweenSegments field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetEncoding) GetLineEndingBetweenSegments() string {
	if o == nil || o.LineEndingBetweenSegments == nil {
		var ret string
		return ret
	}
	return *o.LineEndingBetweenSegments
}

// GetLineEndingBetweenSegmentsOk returns a tuple with the LineEndingBetweenSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetEncoding) GetLineEndingBetweenSegmentsOk() (*string, bool) {
	if o == nil || o.LineEndingBetweenSegments == nil {
		return nil, false
	}
	return o.LineEndingBetweenSegments, true
}

// HasLineEndingBetweenSegments returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetEncoding) HasLineEndingBetweenSegments() bool {
	if o != nil && o.LineEndingBetweenSegments != nil {
		return true
	}

	return false
}

// SetLineEndingBetweenSegments gets a reference to the given string and assigns it to the LineEndingBetweenSegments field.
func (o *EdiFormatConfigurationCharacterSetEncoding) SetLineEndingBetweenSegments(v string) {
	o.LineEndingBetweenSegments = &v
}

// GetComponentElementSeparatorISA16 returns the ComponentElementSeparatorISA16 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetEncoding) GetComponentElementSeparatorISA16() string {
	if o == nil || o.ComponentElementSeparatorISA16 == nil {
		var ret string
		return ret
	}
	return *o.ComponentElementSeparatorISA16
}

// GetComponentElementSeparatorISA16Ok returns a tuple with the ComponentElementSeparatorISA16 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetEncoding) GetComponentElementSeparatorISA16Ok() (*string, bool) {
	if o == nil || o.ComponentElementSeparatorISA16 == nil {
		return nil, false
	}
	return o.ComponentElementSeparatorISA16, true
}

// HasComponentElementSeparatorISA16 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetEncoding) HasComponentElementSeparatorISA16() bool {
	if o != nil && o.ComponentElementSeparatorISA16 != nil {
		return true
	}

	return false
}

// SetComponentElementSeparatorISA16 gets a reference to the given string and assigns it to the ComponentElementSeparatorISA16 field.
func (o *EdiFormatConfigurationCharacterSetEncoding) SetComponentElementSeparatorISA16(v string) {
	o.ComponentElementSeparatorISA16 = &v
}

// GetSubstitutionCharacter returns the SubstitutionCharacter field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetEncoding) GetSubstitutionCharacter() string {
	if o == nil || o.SubstitutionCharacter == nil {
		var ret string
		return ret
	}
	return *o.SubstitutionCharacter
}

// GetSubstitutionCharacterOk returns a tuple with the SubstitutionCharacter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetEncoding) GetSubstitutionCharacterOk() (*string, bool) {
	if o == nil || o.SubstitutionCharacter == nil {
		return nil, false
	}
	return o.SubstitutionCharacter, true
}

// HasSubstitutionCharacter returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetEncoding) HasSubstitutionCharacter() bool {
	if o != nil && o.SubstitutionCharacter != nil {
		return true
	}

	return false
}

// SetSubstitutionCharacter gets a reference to the given string and assigns it to the SubstitutionCharacter field.
func (o *EdiFormatConfigurationCharacterSetEncoding) SetSubstitutionCharacter(v string) {
	o.SubstitutionCharacter = &v
}

// GetWriteUseCRLFLastLine returns the WriteUseCRLFLastLine field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetEncoding) GetWriteUseCRLFLastLine() bool {
	if o == nil || o.WriteUseCRLFLastLine == nil {
		var ret bool
		return ret
	}
	return *o.WriteUseCRLFLastLine
}

// GetWriteUseCRLFLastLineOk returns a tuple with the WriteUseCRLFLastLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetEncoding) GetWriteUseCRLFLastLineOk() (*bool, bool) {
	if o == nil || o.WriteUseCRLFLastLine == nil {
		return nil, false
	}
	return o.WriteUseCRLFLastLine, true
}

// HasWriteUseCRLFLastLine returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetEncoding) HasWriteUseCRLFLastLine() bool {
	if o != nil && o.WriteUseCRLFLastLine != nil {
		return true
	}

	return false
}

// SetWriteUseCRLFLastLine gets a reference to the given bool and assigns it to the WriteUseCRLFLastLine field.
func (o *EdiFormatConfigurationCharacterSetEncoding) SetWriteUseCRLFLastLine(v bool) {
	o.WriteUseCRLFLastLine = &v
}

func (o EdiFormatConfigurationCharacterSetEncoding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CharacterSet != nil {
		toSerialize["characterSet"] = o.CharacterSet
	}
	if o.CharacterEncoding != nil {
		toSerialize["characterEncoding"] = o.CharacterEncoding
	}
	if o.LineEndingBetweenSegments != nil {
		toSerialize["lineEndingBetweenSegments"] = o.LineEndingBetweenSegments
	}
	if o.ComponentElementSeparatorISA16 != nil {
		toSerialize["componentElementSeparatorISA16"] = o.ComponentElementSeparatorISA16
	}
	if o.SubstitutionCharacter != nil {
		toSerialize["substitutionCharacter"] = o.SubstitutionCharacter
	}
	if o.WriteUseCRLFLastLine != nil {
		toSerialize["writeUseCRLFLastLine"] = o.WriteUseCRLFLastLine
	}
	return json.Marshal(toSerialize)
}

type NullableEdiFormatConfigurationCharacterSetEncoding struct {
	value *EdiFormatConfigurationCharacterSetEncoding
	isSet bool
}

func (v NullableEdiFormatConfigurationCharacterSetEncoding) Get() *EdiFormatConfigurationCharacterSetEncoding {
	return v.value
}

func (v *NullableEdiFormatConfigurationCharacterSetEncoding) Set(val *EdiFormatConfigurationCharacterSetEncoding) {
	v.value = val
	v.isSet = true
}

func (v NullableEdiFormatConfigurationCharacterSetEncoding) IsSet() bool {
	return v.isSet
}

func (v *NullableEdiFormatConfigurationCharacterSetEncoding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdiFormatConfigurationCharacterSetEncoding(val *EdiFormatConfigurationCharacterSetEncoding) *NullableEdiFormatConfigurationCharacterSetEncoding {
	return &NullableEdiFormatConfigurationCharacterSetEncoding{value: val, isSet: true}
}

func (v NullableEdiFormatConfigurationCharacterSetEncoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdiFormatConfigurationCharacterSetEncoding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


