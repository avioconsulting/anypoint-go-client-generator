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

// EdiFormatConfigurationCharacterSetAndDelimitersSettings struct for EdiFormatConfigurationCharacterSetAndDelimitersSettings
type EdiFormatConfigurationCharacterSetAndDelimitersSettings struct {
	CharacterEncoding *string `json:"characterEncoding,omitempty"`
	DataSeparator *string `json:"dataSeparator,omitempty"`
	ComponentSeparator *string `json:"componentSeparator,omitempty"`
	RepetitionSeparator *string `json:"repetitionSeparator,omitempty"`
	SegmentTerminator *string `json:"segmentTerminator,omitempty"`
	ReleaseCharacter *string `json:"releaseCharacter,omitempty"`
	LineEnding *string `json:"lineEnding,omitempty"`
	TestIndicator *string `json:"testIndicator,omitempty"`
	WriteUseCRLFLastLine *string `json:"writeUseCRLFLastLine,omitempty"`
}

// NewEdiFormatConfigurationCharacterSetAndDelimitersSettings instantiates a new EdiFormatConfigurationCharacterSetAndDelimitersSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdiFormatConfigurationCharacterSetAndDelimitersSettings() *EdiFormatConfigurationCharacterSetAndDelimitersSettings {
	this := EdiFormatConfigurationCharacterSetAndDelimitersSettings{}
	return &this
}

// NewEdiFormatConfigurationCharacterSetAndDelimitersSettingsWithDefaults instantiates a new EdiFormatConfigurationCharacterSetAndDelimitersSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdiFormatConfigurationCharacterSetAndDelimitersSettingsWithDefaults() *EdiFormatConfigurationCharacterSetAndDelimitersSettings {
	this := EdiFormatConfigurationCharacterSetAndDelimitersSettings{}
	return &this
}

// GetCharacterEncoding returns the CharacterEncoding field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetCharacterEncoding() string {
	if o == nil || o.CharacterEncoding == nil {
		var ret string
		return ret
	}
	return *o.CharacterEncoding
}

// GetCharacterEncodingOk returns a tuple with the CharacterEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetCharacterEncodingOk() (*string, bool) {
	if o == nil || o.CharacterEncoding == nil {
		return nil, false
	}
	return o.CharacterEncoding, true
}

// HasCharacterEncoding returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) HasCharacterEncoding() bool {
	if o != nil && o.CharacterEncoding != nil {
		return true
	}

	return false
}

// SetCharacterEncoding gets a reference to the given string and assigns it to the CharacterEncoding field.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) SetCharacterEncoding(v string) {
	o.CharacterEncoding = &v
}

// GetDataSeparator returns the DataSeparator field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetDataSeparator() string {
	if o == nil || o.DataSeparator == nil {
		var ret string
		return ret
	}
	return *o.DataSeparator
}

// GetDataSeparatorOk returns a tuple with the DataSeparator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetDataSeparatorOk() (*string, bool) {
	if o == nil || o.DataSeparator == nil {
		return nil, false
	}
	return o.DataSeparator, true
}

// HasDataSeparator returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) HasDataSeparator() bool {
	if o != nil && o.DataSeparator != nil {
		return true
	}

	return false
}

// SetDataSeparator gets a reference to the given string and assigns it to the DataSeparator field.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) SetDataSeparator(v string) {
	o.DataSeparator = &v
}

// GetComponentSeparator returns the ComponentSeparator field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetComponentSeparator() string {
	if o == nil || o.ComponentSeparator == nil {
		var ret string
		return ret
	}
	return *o.ComponentSeparator
}

// GetComponentSeparatorOk returns a tuple with the ComponentSeparator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetComponentSeparatorOk() (*string, bool) {
	if o == nil || o.ComponentSeparator == nil {
		return nil, false
	}
	return o.ComponentSeparator, true
}

// HasComponentSeparator returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) HasComponentSeparator() bool {
	if o != nil && o.ComponentSeparator != nil {
		return true
	}

	return false
}

// SetComponentSeparator gets a reference to the given string and assigns it to the ComponentSeparator field.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) SetComponentSeparator(v string) {
	o.ComponentSeparator = &v
}

// GetRepetitionSeparator returns the RepetitionSeparator field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetRepetitionSeparator() string {
	if o == nil || o.RepetitionSeparator == nil {
		var ret string
		return ret
	}
	return *o.RepetitionSeparator
}

// GetRepetitionSeparatorOk returns a tuple with the RepetitionSeparator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetRepetitionSeparatorOk() (*string, bool) {
	if o == nil || o.RepetitionSeparator == nil {
		return nil, false
	}
	return o.RepetitionSeparator, true
}

// HasRepetitionSeparator returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) HasRepetitionSeparator() bool {
	if o != nil && o.RepetitionSeparator != nil {
		return true
	}

	return false
}

// SetRepetitionSeparator gets a reference to the given string and assigns it to the RepetitionSeparator field.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) SetRepetitionSeparator(v string) {
	o.RepetitionSeparator = &v
}

// GetSegmentTerminator returns the SegmentTerminator field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetSegmentTerminator() string {
	if o == nil || o.SegmentTerminator == nil {
		var ret string
		return ret
	}
	return *o.SegmentTerminator
}

// GetSegmentTerminatorOk returns a tuple with the SegmentTerminator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetSegmentTerminatorOk() (*string, bool) {
	if o == nil || o.SegmentTerminator == nil {
		return nil, false
	}
	return o.SegmentTerminator, true
}

// HasSegmentTerminator returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) HasSegmentTerminator() bool {
	if o != nil && o.SegmentTerminator != nil {
		return true
	}

	return false
}

// SetSegmentTerminator gets a reference to the given string and assigns it to the SegmentTerminator field.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) SetSegmentTerminator(v string) {
	o.SegmentTerminator = &v
}

// GetReleaseCharacter returns the ReleaseCharacter field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetReleaseCharacter() string {
	if o == nil || o.ReleaseCharacter == nil {
		var ret string
		return ret
	}
	return *o.ReleaseCharacter
}

// GetReleaseCharacterOk returns a tuple with the ReleaseCharacter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetReleaseCharacterOk() (*string, bool) {
	if o == nil || o.ReleaseCharacter == nil {
		return nil, false
	}
	return o.ReleaseCharacter, true
}

// HasReleaseCharacter returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) HasReleaseCharacter() bool {
	if o != nil && o.ReleaseCharacter != nil {
		return true
	}

	return false
}

// SetReleaseCharacter gets a reference to the given string and assigns it to the ReleaseCharacter field.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) SetReleaseCharacter(v string) {
	o.ReleaseCharacter = &v
}

// GetLineEnding returns the LineEnding field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetLineEnding() string {
	if o == nil || o.LineEnding == nil {
		var ret string
		return ret
	}
	return *o.LineEnding
}

// GetLineEndingOk returns a tuple with the LineEnding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetLineEndingOk() (*string, bool) {
	if o == nil || o.LineEnding == nil {
		return nil, false
	}
	return o.LineEnding, true
}

// HasLineEnding returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) HasLineEnding() bool {
	if o != nil && o.LineEnding != nil {
		return true
	}

	return false
}

// SetLineEnding gets a reference to the given string and assigns it to the LineEnding field.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) SetLineEnding(v string) {
	o.LineEnding = &v
}

// GetTestIndicator returns the TestIndicator field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetTestIndicator() string {
	if o == nil || o.TestIndicator == nil {
		var ret string
		return ret
	}
	return *o.TestIndicator
}

// GetTestIndicatorOk returns a tuple with the TestIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetTestIndicatorOk() (*string, bool) {
	if o == nil || o.TestIndicator == nil {
		return nil, false
	}
	return o.TestIndicator, true
}

// HasTestIndicator returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) HasTestIndicator() bool {
	if o != nil && o.TestIndicator != nil {
		return true
	}

	return false
}

// SetTestIndicator gets a reference to the given string and assigns it to the TestIndicator field.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) SetTestIndicator(v string) {
	o.TestIndicator = &v
}

// GetWriteUseCRLFLastLine returns the WriteUseCRLFLastLine field value if set, zero value otherwise.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetWriteUseCRLFLastLine() string {
	if o == nil || o.WriteUseCRLFLastLine == nil {
		var ret string
		return ret
	}
	return *o.WriteUseCRLFLastLine
}

// GetWriteUseCRLFLastLineOk returns a tuple with the WriteUseCRLFLastLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) GetWriteUseCRLFLastLineOk() (*string, bool) {
	if o == nil || o.WriteUseCRLFLastLine == nil {
		return nil, false
	}
	return o.WriteUseCRLFLastLine, true
}

// HasWriteUseCRLFLastLine returns a boolean if a field has been set.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) HasWriteUseCRLFLastLine() bool {
	if o != nil && o.WriteUseCRLFLastLine != nil {
		return true
	}

	return false
}

// SetWriteUseCRLFLastLine gets a reference to the given string and assigns it to the WriteUseCRLFLastLine field.
func (o *EdiFormatConfigurationCharacterSetAndDelimitersSettings) SetWriteUseCRLFLastLine(v string) {
	o.WriteUseCRLFLastLine = &v
}

func (o EdiFormatConfigurationCharacterSetAndDelimitersSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CharacterEncoding != nil {
		toSerialize["characterEncoding"] = o.CharacterEncoding
	}
	if o.DataSeparator != nil {
		toSerialize["dataSeparator"] = o.DataSeparator
	}
	if o.ComponentSeparator != nil {
		toSerialize["componentSeparator"] = o.ComponentSeparator
	}
	if o.RepetitionSeparator != nil {
		toSerialize["repetitionSeparator"] = o.RepetitionSeparator
	}
	if o.SegmentTerminator != nil {
		toSerialize["segmentTerminator"] = o.SegmentTerminator
	}
	if o.ReleaseCharacter != nil {
		toSerialize["releaseCharacter"] = o.ReleaseCharacter
	}
	if o.LineEnding != nil {
		toSerialize["lineEnding"] = o.LineEnding
	}
	if o.TestIndicator != nil {
		toSerialize["testIndicator"] = o.TestIndicator
	}
	if o.WriteUseCRLFLastLine != nil {
		toSerialize["writeUseCRLFLastLine"] = o.WriteUseCRLFLastLine
	}
	return json.Marshal(toSerialize)
}

type NullableEdiFormatConfigurationCharacterSetAndDelimitersSettings struct {
	value *EdiFormatConfigurationCharacterSetAndDelimitersSettings
	isSet bool
}

func (v NullableEdiFormatConfigurationCharacterSetAndDelimitersSettings) Get() *EdiFormatConfigurationCharacterSetAndDelimitersSettings {
	return v.value
}

func (v *NullableEdiFormatConfigurationCharacterSetAndDelimitersSettings) Set(val *EdiFormatConfigurationCharacterSetAndDelimitersSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableEdiFormatConfigurationCharacterSetAndDelimitersSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableEdiFormatConfigurationCharacterSetAndDelimitersSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdiFormatConfigurationCharacterSetAndDelimitersSettings(val *EdiFormatConfigurationCharacterSetAndDelimitersSettings) *NullableEdiFormatConfigurationCharacterSetAndDelimitersSettings {
	return &NullableEdiFormatConfigurationCharacterSetAndDelimitersSettings{value: val, isSet: true}
}

func (v NullableEdiFormatConfigurationCharacterSetAndDelimitersSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdiFormatConfigurationCharacterSetAndDelimitersSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


