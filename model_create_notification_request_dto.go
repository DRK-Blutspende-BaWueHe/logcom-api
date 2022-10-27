/*
 * LogCom API
 *
 * LogCom Swagger documentation
 *
 * API version: 1.0.63
 * Contact: laborit@blutspende.de
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package logcomapi

import (
	"github.com/google/uuid"
	"time"
)

type CreateNotificationRequestDto struct {
	// The log timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user's ID who created
	CreatedById *uuid.UUID `json:"createdById,omitempty"`
	// The user's name who created
	CreatedByName string `json:"createdByName,omitempty"`
	// The notification event category
	EventCategory string `json:"eventCategory,omitempty"`
	// The log message
	Message string `json:"message"`
	// The service which sent the log
	Service string `json:"service"`
	// The targets who will receive the notification
	Targets map[string][]string `json:"targets,omitempty"`
}
