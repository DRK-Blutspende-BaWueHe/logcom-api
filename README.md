# Go API client for LogCom

LogCom client and Swagger documentation

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
    * [Init the client](#init-the-client)
    * [Service-to-Service authorization](#service-to-service-authorization)
    * [Console logging](#console-logging)
    * [Audit logging](#audit-logging)
    * [Notification](#notification)
- [Builder usage](#builder-usage)
    * [Steps](#steps)
    * [Function descriptions](#function-descriptions)
- [About the author](#author)
- [Disclaimer](#disclaimer)

## Overview

The API documentation can be found [**here**](docs/README.md)

## Installation

```
go get -v github.com/DRK-Blutspende-BaWueHe/logcom-api
```

## Usage

### Init the client

The client initializes itself by default, only the `LOG_COM_URL` environment variable must be set.

For fine-grained usage, call the `logcom.Init` function.

```go
logcom.Init(logcom.Configuration{
    ServiceName: "Your Service (Project) Name",
    LogComURL:   "http://the.logcom.url",
    HeaderProvider: func (ctx context.Context) http.Header {
        if ginCtx, ok := ctx.(*gin.Context); ok {
            return ginCtx.Request.Header
        }
        return http.Header{}
    })
```

### Service-to-Service authorization

For using service-to-service authorization a specific interface has been provided. The API is not intended to be aware
of any authentication provider, so the interface must be implemented and set to obtain the necessary client credentials.

```go
type ClientCredentialProvider interface {
    GetClientCredential() (string, error)
}
```

It is also possible to pass a simple string token (just like it is used for user context aware authorization), when the
interface is not getting implemented.

#### Call when using the interface

```go
UseService2ServiceAuthorization()
```

#### Call when using the simple `string` Bearer JWT token

```go
WithBearerAuthorization(bearerToken string)
```

### Console logging

#### One-shot sending

Functions to use:

- For simple message: `SendConsoleLog(ctx context.Context, logLevel Level, message string)`

    - `ctx`: The context which contains the `Authorization` key and its value
    - `logLevel`: The console log level (Trace=-1, Debug=0, Info=1, Warning=2, Error=3, Fatal=4, Panic=5)
    - `message`: The console log message

    Example:

    ```go
    SendConsoleLog(ctx, zerolog.ErrorLevel, "Something went wrong")
    ```

- For more customized message: `SendConsoleLogWithModel(ctx context.Context, model logcomapi.CreateConsoleLogRequestDto)`

    - `ctx`: The context which contains the `Authorization` key and its value
    - `model`: The data model

    Example:

    ```go
    SendConsoleLogWithModel(ctx, logcomapi.CreateConsoleLogRequestDto{
        CreatedAt:     nil,
        CreatedById:   nil,
        CreatedByName: "",
        Level:         0,
        Message:       "",
        Service:       "",
    })
    ```

#### Builder based

Function to use: `logcom.Log()`

Example:

```go
err := logcom.Log().
	Level(logcom.DebugLevel).
	Message("Debug log")
    WithContext(ctx).
    Send()
```

### Audit logging

#### Creation - OneShot sending

Function to use: `SendAuditLogWithCreation(ctx context.Context, subject, subjectName string, newValue interface{})`

- `ctx`: The context which contains the `Authorization` key and its value
- `subject`: The created subject
- `subjectName`: The unique identifier or name of the subject
- `newValue`: The created object. **Nullable**. Use `nil` when the created object data is not important

Example:

```go
err := logcom.SendAuditLogWithCreation(ctx, "MATERIAL", "ANTIG", newMaterialDTO)
```

The translated result on Bloodlab-UI:

```
ANTIG material was created successfully
```

#### Creation - Builder based

Functions to use:

- `logcom.Audit().Create(subject, subjectName string, newValue interface{})`
- `logcom.AuditCreation(subject, subjectName string, newValue interface{})`

Example:

```go
err := logcom.AuditCreation("MATERIAL", "ANTIG", newMaterialDTO).
    WithContext(ctx).
    OnComplete(func(err error) {
        if err != nil {
            _ = txConn.Rollback()
        }
    }).
    Send()
```

The translated result on Bloodlab-UI:

```
ANTIG material was created successfully
```

#### Creation - Batched

Function to use: `logcom.Audit().BatchCreate(subject string)`

Example:

```go
auditBatch := logcom.Audit().
	BatchCreate("ORDER")
for _, order := range orders {
    auditBatch.CreateItem(order.ID, order)	
}
err := auditBatch.WithContext(ctx).
    Send()
```

#### Modification - OneShot sending

Function to
use: `SendAuditLogWithModification(ctx context.Context, subject, subjectName string, oldValue, newValue interface{})`

- `ctx`: The context which contains the `Authorization` key and its value
- `subject`: The modified subject
- `subjectName`: The unique identifier or name of the subject
- `oldValue`: The original object
- `newValue`: The modified object

Example:

```go
err := logcom.SendAuditLogWithModification(ctx, "MATERIAL", "ANTIG", originalMaterialDTO, updatedMaterialDTO)
```

The translated result in Bloodlab-UI:

- When has only 1 change
  ```
  ANTIG material name was modified successfully from "ANTIG Material" to "Anti-G material"
  ```

- When has more than 1 changes
  ```
  3 properties of ANTIG material were modified successfully

  Material code was modified successfully from "ANTIG" to "ANTI-G"
  Material enabled state was modified successfully from "false" to "true"
  Material name was modified successfully from "ANTIG Material" to "Anti-G material"
  ```

#### Modification - Builder based

Functions to use for builder initialization:

1. `logcom.Audit().Modify(subject, subjectName string, oldValue, newValue interface{})`
2. `logcom.AuditModification(subject, subjectName string, oldValue, newValue interface{})`

Example:

```go
err := logcom.AuditModification("MATERIAL", "ANTIG", originalMaterialDTO, updatedMaterialDTO).
    WithContext(ctx).
    IgnoreChangeOf("modifiedBy", "modifiedAt").
    OnComplete(func (err error) {
        if err != nil {
            _ = txConn.Rollback()
        }
    }).
    Send()
```

The translated result in Bloodlab-UI:

- When has only 1 change
  ```
  ANTIG material name was modified successfully from "ANTIG Material" to "Anti-G material"
  ```

- When has more than 1 changes
  ```
  3 properties of ANTIG material were modified successfully

  Material code was modified successfully from "ANTIG" to "ANTI-G"
  Material enabled state was modified successfully from "false" to "true"
  Material name was modified successfully from "ANTIG Material" to "Anti-G material"
  ```

#### Modification - Batched

Function to use: `logcom.Audit().BatchModify(subject string)`

Example:

```go
err := logcom.Audit().
    BatchModify("ORDER").
    CreateItem(newOrder.ID, newOrder).
    DeleteItem(deletedOrder.ID, deletedOrder).
    WithContext(ctx).
    Send()
```

#### Modification - Grouped

Consider the case of order modification which has basic information and samples with material. They are represented as
separate objects at the end, but still they have to be audit-logged as one change.

Function to use: `SendAuditLogGroup(ctx context.Context, auditLogCollector *AuditLogCollector)`

- `ctx`: The context which contains the `Authorization` key and its value
- `auditLogCollector`: The collection which contains the individual changes belonging to the modified subject

Example:

```go
err := logcom.SendAuditLogGroup(ctx, &auditLogCollector)
```

#### Deletion - OneShot sending

Function to use: `SendAuditLogWithDeletion(ctx context.Context, subject, subjectName string, oldValue interface{})`

- `ctx`: The context which contains the `Authorization` key and its value
- `subject`: The deleted subject
- `subjectName`: The unique identifier or name of the subject
- `oldValue`: The deleted object. **Nullable**. Use `nil` when the deleted object data is not important

Example:

```go
err := logcom.SendAuditLogWithDeletion(ctx, "MATERIAL", "ANTIG", deletedMaterialDTO)
```

The translated result on Bloodlab-UI:

```
ANTIG material was removed successfully
```

#### Deletion - Builder based

Functions to use for builder initialization:

1. `logcom.Audit().Delete(subject, subjectName string, oldValue)`
2. `logcom.AuditDeletion(subject, subjectName string, oldValue interface{})`

Example:

```go
err := logcom.AuditDeletion("MATERIAL", "ANTIG", deletedMaterialDTO).
    WithContext(ctx).
    OnComplete(func(err error) {
        if err != nil {
            _ = txConn.Rollback()
        }
    }).
    Send()
```

The translated result on Bloodlab-UI:

```
ANTIG material was removed successfully
```

#### Deletion - Batched

Function to use: `logcom.Audit().BatchDelete(subject string)`

Example:

```go
auditBatch := logcom.Audit().
	BatchDelete("ORDER")
for _, order := range orders {
    auditBatch.DeleteItem(order.ID, order)	
}
err := auditBatch.WithContext(ctx).
    Send()
```

#### Custom Action - OneShot sending

Function to use: `sendAuditLogWithCustomAction(ctx context.Context, action, subject, subjectName string, oldValue, newValue interface{}, ignoredProperties ...string)`

- `ctx`: The context which contains the `Authorization` key and its value
- `action`: The custom action type as string
- `subject`: The created subject
- `subjectName`: The unique identifier or name of the subject
- `oldValue`: The original object. **Nullable**. Use `nil` when the event's old object data is not important
- `newValue`: The created object. **Nullable**. Use `nil` when the event's new object data is not important

Example:

```go
err := logcom.sendAuditLogWithCustomAction(ctx, "MY_CUSTOM_ACTION", "MATERIAL", "ANTIG", oldMaterialDTO, newMaterialDTO)
```

#### Custom Action - Builder based

Functions to use:

- `logcom.Audit().CustomAction(action, subject, subjectName string, oldValue, newValue interface{})`
- `logcom.AuditCustomAction(action, subject, subjectName string, oldValue, newValue interface{})`

Example:

```go
err := logcom.AuditCustomAction("MY_CUSTOM_ACTION","MATERIAL", "ANTIG", oldMaterialDTO, newMaterialDTO).
    WithContext(ctx).
    OnComplete(func(err error) {
        if err != nil {
            _ = txConn.Rollback()
        }
    }).
    Send()
```

#### Custom Action - Batched

Function to use: `logcom.Audit().BatchCustomAction(action, subject string)`

Example:

```go
auditBatch := logcom.Audit().
BatchCustomAction("MY_CUSTOM_ACTION","ORDER")
for _, order := range orders {
	oldValue := order.SomeField
	.....
    
    auditBatch.CustomActionItem("MY_CUSTOM_ACTION", order.ID, oldValue, newValue)	
}
err := auditBatch.WithContext(ctx).
    Send()
```

### Notification

#### OneShot sending

Function to use: `SendNotification(ctx context.Context, eventCategory string, message string, targets map[string][]string)`

- `ctx`: The context which contains the `Authorization` key and its value
- `eventCategory`: The category of the event (e.g. `NOTIFICATION`)
- `message`: The notification message
- `targets`: The desired range of users / roles / sessions to whom the notification should be sent (Key: `ROLE`, `SESSION`, `USER`; Value: list of targets)

Example:

```go
err := logcom.SendNotification(ctx, "NOTIFICATION", "This is a notification for doctors", map[string][]string{"ROLE": {"doctor"}})
```

#### Builder based

Function to use: `logcom.Notify()`

Example:

```go
err := logcom.Notify().
    Roles("doctor", "admin").
    Message("Test notification").
    WithContext(ctx).
    Send()
```

## Builder Usage

### Steps

1. Initialization
    - `logcom.Audit()`
    - `logcom.AuditCreation(subject, subjectName string, newValue interface{})`
    - `logcom.AuditModification(subject, subjectName string, oldValue, newValue interface{})`
    - `logcom.AuditDeletion(subject, subjectName string, oldValue interface{})`
    - `logcom.AuditCustomAction(action, subject, subjectName string, oldValue interface{}, newValue interface{})`
    - `logcom.Notify()`
    - `logcom.Log()`

2. Initialization sub-actions
    1. Console Log
        - `Level(level Level)`
        - `Message(message string)`
        - `MessageF(format string, params ...any)`
    2. Audit Log
        - `Create(subject, subjectName string, newValue interface{})`
        - `Modify(subject, subjectName string, oldValue, newValue interface{})`
        - `Delete(subject, subjectName string, oldValue interface{})`
        - `CustomAction(action, subject, subjectName string, oldValue interface{}, newValue interface{})`
        - `BatchCreate(subject string)`
            - `CreateItem(subjectName string, newValue interface{})`
        - `BatchModify(subject string)`
            - `CreateItem(subjectName string, newValue interface{})`
            - `ModifyItem(subjectName string, oldValue, newValue interface{})`
            - `DeleteItem(subjectName string, oldValue interface{})`
        - `BatchDelete(subject string)`
            - `DeleteItem(subjectName string, oldValue interface{})`        
        - `BatchCustomAction(action, subject string)`
          - `CustomActionItem(action, subjectName string, oldValue interface{}, newValue interface{})`
        - `GroupedModify(subject, subjectName string)`
            - `AddCreation(subject, subjectName string, newValue interface{})`
            - `AddModification(subject, subjectName string, oldValue, newValue interface{})`
            - `AddDeletion(subject, subjectName string, oldValue interface{})`
    3. Notification
        - `Message(message string)`
        - `Roles(targets ...string)`
        - `Sessions(targets ...string)`
        - `Users(targets ...string)`

3. Configuration
    - `UseService2ServiceAuthorization()`
    - `WithBearerAuthorization(bearerToken string)`
    - `WithContext(ctx context.Context)`
    - `WithTransactionID(transactionID uuid.UUID)`

4. Action
    1. Console Log
        - `OnComplete(onCompleteCallback func(error))`
        - `Send()`
    2. Audit Log
        - `IgnoreChangeOf(propertyNames ...string)`
        - `AndNotify()`
        - `AndLog(logLevel zerolog.Level, message string)`
        - `OnComplete(onCompleteCallback func(error))`
        - `Send()`
    3. Notification
        - `AndLog(logLevel zerolog.Level, message string)`
        - `OnComplete(onCompleteCallback func(error))`
        - `Send()`

### Function descriptions

`AddCreation(subject, subjectName string, newValue interface{})`

- Adds a creation change to the grouped audit log
- `subject`: The created subject
- `subjectName`: The unique identifier or name of the subject
- `newValue`: The created object. **Nullable**. Use `nil` when the created object data is not important

`AddDeletion(subject, subjectName string, oldValue interface{})`

- Adds a deletion change to the grouped audit log
- `subject`: The deleted subject
- `subjectName`: The unique identifier or name of the subject
- `oldValue`: The deleted object. **Nullable**. Use `nil` when the deleted object data is not important

`AddModification(subject, subjectName string, oldValue, newValue interface{})`

- Adds a modification change to the grouped audit log
- `subject`: The modified subject
- `subjectName`: The unique identifier or name of the subject
- `oldValue`: The original object
- `newValue`: The modified object

`AddCustomEvent(action, subject, subjectName string, oldValue interface{}, newValue interface{})`

- Adds a custom eevent to the grouped audit log
- `action`: The custom event
- `subject`: The event subject
- `subjectName`: The unique identifier or name of the subject
- `oldValue`: The original object
- `newValue`: The modified object

`AndNotify()`

- Initializes a **Notification sub-action** as part of the audit log

`AndLog(logLevel zerolog.Level, message string)`

- Sets a *Console Log* action as part of the audit log and / or notification
- `logLevel`: The console log level
- `message`: The console log message

`Audit()`

- Initializes an audit log builder

`AuditCreation(subject, subjectName string, newValue interface{})`

- Initializes an audit log creation builder
- `subject`: The created subject
- `subjectName`: The unique identifier or name of the subject
- `newValue`: The created object. **Nullable**. Use `nil` when the created object data is not important

`AuditDeletion(subject, subjectName string, oldValue interface{})`

- Initializes an audit log deletion builder
- `subject`: The deleted subject
- `subjectName`: The unique identifier or name of the subject
- `oldValue`: The deleted object. **Nullable**. Use `nil` when the deleted object data is not important

`AuditModification(subject, subjectName string, oldValue, newValue interface{})`

- Initializes an audit log modification builder
- `subject`: The modified subject
- `subjectName`: The unique identifier or name of the subject
- `oldValue`: The original object
- `newValue`: The modified object

`BatchCreate(subject string)`

- Specifies the audit log operation (as batched creation)
- `subject`: The main subject of the created items

`BatchDelete(subject string)`

- Specifies the audit log operation (as batched deletion)
- `subject`: The main subject of the deleted items

`BatchModify(subject string)`

- Specifies the audit log operation (as batched modification)
- `subject`: The main subject of the modified items

`Create(subject, subjectName string, newValue interface{})`

- Specifies the audit log operation (as creation)
- `subject`: The created subject
- `subjectName`: The unique identifier or name of the subject
- `newValue`: The created object. **Nullable**. Use `nil` when the created object data is not important

`CreateItem(subjectName string, newValue interface{})`

- Adds a new creation audit log to the initialized batch
- `subjectName`: The unique identifier or name of the subject
- `newValue`: The created object. **Nullable**. Use `nil` when the created object data is not important

`Delete(subject, subjectName string, oldValue interface{})`

- Specifies the audit log operation (as deletion)
- `subject`: The deleted subject
- `subjectName`: The unique identifier or name of the subject
- `oldValue`: The deleted object. **Nullable**. Use `nil` when the deleted object data is not important

`DeleteItem(subjectName string, oldValue interface{})`

- Adds a new deletion audit log to the initialized batch
- `subjectName`: The unique identifier or name of the subject
- `newValue`: The deleted object. **Nullable**. Use `nil` when the deleted object data is not important

`GroupedModify(subject, subjectName string)`

- Groups individual changes (e.g. added / modified / deleted a related subject) belonging to the modified subject
- `subject`: The modified subject
- `subjectName`: The unique identifier or name of the subject

`IgnoreChangeOf(propertyNames ...string)`

- Sets ignored properties of a subject for the modification operation
- `propertyNames`: The ignored property names which are not considered as changes

`Level(logLevel Level)`

- Sets the log level of the console log
- `logLevel`: The console log level (Trace=-1, Debug=0, Info=1, Warning=2, Error=3, Fatal=4, Panic=5)

`Log()`

- Initializes a console log builder

`Message(message string)`

- Sets the message of the console log or notification
- `message`: The console log or notification message

`MessageF(format string, params ...any)`

- Sets the formatted message of the console log
- `format`: The formatted message
- `params`: The message parameters

`Modify(subject, subjectName string, oldValue, newValue interface{})`

- Specifies the audit log operation (as modification)
- `subject`: The modified subject
- `subjectName`: The unique identifier or name of the subject
- `oldValue`: The original object
- `newValue`: The modified object

`ModifyItem(subjectName string, oldValue, newValue interface{})`

- Adds a new modification audit log to the initialized batch
- `subjectName`: The unique identifier or name of the subject
- `newValue`: The modified object. **Nullable**. Use `nil` when the modified object data is not important

`Notify()`

- Initializes a notification builder

`OnComplete(onCompleteCallback func(error))`

- Sets a callback function for handling errors (e.g. rolling back the transaction)
- `onCompleteCallback`: The callback function

`Roles(targets ...string)`

- Specifies the notification targets (as roles)
- `targets`: List of roles

`Send()`

- Sends the audit log and / or notification

`Sessions(targets ...string)`

- Specifies the notification targets (as sessions)
- `targets`: List of sessions

`Users(targets ...string)`

- Specifies the notification targets (as users)
- `targets`: List of users (ID, username, other unique identifier)

`UseService2ServiceAuthorization()`

- Uses the client credentials provided by the calling service

`WithBearerAuthorization(bearerToken string)`

- Sets the bearer JWT token for the request
- `bearerToken`: The bearer token

`WithContext(ctx context.Context)`

- Sets the context for the request
- `ctx`: The context which contains the `Authorization` key and its value

`WithTransactionID(transactionID uuid.UUID)`

- Sets the transaction ID for the request
- `transactionID`: The transaction ID

## Author

laborit@blutspende.de

## Disclaimer

By making use of any information, content and source code in this repository, You agree to the following:

**NO WARRANTIES**

All the information, content and source code provided in this repository is provided "**AS-IS**" and with **NO
WARRANTIES**.

**DISCLAIMER OF LIABILITY**

*DRK Blutspendedienst BaWü u. Hessen gGmbH* specifically DISCLAIMS LIABILITY FOR INCIDENTAL OR CONSEQUENTIAL DAMAGES and
assumes no responsibility or liability for any loss or damage suffered by any person as a result of the use or misuse of
any of the information, content or source code in this repository.

*DRK Blutspendedienst BaWü u. Hessen gGmbH* assumes or undertakes NO LIABILITY for any loss or damage suffered as a
result of the use, misuse or reliance on the information, content and source code in this repository.

**USE AT YOUR OWN RISK**

This repository is for *DRK Blutspendedienst BaWü u. Hessen gGmbH* only.
