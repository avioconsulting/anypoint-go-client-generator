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

// DeploymentStatusMessage struct for DeploymentStatusMessage
type DeploymentStatusMessage struct {
	Status DeploymentStatusType `json:"status"`
	Message *string `json:"message,omitempty"`
}

// NewDeploymentStatusMessage instantiates a new DeploymentStatusMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStatusMessage(status DeploymentStatusType) *DeploymentStatusMessage {
	this := DeploymentStatusMessage{}
	this.Status = status
	return &this
}

// NewDeploymentStatusMessageWithDefaults instantiates a new DeploymentStatusMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStatusMessageWithDefaults() *DeploymentStatusMessage {
	this := DeploymentStatusMessage{}
	return &this
}

// GetStatus returns the Status field value
func (o *DeploymentStatusMessage) GetStatus() DeploymentStatusType {
	if o == nil {
		var ret DeploymentStatusType
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeploymentStatusMessage) GetStatusOk() (*DeploymentStatusType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DeploymentStatusMessage) SetStatus(v DeploymentStatusType) {
	o.Status = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeploymentStatusMessage) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStatusMessage) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeploymentStatusMessage) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DeploymentStatusMessage) SetMessage(v string) {
	o.Message = &v
}

func (o DeploymentStatusMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentStatusMessage struct {
	value *DeploymentStatusMessage
	isSet bool
}

func (v NullableDeploymentStatusMessage) Get() *DeploymentStatusMessage {
	return v.value
}

func (v *NullableDeploymentStatusMessage) Set(val *DeploymentStatusMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStatusMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStatusMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStatusMessage(val *DeploymentStatusMessage) *NullableDeploymentStatusMessage {
	return &NullableDeploymentStatusMessage{value: val, isSet: true}
}

func (v NullableDeploymentStatusMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStatusMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

