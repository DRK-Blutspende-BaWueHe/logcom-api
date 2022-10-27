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

type UpdateNotificationMessageRequestDto struct {
	// The status of the message (only the \"SEEN\" status is supported)
	Status string `json:"status,omitempty"`
}
