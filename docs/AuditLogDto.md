# AuditLogDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | The category of the change | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The creation timestamp | [optional] [default to null]
**CreatedById** | [***uuid.UUID**](uuid.UUID.md) | The user&#39;s ID who created | [optional] [default to null]
**CreatedByName** | **string** | The user&#39;s name who created | [optional] [default to null]
**GroupedChanges** | [**[]AuditLogChangeDto**](AuditLogChangeDTO.md) | The grouped changes | [optional] [default to null]
**Id** | [***uuid.UUID**](uuid.UUID.md) | The ID | [optional] [default to null]
**Message** | **string** | The message | [optional] [default to null]
**NewValue** | **string** | The new value | [optional] [default to null]
**OldValue** | **string** | The old value | [optional] [default to null]
**RequestId** | **string** | The request ID making dependent logs trackable | [optional] [default to null]
**ServiceAffected** | **string** | The service which the change affects | [optional] [default to null]
**ServiceCreated** | **string** | The service which created | [optional] [default to null]
**Subject** | **string** | The subject of the change | [optional] [default to null]
**SubjectName** | **string** | The name of the subject | [optional] [default to null]
**SubjectPropertyName** | **string** | The property name of the subject | [optional] [default to null]

[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


