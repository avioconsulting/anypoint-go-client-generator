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

// AuthModeOAuth This can be used for HTTP Outbound authentication.
type AuthModeOAuth struct {
	AuthModeBaseConfig
	TokenUrl *string `json:"tokenUrl,omitempty"`
	ClientId *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
}

// NewAuthModeOAuth instantiates a new AuthModeOAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthModeOAuth(authType string) *AuthModeOAuth {
	this := AuthModeOAuth{}
	this.AuthType = authType
	return &this
}

// NewAuthModeOAuthWithDefaults instantiates a new AuthModeOAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthModeOAuthWithDefaults() *AuthModeOAuth {
	this := AuthModeOAuth{}
	return &this
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *AuthModeOAuth) GetTokenUrl() string {
	if o == nil || o.TokenUrl == nil {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthModeOAuth) GetTokenUrlOk() (*string, bool) {
	if o == nil || o.TokenUrl == nil {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *AuthModeOAuth) HasTokenUrl() bool {
	if o != nil && o.TokenUrl != nil {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *AuthModeOAuth) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AuthModeOAuth) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthModeOAuth) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AuthModeOAuth) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AuthModeOAuth) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *AuthModeOAuth) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthModeOAuth) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *AuthModeOAuth) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *AuthModeOAuth) SetClientSecret(v string) {
	o.ClientSecret = &v
}

func (o AuthModeOAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthModeBaseConfig, errAuthModeBaseConfig := json.Marshal(o.AuthModeBaseConfig)
	if errAuthModeBaseConfig != nil {
		return []byte{}, errAuthModeBaseConfig
	}
	errAuthModeBaseConfig = json.Unmarshal([]byte(serializedAuthModeBaseConfig), &toSerialize)
	if errAuthModeBaseConfig != nil {
		return []byte{}, errAuthModeBaseConfig
	}
	if o.TokenUrl != nil {
		toSerialize["tokenUrl"] = o.TokenUrl
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	return json.Marshal(toSerialize)
}

type NullableAuthModeOAuth struct {
	value *AuthModeOAuth
	isSet bool
}

func (v NullableAuthModeOAuth) Get() *AuthModeOAuth {
	return v.value
}

func (v *NullableAuthModeOAuth) Set(val *AuthModeOAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthModeOAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthModeOAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthModeOAuth(val *AuthModeOAuth) *NullableAuthModeOAuth {
	return &NullableAuthModeOAuth{value: val, isSet: true}
}

func (v NullableAuthModeOAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthModeOAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


