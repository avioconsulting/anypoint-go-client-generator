# EdiFormatConfigurationParserSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckEndpointId** | Pointer to **string** |  | [optional] 
**CheckDuplicateDays** | Pointer to **int32** |  | [optional] 
**FailDocumentWhenValueLengthOutsideAllowedRange** | Pointer to **bool** |  | [optional] 
**FailDocumentWhenInvalidCharacterInValue** | Pointer to **bool** |  | [optional] 
**FailDocumentIfValueIsRepeatedTooManyTimes** | Pointer to **bool** |  | [optional] 
**FailDocumentIfUnknownSegmentsAreUsed** | Pointer to **bool** |  | [optional] 
**FailDocumentWhenSegmentsAreOutOfOrder** | Pointer to **bool** |  | [optional] 
**FailDocumentWhenTooManyRepeatsOfSegment** | Pointer to **bool** |  | [optional] 
**FailDocumentWhenUnusedSegmentsAreIncluded** | Pointer to **bool** |  | [optional] 
**EnforceConditionalRulesOnParser** | Pointer to **bool** |  | [optional] 
**EnforceCodeSetValidationsParse** | Pointer to **bool** |  | [optional] 
**AcknowledgeEveryTransaction** | Pointer to **bool** |  | [optional] 
**GenerateTA1** | Pointer to **bool** |  | [optional] 
**Generate999** | Pointer to **bool** |  | [optional] 
**Require997** | Pointer to **bool** |  | [optional] 

## Methods

### NewEdiFormatConfigurationParserSettings

`func NewEdiFormatConfigurationParserSettings() *EdiFormatConfigurationParserSettings`

NewEdiFormatConfigurationParserSettings instantiates a new EdiFormatConfigurationParserSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdiFormatConfigurationParserSettingsWithDefaults

`func NewEdiFormatConfigurationParserSettingsWithDefaults() *EdiFormatConfigurationParserSettings`

NewEdiFormatConfigurationParserSettingsWithDefaults instantiates a new EdiFormatConfigurationParserSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckEndpointId

`func (o *EdiFormatConfigurationParserSettings) GetAckEndpointId() string`

GetAckEndpointId returns the AckEndpointId field if non-nil, zero value otherwise.

### GetAckEndpointIdOk

`func (o *EdiFormatConfigurationParserSettings) GetAckEndpointIdOk() (*string, bool)`

GetAckEndpointIdOk returns a tuple with the AckEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckEndpointId

`func (o *EdiFormatConfigurationParserSettings) SetAckEndpointId(v string)`

SetAckEndpointId sets AckEndpointId field to given value.

### HasAckEndpointId

`func (o *EdiFormatConfigurationParserSettings) HasAckEndpointId() bool`

HasAckEndpointId returns a boolean if a field has been set.

### GetCheckDuplicateDays

`func (o *EdiFormatConfigurationParserSettings) GetCheckDuplicateDays() int32`

GetCheckDuplicateDays returns the CheckDuplicateDays field if non-nil, zero value otherwise.

### GetCheckDuplicateDaysOk

`func (o *EdiFormatConfigurationParserSettings) GetCheckDuplicateDaysOk() (*int32, bool)`

GetCheckDuplicateDaysOk returns a tuple with the CheckDuplicateDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckDuplicateDays

`func (o *EdiFormatConfigurationParserSettings) SetCheckDuplicateDays(v int32)`

SetCheckDuplicateDays sets CheckDuplicateDays field to given value.

### HasCheckDuplicateDays

`func (o *EdiFormatConfigurationParserSettings) HasCheckDuplicateDays() bool`

HasCheckDuplicateDays returns a boolean if a field has been set.

### GetFailDocumentWhenValueLengthOutsideAllowedRange

`func (o *EdiFormatConfigurationParserSettings) GetFailDocumentWhenValueLengthOutsideAllowedRange() bool`

GetFailDocumentWhenValueLengthOutsideAllowedRange returns the FailDocumentWhenValueLengthOutsideAllowedRange field if non-nil, zero value otherwise.

### GetFailDocumentWhenValueLengthOutsideAllowedRangeOk

`func (o *EdiFormatConfigurationParserSettings) GetFailDocumentWhenValueLengthOutsideAllowedRangeOk() (*bool, bool)`

GetFailDocumentWhenValueLengthOutsideAllowedRangeOk returns a tuple with the FailDocumentWhenValueLengthOutsideAllowedRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailDocumentWhenValueLengthOutsideAllowedRange

`func (o *EdiFormatConfigurationParserSettings) SetFailDocumentWhenValueLengthOutsideAllowedRange(v bool)`

SetFailDocumentWhenValueLengthOutsideAllowedRange sets FailDocumentWhenValueLengthOutsideAllowedRange field to given value.

### HasFailDocumentWhenValueLengthOutsideAllowedRange

`func (o *EdiFormatConfigurationParserSettings) HasFailDocumentWhenValueLengthOutsideAllowedRange() bool`

HasFailDocumentWhenValueLengthOutsideAllowedRange returns a boolean if a field has been set.

### GetFailDocumentWhenInvalidCharacterInValue

`func (o *EdiFormatConfigurationParserSettings) GetFailDocumentWhenInvalidCharacterInValue() bool`

GetFailDocumentWhenInvalidCharacterInValue returns the FailDocumentWhenInvalidCharacterInValue field if non-nil, zero value otherwise.

### GetFailDocumentWhenInvalidCharacterInValueOk

`func (o *EdiFormatConfigurationParserSettings) GetFailDocumentWhenInvalidCharacterInValueOk() (*bool, bool)`

GetFailDocumentWhenInvalidCharacterInValueOk returns a tuple with the FailDocumentWhenInvalidCharacterInValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailDocumentWhenInvalidCharacterInValue

`func (o *EdiFormatConfigurationParserSettings) SetFailDocumentWhenInvalidCharacterInValue(v bool)`

SetFailDocumentWhenInvalidCharacterInValue sets FailDocumentWhenInvalidCharacterInValue field to given value.

### HasFailDocumentWhenInvalidCharacterInValue

`func (o *EdiFormatConfigurationParserSettings) HasFailDocumentWhenInvalidCharacterInValue() bool`

HasFailDocumentWhenInvalidCharacterInValue returns a boolean if a field has been set.

### GetFailDocumentIfValueIsRepeatedTooManyTimes

`func (o *EdiFormatConfigurationParserSettings) GetFailDocumentIfValueIsRepeatedTooManyTimes() bool`

GetFailDocumentIfValueIsRepeatedTooManyTimes returns the FailDocumentIfValueIsRepeatedTooManyTimes field if non-nil, zero value otherwise.

### GetFailDocumentIfValueIsRepeatedTooManyTimesOk

`func (o *EdiFormatConfigurationParserSettings) GetFailDocumentIfValueIsRepeatedTooManyTimesOk() (*bool, bool)`

GetFailDocumentIfValueIsRepeatedTooManyTimesOk returns a tuple with the FailDocumentIfValueIsRepeatedTooManyTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailDocumentIfValueIsRepeatedTooManyTimes

`func (o *EdiFormatConfigurationParserSettings) SetFailDocumentIfValueIsRepeatedTooManyTimes(v bool)`

SetFailDocumentIfValueIsRepeatedTooManyTimes sets FailDocumentIfValueIsRepeatedTooManyTimes field to given value.

### HasFailDocumentIfValueIsRepeatedTooManyTimes

`func (o *EdiFormatConfigurationParserSettings) HasFailDocumentIfValueIsRepeatedTooManyTimes() bool`

HasFailDocumentIfValueIsRepeatedTooManyTimes returns a boolean if a field has been set.

### GetFailDocumentIfUnknownSegmentsAreUsed

`func (o *EdiFormatConfigurationParserSettings) GetFailDocumentIfUnknownSegmentsAreUsed() bool`

GetFailDocumentIfUnknownSegmentsAreUsed returns the FailDocumentIfUnknownSegmentsAreUsed field if non-nil, zero value otherwise.

### GetFailDocumentIfUnknownSegmentsAreUsedOk

`func (o *EdiFormatConfigurationParserSettings) GetFailDocumentIfUnknownSegmentsAreUsedOk() (*bool, bool)`

GetFailDocumentIfUnknownSegmentsAreUsedOk returns a tuple with the FailDocumentIfUnknownSegmentsAreUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailDocumentIfUnknownSegmentsAreUsed

`func (o *EdiFormatConfigurationParserSettings) SetFailDocumentIfUnknownSegmentsAreUsed(v bool)`

SetFailDocumentIfUnknownSegmentsAreUsed sets FailDocumentIfUnknownSegmentsAreUsed field to given value.

### HasFailDocumentIfUnknownSegmentsAreUsed

`func (o *EdiFormatConfigurationParserSettings) HasFailDocumentIfUnknownSegmentsAreUsed() bool`

HasFailDocumentIfUnknownSegmentsAreUsed returns a boolean if a field has been set.

### GetFailDocumentWhenSegmentsAreOutOfOrder

`func (o *EdiFormatConfigurationParserSettings) GetFailDocumentWhenSegmentsAreOutOfOrder() bool`

GetFailDocumentWhenSegmentsAreOutOfOrder returns the FailDocumentWhenSegmentsAreOutOfOrder field if non-nil, zero value otherwise.

### GetFailDocumentWhenSegmentsAreOutOfOrderOk

`func (o *EdiFormatConfigurationParserSettings) GetFailDocumentWhenSegmentsAreOutOfOrderOk() (*bool, bool)`

GetFailDocumentWhenSegmentsAreOutOfOrderOk returns a tuple with the FailDocumentWhenSegmentsAreOutOfOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailDocumentWhenSegmentsAreOutOfOrder

`func (o *EdiFormatConfigurationParserSettings) SetFailDocumentWhenSegmentsAreOutOfOrder(v bool)`

SetFailDocumentWhenSegmentsAreOutOfOrder sets FailDocumentWhenSegmentsAreOutOfOrder field to given value.

### HasFailDocumentWhenSegmentsAreOutOfOrder

`func (o *EdiFormatConfigurationParserSettings) HasFailDocumentWhenSegmentsAreOutOfOrder() bool`

HasFailDocumentWhenSegmentsAreOutOfOrder returns a boolean if a field has been set.

### GetFailDocumentWhenTooManyRepeatsOfSegment

`func (o *EdiFormatConfigurationParserSettings) GetFailDocumentWhenTooManyRepeatsOfSegment() bool`

GetFailDocumentWhenTooManyRepeatsOfSegment returns the FailDocumentWhenTooManyRepeatsOfSegment field if non-nil, zero value otherwise.

### GetFailDocumentWhenTooManyRepeatsOfSegmentOk

`func (o *EdiFormatConfigurationParserSettings) GetFailDocumentWhenTooManyRepeatsOfSegmentOk() (*bool, bool)`

GetFailDocumentWhenTooManyRepeatsOfSegmentOk returns a tuple with the FailDocumentWhenTooManyRepeatsOfSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailDocumentWhenTooManyRepeatsOfSegment

`func (o *EdiFormatConfigurationParserSettings) SetFailDocumentWhenTooManyRepeatsOfSegment(v bool)`

SetFailDocumentWhenTooManyRepeatsOfSegment sets FailDocumentWhenTooManyRepeatsOfSegment field to given value.

### HasFailDocumentWhenTooManyRepeatsOfSegment

`func (o *EdiFormatConfigurationParserSettings) HasFailDocumentWhenTooManyRepeatsOfSegment() bool`

HasFailDocumentWhenTooManyRepeatsOfSegment returns a boolean if a field has been set.

### GetFailDocumentWhenUnusedSegmentsAreIncluded

`func (o *EdiFormatConfigurationParserSettings) GetFailDocumentWhenUnusedSegmentsAreIncluded() bool`

GetFailDocumentWhenUnusedSegmentsAreIncluded returns the FailDocumentWhenUnusedSegmentsAreIncluded field if non-nil, zero value otherwise.

### GetFailDocumentWhenUnusedSegmentsAreIncludedOk

`func (o *EdiFormatConfigurationParserSettings) GetFailDocumentWhenUnusedSegmentsAreIncludedOk() (*bool, bool)`

GetFailDocumentWhenUnusedSegmentsAreIncludedOk returns a tuple with the FailDocumentWhenUnusedSegmentsAreIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailDocumentWhenUnusedSegmentsAreIncluded

`func (o *EdiFormatConfigurationParserSettings) SetFailDocumentWhenUnusedSegmentsAreIncluded(v bool)`

SetFailDocumentWhenUnusedSegmentsAreIncluded sets FailDocumentWhenUnusedSegmentsAreIncluded field to given value.

### HasFailDocumentWhenUnusedSegmentsAreIncluded

`func (o *EdiFormatConfigurationParserSettings) HasFailDocumentWhenUnusedSegmentsAreIncluded() bool`

HasFailDocumentWhenUnusedSegmentsAreIncluded returns a boolean if a field has been set.

### GetEnforceConditionalRulesOnParser

`func (o *EdiFormatConfigurationParserSettings) GetEnforceConditionalRulesOnParser() bool`

GetEnforceConditionalRulesOnParser returns the EnforceConditionalRulesOnParser field if non-nil, zero value otherwise.

### GetEnforceConditionalRulesOnParserOk

`func (o *EdiFormatConfigurationParserSettings) GetEnforceConditionalRulesOnParserOk() (*bool, bool)`

GetEnforceConditionalRulesOnParserOk returns a tuple with the EnforceConditionalRulesOnParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceConditionalRulesOnParser

`func (o *EdiFormatConfigurationParserSettings) SetEnforceConditionalRulesOnParser(v bool)`

SetEnforceConditionalRulesOnParser sets EnforceConditionalRulesOnParser field to given value.

### HasEnforceConditionalRulesOnParser

`func (o *EdiFormatConfigurationParserSettings) HasEnforceConditionalRulesOnParser() bool`

HasEnforceConditionalRulesOnParser returns a boolean if a field has been set.

### GetEnforceCodeSetValidationsParse

`func (o *EdiFormatConfigurationParserSettings) GetEnforceCodeSetValidationsParse() bool`

GetEnforceCodeSetValidationsParse returns the EnforceCodeSetValidationsParse field if non-nil, zero value otherwise.

### GetEnforceCodeSetValidationsParseOk

`func (o *EdiFormatConfigurationParserSettings) GetEnforceCodeSetValidationsParseOk() (*bool, bool)`

GetEnforceCodeSetValidationsParseOk returns a tuple with the EnforceCodeSetValidationsParse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceCodeSetValidationsParse

`func (o *EdiFormatConfigurationParserSettings) SetEnforceCodeSetValidationsParse(v bool)`

SetEnforceCodeSetValidationsParse sets EnforceCodeSetValidationsParse field to given value.

### HasEnforceCodeSetValidationsParse

`func (o *EdiFormatConfigurationParserSettings) HasEnforceCodeSetValidationsParse() bool`

HasEnforceCodeSetValidationsParse returns a boolean if a field has been set.

### GetAcknowledgeEveryTransaction

`func (o *EdiFormatConfigurationParserSettings) GetAcknowledgeEveryTransaction() bool`

GetAcknowledgeEveryTransaction returns the AcknowledgeEveryTransaction field if non-nil, zero value otherwise.

### GetAcknowledgeEveryTransactionOk

`func (o *EdiFormatConfigurationParserSettings) GetAcknowledgeEveryTransactionOk() (*bool, bool)`

GetAcknowledgeEveryTransactionOk returns a tuple with the AcknowledgeEveryTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeEveryTransaction

`func (o *EdiFormatConfigurationParserSettings) SetAcknowledgeEveryTransaction(v bool)`

SetAcknowledgeEveryTransaction sets AcknowledgeEveryTransaction field to given value.

### HasAcknowledgeEveryTransaction

`func (o *EdiFormatConfigurationParserSettings) HasAcknowledgeEveryTransaction() bool`

HasAcknowledgeEveryTransaction returns a boolean if a field has been set.

### GetGenerateTA1

`func (o *EdiFormatConfigurationParserSettings) GetGenerateTA1() bool`

GetGenerateTA1 returns the GenerateTA1 field if non-nil, zero value otherwise.

### GetGenerateTA1Ok

`func (o *EdiFormatConfigurationParserSettings) GetGenerateTA1Ok() (*bool, bool)`

GetGenerateTA1Ok returns a tuple with the GenerateTA1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateTA1

`func (o *EdiFormatConfigurationParserSettings) SetGenerateTA1(v bool)`

SetGenerateTA1 sets GenerateTA1 field to given value.

### HasGenerateTA1

`func (o *EdiFormatConfigurationParserSettings) HasGenerateTA1() bool`

HasGenerateTA1 returns a boolean if a field has been set.

### GetGenerate999

`func (o *EdiFormatConfigurationParserSettings) GetGenerate999() bool`

GetGenerate999 returns the Generate999 field if non-nil, zero value otherwise.

### GetGenerate999Ok

`func (o *EdiFormatConfigurationParserSettings) GetGenerate999Ok() (*bool, bool)`

GetGenerate999Ok returns a tuple with the Generate999 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerate999

`func (o *EdiFormatConfigurationParserSettings) SetGenerate999(v bool)`

SetGenerate999 sets Generate999 field to given value.

### HasGenerate999

`func (o *EdiFormatConfigurationParserSettings) HasGenerate999() bool`

HasGenerate999 returns a boolean if a field has been set.

### GetRequire997

`func (o *EdiFormatConfigurationParserSettings) GetRequire997() bool`

GetRequire997 returns the Require997 field if non-nil, zero value otherwise.

### GetRequire997Ok

`func (o *EdiFormatConfigurationParserSettings) GetRequire997Ok() (*bool, bool)`

GetRequire997Ok returns a tuple with the Require997 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequire997

`func (o *EdiFormatConfigurationParserSettings) SetRequire997(v bool)`

SetRequire997 sets Require997 field to given value.

### HasRequire997

`func (o *EdiFormatConfigurationParserSettings) HasRequire997() bool`

HasRequire997 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


