/*
LogCom API

LogCom Swagger documentation

API version: 1.2.20
Contact: laborit@blutspende.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logcomapi

import (
	"encoding/json"
	"fmt"
)

// NotificationEventCategory the model 'NotificationEventCategory'
type NotificationEventCategory string

// List of NotificationEventCategory
const (
	Notification NotificationEventCategory = "NOTIFICATION"
)

// All allowed values of NotificationEventCategory enum
var AllowedNotificationEventCategoryEnumValues = []NotificationEventCategory{
	"NOTIFICATION",
}

func (v *NotificationEventCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationEventCategory(value)
	for _, existing := range AllowedNotificationEventCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationEventCategory", value)
}

// NewNotificationEventCategoryFromValue returns a pointer to a valid NotificationEventCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationEventCategoryFromValue(v string) (*NotificationEventCategory, error) {
	ev := NotificationEventCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationEventCategory: valid values are %v", v, AllowedNotificationEventCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationEventCategory) IsValid() bool {
	for _, existing := range AllowedNotificationEventCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationEventCategory value
func (v NotificationEventCategory) Ptr() *NotificationEventCategory {
	return &v
}

type NullableNotificationEventCategory struct {
	value *NotificationEventCategory
	isSet bool
}

func (v NullableNotificationEventCategory) Get() *NotificationEventCategory {
	return v.value
}

func (v *NullableNotificationEventCategory) Set(val *NotificationEventCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationEventCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationEventCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationEventCategory(val *NotificationEventCategory) *NullableNotificationEventCategory {
	return &NullableNotificationEventCategory{value: val, isSet: true}
}

func (v NullableNotificationEventCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationEventCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

