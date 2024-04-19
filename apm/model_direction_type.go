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
	"fmt"
)

// DirectionType the model 'DirectionType'
type DirectionType string

// List of directionType
const (
	INBOUND DirectionType = "INBOUND"
	OUTBOUND DirectionType = "OUTBOUND"
)

func (v *DirectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DirectionType(value)
	for _, existing := range []DirectionType{ "INBOUND", "OUTBOUND",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DirectionType", value)
}

// Ptr returns reference to directionType value
func (v DirectionType) Ptr() *DirectionType {
	return &v
}

type NullableDirectionType struct {
	value *DirectionType
	isSet bool
}

func (v NullableDirectionType) Get() *DirectionType {
	return v.value
}

func (v *NullableDirectionType) Set(val *DirectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectionType(val *DirectionType) *NullableDirectionType {
	return &NullableDirectionType{value: val, isSet: true}
}

func (v NullableDirectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
