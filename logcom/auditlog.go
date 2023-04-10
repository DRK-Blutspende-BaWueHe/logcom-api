package logcom

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/DRK-Blutspende-BaWueHe/logcom-api"
	"github.com/google/uuid"
)

const (
	batchedCreation = iota
	batchedModification
	batchedDeletion
	batchedSubjectModification
	batchedCustomAction
)

var batchedOperationList = []int{
	batchedDeletion,
	batchedModification,
	batchedDeletion,
	batchedSubjectModification,
	batchedCustomAction,
}

type AuditLogAction interface {
	IgnoreChangeOf(propertyNames ...string) AuditLogAction
	AndNotify() NotificationOperation[AuditLogAction]
	AndLog(logLevel logcomapi.LogLevel, message string) AuditLogAction
	OnComplete(onCompleteCallback func(error)) AuditLogAction
	Send() error
}

type AuditLogConfiguration interface {
	UseService2ServiceAuthorization() AuditLogAction
	WithBearerAuthorization(bearerToken string) AuditLogAction
	WithContext(ctx context.Context) AuditLogAction
	WithTransactionID(transactionID uuid.UUID) AuditLogConfiguration
}

type AuditLogOperation interface {
	BatchCreate(subject string) BatchedAuditLogOperation
	BatchDelete(subject string) BatchedAuditLogOperation
	BatchModify(subject string) BatchedAuditLogOperation
	BatchCustomAction(action, subject string) BatchedAuditLogOperation
	Create(subject, subjectName string, newValue interface{}) AuditLogConfiguration
	Modify(subject, subjectName string, oldValue, newValue interface{}) AuditLogConfiguration
	Delete(subject, subjectName string, oldValue interface{}) AuditLogConfiguration
	CustomAction(action, subject, subjectName string, oldValue interface{}, newValue interface{}) AuditLogConfiguration
	GroupedModify(subject, subjectName string) GroupedModificationAuditLogOperation
}

type BatchedAuditLogOperation interface {
	AuditLogConfiguration
	CreateItem(subjectName string, newValue interface{}) BatchedAuditLogOperation
	ModifyItem(subjectName string, oldValue, newValue interface{}) BatchedAuditLogOperation
	ModifyItemWithModelChanges(subjectName string, changes []ModelChange) BatchedAuditLogOperation
	DeleteItem(subjectName string, oldValue interface{}) BatchedAuditLogOperation
	CustomActionItem(action, subjectName string, oldValue interface{}, newValue interface{}) BatchedAuditLogOperation
}

type GroupedModificationAuditLogOperation interface {
	AuditLogConfiguration
	AddCreation(subject, subjectName string, newValue interface{}) GroupedModificationAuditLogOperation
	AddModification(subject, subjectName string, oldValue, newValue interface{}) GroupedModificationAuditLogOperation
	AddModificationWithModelChanges(subject, subjectName string, changes []ModelChange) GroupedModificationAuditLogOperation
	AddDeletion(subject, subjectName string, oldValue interface{}) GroupedModificationAuditLogOperation
	AddCustomEvent(action, subject, subjectName string, oldValue interface{}, newValue interface{}) GroupedModificationAuditLogOperation
}

type AuditLogModelDiff interface {
	GetChanges(model interface{}, ignoredProperties map[string]interface{}) []ModelChange
}

type ModelChange struct {
	PropertyName string
	OldValue     interface{}
	NewValue     interface{}
}

type AuditLogCollector struct {
	parentAuditLog logcomapi.CreateAuditLogRequestDTO
	auditLogs      map[string]map[string][]logcomapi.CreateAuditLogRequestDTO
}

type auditLog[T any] struct {
	ctx                context.Context
	operation          string
	newValue           interface{}
	oldValue           interface{}
	subject            string
	subjectName        string
	ignoredProperties  []string
	consoleLog         *consoleLog
	notification       *notification[T]
	batchedAuditLogMap map[int]*AuditLogCollector
	onCompleteCallback func(error)
}

func Audit() AuditLogOperation {
	return &auditLog[AuditLogAction]{
		ctx:                context.TODO(),
		onCompleteCallback: func(error) {},
	}
}

func AuditCreation(subject, subjectName string, newValue interface{}) AuditLogConfiguration {
	return &auditLog[AuditLogAction]{
		operation:   "CREATION",
		newValue:    newValue,
		subject:     subject,
		subjectName: subjectName,
	}
}

func AuditModification(subject, subjectName string, oldValue, newValue interface{}) AuditLogConfiguration {
	return &auditLog[AuditLogAction]{
		operation:   "MODIFICATION",
		oldValue:    oldValue,
		newValue:    newValue,
		subject:     subject,
		subjectName: subjectName,
	}
}

func AuditDeletion(subject, subjectName string, oldValue interface{}) AuditLogConfiguration {
	return &auditLog[AuditLogAction]{
		operation:   "DELETION",
		oldValue:    oldValue,
		subject:     subject,
		subjectName: subjectName,
	}
}

func AuditCustomAction(action, subject, subjectName string, oldValue interface{}, newValue interface{}) AuditLogConfiguration {
	return &auditLog[AuditLogAction]{
		operation:   action,
		oldValue:    oldValue,
		newValue:    newValue,
		subject:     subject,
		subjectName: subjectName,
	}
}

func NewAuditLogCollector(parentSubject, parentSubjectName string) *AuditLogCollector {
	return &AuditLogCollector{
		parentAuditLog: logcomapi.CreateAuditLogRequestDTO{
			Subject:     parentSubject,
			SubjectName: &parentSubjectName,
		},
		auditLogs: make(map[string]map[string][]logcomapi.CreateAuditLogRequestDTO, 0),
	}
}

func NewAuditLogCollectorWithParent(parentAuditLog logcomapi.CreateAuditLogRequestDTO) *AuditLogCollector {
	return &AuditLogCollector{
		parentAuditLog: parentAuditLog,
		auditLogs:      make(map[string]map[string][]logcomapi.CreateAuditLogRequestDTO, 0),
	}
}

func newAuditLogCollector() *AuditLogCollector {
	return &AuditLogCollector{
		auditLogs: make(map[string]map[string][]logcomapi.CreateAuditLogRequestDTO, 0),
	}
}

func GetModelChanges(oldModel, newModel interface{}, ignoredProperties ...string) ([]ModelChange, error) {
	ignoredPropertySet := make(map[string]struct{}, len(ignoredProperties))
	for i := range ignoredProperties {
		ignoredPropertySet[ignoredProperties[i]] = struct{}{}
	}

	valueOfOldModel := reflect.ValueOf(oldModel)
	typeOfOldModel := valueOfOldModel.Type()
	kindOfOldModel := typeOfOldModel.Kind()
	valueOfNewModel := reflect.ValueOf(newModel)
	typeOfNewModel := valueOfNewModel.Type()
	var fieldCountOfOldModel int

	if kindOfOldModel == reflect.Ptr {
		fieldCountOfOldModel = valueOfOldModel.Elem().NumField()
	} else {
		fieldCountOfOldModel = valueOfOldModel.NumField()
	}

	if kindOfOldModel == reflect.Func || kindOfOldModel == reflect.Chan {
		return nil, errors.New("unsupported type: " + kindOfOldModel.String())
	}

	if typeOfOldModel != typeOfNewModel {
		return nil, errors.New("old and new model have different types")
	}

	changes := make([]ModelChange, 0)
	for i := 0; i < fieldCountOfOldModel; i++ {
		if valueOfOldModel.Field(i).CanInterface() {
			fieldNameOfOldModel := typeOfOldModel.Field(i).Name

			if _, ok := ignoredPropertySet[fieldNameOfOldModel]; ok {
				continue
			}

			fieldOfOldModel := valueOfOldModel.Field(i)
			valueOfOldModelField := fieldOfOldModel.Interface()
			valueOfNewModelField := valueOfNewModel.FieldByName(fieldNameOfOldModel).Interface()

			kindOfOldModelField := fieldOfOldModel.Kind()
			if (kindOfOldModelField == reflect.Func) || (kindOfOldModelField == reflect.Chan) {
				logWarning.Printf("Unsupported type: %s\n", kindOfOldModelField.String())
				continue
			}

			if (kindOfOldModelField >= reflect.Array) && (kindOfOldModelField != reflect.String) {
				if !reflect.DeepEqual(valueOfOldModelField, valueOfNewModelField) {
					changes = append(changes, ModelChange{
						PropertyName: fieldNameOfOldModel,
						OldValue:     valueOfOldModelField,
						NewValue:     valueOfNewModelField,
					})
				}
				continue
			}

			if valueOfOldModelField != valueOfNewModelField {
				changes = append(changes, ModelChange{
					PropertyName: fieldNameOfOldModel,
					OldValue:     valueOfOldModelField,
					NewValue:     valueOfNewModelField,
				})
			}
		}
	}

	return changes, nil
}

func prepareAuditLogRequestDTO(dto *logcomapi.CreateAuditLogRequestDTO) {
	if dto.ServiceCreated == "" {
		dto.ServiceCreated = configuration.ServiceName
	}

	if dto.ServiceAffected == nil {
		dto.ServiceAffected = &configuration.ServiceName
	}

	if dto.CreatedAt != nil {
		utcNow := dto.CreatedAt.UTC()
		dto.CreatedAt = &utcNow
	} else {
		utcNow := time.Now().UTC()
		dto.CreatedAt = &utcNow
	}
}

func (al *auditLog[T]) BatchCreate(subject string) BatchedAuditLogOperation {
	return al.BatchCustomAction("CREATION", subject)
}

func (al *auditLog[T]) BatchModify(subject string) BatchedAuditLogOperation {
	return al.BatchCustomAction("MODIFICATION", subject)
}

func (al *auditLog[T]) BatchDelete(subject string) BatchedAuditLogOperation {
	return al.BatchCustomAction("DELETION", subject)
}

func (al *auditLog[T]) BatchCustomAction(action, subject string) BatchedAuditLogOperation {
	al.operation = strings.ToUpper(strings.TrimSpace(action))
	al.subject = subject
	return al
}

func (al *auditLog[T]) GroupedModify(subject, subjectName string) GroupedModificationAuditLogOperation {
	al.batchedAuditLogMap = make(map[int]*AuditLogCollector, 0)
	al.batchedAuditLogMap[batchedSubjectModification] = NewAuditLogCollector(subject, subjectName)
	return al
}

func (al *auditLog[T]) IgnoreChangeOf(propertyNames ...string) AuditLogAction {
	al.ensureIgnoredProperties()
	al.ignoredProperties = append(al.ignoredProperties, propertyNames...)
	return al
}

func (al *auditLog[T]) AndNotify() NotificationOperation[AuditLogAction] {
	ensureNotification(&al.notification)
	al.notification.eventCategory = logcomapi.Notification
	return interface{}(al).(NotificationOperation[AuditLogAction])
}

func (al *auditLog[T]) AndLog(logLevel logcomapi.LogLevel, message string) AuditLogAction {
	ensureConsoleLog(&al.consoleLog)
	al.consoleLog.logLevel = logLevel
	al.consoleLog.message = message
	return al
}

func (al *auditLog[T]) UseService2ServiceAuthorization() AuditLogAction {
	if configuration.ClientCredentialProvider == nil {
		logFatal.Println("Client credential provider must be set")
		return al
	}

	clientCredential, err := configuration.ClientCredentialProvider.GetClientCredential()
	if err != nil {
		logError.Printf("Failed to get client credential: %v\n", err)
		return al
	}

	return al.WithBearerAuthorization(clientCredential)
}

func (al *auditLog[T]) WithBearerAuthorization(bearerToken string) AuditLogAction {
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		bearerToken = "Bearer " + bearerToken
	}
	al.ctx = context.WithValue(al.ctx, "Authorization", bearerToken)
	return al
}

func (al *auditLog[T]) WithContext(ctx context.Context) AuditLogAction {
	al.ctx = ctx
	return al
}

func (al *auditLog[T]) WithTransactionID(transactionID uuid.UUID) AuditLogConfiguration {
	al.ctx = context.WithValue(al.ctx, "RequestID", transactionID.String())
	return al
}

func (al *auditLog[T]) OnComplete(onCompleteCallback func(error)) AuditLogAction {
	al.onCompleteCallback = onCompleteCallback
	return al
}

func (al *auditLog[T]) Send() error {
	var err error

	var isBatch bool
	for _, operation := range batchedOperationList {
		if al.batchedAuditLogMap[operation] != nil && len(al.batchedAuditLogMap[operation].auditLogs) > 0 {
			err = sendAuditLogGroup(al.ctx, al, al.batchedAuditLogMap[operation])
			if err != nil {
				logError.Println("Failed to send batched audit logs")
				al.onCompleteCallback(err)
				return err
			}
			isBatch = true
		}
	}

	if !isBatch {
		switch al.operation {
		case "CREATION":
			err = sendAuditLogWithCreation(al.ctx, al.subject, al.subjectName, al.newValue)
		case "MODIFICATION":
			err = sendAuditLogWithModification(al.ctx, al.subject, al.subjectName, al.oldValue, al.newValue, al.ignoredProperties...)
		case "DELETION":
			err = sendAuditLogWithDeletion(al.ctx, al.subject, al.subjectName, al.oldValue)
		default:
			err = sendAuditLogWithCustomAction(al.ctx, al.operation, al.subject, al.subjectName, al.oldValue, al.newValue)
		}

		if err != nil {
			logError.Println("Failed to send audit log")
			al.onCompleteCallback(err)
			return err
		}
	}

	if al.notification != nil {
		if err = sendNotification(al.ctx, al.notification.eventCategory, al.notification.message, al.notification.targets); err != nil {
			logError.Printf("Failed to send notification: %v\n", err)
			// Todo: Maybe add it to a list for retry purpose
		}
	}

	if al.consoleLog != nil {
		if err = sendConsoleLog(al.ctx, al.consoleLog.logLevel, al.consoleLog.message); err != nil {
			logError.Printf("Failed to send console log: %v\n", err)
		}
	}

	al.onCompleteCallback(err)

	return nil
}

func (al *auditLog[T]) Create(subject, subjectName string, newValue interface{}) AuditLogConfiguration {
	return al.CustomAction("CREATION", subject, subjectName, nil, newValue)
}

func (al *auditLog[T]) Modify(subject, subjectName string, oldValue, newValue interface{}) AuditLogConfiguration {
	return al.CustomAction("MODIFICATION", subject, subjectName, oldValue, newValue)
}

func (al *auditLog[T]) Delete(subject, subjectName string, oldValue interface{}) AuditLogConfiguration {
	return al.CustomAction("DELETION", subject, subjectName, oldValue, nil)
}

func (al *auditLog[T]) CustomAction(action, subject, subjectName string, oldValue interface{}, newValue interface{}) AuditLogConfiguration {
	al.operation = strings.ToUpper(strings.TrimSpace(action))
	al.oldValue = oldValue
	al.newValue = newValue
	al.subject = subject
	al.subjectName = subjectName
	return al
}

func (al *auditLog[T]) CreateItem(subjectName string, newValue interface{}) BatchedAuditLogOperation {
	al.ensureBatchedAuditLogs(batchedCreation)
	al.batchedAuditLogMap[batchedCreation].AddCreation(al.subject, subjectName, newValue)
	return al
}

func (al *auditLog[T]) ModifyItem(subjectName string, oldValue, newValue interface{}) BatchedAuditLogOperation {
	al.ensureBatchedAuditLogs(batchedModification)
	al.batchedAuditLogMap[batchedModification].AddModification(al.subject, subjectName, oldValue, newValue)
	return al
}

func (al *auditLog[T]) CustomActionItem(action, subjectName string, oldValue interface{}, newValue interface{}) BatchedAuditLogOperation {
	al.ensureBatchedAuditLogs(batchedCustomAction)
	al.batchedAuditLogMap[batchedCustomAction].AddCustomAction(action, al.subject, subjectName, oldValue, newValue)
	return al
}

func (al *auditLog[T]) ModifyItemWithModelChanges(subjectName string, changes []ModelChange) BatchedAuditLogOperation {
	al.ensureBatchedAuditLogs(batchedModification)
	al.batchedAuditLogMap[batchedModification].AddModificationWithModelChanges(al.subject, subjectName, changes)
	return al
}

func (al *auditLog[T]) DeleteItem(subjectName string, oldValue interface{}) BatchedAuditLogOperation {
	al.ensureBatchedAuditLogs(batchedDeletion)
	al.batchedAuditLogMap[batchedDeletion].AddDeletion(al.subject, subjectName, oldValue)
	return al
}

func (al *auditLog[T]) AddCreation(subject, subjectName string, newValue interface{}) GroupedModificationAuditLogOperation {
	al.batchedAuditLogMap[batchedCreation].AddCreation(subject, subjectName, newValue)
	return al
}

func (al *auditLog[T]) AddModification(subject, subjectName string, oldValue, newValue interface{}) GroupedModificationAuditLogOperation {
	al.batchedAuditLogMap[batchedModification].AddModification(subject, subjectName, oldValue, newValue)
	return al
}

func (al *auditLog[T]) AddModificationWithModelChanges(subject, subjectName string, changes []ModelChange) GroupedModificationAuditLogOperation {
	al.batchedAuditLogMap[batchedModification].AddModificationWithModelChanges(subject, subjectName, changes)
	return al
}

func (al *auditLog[T]) AddDeletion(subject, subjectName string, oldValue interface{}) GroupedModificationAuditLogOperation {
	al.batchedAuditLogMap[batchedDeletion].AddDeletion(subject, subjectName, oldValue)
	return al
}

func (al *auditLog[T]) AddCustomEvent(action, subject, subjectName string, oldValue interface{}, newValue interface{}) GroupedModificationAuditLogOperation {
	al.batchedAuditLogMap[batchedCustomAction].AddCustomAction(action, subject, subjectName, oldValue, newValue)
	return al
}

func (al *auditLog[T]) Roles(targets ...string) NotificationMessage[T] {
	al.notification.Roles(targets...)
	return al
}

func (al *auditLog[T]) Sessions(targets ...string) NotificationMessage[T] {
	al.notification.Sessions(targets...)
	return al
}

func (al *auditLog[T]) Users(targets ...string) NotificationMessage[T] {
	al.notification.Users(targets...)
	return al
}

func (al *auditLog[T]) Message(message string) T {
	al.notification.message = message
	return interface{}(al).(T)
}

func (al *auditLog[T]) ensureBatchedAuditLogs(batchType int) {
	if _, ok := al.batchedAuditLogMap[batchType]; !ok {
		al.batchedAuditLogMap = make(map[int]*AuditLogCollector, 0)
		al.batchedAuditLogMap[batchType] = newAuditLogCollector()
	}
}

func (al *auditLog[T]) ensureIgnoredProperties() {
	if al.ignoredProperties == nil {
		al.ignoredProperties = make([]string, 0)
	}
}

func (c *AuditLogCollector) AddCreation(itemSubject, itemSubjectName string, newValue interface{}) {
	c.Add(logcomapi.CreateAuditLogRequestDTO{
		Category:    "CREATION",
		NewValue:    stringify(newValue),
		Subject:     itemSubject,
		SubjectName: &itemSubjectName,
	})
}

func (c *AuditLogCollector) AddModification(itemSubject, itemSubjectName string, oldValue, newValue interface{}) {
	c.Add(logcomapi.CreateAuditLogRequestDTO{
		Category:    "MODIFICATION",
		NewValue:    stringify(newValue),
		OldValue:    stringify(oldValue),
		Subject:     itemSubject,
		SubjectName: &itemSubjectName,
	})
}

func (c *AuditLogCollector) AddCustomAction(action string, itemSubject, itemSubjectName string, oldValue, newValue interface{}) {
	c.Add(logcomapi.CreateAuditLogRequestDTO{
		Category:    strings.ToUpper(strings.TrimSpace(action)),
		NewValue:    stringify(newValue),
		OldValue:    stringify(oldValue),
		Subject:     itemSubject,
		SubjectName: &itemSubjectName,
	})
}

func (c *AuditLogCollector) AddModificationWithModelChanges(itemSubject, itemSubjectName string, changes []ModelChange) {
	dto := logcomapi.CreateAuditLogRequestDTO{
		Category:    "MODIFICATION",
		Subject:     itemSubject,
		SubjectName: &itemSubjectName,
	}

	transformModelChangesToDTO(&dto, changes)

	c.Add(dto)
}

func (c *AuditLogCollector) AddDeletion(itemSubject, itemSubjectName string, oldValue interface{}) {
	if oldValue == nil {
		oldValue = ""
	}
	c.Add(logcomapi.CreateAuditLogRequestDTO{
		Category:    "DELETION",
		OldValue:    stringify(oldValue),
		Subject:     itemSubject,
		SubjectName: &itemSubjectName,
	})
}

func (c *AuditLogCollector) Add(model logcomapi.CreateAuditLogRequestDTO) {
	if subjectNameMap, ok := c.auditLogs[model.Subject]; ok {
		if auditLogList, ok := subjectNameMap[*model.SubjectName]; ok {
			c.auditLogs[model.Subject][*model.SubjectName] = append(auditLogList, model)
		} else {
			c.auditLogs[model.Subject][*model.SubjectName] = []logcomapi.CreateAuditLogRequestDTO{model}
		}
	} else {
		c.auditLogs[model.Subject] = map[string][]logcomapi.CreateAuditLogRequestDTO{*model.SubjectName: {model}}
	}
}

func (c *AuditLogCollector) get() logcomapi.CreateAuditLogRequestDTO {
	c.parentAuditLog.GroupedChanges = make([]logcomapi.CreateAuditLogChangeDTO, 0)
	hasSameCategory := true
	seenCategory := ""

	for subjectAsKey, subjectNameGroupAsValue := range c.auditLogs {
		for subjectNameAsKey, auditLogList := range subjectNameGroupAsValue {
			for _, auditLog := range auditLogList {
				if len(auditLog.GroupedChanges) > 0 {
					for _, nestedAuditLogChange := range auditLog.GroupedChanges {
						c.parentAuditLog.GroupedChanges = append(c.parentAuditLog.GroupedChanges, nestedAuditLogChange)
					}
				} else {
					c.parentAuditLog.GroupedChanges = append(c.parentAuditLog.GroupedChanges, logcomapi.CreateAuditLogChangeDTO{
						Category:            &auditLog.Category,
						Message:             auditLog.Message,
						NewValue:            auditLog.NewValue,
						OldValue:            auditLog.OldValue,
						Subject:             &subjectAsKey,
						SubjectName:         &subjectNameAsKey,
						SubjectPropertyName: auditLog.SubjectPropertyName,
					})
				}

				if hasSameCategory && seenCategory != auditLog.Category && seenCategory != "" {
					hasSameCategory = false
				}

				seenCategory = auditLog.Category
			}
		}
	}

	if c.parentAuditLog.Category == "" {
		if hasSameCategory {
			c.parentAuditLog.Category = seenCategory
		} else {
			c.parentAuditLog.Category = "MODIFICATION"
		}
	}

	return c.parentAuditLog
}
