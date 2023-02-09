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
	"github.com/google/uuid"
	"time"
)

// CreateConsoleLogRequestDTO struct for CreateConsoleLogRequestDTO
type CreateConsoleLogRequestDTO struct {
	// The log timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user's ID who created
	CreatedById *uuid.UUID `json:"createdById,omitempty"`
	// The user's name who created
	CreatedByName *string `json:"createdByName,omitempty"`
	// The log level (Trace=-1, Debug=0, Info=1, Warning=2, Error=3, Fatal=4, Panic=5)
	Level LogLevel `json:"level"`
	// The log message
	Message string `json:"message"`
	// The service which sent the log
	Service string `json:"service"`
}

// NewCreateConsoleLogRequestDTO instantiates a new CreateConsoleLogRequestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConsoleLogRequestDTO(level LogLevel, message string, service string) *CreateConsoleLogRequestDTO {
	this := CreateConsoleLogRequestDTO{}
	this.Level = level
	this.Message = message
	this.Service = service
	return &this
}

// NewCreateConsoleLogRequestDTOWithDefaults instantiates a new CreateConsoleLogRequestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConsoleLogRequestDTOWithDefaults() *CreateConsoleLogRequestDTO {
	this := CreateConsoleLogRequestDTO{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateConsoleLogRequestDTO) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConsoleLogRequestDTO) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateConsoleLogRequestDTO) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CreateConsoleLogRequestDTO) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *CreateConsoleLogRequestDTO) GetCreatedById() uuid.UUID {
	if o == nil || isNil(o.CreatedById) {
		var ret uuid.UUID
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConsoleLogRequestDTO) GetCreatedByIdOk() (*uuid.UUID, bool) {
	if o == nil || isNil(o.CreatedById) {
    return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *CreateConsoleLogRequestDTO) HasCreatedById() bool {
	if o != nil && !isNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given uuid.UUID and assigns it to the CreatedById field.
func (o *CreateConsoleLogRequestDTO) SetCreatedById(v uuid.UUID) {
	o.CreatedById = &v
}

// GetCreatedByName returns the CreatedByName field value if set, zero value otherwise.
func (o *CreateConsoleLogRequestDTO) GetCreatedByName() string {
	if o == nil || isNil(o.CreatedByName) {
		var ret string
		return ret
	}
	return *o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConsoleLogRequestDTO) GetCreatedByNameOk() (*string, bool) {
	if o == nil || isNil(o.CreatedByName) {
    return nil, false
	}
	return o.CreatedByName, true
}

// HasCreatedByName returns a boolean if a field has been set.
func (o *CreateConsoleLogRequestDTO) HasCreatedByName() bool {
	if o != nil && !isNil(o.CreatedByName) {
		return true
	}

	return false
}

// SetCreatedByName gets a reference to the given string and assigns it to the CreatedByName field.
func (o *CreateConsoleLogRequestDTO) SetCreatedByName(v string) {
	o.CreatedByName = &v
}

// GetLevel returns the Level field value
func (o *CreateConsoleLogRequestDTO) GetLevel() LogLevel {
	if o == nil {
		var ret LogLevel
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *CreateConsoleLogRequestDTO) GetLevelOk() (*LogLevel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *CreateConsoleLogRequestDTO) SetLevel(v LogLevel) {
	o.Level = v
}

// GetMessage returns the Message field value
func (o *CreateConsoleLogRequestDTO) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CreateConsoleLogRequestDTO) GetMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CreateConsoleLogRequestDTO) SetMessage(v string) {
	o.Message = v
}

// GetService returns the Service field value
func (o *CreateConsoleLogRequestDTO) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *CreateConsoleLogRequestDTO) GetServiceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *CreateConsoleLogRequestDTO) SetService(v string) {
	o.Service = v
}

func (o CreateConsoleLogRequestDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if !isNil(o.CreatedByName) {
		toSerialize["createdByName"] = o.CreatedByName
	}
	if true {
		toSerialize["level"] = o.Level
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["service"] = o.Service
	}
	return json.Marshal(toSerialize)
}

type NullableCreateConsoleLogRequestDTO struct {
	value *CreateConsoleLogRequestDTO
	isSet bool
}

func (v NullableCreateConsoleLogRequestDTO) Get() *CreateConsoleLogRequestDTO {
	return v.value
}

func (v *NullableCreateConsoleLogRequestDTO) Set(val *CreateConsoleLogRequestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConsoleLogRequestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConsoleLogRequestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConsoleLogRequestDTO(val *CreateConsoleLogRequestDTO) *NullableCreateConsoleLogRequestDTO {
	return &NullableCreateConsoleLogRequestDTO{value: val, isSet: true}
}

func (v NullableCreateConsoleLogRequestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConsoleLogRequestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


