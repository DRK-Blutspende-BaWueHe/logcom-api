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

type ConsoleLogDto struct {
	// The log timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user's ID who created
	CreatedById *uuid.UUID `json:"createdById,omitempty"`
	// The user's name who created
	CreatedByName string `json:"createdByName,omitempty"`
	// The ID
	Id *uuid.UUID `json:"id,omitempty"`
	// The log level (Trace=-1, Debug=0, Info=1, Warning=2, Error=3, Fatal=4, Panic=5)
	Level int32 `json:"level,omitempty"`
	// The log message
	Message string `json:"message,omitempty"`
	// The request ID making dependent logs trackable
	RequestId *uuid.UUID `json:"requestId,omitempty"`
	// The service which sent the log
	Service string `json:"service,omitempty"`
}
