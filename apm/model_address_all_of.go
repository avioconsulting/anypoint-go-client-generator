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

// AddressAllOf struct for AddressAllOf
type AddressAllOf struct {
	Id *string `json:"id,omitempty"`
	AddressLine1 *string `json:"addressLine_1,omitempty"`
	AddressLine2 *string `json:"addressLine_2,omitempty"`
	City *string `json:"city,omitempty"`
	State *string `json:"state,omitempty"`
	Country *string `json:"country,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
}

// NewAddressAllOf instantiates a new AddressAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressAllOf() *AddressAllOf {
	this := AddressAllOf{}
	return &this
}

// NewAddressAllOfWithDefaults instantiates a new AddressAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressAllOfWithDefaults() *AddressAllOf {
	this := AddressAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AddressAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AddressAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AddressAllOf) SetId(v string) {
	o.Id = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *AddressAllOf) GetAddressLine1() string {
	if o == nil || o.AddressLine1 == nil {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf) GetAddressLine1Ok() (*string, bool) {
	if o == nil || o.AddressLine1 == nil {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *AddressAllOf) HasAddressLine1() bool {
	if o != nil && o.AddressLine1 != nil {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *AddressAllOf) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *AddressAllOf) GetAddressLine2() string {
	if o == nil || o.AddressLine2 == nil {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf) GetAddressLine2Ok() (*string, bool) {
	if o == nil || o.AddressLine2 == nil {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *AddressAllOf) HasAddressLine2() bool {
	if o != nil && o.AddressLine2 != nil {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *AddressAllOf) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AddressAllOf) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AddressAllOf) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AddressAllOf) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AddressAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AddressAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AddressAllOf) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AddressAllOf) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AddressAllOf) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *AddressAllOf) SetCountry(v string) {
	o.Country = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AddressAllOf) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AddressAllOf) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AddressAllOf) SetPostalCode(v string) {
	o.PostalCode = &v
}

func (o AddressAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AddressLine1 != nil {
		toSerialize["addressLine_1"] = o.AddressLine1
	}
	if o.AddressLine2 != nil {
		toSerialize["addressLine_2"] = o.AddressLine2
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.PostalCode != nil {
		toSerialize["postalCode"] = o.PostalCode
	}
	return json.Marshal(toSerialize)
}

type NullableAddressAllOf struct {
	value *AddressAllOf
	isSet bool
}

func (v NullableAddressAllOf) Get() *AddressAllOf {
	return v.value
}

func (v *NullableAddressAllOf) Set(val *AddressAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressAllOf(val *AddressAllOf) *NullableAddressAllOf {
	return &NullableAddressAllOf{value: val, isSet: true}
}

func (v NullableAddressAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


