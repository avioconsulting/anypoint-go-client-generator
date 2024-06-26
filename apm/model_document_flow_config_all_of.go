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

// DocumentFlowConfigAllOf struct for DocumentFlowConfigAllOf
type DocumentFlowConfigAllOf struct {
	Id *string `json:"id,omitempty"`
	DocumentFlowId *string `json:"documentFlowId,omitempty"`
	EnvironmentId *string `json:"environmentId,omitempty"`
	Status string `json:"status"`
	Version *int32 `json:"version,omitempty"`
	PreProcessingEndpointId *string `json:"preProcessingEndpointId,omitempty"`
	ReceivingEndpointId *string `json:"receivingEndpointId,omitempty"`
	ReceivingAckEndpointId *string `json:"receivingAckEndpointId,omitempty"`
	TargetEndpointId *string `json:"targetEndpointId,omitempty"`
	SourceDocTypeId *string `json:"sourceDocTypeId,omitempty"`
	TargetDocTypeId *string `json:"targetDocTypeId,omitempty"`
	DocumentMapping *[]DocumentMappingConfig `json:"documentMapping,omitempty"`
	ReceivingAckConfig *BaseAckConfig `json:"receivingAckConfig,omitempty"`
	PartnerIdentifiersConfig *DocumentFlowConfigAllOfPartnerIdentifiersConfig `json:"partnerIdentifiersConfig,omitempty"`
}

// NewDocumentFlowConfigAllOf instantiates a new DocumentFlowConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentFlowConfigAllOf(status string) *DocumentFlowConfigAllOf {
	this := DocumentFlowConfigAllOf{}
	this.Status = status
	return &this
}

// NewDocumentFlowConfigAllOfWithDefaults instantiates a new DocumentFlowConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentFlowConfigAllOfWithDefaults() *DocumentFlowConfigAllOf {
	this := DocumentFlowConfigAllOf{}
	var status string = "DRAFT"
	this.Status = status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocumentFlowConfigAllOf) SetId(v string) {
	o.Id = &v
}

// GetDocumentFlowId returns the DocumentFlowId field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOf) GetDocumentFlowId() string {
	if o == nil || o.DocumentFlowId == nil {
		var ret string
		return ret
	}
	return *o.DocumentFlowId
}

// GetDocumentFlowIdOk returns a tuple with the DocumentFlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOf) GetDocumentFlowIdOk() (*string, bool) {
	if o == nil || o.DocumentFlowId == nil {
		return nil, false
	}
	return o.DocumentFlowId, true
}

// HasDocumentFlowId returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOf) HasDocumentFlowId() bool {
	if o != nil && o.DocumentFlowId != nil {
		return true
	}

	return false
}

// SetDocumentFlowId gets a reference to the given string and assigns it to the DocumentFlowId field.
func (o *DocumentFlowConfigAllOf) SetDocumentFlowId(v string) {
	o.DocumentFlowId = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOf) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOf) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOf) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId != nil {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *DocumentFlowConfigAllOf) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetStatus returns the Status field value
func (o *DocumentFlowConfigAllOf) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOf) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DocumentFlowConfigAllOf) SetStatus(v string) {
	o.Status = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOf) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOf) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *DocumentFlowConfigAllOf) SetVersion(v int32) {
	o.Version = &v
}

// GetPreProcessingEndpointId returns the PreProcessingEndpointId field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOf) GetPreProcessingEndpointId() string {
	if o == nil || o.PreProcessingEndpointId == nil {
		var ret string
		return ret
	}
	return *o.PreProcessingEndpointId
}

// GetPreProcessingEndpointIdOk returns a tuple with the PreProcessingEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOf) GetPreProcessingEndpointIdOk() (*string, bool) {
	if o == nil || o.PreProcessingEndpointId == nil {
		return nil, false
	}
	return o.PreProcessingEndpointId, true
}

// HasPreProcessingEndpointId returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOf) HasPreProcessingEndpointId() bool {
	if o != nil && o.PreProcessingEndpointId != nil {
		return true
	}

	return false
}

// SetPreProcessingEndpointId gets a reference to the given string and assigns it to the PreProcessingEndpointId field.
func (o *DocumentFlowConfigAllOf) SetPreProcessingEndpointId(v string) {
	o.PreProcessingEndpointId = &v
}

// GetReceivingEndpointId returns the ReceivingEndpointId field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOf) GetReceivingEndpointId() string {
	if o == nil || o.ReceivingEndpointId == nil {
		var ret string
		return ret
	}
	return *o.ReceivingEndpointId
}

// GetReceivingEndpointIdOk returns a tuple with the ReceivingEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOf) GetReceivingEndpointIdOk() (*string, bool) {
	if o == nil || o.ReceivingEndpointId == nil {
		return nil, false
	}
	return o.ReceivingEndpointId, true
}

// HasReceivingEndpointId returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOf) HasReceivingEndpointId() bool {
	if o != nil && o.ReceivingEndpointId != nil {
		return true
	}

	return false
}

// SetReceivingEndpointId gets a reference to the given string and assigns it to the ReceivingEndpointId field.
func (o *DocumentFlowConfigAllOf) SetReceivingEndpointId(v string) {
	o.ReceivingEndpointId = &v
}

// GetReceivingAckEndpointId returns the ReceivingAckEndpointId field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOf) GetReceivingAckEndpointId() string {
	if o == nil || o.ReceivingAckEndpointId == nil {
		var ret string
		return ret
	}
	return *o.ReceivingAckEndpointId
}

// GetReceivingAckEndpointIdOk returns a tuple with the ReceivingAckEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOf) GetReceivingAckEndpointIdOk() (*string, bool) {
	if o == nil || o.ReceivingAckEndpointId == nil {
		return nil, false
	}
	return o.ReceivingAckEndpointId, true
}

// HasReceivingAckEndpointId returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOf) HasReceivingAckEndpointId() bool {
	if o != nil && o.ReceivingAckEndpointId != nil {
		return true
	}

	return false
}

// SetReceivingAckEndpointId gets a reference to the given string and assigns it to the ReceivingAckEndpointId field.
func (o *DocumentFlowConfigAllOf) SetReceivingAckEndpointId(v string) {
	o.ReceivingAckEndpointId = &v
}

// GetTargetEndpointId returns the TargetEndpointId field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOf) GetTargetEndpointId() string {
	if o == nil || o.TargetEndpointId == nil {
		var ret string
		return ret
	}
	return *o.TargetEndpointId
}

// GetTargetEndpointIdOk returns a tuple with the TargetEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOf) GetTargetEndpointIdOk() (*string, bool) {
	if o == nil || o.TargetEndpointId == nil {
		return nil, false
	}
	return o.TargetEndpointId, true
}

// HasTargetEndpointId returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOf) HasTargetEndpointId() bool {
	if o != nil && o.TargetEndpointId != nil {
		return true
	}

	return false
}

// SetTargetEndpointId gets a reference to the given string and assigns it to the TargetEndpointId field.
func (o *DocumentFlowConfigAllOf) SetTargetEndpointId(v string) {
	o.TargetEndpointId = &v
}

// GetSourceDocTypeId returns the SourceDocTypeId field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOf) GetSourceDocTypeId() string {
	if o == nil || o.SourceDocTypeId == nil {
		var ret string
		return ret
	}
	return *o.SourceDocTypeId
}

// GetSourceDocTypeIdOk returns a tuple with the SourceDocTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOf) GetSourceDocTypeIdOk() (*string, bool) {
	if o == nil || o.SourceDocTypeId == nil {
		return nil, false
	}
	return o.SourceDocTypeId, true
}

// HasSourceDocTypeId returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOf) HasSourceDocTypeId() bool {
	if o != nil && o.SourceDocTypeId != nil {
		return true
	}

	return false
}

// SetSourceDocTypeId gets a reference to the given string and assigns it to the SourceDocTypeId field.
func (o *DocumentFlowConfigAllOf) SetSourceDocTypeId(v string) {
	o.SourceDocTypeId = &v
}

// GetTargetDocTypeId returns the TargetDocTypeId field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOf) GetTargetDocTypeId() string {
	if o == nil || o.TargetDocTypeId == nil {
		var ret string
		return ret
	}
	return *o.TargetDocTypeId
}

// GetTargetDocTypeIdOk returns a tuple with the TargetDocTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOf) GetTargetDocTypeIdOk() (*string, bool) {
	if o == nil || o.TargetDocTypeId == nil {
		return nil, false
	}
	return o.TargetDocTypeId, true
}

// HasTargetDocTypeId returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOf) HasTargetDocTypeId() bool {
	if o != nil && o.TargetDocTypeId != nil {
		return true
	}

	return false
}

// SetTargetDocTypeId gets a reference to the given string and assigns it to the TargetDocTypeId field.
func (o *DocumentFlowConfigAllOf) SetTargetDocTypeId(v string) {
	o.TargetDocTypeId = &v
}

// GetDocumentMapping returns the DocumentMapping field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOf) GetDocumentMapping() []DocumentMappingConfig {
	if o == nil || o.DocumentMapping == nil {
		var ret []DocumentMappingConfig
		return ret
	}
	return *o.DocumentMapping
}

// GetDocumentMappingOk returns a tuple with the DocumentMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOf) GetDocumentMappingOk() (*[]DocumentMappingConfig, bool) {
	if o == nil || o.DocumentMapping == nil {
		return nil, false
	}
	return o.DocumentMapping, true
}

// HasDocumentMapping returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOf) HasDocumentMapping() bool {
	if o != nil && o.DocumentMapping != nil {
		return true
	}

	return false
}

// SetDocumentMapping gets a reference to the given []DocumentMappingConfig and assigns it to the DocumentMapping field.
func (o *DocumentFlowConfigAllOf) SetDocumentMapping(v []DocumentMappingConfig) {
	o.DocumentMapping = &v
}

// GetReceivingAckConfig returns the ReceivingAckConfig field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOf) GetReceivingAckConfig() BaseAckConfig {
	if o == nil || o.ReceivingAckConfig == nil {
		var ret BaseAckConfig
		return ret
	}
	return *o.ReceivingAckConfig
}

// GetReceivingAckConfigOk returns a tuple with the ReceivingAckConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOf) GetReceivingAckConfigOk() (*BaseAckConfig, bool) {
	if o == nil || o.ReceivingAckConfig == nil {
		return nil, false
	}
	return o.ReceivingAckConfig, true
}

// HasReceivingAckConfig returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOf) HasReceivingAckConfig() bool {
	if o != nil && o.ReceivingAckConfig != nil {
		return true
	}

	return false
}

// SetReceivingAckConfig gets a reference to the given BaseAckConfig and assigns it to the ReceivingAckConfig field.
func (o *DocumentFlowConfigAllOf) SetReceivingAckConfig(v BaseAckConfig) {
	o.ReceivingAckConfig = &v
}

// GetPartnerIdentifiersConfig returns the PartnerIdentifiersConfig field value if set, zero value otherwise.
func (o *DocumentFlowConfigAllOf) GetPartnerIdentifiersConfig() DocumentFlowConfigAllOfPartnerIdentifiersConfig {
	if o == nil || o.PartnerIdentifiersConfig == nil {
		var ret DocumentFlowConfigAllOfPartnerIdentifiersConfig
		return ret
	}
	return *o.PartnerIdentifiersConfig
}

// GetPartnerIdentifiersConfigOk returns a tuple with the PartnerIdentifiersConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentFlowConfigAllOf) GetPartnerIdentifiersConfigOk() (*DocumentFlowConfigAllOfPartnerIdentifiersConfig, bool) {
	if o == nil || o.PartnerIdentifiersConfig == nil {
		return nil, false
	}
	return o.PartnerIdentifiersConfig, true
}

// HasPartnerIdentifiersConfig returns a boolean if a field has been set.
func (o *DocumentFlowConfigAllOf) HasPartnerIdentifiersConfig() bool {
	if o != nil && o.PartnerIdentifiersConfig != nil {
		return true
	}

	return false
}

// SetPartnerIdentifiersConfig gets a reference to the given DocumentFlowConfigAllOfPartnerIdentifiersConfig and assigns it to the PartnerIdentifiersConfig field.
func (o *DocumentFlowConfigAllOf) SetPartnerIdentifiersConfig(v DocumentFlowConfigAllOfPartnerIdentifiersConfig) {
	o.PartnerIdentifiersConfig = &v
}

func (o DocumentFlowConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DocumentFlowId != nil {
		toSerialize["documentFlowId"] = o.DocumentFlowId
	}
	if o.EnvironmentId != nil {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.PreProcessingEndpointId != nil {
		toSerialize["preProcessingEndpointId"] = o.PreProcessingEndpointId
	}
	if o.ReceivingEndpointId != nil {
		toSerialize["receivingEndpointId"] = o.ReceivingEndpointId
	}
	if o.ReceivingAckEndpointId != nil {
		toSerialize["receivingAckEndpointId"] = o.ReceivingAckEndpointId
	}
	if o.TargetEndpointId != nil {
		toSerialize["targetEndpointId"] = o.TargetEndpointId
	}
	if o.SourceDocTypeId != nil {
		toSerialize["sourceDocTypeId"] = o.SourceDocTypeId
	}
	if o.TargetDocTypeId != nil {
		toSerialize["targetDocTypeId"] = o.TargetDocTypeId
	}
	if o.DocumentMapping != nil {
		toSerialize["documentMapping"] = o.DocumentMapping
	}
	if o.ReceivingAckConfig != nil {
		toSerialize["receivingAckConfig"] = o.ReceivingAckConfig
	}
	if o.PartnerIdentifiersConfig != nil {
		toSerialize["partnerIdentifiersConfig"] = o.PartnerIdentifiersConfig
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentFlowConfigAllOf struct {
	value *DocumentFlowConfigAllOf
	isSet bool
}

func (v NullableDocumentFlowConfigAllOf) Get() *DocumentFlowConfigAllOf {
	return v.value
}

func (v *NullableDocumentFlowConfigAllOf) Set(val *DocumentFlowConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentFlowConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentFlowConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentFlowConfigAllOf(val *DocumentFlowConfigAllOf) *NullableDocumentFlowConfigAllOf {
	return &NullableDocumentFlowConfigAllOf{value: val, isSet: true}
}

func (v NullableDocumentFlowConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentFlowConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


