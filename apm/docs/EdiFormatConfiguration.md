# EdiFormatConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ConfigType** | Pointer to **string** |  | [optional] 
**FormatType** | Pointer to **string** |  | [optional] 
**FormatTypeId** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **string** |  | [optional] 
**AckEndpointId** | Pointer to **string** |  | [optional] 
**ConfigId** | Pointer to **string** |  | [optional] 
**IsTemplate** | Pointer to **bool** |  | [optional] 
**CharacterSetAndDelimitersSettings** | Pointer to [**EdiFormatConfigurationCharacterSetAndDelimitersSettings**](EdiFormatConfigurationCharacterSetAndDelimitersSettings.md) |  | [optional] 
**WriteSettings** | Pointer to [**EdiFormatConfigurationWriteSettings**](EdiFormatConfigurationWriteSettings.md) |  | [optional] 
**EdifactControlNumberSettings** | Pointer to [**EdiFormatConfigurationEdifactControlNumberSettings**](EdiFormatConfigurationEdifactControlNumberSettings.md) |  | [optional] 
**EdifactParserSettings** | Pointer to [**EdiFormatConfigurationEdifactParserSettings**](EdiFormatConfigurationEdifactParserSettings.md) |  | [optional] 
**Separators** | Pointer to [**EdiFormatConfigurationSeparators**](EdiFormatConfigurationSeparators.md) |  | [optional] 
**AcknowledgementSettings** | Pointer to [**EdiFormatConfigurationAcknowledgementSettings**](EdiFormatConfigurationAcknowledgementSettings.md) |  | [optional] 
**WriterSettings** | Pointer to [**EdiFormatConfigurationWriterSettings**](EdiFormatConfigurationWriterSettings.md) |  | [optional] 
**CharacterSetEncoding** | Pointer to [**EdiFormatConfigurationCharacterSetEncoding**](EdiFormatConfigurationCharacterSetEncoding.md) |  | [optional] 
**TerminatorDelimiter** | Pointer to [**EdiFormatConfigurationTerminatorDelimiter**](EdiFormatConfigurationTerminatorDelimiter.md) |  | [optional] 
**ControlNumberSettings** | Pointer to [**EdiFormatConfigurationControlNumberSettings**](EdiFormatConfigurationControlNumberSettings.md) |  | [optional] 
**CharacterSetAndEncoding** | Pointer to [**EdiFormatConfigurationCharacterSetAndEncoding**](EdiFormatConfigurationCharacterSetAndEncoding.md) |  | [optional] 
**ParserSettings** | Pointer to [**EdiFormatConfigurationParserSettings**](EdiFormatConfigurationParserSettings.md) |  | [optional] 
**EnvelopeHeaders** | Pointer to [**EdiFormatConfigurationEnvelopeHeaders**](EdiFormatConfigurationEnvelopeHeaders.md) |  | [optional] 

## Methods

### NewEdiFormatConfiguration

`func NewEdiFormatConfiguration() *EdiFormatConfiguration`

NewEdiFormatConfiguration instantiates a new EdiFormatConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdiFormatConfigurationWithDefaults

`func NewEdiFormatConfigurationWithDefaults() *EdiFormatConfiguration`

NewEdiFormatConfigurationWithDefaults instantiates a new EdiFormatConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdiFormatConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdiFormatConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdiFormatConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EdiFormatConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfigType

`func (o *EdiFormatConfiguration) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *EdiFormatConfiguration) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *EdiFormatConfiguration) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.

### HasConfigType

`func (o *EdiFormatConfiguration) HasConfigType() bool`

HasConfigType returns a boolean if a field has been set.

### GetFormatType

`func (o *EdiFormatConfiguration) GetFormatType() string`

GetFormatType returns the FormatType field if non-nil, zero value otherwise.

### GetFormatTypeOk

`func (o *EdiFormatConfiguration) GetFormatTypeOk() (*string, bool)`

GetFormatTypeOk returns a tuple with the FormatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatType

`func (o *EdiFormatConfiguration) SetFormatType(v string)`

SetFormatType sets FormatType field to given value.

### HasFormatType

`func (o *EdiFormatConfiguration) HasFormatType() bool`

HasFormatType returns a boolean if a field has been set.

### GetFormatTypeId

`func (o *EdiFormatConfiguration) GetFormatTypeId() string`

GetFormatTypeId returns the FormatTypeId field if non-nil, zero value otherwise.

### GetFormatTypeIdOk

`func (o *EdiFormatConfiguration) GetFormatTypeIdOk() (*string, bool)`

GetFormatTypeIdOk returns a tuple with the FormatTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatTypeId

`func (o *EdiFormatConfiguration) SetFormatTypeId(v string)`

SetFormatTypeId sets FormatTypeId field to given value.

### HasFormatTypeId

`func (o *EdiFormatConfiguration) HasFormatTypeId() bool`

HasFormatTypeId returns a boolean if a field has been set.

### GetPartnerId

`func (o *EdiFormatConfiguration) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *EdiFormatConfiguration) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *EdiFormatConfiguration) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *EdiFormatConfiguration) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetAckEndpointId

`func (o *EdiFormatConfiguration) GetAckEndpointId() string`

GetAckEndpointId returns the AckEndpointId field if non-nil, zero value otherwise.

### GetAckEndpointIdOk

`func (o *EdiFormatConfiguration) GetAckEndpointIdOk() (*string, bool)`

GetAckEndpointIdOk returns a tuple with the AckEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckEndpointId

`func (o *EdiFormatConfiguration) SetAckEndpointId(v string)`

SetAckEndpointId sets AckEndpointId field to given value.

### HasAckEndpointId

`func (o *EdiFormatConfiguration) HasAckEndpointId() bool`

HasAckEndpointId returns a boolean if a field has been set.

### GetConfigId

`func (o *EdiFormatConfiguration) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *EdiFormatConfiguration) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *EdiFormatConfiguration) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *EdiFormatConfiguration) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetIsTemplate

`func (o *EdiFormatConfiguration) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *EdiFormatConfiguration) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *EdiFormatConfiguration) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *EdiFormatConfiguration) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetCharacterSetAndDelimitersSettings

`func (o *EdiFormatConfiguration) GetCharacterSetAndDelimitersSettings() EdiFormatConfigurationCharacterSetAndDelimitersSettings`

GetCharacterSetAndDelimitersSettings returns the CharacterSetAndDelimitersSettings field if non-nil, zero value otherwise.

### GetCharacterSetAndDelimitersSettingsOk

`func (o *EdiFormatConfiguration) GetCharacterSetAndDelimitersSettingsOk() (*EdiFormatConfigurationCharacterSetAndDelimitersSettings, bool)`

GetCharacterSetAndDelimitersSettingsOk returns a tuple with the CharacterSetAndDelimitersSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterSetAndDelimitersSettings

`func (o *EdiFormatConfiguration) SetCharacterSetAndDelimitersSettings(v EdiFormatConfigurationCharacterSetAndDelimitersSettings)`

SetCharacterSetAndDelimitersSettings sets CharacterSetAndDelimitersSettings field to given value.

### HasCharacterSetAndDelimitersSettings

`func (o *EdiFormatConfiguration) HasCharacterSetAndDelimitersSettings() bool`

HasCharacterSetAndDelimitersSettings returns a boolean if a field has been set.

### GetWriteSettings

`func (o *EdiFormatConfiguration) GetWriteSettings() EdiFormatConfigurationWriteSettings`

GetWriteSettings returns the WriteSettings field if non-nil, zero value otherwise.

### GetWriteSettingsOk

`func (o *EdiFormatConfiguration) GetWriteSettingsOk() (*EdiFormatConfigurationWriteSettings, bool)`

GetWriteSettingsOk returns a tuple with the WriteSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteSettings

`func (o *EdiFormatConfiguration) SetWriteSettings(v EdiFormatConfigurationWriteSettings)`

SetWriteSettings sets WriteSettings field to given value.

### HasWriteSettings

`func (o *EdiFormatConfiguration) HasWriteSettings() bool`

HasWriteSettings returns a boolean if a field has been set.

### GetEdifactControlNumberSettings

`func (o *EdiFormatConfiguration) GetEdifactControlNumberSettings() EdiFormatConfigurationEdifactControlNumberSettings`

GetEdifactControlNumberSettings returns the EdifactControlNumberSettings field if non-nil, zero value otherwise.

### GetEdifactControlNumberSettingsOk

`func (o *EdiFormatConfiguration) GetEdifactControlNumberSettingsOk() (*EdiFormatConfigurationEdifactControlNumberSettings, bool)`

GetEdifactControlNumberSettingsOk returns a tuple with the EdifactControlNumberSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdifactControlNumberSettings

`func (o *EdiFormatConfiguration) SetEdifactControlNumberSettings(v EdiFormatConfigurationEdifactControlNumberSettings)`

SetEdifactControlNumberSettings sets EdifactControlNumberSettings field to given value.

### HasEdifactControlNumberSettings

`func (o *EdiFormatConfiguration) HasEdifactControlNumberSettings() bool`

HasEdifactControlNumberSettings returns a boolean if a field has been set.

### GetEdifactParserSettings

`func (o *EdiFormatConfiguration) GetEdifactParserSettings() EdiFormatConfigurationEdifactParserSettings`

GetEdifactParserSettings returns the EdifactParserSettings field if non-nil, zero value otherwise.

### GetEdifactParserSettingsOk

`func (o *EdiFormatConfiguration) GetEdifactParserSettingsOk() (*EdiFormatConfigurationEdifactParserSettings, bool)`

GetEdifactParserSettingsOk returns a tuple with the EdifactParserSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdifactParserSettings

`func (o *EdiFormatConfiguration) SetEdifactParserSettings(v EdiFormatConfigurationEdifactParserSettings)`

SetEdifactParserSettings sets EdifactParserSettings field to given value.

### HasEdifactParserSettings

`func (o *EdiFormatConfiguration) HasEdifactParserSettings() bool`

HasEdifactParserSettings returns a boolean if a field has been set.

### GetSeparators

`func (o *EdiFormatConfiguration) GetSeparators() EdiFormatConfigurationSeparators`

GetSeparators returns the Separators field if non-nil, zero value otherwise.

### GetSeparatorsOk

`func (o *EdiFormatConfiguration) GetSeparatorsOk() (*EdiFormatConfigurationSeparators, bool)`

GetSeparatorsOk returns a tuple with the Separators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparators

`func (o *EdiFormatConfiguration) SetSeparators(v EdiFormatConfigurationSeparators)`

SetSeparators sets Separators field to given value.

### HasSeparators

`func (o *EdiFormatConfiguration) HasSeparators() bool`

HasSeparators returns a boolean if a field has been set.

### GetAcknowledgementSettings

`func (o *EdiFormatConfiguration) GetAcknowledgementSettings() EdiFormatConfigurationAcknowledgementSettings`

GetAcknowledgementSettings returns the AcknowledgementSettings field if non-nil, zero value otherwise.

### GetAcknowledgementSettingsOk

`func (o *EdiFormatConfiguration) GetAcknowledgementSettingsOk() (*EdiFormatConfigurationAcknowledgementSettings, bool)`

GetAcknowledgementSettingsOk returns a tuple with the AcknowledgementSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgementSettings

`func (o *EdiFormatConfiguration) SetAcknowledgementSettings(v EdiFormatConfigurationAcknowledgementSettings)`

SetAcknowledgementSettings sets AcknowledgementSettings field to given value.

### HasAcknowledgementSettings

`func (o *EdiFormatConfiguration) HasAcknowledgementSettings() bool`

HasAcknowledgementSettings returns a boolean if a field has been set.

### GetWriterSettings

`func (o *EdiFormatConfiguration) GetWriterSettings() EdiFormatConfigurationWriterSettings`

GetWriterSettings returns the WriterSettings field if non-nil, zero value otherwise.

### GetWriterSettingsOk

`func (o *EdiFormatConfiguration) GetWriterSettingsOk() (*EdiFormatConfigurationWriterSettings, bool)`

GetWriterSettingsOk returns a tuple with the WriterSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriterSettings

`func (o *EdiFormatConfiguration) SetWriterSettings(v EdiFormatConfigurationWriterSettings)`

SetWriterSettings sets WriterSettings field to given value.

### HasWriterSettings

`func (o *EdiFormatConfiguration) HasWriterSettings() bool`

HasWriterSettings returns a boolean if a field has been set.

### GetCharacterSetEncoding

`func (o *EdiFormatConfiguration) GetCharacterSetEncoding() EdiFormatConfigurationCharacterSetEncoding`

GetCharacterSetEncoding returns the CharacterSetEncoding field if non-nil, zero value otherwise.

### GetCharacterSetEncodingOk

`func (o *EdiFormatConfiguration) GetCharacterSetEncodingOk() (*EdiFormatConfigurationCharacterSetEncoding, bool)`

GetCharacterSetEncodingOk returns a tuple with the CharacterSetEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterSetEncoding

`func (o *EdiFormatConfiguration) SetCharacterSetEncoding(v EdiFormatConfigurationCharacterSetEncoding)`

SetCharacterSetEncoding sets CharacterSetEncoding field to given value.

### HasCharacterSetEncoding

`func (o *EdiFormatConfiguration) HasCharacterSetEncoding() bool`

HasCharacterSetEncoding returns a boolean if a field has been set.

### GetTerminatorDelimiter

`func (o *EdiFormatConfiguration) GetTerminatorDelimiter() EdiFormatConfigurationTerminatorDelimiter`

GetTerminatorDelimiter returns the TerminatorDelimiter field if non-nil, zero value otherwise.

### GetTerminatorDelimiterOk

`func (o *EdiFormatConfiguration) GetTerminatorDelimiterOk() (*EdiFormatConfigurationTerminatorDelimiter, bool)`

GetTerminatorDelimiterOk returns a tuple with the TerminatorDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatorDelimiter

`func (o *EdiFormatConfiguration) SetTerminatorDelimiter(v EdiFormatConfigurationTerminatorDelimiter)`

SetTerminatorDelimiter sets TerminatorDelimiter field to given value.

### HasTerminatorDelimiter

`func (o *EdiFormatConfiguration) HasTerminatorDelimiter() bool`

HasTerminatorDelimiter returns a boolean if a field has been set.

### GetControlNumberSettings

`func (o *EdiFormatConfiguration) GetControlNumberSettings() EdiFormatConfigurationControlNumberSettings`

GetControlNumberSettings returns the ControlNumberSettings field if non-nil, zero value otherwise.

### GetControlNumberSettingsOk

`func (o *EdiFormatConfiguration) GetControlNumberSettingsOk() (*EdiFormatConfigurationControlNumberSettings, bool)`

GetControlNumberSettingsOk returns a tuple with the ControlNumberSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlNumberSettings

`func (o *EdiFormatConfiguration) SetControlNumberSettings(v EdiFormatConfigurationControlNumberSettings)`

SetControlNumberSettings sets ControlNumberSettings field to given value.

### HasControlNumberSettings

`func (o *EdiFormatConfiguration) HasControlNumberSettings() bool`

HasControlNumberSettings returns a boolean if a field has been set.

### GetCharacterSetAndEncoding

`func (o *EdiFormatConfiguration) GetCharacterSetAndEncoding() EdiFormatConfigurationCharacterSetAndEncoding`

GetCharacterSetAndEncoding returns the CharacterSetAndEncoding field if non-nil, zero value otherwise.

### GetCharacterSetAndEncodingOk

`func (o *EdiFormatConfiguration) GetCharacterSetAndEncodingOk() (*EdiFormatConfigurationCharacterSetAndEncoding, bool)`

GetCharacterSetAndEncodingOk returns a tuple with the CharacterSetAndEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterSetAndEncoding

`func (o *EdiFormatConfiguration) SetCharacterSetAndEncoding(v EdiFormatConfigurationCharacterSetAndEncoding)`

SetCharacterSetAndEncoding sets CharacterSetAndEncoding field to given value.

### HasCharacterSetAndEncoding

`func (o *EdiFormatConfiguration) HasCharacterSetAndEncoding() bool`

HasCharacterSetAndEncoding returns a boolean if a field has been set.

### GetParserSettings

`func (o *EdiFormatConfiguration) GetParserSettings() EdiFormatConfigurationParserSettings`

GetParserSettings returns the ParserSettings field if non-nil, zero value otherwise.

### GetParserSettingsOk

`func (o *EdiFormatConfiguration) GetParserSettingsOk() (*EdiFormatConfigurationParserSettings, bool)`

GetParserSettingsOk returns a tuple with the ParserSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserSettings

`func (o *EdiFormatConfiguration) SetParserSettings(v EdiFormatConfigurationParserSettings)`

SetParserSettings sets ParserSettings field to given value.

### HasParserSettings

`func (o *EdiFormatConfiguration) HasParserSettings() bool`

HasParserSettings returns a boolean if a field has been set.

### GetEnvelopeHeaders

`func (o *EdiFormatConfiguration) GetEnvelopeHeaders() EdiFormatConfigurationEnvelopeHeaders`

GetEnvelopeHeaders returns the EnvelopeHeaders field if non-nil, zero value otherwise.

### GetEnvelopeHeadersOk

`func (o *EdiFormatConfiguration) GetEnvelopeHeadersOk() (*EdiFormatConfigurationEnvelopeHeaders, bool)`

GetEnvelopeHeadersOk returns a tuple with the EnvelopeHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvelopeHeaders

`func (o *EdiFormatConfiguration) SetEnvelopeHeaders(v EdiFormatConfigurationEnvelopeHeaders)`

SetEnvelopeHeaders sets EnvelopeHeaders field to given value.

### HasEnvelopeHeaders

`func (o *EdiFormatConfiguration) HasEnvelopeHeaders() bool`

HasEnvelopeHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


