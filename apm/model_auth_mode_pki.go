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

// AuthModePKI This is a possible future extension for SFTP security.
type AuthModePKI struct {
	AuthModeBaseConfig
	IdentityFilePath *string `json:"identityFilePath,omitempty"`
	Passphrase *string `json:"passphrase,omitempty"`
	Algorithm *string `json:"algorithm,omitempty"`
}

// NewAuthModePKI instantiates a new AuthModePKI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthModePKI(authType string) *AuthModePKI {
	this := AuthModePKI{}
	this.AuthType = authType
	return &this
}

// NewAuthModePKIWithDefaults instantiates a new AuthModePKI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthModePKIWithDefaults() *AuthModePKI {
	this := AuthModePKI{}
	return &this
}

// GetIdentityFilePath returns the IdentityFilePath field value if set, zero value otherwise.
func (o *AuthModePKI) GetIdentityFilePath() string {
	if o == nil || o.IdentityFilePath == nil {
		var ret string
		return ret
	}
	return *o.IdentityFilePath
}

// GetIdentityFilePathOk returns a tuple with the IdentityFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthModePKI) GetIdentityFilePathOk() (*string, bool) {
	if o == nil || o.IdentityFilePath == nil {
		return nil, false
	}
	return o.IdentityFilePath, true
}

// HasIdentityFilePath returns a boolean if a field has been set.
func (o *AuthModePKI) HasIdentityFilePath() bool {
	if o != nil && o.IdentityFilePath != nil {
		return true
	}

	return false
}

// SetIdentityFilePath gets a reference to the given string and assigns it to the IdentityFilePath field.
func (o *AuthModePKI) SetIdentityFilePath(v string) {
	o.IdentityFilePath = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *AuthModePKI) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthModePKI) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *AuthModePKI) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *AuthModePKI) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *AuthModePKI) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthModePKI) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *AuthModePKI) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *AuthModePKI) SetAlgorithm(v string) {
	o.Algorithm = &v
}

func (o AuthModePKI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthModeBaseConfig, errAuthModeBaseConfig := json.Marshal(o.AuthModeBaseConfig)
	if errAuthModeBaseConfig != nil {
		return []byte{}, errAuthModeBaseConfig
	}
	errAuthModeBaseConfig = json.Unmarshal([]byte(serializedAuthModeBaseConfig), &toSerialize)
	if errAuthModeBaseConfig != nil {
		return []byte{}, errAuthModeBaseConfig
	}
	if o.IdentityFilePath != nil {
		toSerialize["identityFilePath"] = o.IdentityFilePath
	}
	if o.Passphrase != nil {
		toSerialize["passphrase"] = o.Passphrase
	}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}
	return json.Marshal(toSerialize)
}

type NullableAuthModePKI struct {
	value *AuthModePKI
	isSet bool
}

func (v NullableAuthModePKI) Get() *AuthModePKI {
	return v.value
}

func (v *NullableAuthModePKI) Set(val *AuthModePKI) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthModePKI) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthModePKI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthModePKI(val *AuthModePKI) *NullableAuthModePKI {
	return &NullableAuthModePKI{value: val, isSet: true}
}

func (v NullableAuthModePKI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthModePKI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


