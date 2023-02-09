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

// CreateNotificationRequestDTO struct for CreateNotificationRequestDTO
type CreateNotificationRequestDTO struct {
	// The log timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user's ID who created
	CreatedById *uuid.UUID `json:"createdById,omitempty"`
	// The user's name who created
	CreatedByName *string `json:"createdByName,omitempty"`
	// The notification event category
	EventCategory *NotificationEventCategory `json:"eventCategory,omitempty"`
	// The log message
	Message string `json:"message"`
	// The service which sent the log
	Service string `json:"service"`
	// The targets who will receive the notification
	Targets *map[string][]string `json:"targets,omitempty"`
}

// NewCreateNotificationRequestDTO instantiates a new CreateNotificationRequestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNotificationRequestDTO(message string, service string) *CreateNotificationRequestDTO {
	this := CreateNotificationRequestDTO{}
	this.Message = message
	this.Service = service
	return &this
}

// NewCreateNotificationRequestDTOWithDefaults instantiates a new CreateNotificationRequestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNotificationRequestDTOWithDefaults() *CreateNotificationRequestDTO {
	this := CreateNotificationRequestDTO{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateNotificationRequestDTO) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequestDTO) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateNotificationRequestDTO) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CreateNotificationRequestDTO) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *CreateNotificationRequestDTO) GetCreatedById() uuid.UUID {
	if o == nil || isNil(o.CreatedById) {
		var ret uuid.UUID
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequestDTO) GetCreatedByIdOk() (*uuid.UUID, bool) {
	if o == nil || isNil(o.CreatedById) {
    return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *CreateNotificationRequestDTO) HasCreatedById() bool {
	if o != nil && !isNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given uuid.UUID and assigns it to the CreatedById field.
func (o *CreateNotificationRequestDTO) SetCreatedById(v uuid.UUID) {
	o.CreatedById = &v
}

// GetCreatedByName returns the CreatedByName field value if set, zero value otherwise.
func (o *CreateNotificationRequestDTO) GetCreatedByName() string {
	if o == nil || isNil(o.CreatedByName) {
		var ret string
		return ret
	}
	return *o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequestDTO) GetCreatedByNameOk() (*string, bool) {
	if o == nil || isNil(o.CreatedByName) {
    return nil, false
	}
	return o.CreatedByName, true
}

// HasCreatedByName returns a boolean if a field has been set.
func (o *CreateNotificationRequestDTO) HasCreatedByName() bool {
	if o != nil && !isNil(o.CreatedByName) {
		return true
	}

	return false
}

// SetCreatedByName gets a reference to the given string and assigns it to the CreatedByName field.
func (o *CreateNotificationRequestDTO) SetCreatedByName(v string) {
	o.CreatedByName = &v
}

// GetEventCategory returns the EventCategory field value if set, zero value otherwise.
func (o *CreateNotificationRequestDTO) GetEventCategory() NotificationEventCategory {
	if o == nil || isNil(o.EventCategory) {
		var ret NotificationEventCategory
		return ret
	}
	return *o.EventCategory
}

// GetEventCategoryOk returns a tuple with the EventCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequestDTO) GetEventCategoryOk() (*NotificationEventCategory, bool) {
	if o == nil || isNil(o.EventCategory) {
    return nil, false
	}
	return o.EventCategory, true
}

// HasEventCategory returns a boolean if a field has been set.
func (o *CreateNotificationRequestDTO) HasEventCategory() bool {
	if o != nil && !isNil(o.EventCategory) {
		return true
	}

	return false
}

// SetEventCategory gets a reference to the given NotificationEventCategory and assigns it to the EventCategory field.
func (o *CreateNotificationRequestDTO) SetEventCategory(v NotificationEventCategory) {
	o.EventCategory = &v
}

// GetMessage returns the Message field value
func (o *CreateNotificationRequestDTO) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequestDTO) GetMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CreateNotificationRequestDTO) SetMessage(v string) {
	o.Message = v
}

// GetService returns the Service field value
func (o *CreateNotificationRequestDTO) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequestDTO) GetServiceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *CreateNotificationRequestDTO) SetService(v string) {
	o.Service = v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *CreateNotificationRequestDTO) GetTargets() map[string][]string {
	if o == nil || isNil(o.Targets) {
		var ret map[string][]string
		return ret
	}
	return *o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequestDTO) GetTargetsOk() (*map[string][]string, bool) {
	if o == nil || isNil(o.Targets) {
    return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *CreateNotificationRequestDTO) HasTargets() bool {
	if o != nil && !isNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given map[string][]string and assigns it to the Targets field.
func (o *CreateNotificationRequestDTO) SetTargets(v map[string][]string) {
	o.Targets = &v
}

func (o CreateNotificationRequestDTO) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.EventCategory) {
		toSerialize["eventCategory"] = o.EventCategory
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["service"] = o.Service
	}
	if !isNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNotificationRequestDTO struct {
	value *CreateNotificationRequestDTO
	isSet bool
}

func (v NullableCreateNotificationRequestDTO) Get() *CreateNotificationRequestDTO {
	return v.value
}

func (v *NullableCreateNotificationRequestDTO) Set(val *CreateNotificationRequestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNotificationRequestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNotificationRequestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNotificationRequestDTO(val *CreateNotificationRequestDTO) *NullableCreateNotificationRequestDTO {
	return &NullableCreateNotificationRequestDTO{value: val, isSet: true}
}

func (v NullableCreateNotificationRequestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNotificationRequestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


