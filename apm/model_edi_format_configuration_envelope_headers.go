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

// EdiFormatConfigurationEnvelopeHeaders struct for EdiFormatConfigurationEnvelopeHeaders
type EdiFormatConfigurationEnvelopeHeaders struct {
	AuthInfoQualifierISA01 *string `json:"authInfoQualifierISA01,omitempty"`
	AuthInfoISA02 *string `json:"authInfoISA02,omitempty"`
	SecurityInfoQualifierISA03 *string `json:"securityInfoQualifierISA03,omitempty"`
	SecurityInfoISA04 *string `json:"securityInfoISA04,omitempty"`
	InterchangeReceiverIdQualifierISA07 *string `json:"interchangeReceiverIdQualifierISA07,omitempty"`
	InterchangeReceiverIdISA07 *string `json:"interchangeReceiverIdISA07,omitempty"`
	RepetitionSeparatorCharacterISA11 *string `json:"repetitionSeparatorCharacterISA11,omitempty"`
	RepetitionInterchangeAcknowledmentsISA14 *string `json:"repetitionInterchangeAcknowledmentsISA14,omitempty"`
	DefaultInterchangeUsageIndicatorISA15 *string `json:"defaultInterchangeUsageIndicatorISA15,omitempty"`
	ComponentElementSeparator *string `json:"componentElementSeparator,omitempty"`
	InterchangeControlVersionNumberISA12 *string `json:"interchangeControlVersionNumberISA12,omitempty"`
	InterchangeControlStandardsIdentifierISA11 *string `json:"interchangeControlStandardsIdentifierISA11,omitempty"`
	ResponsibleAgencyCodeGS07 *string `json:"responsibleAgencyCodeGS07,omitempty"`
	ReleaseNumberGS08 *string `json:"releaseNumberGS08,omitempty"`
	ImplementationConventionReferenceST03 *string `json:"implementationConventionReferenceST03,omitempty"`
}

// NewEdiFormatConfigurationEnvelopeHeaders instantiates a new EdiFormatConfigurationEnvelopeHeaders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdiFormatConfigurationEnvelopeHeaders() *EdiFormatConfigurationEnvelopeHeaders {
	this := EdiFormatConfigurationEnvelopeHeaders{}
	return &this
}

// NewEdiFormatConfigurationEnvelopeHeadersWithDefaults instantiates a new EdiFormatConfigurationEnvelopeHeaders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdiFormatConfigurationEnvelopeHeadersWithDefaults() *EdiFormatConfigurationEnvelopeHeaders {
	this := EdiFormatConfigurationEnvelopeHeaders{}
	return &this
}

// GetAuthInfoQualifierISA01 returns the AuthInfoQualifierISA01 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetAuthInfoQualifierISA01() string {
	if o == nil || o.AuthInfoQualifierISA01 == nil {
		var ret string
		return ret
	}
	return *o.AuthInfoQualifierISA01
}

// GetAuthInfoQualifierISA01Ok returns a tuple with the AuthInfoQualifierISA01 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetAuthInfoQualifierISA01Ok() (*string, bool) {
	if o == nil || o.AuthInfoQualifierISA01 == nil {
		return nil, false
	}
	return o.AuthInfoQualifierISA01, true
}

// HasAuthInfoQualifierISA01 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasAuthInfoQualifierISA01() bool {
	if o != nil && o.AuthInfoQualifierISA01 != nil {
		return true
	}

	return false
}

// SetAuthInfoQualifierISA01 gets a reference to the given string and assigns it to the AuthInfoQualifierISA01 field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetAuthInfoQualifierISA01(v string) {
	o.AuthInfoQualifierISA01 = &v
}

// GetAuthInfoISA02 returns the AuthInfoISA02 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetAuthInfoISA02() string {
	if o == nil || o.AuthInfoISA02 == nil {
		var ret string
		return ret
	}
	return *o.AuthInfoISA02
}

// GetAuthInfoISA02Ok returns a tuple with the AuthInfoISA02 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetAuthInfoISA02Ok() (*string, bool) {
	if o == nil || o.AuthInfoISA02 == nil {
		return nil, false
	}
	return o.AuthInfoISA02, true
}

// HasAuthInfoISA02 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasAuthInfoISA02() bool {
	if o != nil && o.AuthInfoISA02 != nil {
		return true
	}

	return false
}

// SetAuthInfoISA02 gets a reference to the given string and assigns it to the AuthInfoISA02 field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetAuthInfoISA02(v string) {
	o.AuthInfoISA02 = &v
}

// GetSecurityInfoQualifierISA03 returns the SecurityInfoQualifierISA03 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetSecurityInfoQualifierISA03() string {
	if o == nil || o.SecurityInfoQualifierISA03 == nil {
		var ret string
		return ret
	}
	return *o.SecurityInfoQualifierISA03
}

// GetSecurityInfoQualifierISA03Ok returns a tuple with the SecurityInfoQualifierISA03 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetSecurityInfoQualifierISA03Ok() (*string, bool) {
	if o == nil || o.SecurityInfoQualifierISA03 == nil {
		return nil, false
	}
	return o.SecurityInfoQualifierISA03, true
}

// HasSecurityInfoQualifierISA03 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasSecurityInfoQualifierISA03() bool {
	if o != nil && o.SecurityInfoQualifierISA03 != nil {
		return true
	}

	return false
}

// SetSecurityInfoQualifierISA03 gets a reference to the given string and assigns it to the SecurityInfoQualifierISA03 field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetSecurityInfoQualifierISA03(v string) {
	o.SecurityInfoQualifierISA03 = &v
}

// GetSecurityInfoISA04 returns the SecurityInfoISA04 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetSecurityInfoISA04() string {
	if o == nil || o.SecurityInfoISA04 == nil {
		var ret string
		return ret
	}
	return *o.SecurityInfoISA04
}

// GetSecurityInfoISA04Ok returns a tuple with the SecurityInfoISA04 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetSecurityInfoISA04Ok() (*string, bool) {
	if o == nil || o.SecurityInfoISA04 == nil {
		return nil, false
	}
	return o.SecurityInfoISA04, true
}

// HasSecurityInfoISA04 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasSecurityInfoISA04() bool {
	if o != nil && o.SecurityInfoISA04 != nil {
		return true
	}

	return false
}

// SetSecurityInfoISA04 gets a reference to the given string and assigns it to the SecurityInfoISA04 field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetSecurityInfoISA04(v string) {
	o.SecurityInfoISA04 = &v
}

// GetInterchangeReceiverIdQualifierISA07 returns the InterchangeReceiverIdQualifierISA07 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetInterchangeReceiverIdQualifierISA07() string {
	if o == nil || o.InterchangeReceiverIdQualifierISA07 == nil {
		var ret string
		return ret
	}
	return *o.InterchangeReceiverIdQualifierISA07
}

// GetInterchangeReceiverIdQualifierISA07Ok returns a tuple with the InterchangeReceiverIdQualifierISA07 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetInterchangeReceiverIdQualifierISA07Ok() (*string, bool) {
	if o == nil || o.InterchangeReceiverIdQualifierISA07 == nil {
		return nil, false
	}
	return o.InterchangeReceiverIdQualifierISA07, true
}

// HasInterchangeReceiverIdQualifierISA07 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasInterchangeReceiverIdQualifierISA07() bool {
	if o != nil && o.InterchangeReceiverIdQualifierISA07 != nil {
		return true
	}

	return false
}

// SetInterchangeReceiverIdQualifierISA07 gets a reference to the given string and assigns it to the InterchangeReceiverIdQualifierISA07 field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetInterchangeReceiverIdQualifierISA07(v string) {
	o.InterchangeReceiverIdQualifierISA07 = &v
}

// GetInterchangeReceiverIdISA07 returns the InterchangeReceiverIdISA07 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetInterchangeReceiverIdISA07() string {
	if o == nil || o.InterchangeReceiverIdISA07 == nil {
		var ret string
		return ret
	}
	return *o.InterchangeReceiverIdISA07
}

// GetInterchangeReceiverIdISA07Ok returns a tuple with the InterchangeReceiverIdISA07 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetInterchangeReceiverIdISA07Ok() (*string, bool) {
	if o == nil || o.InterchangeReceiverIdISA07 == nil {
		return nil, false
	}
	return o.InterchangeReceiverIdISA07, true
}

// HasInterchangeReceiverIdISA07 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasInterchangeReceiverIdISA07() bool {
	if o != nil && o.InterchangeReceiverIdISA07 != nil {
		return true
	}

	return false
}

// SetInterchangeReceiverIdISA07 gets a reference to the given string and assigns it to the InterchangeReceiverIdISA07 field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetInterchangeReceiverIdISA07(v string) {
	o.InterchangeReceiverIdISA07 = &v
}

// GetRepetitionSeparatorCharacterISA11 returns the RepetitionSeparatorCharacterISA11 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetRepetitionSeparatorCharacterISA11() string {
	if o == nil || o.RepetitionSeparatorCharacterISA11 == nil {
		var ret string
		return ret
	}
	return *o.RepetitionSeparatorCharacterISA11
}

// GetRepetitionSeparatorCharacterISA11Ok returns a tuple with the RepetitionSeparatorCharacterISA11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetRepetitionSeparatorCharacterISA11Ok() (*string, bool) {
	if o == nil || o.RepetitionSeparatorCharacterISA11 == nil {
		return nil, false
	}
	return o.RepetitionSeparatorCharacterISA11, true
}

// HasRepetitionSeparatorCharacterISA11 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasRepetitionSeparatorCharacterISA11() bool {
	if o != nil && o.RepetitionSeparatorCharacterISA11 != nil {
		return true
	}

	return false
}

// SetRepetitionSeparatorCharacterISA11 gets a reference to the given string and assigns it to the RepetitionSeparatorCharacterISA11 field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetRepetitionSeparatorCharacterISA11(v string) {
	o.RepetitionSeparatorCharacterISA11 = &v
}

// GetRepetitionInterchangeAcknowledmentsISA14 returns the RepetitionInterchangeAcknowledmentsISA14 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetRepetitionInterchangeAcknowledmentsISA14() string {
	if o == nil || o.RepetitionInterchangeAcknowledmentsISA14 == nil {
		var ret string
		return ret
	}
	return *o.RepetitionInterchangeAcknowledmentsISA14
}

// GetRepetitionInterchangeAcknowledmentsISA14Ok returns a tuple with the RepetitionInterchangeAcknowledmentsISA14 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetRepetitionInterchangeAcknowledmentsISA14Ok() (*string, bool) {
	if o == nil || o.RepetitionInterchangeAcknowledmentsISA14 == nil {
		return nil, false
	}
	return o.RepetitionInterchangeAcknowledmentsISA14, true
}

// HasRepetitionInterchangeAcknowledmentsISA14 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasRepetitionInterchangeAcknowledmentsISA14() bool {
	if o != nil && o.RepetitionInterchangeAcknowledmentsISA14 != nil {
		return true
	}

	return false
}

// SetRepetitionInterchangeAcknowledmentsISA14 gets a reference to the given string and assigns it to the RepetitionInterchangeAcknowledmentsISA14 field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetRepetitionInterchangeAcknowledmentsISA14(v string) {
	o.RepetitionInterchangeAcknowledmentsISA14 = &v
}

// GetDefaultInterchangeUsageIndicatorISA15 returns the DefaultInterchangeUsageIndicatorISA15 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetDefaultInterchangeUsageIndicatorISA15() string {
	if o == nil || o.DefaultInterchangeUsageIndicatorISA15 == nil {
		var ret string
		return ret
	}
	return *o.DefaultInterchangeUsageIndicatorISA15
}

// GetDefaultInterchangeUsageIndicatorISA15Ok returns a tuple with the DefaultInterchangeUsageIndicatorISA15 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetDefaultInterchangeUsageIndicatorISA15Ok() (*string, bool) {
	if o == nil || o.DefaultInterchangeUsageIndicatorISA15 == nil {
		return nil, false
	}
	return o.DefaultInterchangeUsageIndicatorISA15, true
}

// HasDefaultInterchangeUsageIndicatorISA15 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasDefaultInterchangeUsageIndicatorISA15() bool {
	if o != nil && o.DefaultInterchangeUsageIndicatorISA15 != nil {
		return true
	}

	return false
}

// SetDefaultInterchangeUsageIndicatorISA15 gets a reference to the given string and assigns it to the DefaultInterchangeUsageIndicatorISA15 field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetDefaultInterchangeUsageIndicatorISA15(v string) {
	o.DefaultInterchangeUsageIndicatorISA15 = &v
}

// GetComponentElementSeparator returns the ComponentElementSeparator field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetComponentElementSeparator() string {
	if o == nil || o.ComponentElementSeparator == nil {
		var ret string
		return ret
	}
	return *o.ComponentElementSeparator
}

// GetComponentElementSeparatorOk returns a tuple with the ComponentElementSeparator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetComponentElementSeparatorOk() (*string, bool) {
	if o == nil || o.ComponentElementSeparator == nil {
		return nil, false
	}
	return o.ComponentElementSeparator, true
}

// HasComponentElementSeparator returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasComponentElementSeparator() bool {
	if o != nil && o.ComponentElementSeparator != nil {
		return true
	}

	return false
}

// SetComponentElementSeparator gets a reference to the given string and assigns it to the ComponentElementSeparator field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetComponentElementSeparator(v string) {
	o.ComponentElementSeparator = &v
}

// GetInterchangeControlVersionNumberISA12 returns the InterchangeControlVersionNumberISA12 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetInterchangeControlVersionNumberISA12() string {
	if o == nil || o.InterchangeControlVersionNumberISA12 == nil {
		var ret string
		return ret
	}
	return *o.InterchangeControlVersionNumberISA12
}

// GetInterchangeControlVersionNumberISA12Ok returns a tuple with the InterchangeControlVersionNumberISA12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetInterchangeControlVersionNumberISA12Ok() (*string, bool) {
	if o == nil || o.InterchangeControlVersionNumberISA12 == nil {
		return nil, false
	}
	return o.InterchangeControlVersionNumberISA12, true
}

// HasInterchangeControlVersionNumberISA12 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasInterchangeControlVersionNumberISA12() bool {
	if o != nil && o.InterchangeControlVersionNumberISA12 != nil {
		return true
	}

	return false
}

// SetInterchangeControlVersionNumberISA12 gets a reference to the given string and assigns it to the InterchangeControlVersionNumberISA12 field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetInterchangeControlVersionNumberISA12(v string) {
	o.InterchangeControlVersionNumberISA12 = &v
}

// GetInterchangeControlStandardsIdentifierISA11 returns the InterchangeControlStandardsIdentifierISA11 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetInterchangeControlStandardsIdentifierISA11() string {
	if o == nil || o.InterchangeControlStandardsIdentifierISA11 == nil {
		var ret string
		return ret
	}
	return *o.InterchangeControlStandardsIdentifierISA11
}

// GetInterchangeControlStandardsIdentifierISA11Ok returns a tuple with the InterchangeControlStandardsIdentifierISA11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetInterchangeControlStandardsIdentifierISA11Ok() (*string, bool) {
	if o == nil || o.InterchangeControlStandardsIdentifierISA11 == nil {
		return nil, false
	}
	return o.InterchangeControlStandardsIdentifierISA11, true
}

// HasInterchangeControlStandardsIdentifierISA11 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasInterchangeControlStandardsIdentifierISA11() bool {
	if o != nil && o.InterchangeControlStandardsIdentifierISA11 != nil {
		return true
	}

	return false
}

// SetInterchangeControlStandardsIdentifierISA11 gets a reference to the given string and assigns it to the InterchangeControlStandardsIdentifierISA11 field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetInterchangeControlStandardsIdentifierISA11(v string) {
	o.InterchangeControlStandardsIdentifierISA11 = &v
}

// GetResponsibleAgencyCodeGS07 returns the ResponsibleAgencyCodeGS07 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetResponsibleAgencyCodeGS07() string {
	if o == nil || o.ResponsibleAgencyCodeGS07 == nil {
		var ret string
		return ret
	}
	return *o.ResponsibleAgencyCodeGS07
}

// GetResponsibleAgencyCodeGS07Ok returns a tuple with the ResponsibleAgencyCodeGS07 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetResponsibleAgencyCodeGS07Ok() (*string, bool) {
	if o == nil || o.ResponsibleAgencyCodeGS07 == nil {
		return nil, false
	}
	return o.ResponsibleAgencyCodeGS07, true
}

// HasResponsibleAgencyCodeGS07 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasResponsibleAgencyCodeGS07() bool {
	if o != nil && o.ResponsibleAgencyCodeGS07 != nil {
		return true
	}

	return false
}

// SetResponsibleAgencyCodeGS07 gets a reference to the given string and assigns it to the ResponsibleAgencyCodeGS07 field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetResponsibleAgencyCodeGS07(v string) {
	o.ResponsibleAgencyCodeGS07 = &v
}

// GetReleaseNumberGS08 returns the ReleaseNumberGS08 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetReleaseNumberGS08() string {
	if o == nil || o.ReleaseNumberGS08 == nil {
		var ret string
		return ret
	}
	return *o.ReleaseNumberGS08
}

// GetReleaseNumberGS08Ok returns a tuple with the ReleaseNumberGS08 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetReleaseNumberGS08Ok() (*string, bool) {
	if o == nil || o.ReleaseNumberGS08 == nil {
		return nil, false
	}
	return o.ReleaseNumberGS08, true
}

// HasReleaseNumberGS08 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasReleaseNumberGS08() bool {
	if o != nil && o.ReleaseNumberGS08 != nil {
		return true
	}

	return false
}

// SetReleaseNumberGS08 gets a reference to the given string and assigns it to the ReleaseNumberGS08 field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetReleaseNumberGS08(v string) {
	o.ReleaseNumberGS08 = &v
}

// GetImplementationConventionReferenceST03 returns the ImplementationConventionReferenceST03 field value if set, zero value otherwise.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetImplementationConventionReferenceST03() string {
	if o == nil || o.ImplementationConventionReferenceST03 == nil {
		var ret string
		return ret
	}
	return *o.ImplementationConventionReferenceST03
}

// GetImplementationConventionReferenceST03Ok returns a tuple with the ImplementationConventionReferenceST03 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) GetImplementationConventionReferenceST03Ok() (*string, bool) {
	if o == nil || o.ImplementationConventionReferenceST03 == nil {
		return nil, false
	}
	return o.ImplementationConventionReferenceST03, true
}

// HasImplementationConventionReferenceST03 returns a boolean if a field has been set.
func (o *EdiFormatConfigurationEnvelopeHeaders) HasImplementationConventionReferenceST03() bool {
	if o != nil && o.ImplementationConventionReferenceST03 != nil {
		return true
	}

	return false
}

// SetImplementationConventionReferenceST03 gets a reference to the given string and assigns it to the ImplementationConventionReferenceST03 field.
func (o *EdiFormatConfigurationEnvelopeHeaders) SetImplementationConventionReferenceST03(v string) {
	o.ImplementationConventionReferenceST03 = &v
}

func (o EdiFormatConfigurationEnvelopeHeaders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthInfoQualifierISA01 != nil {
		toSerialize["authInfoQualifierISA01"] = o.AuthInfoQualifierISA01
	}
	if o.AuthInfoISA02 != nil {
		toSerialize["authInfoISA02"] = o.AuthInfoISA02
	}
	if o.SecurityInfoQualifierISA03 != nil {
		toSerialize["securityInfoQualifierISA03"] = o.SecurityInfoQualifierISA03
	}
	if o.SecurityInfoISA04 != nil {
		toSerialize["securityInfoISA04"] = o.SecurityInfoISA04
	}
	if o.InterchangeReceiverIdQualifierISA07 != nil {
		toSerialize["interchangeReceiverIdQualifierISA07"] = o.InterchangeReceiverIdQualifierISA07
	}
	if o.InterchangeReceiverIdISA07 != nil {
		toSerialize["interchangeReceiverIdISA07"] = o.InterchangeReceiverIdISA07
	}
	if o.RepetitionSeparatorCharacterISA11 != nil {
		toSerialize["repetitionSeparatorCharacterISA11"] = o.RepetitionSeparatorCharacterISA11
	}
	if o.RepetitionInterchangeAcknowledmentsISA14 != nil {
		toSerialize["repetitionInterchangeAcknowledmentsISA14"] = o.RepetitionInterchangeAcknowledmentsISA14
	}
	if o.DefaultInterchangeUsageIndicatorISA15 != nil {
		toSerialize["defaultInterchangeUsageIndicatorISA15"] = o.DefaultInterchangeUsageIndicatorISA15
	}
	if o.ComponentElementSeparator != nil {
		toSerialize["componentElementSeparator"] = o.ComponentElementSeparator
	}
	if o.InterchangeControlVersionNumberISA12 != nil {
		toSerialize["interchangeControlVersionNumberISA12"] = o.InterchangeControlVersionNumberISA12
	}
	if o.InterchangeControlStandardsIdentifierISA11 != nil {
		toSerialize["interchangeControlStandardsIdentifierISA11"] = o.InterchangeControlStandardsIdentifierISA11
	}
	if o.ResponsibleAgencyCodeGS07 != nil {
		toSerialize["responsibleAgencyCodeGS07"] = o.ResponsibleAgencyCodeGS07
	}
	if o.ReleaseNumberGS08 != nil {
		toSerialize["releaseNumberGS08"] = o.ReleaseNumberGS08
	}
	if o.ImplementationConventionReferenceST03 != nil {
		toSerialize["implementationConventionReferenceST03"] = o.ImplementationConventionReferenceST03
	}
	return json.Marshal(toSerialize)
}

type NullableEdiFormatConfigurationEnvelopeHeaders struct {
	value *EdiFormatConfigurationEnvelopeHeaders
	isSet bool
}

func (v NullableEdiFormatConfigurationEnvelopeHeaders) Get() *EdiFormatConfigurationEnvelopeHeaders {
	return v.value
}

func (v *NullableEdiFormatConfigurationEnvelopeHeaders) Set(val *EdiFormatConfigurationEnvelopeHeaders) {
	v.value = val
	v.isSet = true
}

func (v NullableEdiFormatConfigurationEnvelopeHeaders) IsSet() bool {
	return v.isSet
}

func (v *NullableEdiFormatConfigurationEnvelopeHeaders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdiFormatConfigurationEnvelopeHeaders(val *EdiFormatConfigurationEnvelopeHeaders) *NullableEdiFormatConfigurationEnvelopeHeaders {
	return &NullableEdiFormatConfigurationEnvelopeHeaders{value: val, isSet: true}
}

func (v NullableEdiFormatConfigurationEnvelopeHeaders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdiFormatConfigurationEnvelopeHeaders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

