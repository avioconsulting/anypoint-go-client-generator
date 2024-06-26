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

// BaseEndpointConfig struct for BaseEndpointConfig
type BaseEndpointConfig struct {
	ConfigName string `json:"configName"`
	// Defines Server/Host Address
	ServerAddress *string `json:"serverAddress,omitempty"`
	ServerPort *int32 `json:"serverPort,omitempty"`
	AuthMode *AuthModeBaseConfig `json:"authMode,omitempty"`
	Path *string `json:"path,omitempty"`
	MovedPath NullableString `json:"movedPath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseEndpointConfig BaseEndpointConfig

// NewBaseEndpointConfig instantiates a new BaseEndpointConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEndpointConfig(configName string) *BaseEndpointConfig {
	this := BaseEndpointConfig{}
	this.ConfigName = configName
	return &this
}

// NewBaseEndpointConfigWithDefaults instantiates a new BaseEndpointConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEndpointConfigWithDefaults() *BaseEndpointConfig {
	this := BaseEndpointConfig{}
	return &this
}

// GetConfigName returns the ConfigName field value
func (o *BaseEndpointConfig) GetConfigName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigName
}

// GetConfigNameOk returns a tuple with the ConfigName field value
// and a boolean to check if the value has been set.
func (o *BaseEndpointConfig) GetConfigNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigName, true
}

// SetConfigName sets field value
func (o *BaseEndpointConfig) SetConfigName(v string) {
	o.ConfigName = v
}

// GetServerAddress returns the ServerAddress field value if set, zero value otherwise.
func (o *BaseEndpointConfig) GetServerAddress() string {
	if o == nil || o.ServerAddress == nil {
		var ret string
		return ret
	}
	return *o.ServerAddress
}

// GetServerAddressOk returns a tuple with the ServerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEndpointConfig) GetServerAddressOk() (*string, bool) {
	if o == nil || o.ServerAddress == nil {
		return nil, false
	}
	return o.ServerAddress, true
}

// HasServerAddress returns a boolean if a field has been set.
func (o *BaseEndpointConfig) HasServerAddress() bool {
	if o != nil && o.ServerAddress != nil {
		return true
	}

	return false
}

// SetServerAddress gets a reference to the given string and assigns it to the ServerAddress field.
func (o *BaseEndpointConfig) SetServerAddress(v string) {
	o.ServerAddress = &v
}

// GetServerPort returns the ServerPort field value if set, zero value otherwise.
func (o *BaseEndpointConfig) GetServerPort() int32 {
	if o == nil || o.ServerPort == nil {
		var ret int32
		return ret
	}
	return *o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEndpointConfig) GetServerPortOk() (*int32, bool) {
	if o == nil || o.ServerPort == nil {
		return nil, false
	}
	return o.ServerPort, true
}

// HasServerPort returns a boolean if a field has been set.
func (o *BaseEndpointConfig) HasServerPort() bool {
	if o != nil && o.ServerPort != nil {
		return true
	}

	return false
}

// SetServerPort gets a reference to the given int32 and assigns it to the ServerPort field.
func (o *BaseEndpointConfig) SetServerPort(v int32) {
	o.ServerPort = &v
}

// GetAuthMode returns the AuthMode field value if set, zero value otherwise.
func (o *BaseEndpointConfig) GetAuthMode() AuthModeBaseConfig {
	if o == nil || o.AuthMode == nil {
		var ret AuthModeBaseConfig
		return ret
	}
	return *o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEndpointConfig) GetAuthModeOk() (*AuthModeBaseConfig, bool) {
	if o == nil || o.AuthMode == nil {
		return nil, false
	}
	return o.AuthMode, true
}

// HasAuthMode returns a boolean if a field has been set.
func (o *BaseEndpointConfig) HasAuthMode() bool {
	if o != nil && o.AuthMode != nil {
		return true
	}

	return false
}

// SetAuthMode gets a reference to the given AuthModeBaseConfig and assigns it to the AuthMode field.
func (o *BaseEndpointConfig) SetAuthMode(v AuthModeBaseConfig) {
	o.AuthMode = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BaseEndpointConfig) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEndpointConfig) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BaseEndpointConfig) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *BaseEndpointConfig) SetPath(v string) {
	o.Path = &v
}

// GetMovedPath returns the MovedPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseEndpointConfig) GetMovedPath() string {
	if o == nil || o.MovedPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.MovedPath.Get()
}

// GetMovedPathOk returns a tuple with the MovedPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseEndpointConfig) GetMovedPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MovedPath.Get(), o.MovedPath.IsSet()
}

// HasMovedPath returns a boolean if a field has been set.
func (o *BaseEndpointConfig) HasMovedPath() bool {
	if o != nil && o.MovedPath.IsSet() {
		return true
	}

	return false
}

// SetMovedPath gets a reference to the given NullableString and assigns it to the MovedPath field.
func (o *BaseEndpointConfig) SetMovedPath(v string) {
	o.MovedPath.Set(&v)
}
// SetMovedPathNil sets the value for MovedPath to be an explicit nil
func (o *BaseEndpointConfig) SetMovedPathNil() {
	o.MovedPath.Set(nil)
}

// UnsetMovedPath ensures that no value is present for MovedPath, not even an explicit nil
func (o *BaseEndpointConfig) UnsetMovedPath() {
	o.MovedPath.Unset()
}

func (o BaseEndpointConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["configName"] = o.ConfigName
	}
	if o.ServerAddress != nil {
		toSerialize["serverAddress"] = o.ServerAddress
	}
	if o.ServerPort != nil {
		toSerialize["serverPort"] = o.ServerPort
	}
	if o.AuthMode != nil {
		toSerialize["authMode"] = o.AuthMode
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.MovedPath.IsSet() {
		toSerialize["movedPath"] = o.MovedPath.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BaseEndpointConfig) UnmarshalJSON(bytes []byte) (err error) {
	varBaseEndpointConfig := _BaseEndpointConfig{}

	if err = json.Unmarshal(bytes, &varBaseEndpointConfig); err == nil {
		*o = BaseEndpointConfig(varBaseEndpointConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "configName")
		delete(additionalProperties, "serverAddress")
		delete(additionalProperties, "serverPort")
		delete(additionalProperties, "authMode")
		delete(additionalProperties, "path")
		delete(additionalProperties, "movedPath")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseEndpointConfig struct {
	value *BaseEndpointConfig
	isSet bool
}

func (v NullableBaseEndpointConfig) Get() *BaseEndpointConfig {
	return v.value
}

func (v *NullableBaseEndpointConfig) Set(val *BaseEndpointConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEndpointConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEndpointConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEndpointConfig(val *BaseEndpointConfig) *NullableBaseEndpointConfig {
	return &NullableBaseEndpointConfig{value: val, isSet: true}
}

func (v NullableBaseEndpointConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEndpointConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


