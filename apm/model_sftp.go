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

// Sftp SFTP Endpoint Configuration
type Sftp struct {
	BaseEndpoint
	Config SftpConfig `json:"config"`
}

// NewSftp instantiates a new Sftp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSftp(name string, description string, environmentId string, endpointTypeId string, endpointRole EndpointRoleType, config SftpConfig) *Sftp {
	this := Sftp{}
	this.Name = name
	this.Description = description
	this.EnvironmentId = environmentId
	this.EndpointTypeId = endpointTypeId
	this.EndpointRole = endpointRole
	this.Config = config
	return &this
}

// NewSftpWithDefaults instantiates a new Sftp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSftpWithDefaults() *Sftp {
	this := Sftp{}
	return &this
}

// GetConfig returns the Config field value
func (o *Sftp) GetConfig() SftpConfig {
	if o == nil {
		var ret SftpConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *Sftp) GetConfigOk() (*SftpConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *Sftp) SetConfig(v SftpConfig) {
	o.Config = v
}

func (o Sftp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseEndpoint, errBaseEndpoint := json.Marshal(o.BaseEndpoint)
	if errBaseEndpoint != nil {
		return []byte{}, errBaseEndpoint
	}
	errBaseEndpoint = json.Unmarshal([]byte(serializedBaseEndpoint), &toSerialize)
	if errBaseEndpoint != nil {
		return []byte{}, errBaseEndpoint
	}
	if true {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableSftp struct {
	value *Sftp
	isSet bool
}

func (v NullableSftp) Get() *Sftp {
	return v.value
}

func (v *NullableSftp) Set(val *Sftp) {
	v.value = val
	v.isSet = true
}

func (v NullableSftp) IsSet() bool {
	return v.isSet
}

func (v *NullableSftp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSftp(val *Sftp) *NullableSftp {
	return &NullableSftp{value: val, isSet: true}
}

func (v NullableSftp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSftp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


