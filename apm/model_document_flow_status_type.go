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

// DocumentFlowStatusType the model 'DocumentFlowStatusType'
type DocumentFlowStatusType string

// List of documentFlowStatusType
const (
	DRAFT DocumentFlowStatusType = "DRAFT"
	DEPLOYED DocumentFlowStatusType = "DEPLOYED"
	ALL DocumentFlowStatusType = "ALL"
)

func (v *DocumentFlowStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DocumentFlowStatusType(value)
	for _, existing := range []DocumentFlowStatusType{ "DRAFT", "DEPLOYED", "ALL",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DocumentFlowStatusType", value)
}

// Ptr returns reference to documentFlowStatusType value
func (v DocumentFlowStatusType) Ptr() *DocumentFlowStatusType {
	return &v
}

type NullableDocumentFlowStatusType struct {
	value *DocumentFlowStatusType
	isSet bool
}

func (v NullableDocumentFlowStatusType) Get() *DocumentFlowStatusType {
	return v.value
}

func (v *NullableDocumentFlowStatusType) Set(val *DocumentFlowStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentFlowStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentFlowStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentFlowStatusType(val *DocumentFlowStatusType) *NullableDocumentFlowStatusType {
	return &NullableDocumentFlowStatusType{value: val, isSet: true}
}

func (v NullableDocumentFlowStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentFlowStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

