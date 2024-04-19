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
	"time"
)

// DocumentFlowSummary struct for DocumentFlowSummary
type DocumentFlowSummary struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Direction *DirectionType `json:"direction,omitempty"`
	PartnerFrom *Partner `json:"partnerFrom,omitempty"`
	PartnerTo *Partner `json:"partnerTo,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	UpdatedBy *string `json:"updatedBy,omitempty"`
	DeploymentSummary *[]DocumentFlowDeploymentSummary `json:"deploymentSummary,omitempty"`
}

// NewDocumentFlowSummary instantiates a new DocumentFlowSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentFlowSummary() *DocumentFlowSummary {
	this := DocumentFlowSummary{}
	return &this
}

// NewDocumentFlowSummaryWithDefaults instantiates a new DocumentFlowSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentFlowSummaryWithDefaults() *DocumentFlowSummary {
	this := DocumentFlowSummary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentFlowSummary) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowSummary) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentFlowSummary) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocumentFlowSummary) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DocumentFlowSummary) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowSummary) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DocumentFlowSummary) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DocumentFlowSummary) SetName(v string) {
	o.Name = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *DocumentFlowSummary) GetDirection() DirectionType {
	if o == nil || o.Direction == nil {
		var ret DirectionType
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowSummary) GetDirectionOk() (*DirectionType, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *DocumentFlowSummary) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given DirectionType and assigns it to the Direction field.
func (o *DocumentFlowSummary) SetDirection(v DirectionType) {
	o.Direction = &v
}

// GetPartnerFrom returns the PartnerFrom field value if set, zero value otherwise.
func (o *DocumentFlowSummary) GetPartnerFrom() Partner {
	if o == nil || o.PartnerFrom == nil {
		var ret Partner
		return ret
	}
	return *o.PartnerFrom
}

// GetPartnerFromOk returns a tuple with the PartnerFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowSummary) GetPartnerFromOk() (*Partner, bool) {
	if o == nil || o.PartnerFrom == nil {
		return nil, false
	}
	return o.PartnerFrom, true
}

// HasPartnerFrom returns a boolean if a field has been set.
func (o *DocumentFlowSummary) HasPartnerFrom() bool {
	if o != nil && o.PartnerFrom != nil {
		return true
	}

	return false
}

// SetPartnerFrom gets a reference to the given Partner and assigns it to the PartnerFrom field.
func (o *DocumentFlowSummary) SetPartnerFrom(v Partner) {
	o.PartnerFrom = &v
}

// GetPartnerTo returns the PartnerTo field value if set, zero value otherwise.
func (o *DocumentFlowSummary) GetPartnerTo() Partner {
	if o == nil || o.PartnerTo == nil {
		var ret Partner
		return ret
	}
	return *o.PartnerTo
}

// GetPartnerToOk returns a tuple with the PartnerTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowSummary) GetPartnerToOk() (*Partner, bool) {
	if o == nil || o.PartnerTo == nil {
		return nil, false
	}
	return o.PartnerTo, true
}

// HasPartnerTo returns a boolean if a field has been set.
func (o *DocumentFlowSummary) HasPartnerTo() bool {
	if o != nil && o.PartnerTo != nil {
		return true
	}

	return false
}

// SetPartnerTo gets a reference to the given Partner and assigns it to the PartnerTo field.
func (o *DocumentFlowSummary) SetPartnerTo(v Partner) {
	o.PartnerTo = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DocumentFlowSummary) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowSummary) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DocumentFlowSummary) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DocumentFlowSummary) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DocumentFlowSummary) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowSummary) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DocumentFlowSummary) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *DocumentFlowSummary) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DocumentFlowSummary) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowSummary) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DocumentFlowSummary) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DocumentFlowSummary) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *DocumentFlowSummary) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowSummary) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *DocumentFlowSummary) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *DocumentFlowSummary) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetDeploymentSummary returns the DeploymentSummary field value if set, zero value otherwise.
func (o *DocumentFlowSummary) GetDeploymentSummary() []DocumentFlowDeploymentSummary {
	if o == nil || o.DeploymentSummary == nil {
		var ret []DocumentFlowDeploymentSummary
		return ret
	}
	return *o.DeploymentSummary
}

// GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowSummary) GetDeploymentSummaryOk() (*[]DocumentFlowDeploymentSummary, bool) {
	if o == nil || o.DeploymentSummary == nil {
		return nil, false
	}
	return o.DeploymentSummary, true
}

// HasDeploymentSummary returns a boolean if a field has been set.
func (o *DocumentFlowSummary) HasDeploymentSummary() bool {
	if o != nil && o.DeploymentSummary != nil {
		return true
	}

	return false
}

// SetDeploymentSummary gets a reference to the given []DocumentFlowDeploymentSummary and assigns it to the DeploymentSummary field.
func (o *DocumentFlowSummary) SetDeploymentSummary(v []DocumentFlowDeploymentSummary) {
	o.DeploymentSummary = &v
}

func (o DocumentFlowSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.PartnerFrom != nil {
		toSerialize["partnerFrom"] = o.PartnerFrom
	}
	if o.PartnerTo != nil {
		toSerialize["partnerTo"] = o.PartnerTo
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if o.DeploymentSummary != nil {
		toSerialize["deploymentSummary"] = o.DeploymentSummary
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentFlowSummary struct {
	value *DocumentFlowSummary
	isSet bool
}

func (v NullableDocumentFlowSummary) Get() *DocumentFlowSummary {
	return v.value
}

func (v *NullableDocumentFlowSummary) Set(val *DocumentFlowSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentFlowSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentFlowSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentFlowSummary(val *DocumentFlowSummary) *NullableDocumentFlowSummary {
	return &NullableDocumentFlowSummary{value: val, isSet: true}
}

func (v NullableDocumentFlowSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentFlowSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


