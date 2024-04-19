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

// TLSContext This can be used to configure the TLS context of HTTPS endpoints
type TLSContext struct {
	Insecure bool `json:"insecure"`
	NeedCertificate bool `json:"needCertificate"`
	CertificateId *string `json:"certificateId,omitempty"`
}

// NewTLSContext instantiates a new TLSContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSContext(insecure bool, needCertificate bool) *TLSContext {
	this := TLSContext{}
	this.Insecure = insecure
	this.NeedCertificate = needCertificate
	var certificateId string = "true"
	this.CertificateId = &certificateId
	return &this
}

// NewTLSContextWithDefaults instantiates a new TLSContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSContextWithDefaults() *TLSContext {
	this := TLSContext{}
	var insecure bool = false
	this.Insecure = insecure
	var needCertificate bool = true
	this.NeedCertificate = needCertificate
	var certificateId string = "true"
	this.CertificateId = &certificateId
	return &this
}

// GetInsecure returns the Insecure field value
func (o *TLSContext) GetInsecure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Insecure
}

// GetInsecureOk returns a tuple with the Insecure field value
// and a boolean to check if the value has been set.
func (o *TLSContext) GetInsecureOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Insecure, true
}

// SetInsecure sets field value
func (o *TLSContext) SetInsecure(v bool) {
	o.Insecure = v
}

// GetNeedCertificate returns the NeedCertificate field value
func (o *TLSContext) GetNeedCertificate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NeedCertificate
}

// GetNeedCertificateOk returns a tuple with the NeedCertificate field value
// and a boolean to check if the value has been set.
func (o *TLSContext) GetNeedCertificateOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NeedCertificate, true
}

// SetNeedCertificate sets field value
func (o *TLSContext) SetNeedCertificate(v bool) {
	o.NeedCertificate = v
}

// GetCertificateId returns the CertificateId field value if set, zero value otherwise.
func (o *TLSContext) GetCertificateId() string {
	if o == nil || o.CertificateId == nil {
		var ret string
		return ret
	}
	return *o.CertificateId
}

// GetCertificateIdOk returns a tuple with the CertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSContext) GetCertificateIdOk() (*string, bool) {
	if o == nil || o.CertificateId == nil {
		return nil, false
	}
	return o.CertificateId, true
}

// HasCertificateId returns a boolean if a field has been set.
func (o *TLSContext) HasCertificateId() bool {
	if o != nil && o.CertificateId != nil {
		return true
	}

	return false
}

// SetCertificateId gets a reference to the given string and assigns it to the CertificateId field.
func (o *TLSContext) SetCertificateId(v string) {
	o.CertificateId = &v
}

func (o TLSContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["insecure"] = o.Insecure
	}
	if true {
		toSerialize["needCertificate"] = o.NeedCertificate
	}
	if o.CertificateId != nil {
		toSerialize["certificateId"] = o.CertificateId
	}
	return json.Marshal(toSerialize)
}

type NullableTLSContext struct {
	value *TLSContext
	isSet bool
}

func (v NullableTLSContext) Get() *TLSContext {
	return v.value
}

func (v *NullableTLSContext) Set(val *TLSContext) {
	v.value = val
	v.isSet = true
}

func (v NullableTLSContext) IsSet() bool {
	return v.isSet
}

func (v *NullableTLSContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTLSContext(val *TLSContext) *NullableTLSContext {
	return &NullableTLSContext{value: val, isSet: true}
}

func (v NullableTLSContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTLSContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


