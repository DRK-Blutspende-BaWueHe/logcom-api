/*
 * LogCom API
 *
 * LogCom Swagger documentation
 *
 * API version: 1.0.58
 * Contact: laborit@blutspende.de
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package logcomapi

import (
	"github.com/google/uuid"
	"time"
)

type AuditLogDto struct {
	// The category of the change
	Category string `json:"category,omitempty"`
	// The creation timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user's ID who created
	CreatedById *uuid.UUID `json:"createdById,omitempty"`
	// The user's name who created
	CreatedByName string `json:"createdByName,omitempty"`
	// The grouped changes
	GroupedChanges []AuditLogChangeDto `json:"groupedChanges,omitempty"`
	// The ID
	Id *uuid.UUID `json:"id,omitempty"`
	// The message
	Message string `json:"message,omitempty"`
	// The new value
	NewValue string `json:"newValue,omitempty"`
	// The old value
	OldValue string `json:"oldValue,omitempty"`
	// The request ID making dependent logs trackable
	RequestId string `json:"requestId,omitempty"`
	// The service which the change affects
	ServiceAffected string `json:"serviceAffected,omitempty"`
	// The service which created
	ServiceCreated string `json:"serviceCreated,omitempty"`
	// The subject of the change
	Subject string `json:"subject,omitempty"`
	// The name of the subject
	SubjectName string `json:"subjectName,omitempty"`
	// The property name of the subject
	SubjectPropertyName string `json:"subjectPropertyName,omitempty"`
}
