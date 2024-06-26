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

// EdiFormatConfigurationCharacterSetAndEncoding struct for EdiFormatConfigurationCharacterSetAndEncoding
type EdiFormatConfigurationCharacterSetAndEncoding struct {
	CharacterSet *string `json:"characterSet,omitempty"`
	CharacterEncoding *string `json:"characterEncoding,omitempty"`
	LineEndingBetweenSegments *string `json:"lineEndingBetweenSegments,omitempty"`
	ComponentElementSeparatorISA16 *string `json:"componentElementSeparatorISA16,omitempty"`
	WriteUseCRLFLastLine *bool `json:"writeUseCRLFLastLine,omitempty"`
	SubstitutionCharacter *string `json:"substitutionCharacter,omitempty"`
}

// NewEdiFormatConfigurationCharacterSetAndEncoding instantiates a new EdiFormatConfigurationCharacterSetAndEncoding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdiFormatConfigurationCharacterSetAndEncoding() *EdiFormatConfigurationCharacterSetAndEncoding {
	this := EdiFormatConfigurationCharacterSetAndEncoding{}
	return &this
}

// NewEdiFormatConfigurationCharacterSetAndEncodingWithDefaults instantiates a new EdiFormatConfigurationCharacterSetAndEncoding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdiFormatConfigurationCharacterSetAndEncodingWithDefaults() *EdiFormatConfigurationCharacterSetAndEncoding {
	this := EdiFormatConfigurationCharacterSetAndEncoding{}
	return &this
}

// GetCharacterSet returns the CharacterSet field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) GetCharacterSet() string {
	if o == nil || o.CharacterSet == nil {
		var ret string
		return ret
	}
	return *o.CharacterSet
}

// GetCharacterSetOk returns a tuple with the CharacterSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) GetCharacterSetOk() (*string, bool) {
	if o == nil || o.CharacterSet == nil {
		return nil, false
	}
	return o.CharacterSet, true
}

// HasCharacterSet returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) HasCharacterSet() bool {
	if o != nil && o.CharacterSet != nil {
		return true
	}

	return false
}

// SetCharacterSet gets a reference to the given string and assigns it to the CharacterSet field.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) SetCharacterSet(v string) {
	o.CharacterSet = &v
}

// GetCharacterEncoding returns the CharacterEncoding field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) GetCharacterEncoding() string {
	if o == nil || o.CharacterEncoding == nil {
		var ret string
		return ret
	}
	return *o.CharacterEncoding
}

// GetCharacterEncodingOk returns a tuple with the CharacterEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) GetCharacterEncodingOk() (*string, bool) {
	if o == nil || o.CharacterEncoding == nil {
		return nil, false
	}
	return o.CharacterEncoding, true
}

// HasCharacterEncoding returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) HasCharacterEncoding() bool {
	if o != nil && o.CharacterEncoding != nil {
		return true
	}

	return false
}

// SetCharacterEncoding gets a reference to the given string and assigns it to the CharacterEncoding field.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) SetCharacterEncoding(v string) {
	o.CharacterEncoding = &v
}

// GetLineEndingBetweenSegments returns the LineEndingBetweenSegments field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) GetLineEndingBetweenSegments() string {
	if o == nil || o.LineEndingBetweenSegments == nil {
		var ret string
		return ret
	}
	return *o.LineEndingBetweenSegments
}

// GetLineEndingBetweenSegmentsOk returns a tuple with the LineEndingBetweenSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) GetLineEndingBetweenSegmentsOk() (*string, bool) {
	if o == nil || o.LineEndingBetweenSegments == nil {
		return nil, false
	}
	return o.LineEndingBetweenSegments, true
}

// HasLineEndingBetweenSegments returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) HasLineEndingBetweenSegments() bool {
	if o != nil && o.LineEndingBetweenSegments != nil {
		return true
	}

	return false
}

// SetLineEndingBetweenSegments gets a reference to the given string and assigns it to the LineEndingBetweenSegments field.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) SetLineEndingBetweenSegments(v string) {
	o.LineEndingBetweenSegments = &v
}

// GetComponentElementSeparatorISA16 returns the ComponentElementSeparatorISA16 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) GetComponentElementSeparatorISA16() string {
	if o == nil || o.ComponentElementSeparatorISA16 == nil {
		var ret string
		return ret
	}
	return *o.ComponentElementSeparatorISA16
}

// GetComponentElementSeparatorISA16Ok returns a tuple with the ComponentElementSeparatorISA16 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) GetComponentElementSeparatorISA16Ok() (*string, bool) {
	if o == nil || o.ComponentElementSeparatorISA16 == nil {
		return nil, false
	}
	return o.ComponentElementSeparatorISA16, true
}

// HasComponentElementSeparatorISA16 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) HasComponentElementSeparatorISA16() bool {
	if o != nil && o.ComponentElementSeparatorISA16 != nil {
		return true
	}

	return false
}

// SetComponentElementSeparatorISA16 gets a reference to the given string and assigns it to the ComponentElementSeparatorISA16 field.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) SetComponentElementSeparatorISA16(v string) {
	o.ComponentElementSeparatorISA16 = &v
}

// GetWriteUseCRLFLastLine returns the WriteUseCRLFLastLine field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) GetWriteUseCRLFLastLine() bool {
	if o == nil || o.WriteUseCRLFLastLine == nil {
		var ret bool
		return ret
	}
	return *o.WriteUseCRLFLastLine
}

// GetWriteUseCRLFLastLineOk returns a tuple with the WriteUseCRLFLastLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) GetWriteUseCRLFLastLineOk() (*bool, bool) {
	if o == nil || o.WriteUseCRLFLastLine == nil {
		return nil, false
	}
	return o.WriteUseCRLFLastLine, true
}

// HasWriteUseCRLFLastLine returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) HasWriteUseCRLFLastLine() bool {
	if o != nil && o.WriteUseCRLFLastLine != nil {
		return true
	}

	return false
}

// SetWriteUseCRLFLastLine gets a reference to the given bool and assigns it to the WriteUseCRLFLastLine field.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) SetWriteUseCRLFLastLine(v bool) {
	o.WriteUseCRLFLastLine = &v
}

// GetSubstitutionCharacter returns the SubstitutionCharacter field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) GetSubstitutionCharacter() string {
	if o == nil || o.SubstitutionCharacter == nil {
		var ret string
		return ret
	}
	return *o.SubstitutionCharacter
}

// GetSubstitutionCharacterOk returns a tuple with the SubstitutionCharacter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) GetSubstitutionCharacterOk() (*string, bool) {
	if o == nil || o.SubstitutionCharacter == nil {
		return nil, false
	}
	return o.SubstitutionCharacter, true
}

// HasSubstitutionCharacter returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) HasSubstitutionCharacter() bool {
	if o != nil && o.SubstitutionCharacter != nil {
		return true
	}

	return false
}

// SetSubstitutionCharacter gets a reference to the given string and assigns it to the SubstitutionCharacter field.
func (o *EdiFormatConfigurationCharacterSetAndEncoding) SetSubstitutionCharacter(v string) {
	o.SubstitutionCharacter = &v
}

func (o EdiFormatConfigurationCharacterSetAndEncoding) MarshalJSON() ([]byte, error) {
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
	if o.WriteUseCRLFLastLine != nil {
		toSerialize["writeUseCRLFLastLine"] = o.WriteUseCRLFLastLine
	}
	if o.SubstitutionCharacter != nil {
		toSerialize["substitutionCharacter"] = o.SubstitutionCharacter
	}
	return json.Marshal(toSerialize)
}

type NullableEdiFormatConfigurationCharacterSetAndEncoding struct {
	value *EdiFormatConfigurationCharacterSetAndEncoding
	isSet bool
}

func (v NullableEdiFormatConfigurationCharacterSetAndEncoding) Get() *EdiFormatConfigurationCharacterSetAndEncoding {
	return v.value
}

func (v *NullableEdiFormatConfigurationCharacterSetAndEncoding) Set(val *EdiFormatConfigurationCharacterSetAndEncoding) {
	v.value = val
	v.isSet = true
}

func (v NullableEdiFormatConfigurationCharacterSetAndEncoding) IsSet() bool {
	return v.isSet
}

func (v *NullableEdiFormatConfigurationCharacterSetAndEncoding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdiFormatConfigurationCharacterSetAndEncoding(val *EdiFormatConfigurationCharacterSetAndEncoding) *NullableEdiFormatConfigurationCharacterSetAndEncoding {
	return &NullableEdiFormatConfigurationCharacterSetAndEncoding{value: val, isSet: true}
}

func (v NullableEdiFormatConfigurationCharacterSetAndEncoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdiFormatConfigurationCharacterSetAndEncoding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


