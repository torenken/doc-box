/*
Doc Box API

The intent of this API is to provide a consistent/standardized mechanism to create new documents, delete existing documents, retrieving information about uploaded documents. This Specification is based on TMF667 - Document Management Release 17.0.1 from November 2017.

API version: 0.0.1
Contact: https://github.com/torenken/doc-box
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmf

import (
	"encoding/json"
)

// AttachmentResp struct for AttachmentResp
type AttachmentResp struct {
	// The preSigned url the
	PreSignedUrl *string `json:"preSignedUrl,omitempty"`
}

// NewAttachmentResp instantiates a new AttachmentResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentResp() *AttachmentResp {
	this := AttachmentResp{}
	return &this
}

// NewAttachmentRespWithDefaults instantiates a new AttachmentResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentRespWithDefaults() *AttachmentResp {
	this := AttachmentResp{}
	return &this
}

// GetPreSignedUrl returns the PreSignedUrl field value if set, zero value otherwise.
func (o *AttachmentResp) GetPreSignedUrl() string {
	if o == nil || o.PreSignedUrl == nil {
		var ret string
		return ret
	}
	return *o.PreSignedUrl
}

// GetPreSignedUrlOk returns a tuple with the PreSignedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentResp) GetPreSignedUrlOk() (*string, bool) {
	if o == nil || o.PreSignedUrl == nil {
		return nil, false
	}
	return o.PreSignedUrl, true
}

// HasPreSignedUrl returns a boolean if a field has been set.
func (o *AttachmentResp) HasPreSignedUrl() bool {
	if o != nil && o.PreSignedUrl != nil {
		return true
	}

	return false
}

// SetPreSignedUrl gets a reference to the given string and assigns it to the PreSignedUrl field.
func (o *AttachmentResp) SetPreSignedUrl(v string) {
	o.PreSignedUrl = &v
}

func (o AttachmentResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PreSignedUrl != nil {
		toSerialize["preSignedUrl"] = o.PreSignedUrl
	}
	return json.Marshal(toSerialize)
}

type NullableAttachmentResp struct {
	value *AttachmentResp
	isSet bool
}

func (v NullableAttachmentResp) Get() *AttachmentResp {
	return v.value
}

func (v *NullableAttachmentResp) Set(val *AttachmentResp) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentResp) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentResp(val *AttachmentResp) *NullableAttachmentResp {
	return &NullableAttachmentResp{value: val, isSet: true}
}

func (v NullableAttachmentResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
