/*
LogCom API

LogCom Swagger documentation

API version: 1.2.14
Contact: laborit@blutspende.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logcomapi

import (
	"encoding/json"
)

// HealthCheck struct for HealthCheck
type HealthCheck struct {
	ApiVersion []string `json:"apiVersion,omitempty"`
	Service *string `json:"service,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewHealthCheck instantiates a new HealthCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheck() *HealthCheck {
	this := HealthCheck{}
	return &this
}

// NewHealthCheckWithDefaults instantiates a new HealthCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckWithDefaults() *HealthCheck {
	this := HealthCheck{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *HealthCheck) GetApiVersion() []string {
	if o == nil || isNil(o.ApiVersion) {
		var ret []string
		return ret
	}
	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetApiVersionOk() ([]string, bool) {
	if o == nil || isNil(o.ApiVersion) {
    return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *HealthCheck) HasApiVersion() bool {
	if o != nil && !isNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given []string and assigns it to the ApiVersion field.
func (o *HealthCheck) SetApiVersion(v []string) {
	o.ApiVersion = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *HealthCheck) GetService() string {
	if o == nil || isNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetServiceOk() (*string, bool) {
	if o == nil || isNil(o.Service) {
    return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *HealthCheck) HasService() bool {
	if o != nil && !isNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *HealthCheck) SetService(v string) {
	o.Service = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HealthCheck) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HealthCheck) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HealthCheck) SetStatus(v string) {
	o.Status = &v
}

func (o HealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !isNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableHealthCheck struct {
	value *HealthCheck
	isSet bool
}

func (v NullableHealthCheck) Get() *HealthCheck {
	return v.value
}

func (v *NullableHealthCheck) Set(val *HealthCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheck(val *HealthCheck) *NullableHealthCheck {
	return &NullableHealthCheck{value: val, isSet: true}
}

func (v NullableHealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


