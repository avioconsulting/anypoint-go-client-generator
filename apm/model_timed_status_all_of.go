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

// TimedStatusAllOf struct for TimedStatusAllOf
type TimedStatusAllOf struct {
	Id *string `json:"id,omitempty"`
	StartDate NullableString `json:"startDate,omitempty"`
	EndDate NullableString `json:"endDate,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewTimedStatusAllOf instantiates a new TimedStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimedStatusAllOf() *TimedStatusAllOf {
	this := TimedStatusAllOf{}
	var status string = "DRAFT"
	this.Status = &status
	return &this
}

// NewTimedStatusAllOfWithDefaults instantiates a new TimedStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimedStatusAllOfWithDefaults() *TimedStatusAllOf {
	this := TimedStatusAllOf{}
	var status string = "DRAFT"
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TimedStatusAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimedStatusAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TimedStatusAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TimedStatusAllOf) SetId(v string) {
	o.Id = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimedStatusAllOf) GetStartDate() string {
	if o == nil || o.StartDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimedStatusAllOf) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *TimedStatusAllOf) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableString and assigns it to the StartDate field.
func (o *TimedStatusAllOf) SetStartDate(v string) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *TimedStatusAllOf) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *TimedStatusAllOf) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimedStatusAllOf) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimedStatusAllOf) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *TimedStatusAllOf) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *TimedStatusAllOf) SetEndDate(v string) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *TimedStatusAllOf) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *TimedStatusAllOf) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TimedStatusAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimedStatusAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TimedStatusAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TimedStatusAllOf) SetStatus(v string) {
	o.Status = &v
}

func (o TimedStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.StartDate.IsSet() {
		toSerialize["startDate"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableTimedStatusAllOf struct {
	value *TimedStatusAllOf
	isSet bool
}

func (v NullableTimedStatusAllOf) Get() *TimedStatusAllOf {
	return v.value
}

func (v *NullableTimedStatusAllOf) Set(val *TimedStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTimedStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTimedStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimedStatusAllOf(val *TimedStatusAllOf) *NullableTimedStatusAllOf {
	return &NullableTimedStatusAllOf{value: val, isSet: true}
}

func (v NullableTimedStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimedStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


