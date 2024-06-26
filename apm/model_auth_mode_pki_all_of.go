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

// AuthModePKIAllOf struct for AuthModePKIAllOf
type AuthModePKIAllOf struct {
	IdentityFilePath *string `json:"identityFilePath,omitempty"`
	Passphrase *string `json:"passphrase,omitempty"`
	Algorithm *string `json:"algorithm,omitempty"`
}

// NewAuthModePKIAllOf instantiates a new AuthModePKIAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthModePKIAllOf() *AuthModePKIAllOf {
	this := AuthModePKIAllOf{}
	return &this
}

// NewAuthModePKIAllOfWithDefaults instantiates a new AuthModePKIAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthModePKIAllOfWithDefaults() *AuthModePKIAllOf {
	this := AuthModePKIAllOf{}
	return &this
}

// GetIdentityFilePath returns the IdentityFilePath field value if set, zero value otherwise.
func (o *AuthModePKIAllOf) GetIdentityFilePath() string {
	if o == nil || o.IdentityFilePath == nil {
		var ret string
		return ret
	}
	return *o.IdentityFilePath
}

// GetIdentityFilePathOk returns a tuple with the IdentityFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthModePKIAllOf) GetIdentityFilePathOk() (*string, bool) {
	if o == nil || o.IdentityFilePath == nil {
		return nil, false
	}
	return o.IdentityFilePath, true
}

// HasIdentityFilePath returns a boolean if a field has been set.
func (o *AuthModePKIAllOf) HasIdentityFilePath() bool {
	if o != nil && o.IdentityFilePath != nil {
		return true
	}

	return false
}

// SetIdentityFilePath gets a reference to the given string and assigns it to the IdentityFilePath field.
func (o *AuthModePKIAllOf) SetIdentityFilePath(v string) {
	o.IdentityFilePath = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *AuthModePKIAllOf) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthModePKIAllOf) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *AuthModePKIAllOf) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *AuthModePKIAllOf) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *AuthModePKIAllOf) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthModePKIAllOf) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *AuthModePKIAllOf) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *AuthModePKIAllOf) SetAlgorithm(v string) {
	o.Algorithm = &v
}

func (o AuthModePKIAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableAuthModePKIAllOf struct {
	value *AuthModePKIAllOf
	isSet bool
}

func (v NullableAuthModePKIAllOf) Get() *AuthModePKIAllOf {
	return v.value
}

func (v *NullableAuthModePKIAllOf) Set(val *AuthModePKIAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthModePKIAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthModePKIAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthModePKIAllOf(val *AuthModePKIAllOf) *NullableAuthModePKIAllOf {
	return &NullableAuthModePKIAllOf{value: val, isSet: true}
}

func (v NullableAuthModePKIAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthModePKIAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


