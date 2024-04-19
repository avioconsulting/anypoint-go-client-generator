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

// EdiFormatConfigurationWriteSettings struct for EdiFormatConfigurationWriteSettings
type EdiFormatConfigurationWriteSettings struct {
	SendSyntaxVersion *string `json:"sendSyntaxVersion,omitempty"`
	AlwaysSendUNA *bool `json:"alwaysSendUNA,omitempty"`
	EnforceWriteCharacters *bool `json:"enforceWriteCharacters,omitempty"`
	WriteEnforceLengthLimits *bool `json:"writeEnforceLengthLimits,omitempty"`
	RequestAcks *bool `json:"requestAcks,omitempty"`
	EnforceCodeSetValidationsWrite *bool `json:"enforceCodeSetValidationsWrite,omitempty"`
}

// NewEdiFormatConfigurationWriteSettings instantiates a new EdiFormatConfigurationWriteSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdiFormatConfigurationWriteSettings() *EdiFormatConfigurationWriteSettings {
	this := EdiFormatConfigurationWriteSettings{}
	return &this
}

// NewEdiFormatConfigurationWriteSettingsWithDefaults instantiates a new EdiFormatConfigurationWriteSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdiFormatConfigurationWriteSettingsWithDefaults() *EdiFormatConfigurationWriteSettings {
	this := EdiFormatConfigurationWriteSettings{}
	return &this
}

// GetSendSyntaxVersion returns the SendSyntaxVersion field value if set, zero value otherwise.
func (o *EdiFormatConfigurationWriteSettings) GetSendSyntaxVersion() string {
	if o == nil || o.SendSyntaxVersion == nil {
		var ret string
		return ret
	}
	return *o.SendSyntaxVersion
}

// GetSendSyntaxVersionOk returns a tuple with the SendSyntaxVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationWriteSettings) GetSendSyntaxVersionOk() (*string, bool) {
	if o == nil || o.SendSyntaxVersion == nil {
		return nil, false
	}
	return o.SendSyntaxVersion, true
}

// HasSendSyntaxVersion returns a boolean if a field has been set.
func (o *EdiFormatConfigurationWriteSettings) HasSendSyntaxVersion() bool {
	if o != nil && o.SendSyntaxVersion != nil {
		return true
	}

	return false
}

// SetSendSyntaxVersion gets a reference to the given string and assigns it to the SendSyntaxVersion field.
func (o *EdiFormatConfigurationWriteSettings) SetSendSyntaxVersion(v string) {
	o.SendSyntaxVersion = &v
}

// GetAlwaysSendUNA returns the AlwaysSendUNA field value if set, zero value otherwise.
func (o *EdiFormatConfigurationWriteSettings) GetAlwaysSendUNA() bool {
	if o == nil || o.AlwaysSendUNA == nil {
		var ret bool
		return ret
	}
	return *o.AlwaysSendUNA
}

// GetAlwaysSendUNAOk returns a tuple with the AlwaysSendUNA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationWriteSettings) GetAlwaysSendUNAOk() (*bool, bool) {
	if o == nil || o.AlwaysSendUNA == nil {
		return nil, false
	}
	return o.AlwaysSendUNA, true
}

// HasAlwaysSendUNA returns a boolean if a field has been set.
func (o *EdiFormatConfigurationWriteSettings) HasAlwaysSendUNA() bool {
	if o != nil && o.AlwaysSendUNA != nil {
		return true
	}

	return false
}

// SetAlwaysSendUNA gets a reference to the given bool and assigns it to the AlwaysSendUNA field.
func (o *EdiFormatConfigurationWriteSettings) SetAlwaysSendUNA(v bool) {
	o.AlwaysSendUNA = &v
}

// GetEnforceWriteCharacters returns the EnforceWriteCharacters field value if set, zero value otherwise.
func (o *EdiFormatConfigurationWriteSettings) GetEnforceWriteCharacters() bool {
	if o == nil || o.EnforceWriteCharacters == nil {
		var ret bool
		return ret
	}
	return *o.EnforceWriteCharacters
}

// GetEnforceWriteCharactersOk returns a tuple with the EnforceWriteCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationWriteSettings) GetEnforceWriteCharactersOk() (*bool, bool) {
	if o == nil || o.EnforceWriteCharacters == nil {
		return nil, false
	}
	return o.EnforceWriteCharacters, true
}

// HasEnforceWriteCharacters returns a boolean if a field has been set.
func (o *EdiFormatConfigurationWriteSettings) HasEnforceWriteCharacters() bool {
	if o != nil && o.EnforceWriteCharacters != nil {
		return true
	}

	return false
}

// SetEnforceWriteCharacters gets a reference to the given bool and assigns it to the EnforceWriteCharacters field.
func (o *EdiFormatConfigurationWriteSettings) SetEnforceWriteCharacters(v bool) {
	o.EnforceWriteCharacters = &v
}

// GetWriteEnforceLengthLimits returns the WriteEnforceLengthLimits field value if set, zero value otherwise.
func (o *EdiFormatConfigurationWriteSettings) GetWriteEnforceLengthLimits() bool {
	if o == nil || o.WriteEnforceLengthLimits == nil {
		var ret bool
		return ret
	}
	return *o.WriteEnforceLengthLimits
}

// GetWriteEnforceLengthLimitsOk returns a tuple with the WriteEnforceLengthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationWriteSettings) GetWriteEnforceLengthLimitsOk() (*bool, bool) {
	if o == nil || o.WriteEnforceLengthLimits == nil {
		return nil, false
	}
	return o.WriteEnforceLengthLimits, true
}

// HasWriteEnforceLengthLimits returns a boolean if a field has been set.
func (o *EdiFormatConfigurationWriteSettings) HasWriteEnforceLengthLimits() bool {
	if o != nil && o.WriteEnforceLengthLimits != nil {
		return true
	}

	return false
}

// SetWriteEnforceLengthLimits gets a reference to the given bool and assigns it to the WriteEnforceLengthLimits field.
func (o *EdiFormatConfigurationWriteSettings) SetWriteEnforceLengthLimits(v bool) {
	o.WriteEnforceLengthLimits = &v
}

// GetRequestAcks returns the RequestAcks field value if set, zero value otherwise.
func (o *EdiFormatConfigurationWriteSettings) GetRequestAcks() bool {
	if o == nil || o.RequestAcks == nil {
		var ret bool
		return ret
	}
	return *o.RequestAcks
}

// GetRequestAcksOk returns a tuple with the RequestAcks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationWriteSettings) GetRequestAcksOk() (*bool, bool) {
	if o == nil || o.RequestAcks == nil {
		return nil, false
	}
	return o.RequestAcks, true
}

// HasRequestAcks returns a boolean if a field has been set.
func (o *EdiFormatConfigurationWriteSettings) HasRequestAcks() bool {
	if o != nil && o.RequestAcks != nil {
		return true
	}

	return false
}

// SetRequestAcks gets a reference to the given bool and assigns it to the RequestAcks field.
func (o *EdiFormatConfigurationWriteSettings) SetRequestAcks(v bool) {
	o.RequestAcks = &v
}

// GetEnforceCodeSetValidationsWrite returns the EnforceCodeSetValidationsWrite field value if set, zero value otherwise.
func (o *EdiFormatConfigurationWriteSettings) GetEnforceCodeSetValidationsWrite() bool {
	if o == nil || o.EnforceCodeSetValidationsWrite == nil {
		var ret bool
		return ret
	}
	return *o.EnforceCodeSetValidationsWrite
}

// GetEnforceCodeSetValidationsWriteOk returns a tuple with the EnforceCodeSetValidationsWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationWriteSettings) GetEnforceCodeSetValidationsWriteOk() (*bool, bool) {
	if o == nil || o.EnforceCodeSetValidationsWrite == nil {
		return nil, false
	}
	return o.EnforceCodeSetValidationsWrite, true
}

// HasEnforceCodeSetValidationsWrite returns a boolean if a field has been set.
func (o *EdiFormatConfigurationWriteSettings) HasEnforceCodeSetValidationsWrite() bool {
	if o != nil && o.EnforceCodeSetValidationsWrite != nil {
		return true
	}

	return false
}

// SetEnforceCodeSetValidationsWrite gets a reference to the given bool and assigns it to the EnforceCodeSetValidationsWrite field.
func (o *EdiFormatConfigurationWriteSettings) SetEnforceCodeSetValidationsWrite(v bool) {
	o.EnforceCodeSetValidationsWrite = &v
}

func (o EdiFormatConfigurationWriteSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SendSyntaxVersion != nil {
		toSerialize["sendSyntaxVersion"] = o.SendSyntaxVersion
	}
	if o.AlwaysSendUNA != nil {
		toSerialize["alwaysSendUNA"] = o.AlwaysSendUNA
	}
	if o.EnforceWriteCharacters != nil {
		toSerialize["enforceWriteCharacters"] = o.EnforceWriteCharacters
	}
	if o.WriteEnforceLengthLimits != nil {
		toSerialize["writeEnforceLengthLimits"] = o.WriteEnforceLengthLimits
	}
	if o.RequestAcks != nil {
		toSerialize["requestAcks"] = o.RequestAcks
	}
	if o.EnforceCodeSetValidationsWrite != nil {
		toSerialize["enforceCodeSetValidationsWrite"] = o.EnforceCodeSetValidationsWrite
	}
	return json.Marshal(toSerialize)
}

type NullableEdiFormatConfigurationWriteSettings struct {
	value *EdiFormatConfigurationWriteSettings
	isSet bool
}

func (v NullableEdiFormatConfigurationWriteSettings) Get() *EdiFormatConfigurationWriteSettings {
	return v.value
}

func (v *NullableEdiFormatConfigurationWriteSettings) Set(val *EdiFormatConfigurationWriteSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableEdiFormatConfigurationWriteSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableEdiFormatConfigurationWriteSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdiFormatConfigurationWriteSettings(val *EdiFormatConfigurationWriteSettings) *NullableEdiFormatConfigurationWriteSettings {
	return &NullableEdiFormatConfigurationWriteSettings{value: val, isSet: true}
}

func (v NullableEdiFormatConfigurationWriteSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdiFormatConfigurationWriteSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

