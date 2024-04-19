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

// EndpointDeployment struct for EndpointDeployment
type EndpointDeployment struct {
	Id *string `json:"id,omitempty"`
	Status string `json:"status"`
	Name string `json:"name"`
	RuntimeDeploymentId int32 `json:"runtimeDeploymentId"`
	Port int32 `json:"port"`
}

// NewEndpointDeployment instantiates a new EndpointDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointDeployment(status string, name string, runtimeDeploymentId int32, port int32) *EndpointDeployment {
	this := EndpointDeployment{}
	this.Status = status
	this.Name = name
	this.RuntimeDeploymentId = runtimeDeploymentId
	this.Port = port
	return &this
}

// NewEndpointDeploymentWithDefaults instantiates a new EndpointDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointDeploymentWithDefaults() *EndpointDeployment {
	this := EndpointDeployment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EndpointDeployment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointDeployment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EndpointDeployment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EndpointDeployment) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value
func (o *EndpointDeployment) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EndpointDeployment) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EndpointDeployment) SetStatus(v string) {
	o.Status = v
}

// GetName returns the Name field value
func (o *EndpointDeployment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EndpointDeployment) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EndpointDeployment) SetName(v string) {
	o.Name = v
}

// GetRuntimeDeploymentId returns the RuntimeDeploymentId field value
func (o *EndpointDeployment) GetRuntimeDeploymentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RuntimeDeploymentId
}

// GetRuntimeDeploymentIdOk returns a tuple with the RuntimeDeploymentId field value
// and a boolean to check if the value has been set.
func (o *EndpointDeployment) GetRuntimeDeploymentIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RuntimeDeploymentId, true
}

// SetRuntimeDeploymentId sets field value
func (o *EndpointDeployment) SetRuntimeDeploymentId(v int32) {
	o.RuntimeDeploymentId = v
}

// GetPort returns the Port field value
func (o *EndpointDeployment) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *EndpointDeployment) GetPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *EndpointDeployment) SetPort(v int32) {
	o.Port = v
}

func (o EndpointDeployment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["runtimeDeploymentId"] = o.RuntimeDeploymentId
	}
	if true {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointDeployment struct {
	value *EndpointDeployment
	isSet bool
}

func (v NullableEndpointDeployment) Get() *EndpointDeployment {
	return v.value
}

func (v *NullableEndpointDeployment) Set(val *EndpointDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointDeployment(val *EndpointDeployment) *NullableEndpointDeployment {
	return &NullableEndpointDeployment{value: val, isSet: true}
}

func (v NullableEndpointDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


