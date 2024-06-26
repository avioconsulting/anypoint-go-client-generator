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

// AuthModeClientCredentialsAllOf struct for AuthModeClientCredentialsAllOf
type AuthModeClientCredentialsAllOf struct {
	ClientId *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
}

// NewAuthModeClientCredentialsAllOf instantiates a new AuthModeClientCredentialsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthModeClientCredentialsAllOf() *AuthModeClientCredentialsAllOf {
	this := AuthModeClientCredentialsAllOf{}
	return &this
}

// NewAuthModeClientCredentialsAllOfWithDefaults instantiates a new AuthModeClientCredentialsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthModeClientCredentialsAllOfWithDefaults() *AuthModeClientCredentialsAllOf {
	this := AuthModeClientCredentialsAllOf{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AuthModeClientCredentialsAllOf) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthModeClientCredentialsAllOf) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AuthModeClientCredentialsAllOf) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AuthModeClientCredentialsAllOf) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *AuthModeClientCredentialsAllOf) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthModeClientCredentialsAllOf) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *AuthModeClientCredentialsAllOf) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *AuthModeClientCredentialsAllOf) SetClientSecret(v string) {
	o.ClientSecret = &v
}

func (o AuthModeClientCredentialsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	return json.Marshal(toSerialize)
}

type NullableAuthModeClientCredentialsAllOf struct {
	value *AuthModeClientCredentialsAllOf
	isSet bool
}

func (v NullableAuthModeClientCredentialsAllOf) Get() *AuthModeClientCredentialsAllOf {
	return v.value
}

func (v *NullableAuthModeClientCredentialsAllOf) Set(val *AuthModeClientCredentialsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthModeClientCredentialsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthModeClientCredentialsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthModeClientCredentialsAllOf(val *AuthModeClientCredentialsAllOf) *NullableAuthModeClientCredentialsAllOf {
	return &NullableAuthModeClientCredentialsAllOf{value: val, isSet: true}
}

func (v NullableAuthModeClientCredentialsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthModeClientCredentialsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


