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

// EdiDocumentType struct for EdiDocumentType
type EdiDocumentType struct {
	Id *string `json:"id,omitempty"`
	DocumentName string `json:"documentName"`
	Label string `json:"label"`
	SchemaPath string `json:"schemaPath"`
	Description string `json:"description"`
	FormatType string `json:"formatType"`
	Version string `json:"version"`
}

// NewEdiDocumentType instantiates a new EdiDocumentType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdiDocumentType(documentName string, label string, schemaPath string, description string, formatType string, version string) *EdiDocumentType {
	this := EdiDocumentType{}
	this.DocumentName = documentName
	this.Label = label
	this.SchemaPath = schemaPath
	this.Description = description
	this.FormatType = formatType
	this.Version = version
	return &this
}

// NewEdiDocumentTypeWithDefaults instantiates a new EdiDocumentType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdiDocumentTypeWithDefaults() *EdiDocumentType {
	this := EdiDocumentType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EdiDocumentType) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiDocumentType) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EdiDocumentType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EdiDocumentType) SetId(v string) {
	o.Id = &v
}

// GetDocumentName returns the DocumentName field value
func (o *EdiDocumentType) GetDocumentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentName
}

// GetDocumentNameOk returns a tuple with the DocumentName field value
// and a boolean to check if the value has been set.
func (o *EdiDocumentType) GetDocumentNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentName, true
}

// SetDocumentName sets field value
func (o *EdiDocumentType) SetDocumentName(v string) {
	o.DocumentName = v
}

// GetLabel returns the Label field value
func (o *EdiDocumentType) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *EdiDocumentType) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *EdiDocumentType) SetLabel(v string) {
	o.Label = v
}

// GetSchemaPath returns the SchemaPath field value
func (o *EdiDocumentType) GetSchemaPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaPath
}

// GetSchemaPathOk returns a tuple with the SchemaPath field value
// and a boolean to check if the value has been set.
func (o *EdiDocumentType) GetSchemaPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SchemaPath, true
}

// SetSchemaPath sets field value
func (o *EdiDocumentType) SetSchemaPath(v string) {
	o.SchemaPath = v
}

// GetDescription returns the Description field value
func (o *EdiDocumentType) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EdiDocumentType) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EdiDocumentType) SetDescription(v string) {
	o.Description = v
}

// GetFormatType returns the FormatType field value
func (o *EdiDocumentType) GetFormatType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormatType
}

// GetFormatTypeOk returns a tuple with the FormatType field value
// and a boolean to check if the value has been set.
func (o *EdiDocumentType) GetFormatTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FormatType, true
}

// SetFormatType sets field value
func (o *EdiDocumentType) SetFormatType(v string) {
	o.FormatType = v
}

// GetVersion returns the Version field value
func (o *EdiDocumentType) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EdiDocumentType) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *EdiDocumentType) SetVersion(v string) {
	o.Version = v
}

func (o EdiDocumentType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["documentName"] = o.DocumentName
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["schemaPath"] = o.SchemaPath
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["formatType"] = o.FormatType
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableEdiDocumentType struct {
	value *EdiDocumentType
	isSet bool
}

func (v NullableEdiDocumentType) Get() *EdiDocumentType {
	return v.value
}

func (v *NullableEdiDocumentType) Set(val *EdiDocumentType) {
	v.value = val
	v.isSet = true
}

func (v NullableEdiDocumentType) IsSet() bool {
	return v.isSet
}

func (v *NullableEdiDocumentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdiDocumentType(val *EdiDocumentType) *NullableEdiDocumentType {
	return &NullableEdiDocumentType{value: val, isSet: true}
}

func (v NullableEdiDocumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdiDocumentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


