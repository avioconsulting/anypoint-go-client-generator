# TLSContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Insecure** | **bool** |  | [default to false]
**NeedCertificate** | **bool** |  | [default to true]
**CertificateId** | Pointer to **string** |  | [optional] [default to "true"]

## Methods

### NewTLSContext

`func NewTLSContext(insecure bool, needCertificate bool, ) *TLSContext`

NewTLSContext instantiates a new TLSContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSContextWithDefaults

`func NewTLSContextWithDefaults() *TLSContext`

NewTLSContextWithDefaults instantiates a new TLSContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsecure

`func (o *TLSContext) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *TLSContext) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *TLSContext) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.


### GetNeedCertificate

`func (o *TLSContext) GetNeedCertificate() bool`

GetNeedCertificate returns the NeedCertificate field if non-nil, zero value otherwise.

### GetNeedCertificateOk

`func (o *TLSContext) GetNeedCertificateOk() (*bool, bool)`

GetNeedCertificateOk returns a tuple with the NeedCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedCertificate

`func (o *TLSContext) SetNeedCertificate(v bool)`

SetNeedCertificate sets NeedCertificate field to given value.


### GetCertificateId

`func (o *TLSContext) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *TLSContext) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *TLSContext) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *TLSContext) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


