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

// AuditLogDTO struct for AuditLogDTO
type AuditLogDTO struct {
	// The category of the change
	Category *string `json:"category,omitempty"`
	// The creation timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user's ID who created
	CreatedById *uuid.UUID `json:"createdById,omitempty"`
	// The user's name who created
	CreatedByName *string `json:"createdByName,omitempty"`
	// The grouped changes
	GroupedChanges []AuditLogChangeDTO `json:"groupedChanges,omitempty"`
	// The ID
	Id *uuid.UUID `json:"id,omitempty"`
	// The message
	Message *string `json:"message,omitempty"`
	// The new value
	NewValue *string `json:"newValue,omitempty"`
	// The old value
	OldValue *string `json:"oldValue,omitempty"`
	// The request ID making dependent logs trackable
	RequestId *string `json:"requestId,omitempty"`
	// The service which the change affects
	ServiceAffected *string `json:"serviceAffected,omitempty"`
	// The service which created
	ServiceCreated *string `json:"serviceCreated,omitempty"`
	// The subject of the change
	Subject *string `json:"subject,omitempty"`
	// The name of the subject
	SubjectName *string `json:"subjectName,omitempty"`
	// The property name of the subject
	SubjectPropertyName *string `json:"subjectPropertyName,omitempty"`
}

// NewAuditLogDTO instantiates a new AuditLogDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogDTO() *AuditLogDTO {
	this := AuditLogDTO{}
	return &this
}

// NewAuditLogDTOWithDefaults instantiates a new AuditLogDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogDTOWithDefaults() *AuditLogDTO {
	this := AuditLogDTO{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AuditLogDTO) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AuditLogDTO) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *AuditLogDTO) SetCategory(v string) {
	o.Category = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuditLogDTO) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuditLogDTO) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AuditLogDTO) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *AuditLogDTO) GetCreatedById() uuid.UUID {
	if o == nil || isNil(o.CreatedById) {
		var ret uuid.UUID
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetCreatedByIdOk() (*uuid.UUID, bool) {
	if o == nil || isNil(o.CreatedById) {
    return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *AuditLogDTO) HasCreatedById() bool {
	if o != nil && !isNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given uuid.UUID and assigns it to the CreatedById field.
func (o *AuditLogDTO) SetCreatedById(v uuid.UUID) {
	o.CreatedById = &v
}

// GetCreatedByName returns the CreatedByName field value if set, zero value otherwise.
func (o *AuditLogDTO) GetCreatedByName() string {
	if o == nil || isNil(o.CreatedByName) {
		var ret string
		return ret
	}
	return *o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetCreatedByNameOk() (*string, bool) {
	if o == nil || isNil(o.CreatedByName) {
    return nil, false
	}
	return o.CreatedByName, true
}

// HasCreatedByName returns a boolean if a field has been set.
func (o *AuditLogDTO) HasCreatedByName() bool {
	if o != nil && !isNil(o.CreatedByName) {
		return true
	}

	return false
}

// SetCreatedByName gets a reference to the given string and assigns it to the CreatedByName field.
func (o *AuditLogDTO) SetCreatedByName(v string) {
	o.CreatedByName = &v
}

// GetGroupedChanges returns the GroupedChanges field value if set, zero value otherwise.
func (o *AuditLogDTO) GetGroupedChanges() []AuditLogChangeDTO {
	if o == nil || isNil(o.GroupedChanges) {
		var ret []AuditLogChangeDTO
		return ret
	}
	return o.GroupedChanges
}

// GetGroupedChangesOk returns a tuple with the GroupedChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetGroupedChangesOk() ([]AuditLogChangeDTO, bool) {
	if o == nil || isNil(o.GroupedChanges) {
    return nil, false
	}
	return o.GroupedChanges, true
}

// HasGroupedChanges returns a boolean if a field has been set.
func (o *AuditLogDTO) HasGroupedChanges() bool {
	if o != nil && !isNil(o.GroupedChanges) {
		return true
	}

	return false
}

// SetGroupedChanges gets a reference to the given []AuditLogChangeDTO and assigns it to the GroupedChanges field.
func (o *AuditLogDTO) SetGroupedChanges(v []AuditLogChangeDTO) {
	o.GroupedChanges = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditLogDTO) GetId() uuid.UUID {
	if o == nil || isNil(o.Id) {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditLogDTO) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *AuditLogDTO) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AuditLogDTO) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AuditLogDTO) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AuditLogDTO) SetMessage(v string) {
	o.Message = &v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *AuditLogDTO) GetNewValue() string {
	if o == nil || isNil(o.NewValue) {
		var ret string
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetNewValueOk() (*string, bool) {
	if o == nil || isNil(o.NewValue) {
    return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *AuditLogDTO) HasNewValue() bool {
	if o != nil && !isNil(o.NewValue) {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given string and assigns it to the NewValue field.
func (o *AuditLogDTO) SetNewValue(v string) {
	o.NewValue = &v
}

// GetOldValue returns the OldValue field value if set, zero value otherwise.
func (o *AuditLogDTO) GetOldValue() string {
	if o == nil || isNil(o.OldValue) {
		var ret string
		return ret
	}
	return *o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetOldValueOk() (*string, bool) {
	if o == nil || isNil(o.OldValue) {
    return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *AuditLogDTO) HasOldValue() bool {
	if o != nil && !isNil(o.OldValue) {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given string and assigns it to the OldValue field.
func (o *AuditLogDTO) SetOldValue(v string) {
	o.OldValue = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *AuditLogDTO) GetRequestId() string {
	if o == nil || isNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetRequestIdOk() (*string, bool) {
	if o == nil || isNil(o.RequestId) {
    return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *AuditLogDTO) HasRequestId() bool {
	if o != nil && !isNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *AuditLogDTO) SetRequestId(v string) {
	o.RequestId = &v
}

// GetServiceAffected returns the ServiceAffected field value if set, zero value otherwise.
func (o *AuditLogDTO) GetServiceAffected() string {
	if o == nil || isNil(o.ServiceAffected) {
		var ret string
		return ret
	}
	return *o.ServiceAffected
}

// GetServiceAffectedOk returns a tuple with the ServiceAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetServiceAffectedOk() (*string, bool) {
	if o == nil || isNil(o.ServiceAffected) {
    return nil, false
	}
	return o.ServiceAffected, true
}

// HasServiceAffected returns a boolean if a field has been set.
func (o *AuditLogDTO) HasServiceAffected() bool {
	if o != nil && !isNil(o.ServiceAffected) {
		return true
	}

	return false
}

// SetServiceAffected gets a reference to the given string and assigns it to the ServiceAffected field.
func (o *AuditLogDTO) SetServiceAffected(v string) {
	o.ServiceAffected = &v
}

// GetServiceCreated returns the ServiceCreated field value if set, zero value otherwise.
func (o *AuditLogDTO) GetServiceCreated() string {
	if o == nil || isNil(o.ServiceCreated) {
		var ret string
		return ret
	}
	return *o.ServiceCreated
}

// GetServiceCreatedOk returns a tuple with the ServiceCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetServiceCreatedOk() (*string, bool) {
	if o == nil || isNil(o.ServiceCreated) {
    return nil, false
	}
	return o.ServiceCreated, true
}

// HasServiceCreated returns a boolean if a field has been set.
func (o *AuditLogDTO) HasServiceCreated() bool {
	if o != nil && !isNil(o.ServiceCreated) {
		return true
	}

	return false
}

// SetServiceCreated gets a reference to the given string and assigns it to the ServiceCreated field.
func (o *AuditLogDTO) SetServiceCreated(v string) {
	o.ServiceCreated = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *AuditLogDTO) GetSubject() string {
	if o == nil || isNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetSubjectOk() (*string, bool) {
	if o == nil || isNil(o.Subject) {
    return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *AuditLogDTO) HasSubject() bool {
	if o != nil && !isNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *AuditLogDTO) SetSubject(v string) {
	o.Subject = &v
}

// GetSubjectName returns the SubjectName field value if set, zero value otherwise.
func (o *AuditLogDTO) GetSubjectName() string {
	if o == nil || isNil(o.SubjectName) {
		var ret string
		return ret
	}
	return *o.SubjectName
}

// GetSubjectNameOk returns a tuple with the SubjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetSubjectNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectName) {
    return nil, false
	}
	return o.SubjectName, true
}

// HasSubjectName returns a boolean if a field has been set.
func (o *AuditLogDTO) HasSubjectName() bool {
	if o != nil && !isNil(o.SubjectName) {
		return true
	}

	return false
}

// SetSubjectName gets a reference to the given string and assigns it to the SubjectName field.
func (o *AuditLogDTO) SetSubjectName(v string) {
	o.SubjectName = &v
}

// GetSubjectPropertyName returns the SubjectPropertyName field value if set, zero value otherwise.
func (o *AuditLogDTO) GetSubjectPropertyName() string {
	if o == nil || isNil(o.SubjectPropertyName) {
		var ret string
		return ret
	}
	return *o.SubjectPropertyName
}

// GetSubjectPropertyNameOk returns a tuple with the SubjectPropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogDTO) GetSubjectPropertyNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectPropertyName) {
    return nil, false
	}
	return o.SubjectPropertyName, true
}

// HasSubjectPropertyName returns a boolean if a field has been set.
func (o *AuditLogDTO) HasSubjectPropertyName() bool {
	if o != nil && !isNil(o.SubjectPropertyName) {
		return true
	}

	return false
}

// SetSubjectPropertyName gets a reference to the given string and assigns it to the SubjectPropertyName field.
func (o *AuditLogDTO) SetSubjectPropertyName(v string) {
	o.SubjectPropertyName = &v
}

func (o AuditLogDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if !isNil(o.CreatedByName) {
		toSerialize["createdByName"] = o.CreatedByName
	}
	if !isNil(o.GroupedChanges) {
		toSerialize["groupedChanges"] = o.GroupedChanges
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.NewValue) {
		toSerialize["newValue"] = o.NewValue
	}
	if !isNil(o.OldValue) {
		toSerialize["oldValue"] = o.OldValue
	}
	if !isNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !isNil(o.ServiceAffected) {
		toSerialize["serviceAffected"] = o.ServiceAffected
	}
	if !isNil(o.ServiceCreated) {
		toSerialize["serviceCreated"] = o.ServiceCreated
	}
	if !isNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !isNil(o.SubjectName) {
		toSerialize["subjectName"] = o.SubjectName
	}
	if !isNil(o.SubjectPropertyName) {
		toSerialize["subjectPropertyName"] = o.SubjectPropertyName
	}
	return json.Marshal(toSerialize)
}

type NullableAuditLogDTO struct {
	value *AuditLogDTO
	isSet bool
}

func (v NullableAuditLogDTO) Get() *AuditLogDTO {
	return v.value
}

func (v *NullableAuditLogDTO) Set(val *AuditLogDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogDTO(val *AuditLogDTO) *NullableAuditLogDTO {
	return &NullableAuditLogDTO{value: val, isSet: true}
}

func (v NullableAuditLogDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


