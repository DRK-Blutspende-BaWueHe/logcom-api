package logcomapi

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/DRK-Blutspende-BaWueHe/zerolog-for-logcom/log"
)

type ModelChange struct {
	PropertyName string
	OldValue     interface{}
	NewValue     interface{}
}

type AuditLogModelDiff interface {
	GetChanges(model interface{}, ignoredProperties map[string]interface{}) []ModelChange
}

type AuditLogCollector struct {
	parentAuditLog CreateAuditLogRequestDto
	auditLogs      map[string]map[string][]CreateAuditLogRequestDto
}

func NewAuditLogCollector(subject, subjectName string) *AuditLogCollector {
	return &AuditLogCollector{
		parentAuditLog: CreateAuditLogRequestDto{Subject: subject},
		auditLogs:      make(map[string]map[string][]CreateAuditLogRequestDto, 0),
	}
}

func NewAuditLogCollectorWithParent(parentAuditLog CreateAuditLogRequestDto) *AuditLogCollector {
	return &AuditLogCollector{
		parentAuditLog: parentAuditLog,
		auditLogs:      make(map[string]map[string][]CreateAuditLogRequestDto, 0),
	}
}

func (c *AuditLogCollector) AddCreation(subject string, newValue interface{}) {
	c.AddCreationForSubject(subject, "", newValue)
}

func (c *AuditLogCollector) AddCreationForSubject(subject, subjectName string, newValue interface{}) {
	c.Add(CreateAuditLogRequestDto{
		Category:    "CREATION",
		NewValue:    stringify(newValue),
		Subject:     subject,
		SubjectName: subjectName,
	})
}

func (c *AuditLogCollector) Add(model CreateAuditLogRequestDto) {
	if subjectNameMap, ok := c.auditLogs[model.Subject]; ok {
		if auditLogList, ok := subjectNameMap[model.SubjectName]; ok {
			c.auditLogs[model.Subject][model.SubjectName] = append(auditLogList, model)
		} else {
			c.auditLogs[model.Subject][model.SubjectName] = []CreateAuditLogRequestDto{model}
		}
	} else {
		c.auditLogs[model.Subject] = map[string][]CreateAuditLogRequestDto{model.SubjectName: {model}}
	}
}

func GetModelChanges(ctx context.Context, oldModel, newModel interface{}) ([]ModelChange, error) {
	valueOfOldModel := reflect.ValueOf(oldModel)
	typeOfOldModel := valueOfOldModel.Type()
	fieldCountOfOldModel := valueOfOldModel.NumField()
	valueOfNewModel := reflect.ValueOf(newModel)
	typeOfNewModel := valueOfNewModel.Type()

	if typeOfOldModel != typeOfNewModel {
		log.Fatal().MsgContext(ctx, "Old and new model have different types")
		return nil, errors.New("old and new model have different types")
	}

	changes := make([]ModelChange, 0)
	for i := 0; i < fieldCountOfOldModel; i++ {
		if valueOfOldModel.Field(i).CanInterface() {
			fieldNameOfOldModel := typeOfOldModel.Field(i).Name

			realValueOfOldModel := valueOfOldModel.Field(i).Interface()
			realValueOfNewModel := valueOfNewModel.FieldByName(fieldNameOfOldModel).Interface()

			if realValueOfOldModel != realValueOfNewModel {
				changes = append(changes, ModelChange{
					PropertyName: fieldNameOfOldModel,
					OldValue:     realValueOfOldModel,
					NewValue:     realValueOfNewModel,
				})
			}
		}
	}

	return changes, nil
}

func (c *AuditLogCollector) get() CreateAuditLogRequestDto {
	c.parentAuditLog.GroupedChanges = make([]NewAuditLogChangeDto, 0)
	hasSameCategory := true
	seenCategory := ""

	for subjectAsKey, subjectNameGroupAsValue := range c.auditLogs {
		for subjectNameAsKey, auditLogList := range subjectNameGroupAsValue {
			for _, auditLog := range auditLogList {
				c.parentAuditLog.GroupedChanges = append(c.parentAuditLog.GroupedChanges, NewAuditLogChangeDto{
					Category:            auditLog.Category,
					Message:             auditLog.Message,
					NewValue:            auditLog.NewValue,
					OldValue:            auditLog.OldValue,
					Subject:             subjectAsKey,
					SubjectName:         subjectNameAsKey,
					SubjectPropertyName: auditLog.SubjectPropertyName,
				})

				if hasSameCategory && seenCategory != auditLog.Category && seenCategory != "" {
					hasSameCategory = false
				}

				seenCategory = auditLog.Category
			}
		}
	}

	if hasSameCategory {
		c.parentAuditLog.Category = seenCategory
	} else {
		c.parentAuditLog.Category = "MODIFICATION"
	}

	return c.parentAuditLog
}

func prepareAuditLogRequestDTO(dto *CreateAuditLogRequestDto) {
	if dto.ServiceCreated == "" {
		dto.ServiceCreated = configuration.ServiceName
	}

	if dto.ServiceAffected == "" {
		dto.ServiceAffected = configuration.ServiceName
	}

	if dto.CreatedAt != nil {
		utcNow := dto.CreatedAt.UTC()
		dto.CreatedAt = &utcNow
	} else {
		utcNow := time.Now().UTC()
		dto.CreatedAt = &utcNow
	}
}
