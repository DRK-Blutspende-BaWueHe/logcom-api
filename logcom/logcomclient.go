package logcom

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/DRK-Blutspende-BaWueHe/logcom-api"
	"github.com/google/uuid"
)

var (
	configuration     Configuration
	apiClientInstance *logcomapi.APIClient
	once              sync.Once

	logInfo    *log.Logger
	logWarning *log.Logger
	logError   *log.Logger
	logFatal   *log.Logger
)

type ClientCredentialProvider interface {
	GetClientCredential() (string, error)
}

type HeaderProviderFunc func(ctx context.Context) http.Header

type Configuration struct {
	ServiceName              string
	LogComURL                string
	HeaderProvider           HeaderProviderFunc
	ClientCredentialProvider ClientCredentialProvider
}

func init() {
	logFormat := &logFormat{
		message: "timestamp level file [LogCom] > message",
		time:    "2006-01-02T15:04:05Z07:00",
	}

	infoHandle := logWriter{
		writer:    os.Stdout,
		logFormat: logFormat,
		level:     "INFO",
	}
	warningHandle := logWriter{
		writer:    os.Stdout,
		logFormat: logFormat,
		level:     "WARN",
	}
	errorHandle := logWriter{
		writer:    os.Stderr,
		logFormat: logFormat,
		level:     "ERR",
	}
	fatalHandle := logWriter{
		writer:    os.Stderr,
		logFormat: logFormat,
		level:     "FTL",
	}

	logInfo = log.New(infoHandle, "", 0)
	logWarning = log.New(warningHandle, "", 0)
	logError = log.New(errorHandle, "", 0)
	logFatal = log.New(fatalHandle, "", 0)

	logcomURL := os.Getenv("LOG_COM_URL")

	config := Configuration{
		ServiceName: "Unknown",
		LogComURL:   logcomURL,
		HeaderProvider: func(ctx context.Context) http.Header {
			return http.Header{}
		},
	}

	configuration = config

	logcomAPIConfig := logcomapi.NewConfiguration()
	logcomAPIConfig.Servers = logcomapi.ServerConfigurations{
		{
			URL:         configuration.LogComURL + logcomAPIConfig.Servers[0].URL,
			Description: logcomAPIConfig.Servers[0].Description,
			Variables:   logcomAPIConfig.Servers[0].Variables,
		},
	}

	parsedUrl, err := url.Parse(configuration.LogComURL)
	if err != nil {
		logError.Printf("Failed to get LogCom URL scheme, falling back to default (%s): %v\n", "https", err)

		parsedUrl = &url.URL{
			Scheme: "https",
		}
	}

	logcomAPIConfig.Scheme = parsedUrl.Scheme

	if logcomURL == "" {
		logError.Println("LogCom URL is missing thus functionalities are not available")
		return
	}

	apiClientInstance = logcomapi.NewAPIClient(logcomAPIConfig)

	apiClientInstance.GetConfig().HTTPClient.Transport = &headerMiddleware{
		originalRoundTripper: apiClientInstance.GetConfig().HTTPClient.Transport,
	}
}

func Init(config Configuration) {
	if config.LogComURL == "" {
		config.LogComURL = os.Getenv("LOG_COM_URL")
		if config.LogComURL == "" {
			logFatal.Println("LogCom URL is missing")
			return
		}
	}

	once.Do(func() {
		configuration.LogComURL = config.LogComURL

		if config.ServiceName != "" {
			configuration.ServiceName = config.ServiceName
		}

		if config.HeaderProvider != nil {
			configuration.HeaderProvider = config.HeaderProvider
		}

		if config.ClientCredentialProvider != nil {
			configuration.ClientCredentialProvider = config.ClientCredentialProvider
		}

		logcomAPIConfig := logcomapi.NewConfiguration()
		logcomAPIConfig.Servers = logcomapi.ServerConfigurations{
			{
				URL:         configuration.LogComURL + logcomAPIConfig.Servers[0].URL,
				Description: logcomAPIConfig.Servers[0].Description,
				Variables:   logcomAPIConfig.Servers[0].Variables,
			},
		}

		parsedUrl, err := url.Parse(configuration.LogComURL)
		if err != nil {
			logError.Printf("Failed to get LogCom URL scheme, falling back to default (%s): %v\n", "https", err)

			parsedUrl = &url.URL{
				Scheme: "https",
			}
		}

		logcomAPIConfig.Scheme = parsedUrl.Scheme

		apiClientInstance = logcomapi.NewAPIClient(logcomAPIConfig)

		apiClientInstance.GetConfig().HTTPClient.Transport = &headerMiddleware{
			originalRoundTripper: apiClientInstance.GetConfig().HTTPClient.Transport,
		}
	})
}

func SetClientCredentialProvider(provider ClientCredentialProvider) {
	configuration.ClientCredentialProvider = provider
}

func IsEnabled() bool {
	return apiClientInstance != nil
}

func SendConsoleLog(ctx context.Context, logLevel logcomapi.LogLevel, message string) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendConsoleLog(ctx, logLevel, message)
}

func SendConsoleLogWithModel(ctx context.Context, model logcomapi.CreateConsoleLogRequestDTO) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendConsoleLogWithModel(ctx, model)
}

func SendAuditLogWithCreation(ctx context.Context, subject, subjectName string, newValue interface{}) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendAuditLogWithCreation(ctx, subject, subjectName, newValue)
}

func SendAuditLogWithModification(ctx context.Context, subject, subjectName string, oldValue, newValue interface{}, ignoredProperties ...string) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendAuditLogWithModification(ctx, subject, subjectName, oldValue, newValue, ignoredProperties...)
}

func SendAuditLogWithCustomAction(ctx context.Context, action, subject, subjectName string, oldValue, newValue interface{}, ignoredProperties ...string) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendAuditLogWithCustomAction(ctx, action, subject, subjectName, oldValue, newValue, ignoredProperties...)
}

func SendAuditLogWithModificationModelChanges(ctx context.Context, subject, subjectName string, changes []ModelChange) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendAuditLogWithModificationModelChanges(ctx, subject, subjectName, changes)
}

func SendAuditLogWithDeletion(ctx context.Context, subject, subjectName string, oldValue interface{}) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendAuditLogWithDeletion(ctx, subject, subjectName, oldValue)
}

func SendAuditLog(ctx context.Context, model logcomapi.CreateAuditLogRequestDTO) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendAuditLog(ctx, model)
}

func SendAuditLogGroup(ctx context.Context, auditLogCollector *AuditLogCollector) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendAuditLog(ctx, auditLogCollector.get())
}

func SendNotification(ctx context.Context, eventCategory logcomapi.NotificationEventCategory, message string, targets map[string][]string) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendNotification(ctx, eventCategory, message, targets)
}

func SendNotificationWithModel(ctx context.Context, model logcomapi.CreateNotificationRequestDTO) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendNotificationWithModel(ctx, model)
}

func sendConsoleLog(ctx context.Context, logLevel logcomapi.LogLevel, message string) error {
	return sendConsoleLogWithModel(ctx, logcomapi.CreateConsoleLogRequestDTO{
		Level:   logLevel,
		Message: message,
		Service: configuration.ServiceName,
	})
}

func sendConsoleLogWithModel(ctx context.Context, model logcomapi.CreateConsoleLogRequestDTO) error {
	if !IsEnabled() {
		logInfo.Println("LogCom is disabled")
		return nil
	}

	prepareConsoleLogRequestDTO(&model)

	result, err := apiClientInstance.ConsoleLogApi.CreateConsoleLogV1Int(ctx).Model(model).Execute()
	if err != nil {
		return err
	}

	if !isHTTPStatusSuccess(result.StatusCode) {
		return errors.New(result.Status)
	}

	return nil
}

func sendAuditLogWithCreation(ctx context.Context, subject, subjectName string, newValue interface{}) error {
	if newValue == nil {
		newValue = ""
	}
	return sendAuditLog(ctx, logcomapi.CreateAuditLogRequestDTO{
		Category:    "CREATION",
		NewValue:    stringify(newValue),
		Subject:     subject,
		SubjectName: &subjectName,
	})
}

func sendAuditLogWithModification(ctx context.Context, subject, subjectName string, oldValue, newValue interface{}, ignoredProperties ...string) error {
	if oldValue == nil || newValue == nil {
		return errors.New("oldValue and newValue are required")
	}

	if isPrimitiveType(oldValue) {
		return SendAuditLog(ctx, logcomapi.CreateAuditLogRequestDTO{
			Category:    "MODIFICATION",
			NewValue:    stringify(newValue),
			OldValue:    stringify(oldValue),
			Subject:     subject,
			SubjectName: &subjectName,
		})
	}

	changes, err := GetModelChanges(oldValue, newValue, ignoredProperties...)
	if err != nil {
		err = sendConsoleLog(ctx, logcomapi.Error, "Failed to send audit log: "+err.Error())
		if err != nil {
			logError.Printf("Failed to send console log: %v\n", err)
		}
		return err
	}

	if len(changes) < 1 {
		err = sendConsoleLog(ctx, logcomapi.Debug, "No changes")
		if err != nil {
			logError.Printf("Failed to send console log: %v\n", err)
		}
		return nil
	}

	return sendAuditLogWithModificationModelChanges(ctx, subject, subjectName, changes)
}

func sendAuditLogWithModificationModelChanges(ctx context.Context, subject, subjectName string, changes []ModelChange) error {
	dto := logcomapi.CreateAuditLogRequestDTO{
		Category:    "MODIFICATION",
		Subject:     subject,
		SubjectName: &subjectName,
	}

	transformModelChangesToDTO(&dto, changes)

	return sendAuditLog(ctx, dto)
}

func sendAuditLogWithDeletion(ctx context.Context, subject, subjectName string, oldValue interface{}) error {
	if oldValue == nil {
		oldValue = ""
	}
	return sendAuditLog(ctx, logcomapi.CreateAuditLogRequestDTO{
		Category:    "DELETION",
		OldValue:    stringify(oldValue),
		Subject:     subject,
		SubjectName: &subjectName,
	})
}

func sendAuditLogWithCustomAction(ctx context.Context, action, subject, subjectName string, oldValue, newValue interface{}, ignoredProperties ...string) error {
	if newValue == nil {
		newValue = ""
	}

	if oldValue == nil {
		oldValue = ""
	}

	if isPrimitiveType(oldValue) {
		return SendAuditLog(ctx, logcomapi.CreateAuditLogRequestDTO{
			Category:    action,
			NewValue:    stringify(newValue),
			OldValue:    stringify(oldValue),
			Subject:     subject,
			SubjectName: &subjectName,
		})
	}

	changes, err := GetModelChanges(oldValue, newValue, ignoredProperties...)
	if err != nil {
		err = sendConsoleLog(ctx, logcomapi.Error, "Failed to send audit log: "+err.Error())
		if err != nil {
			logError.Printf("Failed to send console log: %v\n", err)
		}
		return err
	}

	if len(changes) < 1 {
		err = sendConsoleLog(ctx, logcomapi.Debug, "No changes")
		if err != nil {
			logError.Printf("Failed to send console log: %v\n", err)
		}
		return nil
	}

	return sendAuditLogWithCustomActionModelChanges(ctx, action, subject, subjectName, changes)
}

func sendAuditLogWithCustomActionModelChanges(ctx context.Context, action, subject, subjectName string, changes []ModelChange) error {
	dto := logcomapi.CreateAuditLogRequestDTO{
		Category:    action,
		Subject:     subject,
		SubjectName: &subjectName,
	}

	transformModelChangesToDTO(&dto, changes)

	return sendAuditLog(ctx, dto)
}

func sendAuditLog(ctx context.Context, model logcomapi.CreateAuditLogRequestDTO) error {
	if !IsEnabled() {
		logInfo.Println("LogCom is disabled")
		return nil
	}

	prepareAuditLogRequestDTO(&model)

	result, err := apiClientInstance.AuditLogApi.CreateAuditLogV1Int(ctx).Model(model).Execute()
	if err != nil {
		return err
	}

	if !isHTTPStatusSuccess(result.StatusCode) {
		return errors.New(result.Status)
	}

	return nil
}

func sendAuditLogGroup[T any](ctx context.Context, auditLog *auditLog[T], auditLogCollector *AuditLogCollector) error {
	auditLogCollector.parentAuditLog = logcomapi.CreateAuditLogRequestDTO{
		Category:    auditLog.operation,
		Subject:     auditLog.subject,
		SubjectName: &auditLog.subjectName,
	}
	return sendAuditLog(ctx, auditLogCollector.get())
}

func sendNotification(ctx context.Context, eventCategory logcomapi.NotificationEventCategory, message string, targets map[string][]string) error {
	return sendNotificationWithModel(ctx, logcomapi.CreateNotificationRequestDTO{
		EventCategory: &eventCategory,
		Message:       message,
		Service:       configuration.ServiceName,
		Targets:       &targets,
	})
}

func sendNotificationWithModel(ctx context.Context, model logcomapi.CreateNotificationRequestDTO) error {
	if !IsEnabled() {
		logInfo.Println("LogCom is disabled")
		return nil
	}

	prepareNotificationRequestDTO(&model)

	result, err := apiClientInstance.NotificationApi.CreateNotificationV1(ctx).Model(model).Execute()
	if err != nil {
		return err
	}

	if !isHTTPStatusSuccess(result.StatusCode) {
		return errors.New(result.Status)
	}

	return nil
}

func isHTTPStatusSuccess(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices
}

func transformModelChangesToDTO(targetDTO *logcomapi.CreateAuditLogRequestDTO, changes []ModelChange) {
	if changesCount := len(changes); changesCount > 1 {
		changesDTO := make([]logcomapi.CreateAuditLogChangeDTO, changesCount)

		for i, change := range changes {
			changesDTO[i] = logcomapi.CreateAuditLogChangeDTO{
				Category:            &targetDTO.Category,
				NewValue:            stringify(change.NewValue),
				OldValue:            stringify(change.OldValue),
				Subject:             &targetDTO.Subject,
				SubjectName:         targetDTO.SubjectName,
				SubjectPropertyName: &change.PropertyName,
			}
		}

		targetDTO.GroupedChanges = changesDTO
	} else if changesCount > 0 {
		targetDTO.NewValue = stringify(changes[0].NewValue)
		targetDTO.OldValue = stringify(changes[0].OldValue)
		targetDTO.SubjectPropertyName = &changes[0].PropertyName
	}
}

func stringify(value interface{}) *string {
	if value == nil {
		return nil
	}

	if isPrimitiveType(value) {
		return logcomapi.PtrString(fmt.Sprintf("%v", value))
	}

	jsonData, err := json.Marshal(value)
	if err != nil {
		return logcomapi.PtrString(fmt.Sprintf("%+v", value))
	}

	return logcomapi.PtrString(string(jsonData))
}

func isPrimitiveType(value interface{}) bool {
	valueKind := reflect.TypeOf(value).Kind()
	return ((valueKind < reflect.Array) && (valueKind > 0) && (valueKind != reflect.Uintptr)) || (valueKind == reflect.String)
}

func ensureConsoleLog(consoleLogPointer **consoleLog) {
	if *consoleLogPointer == nil {
		*consoleLogPointer = &consoleLog{}
	}
}

func ensureNotification[T any](notificationPointer **notification[T]) {
	if *notificationPointer == nil {
		*notificationPointer = &notification[T]{}
	}
}

type logFormat struct {
	message string
	time    string
}

type logWriter struct {
	writer    io.Writer
	level     string
	logFormat *logFormat
}

func (w logWriter) Write(bytes []byte) (int, error) {
	messageFormat := w.logFormat.message
	formatKeys := strings.Split(messageFormat, " ")
	orderedValues := make([]interface{}, 0)
	for _, formatKey := range formatKeys {
		var value interface{}

		if strings.Contains(formatKey, "timestamp") {
			messageFormat = strings.Replace(messageFormat, "timestamp", "%s", 1)
			value = time.Now().UTC().Format(w.logFormat.time)
		} else if strings.Contains(formatKey, "level") {
			messageFormat = strings.Replace(messageFormat, "level", "%s", 1)
			value = w.level
		} else if strings.Contains(formatKey, "message") {
			messageFormat = strings.Replace(messageFormat, "message", "%s", 1)
			value = string(bytes)
		} else if strings.Contains(formatKey, "file") {
			_, filename, line, _ := runtime.Caller(3)
			messageFormat = strings.Replace(messageFormat, "file", "%s", 1)
			value = filename + ":" + strconv.Itoa(line)
		} else {
			continue
		}

		orderedValues = append(orderedValues, value)
	}

	return fmt.Fprintf(w.writer, messageFormat, orderedValues...)
}

type headerMiddleware struct {
	originalRoundTripper http.RoundTripper
}

func (t *headerMiddleware) RoundTrip(r *http.Request) (*http.Response, error) {
	requestCopy := r.Clone(r.Context())

	for headerKey, headerValue := range configuration.HeaderProvider(requestCopy.Context()) {
		for _, header := range headerValue {
			requestCopy.Header.Add(headerKey, header)
		}
	}

	authToken := requestCopy.Context().Value("Authorization")
	if authToken == nil && requestCopy.Header.Get("Authorization") == "" {
		return nil, errors.New("authorization header cannot be extracted from context")
	}

	requestID := requestCopy.Context().Value("RequestID")
	if requestID == nil && requestCopy.Header.Get("X-Request-ID") == "" {
		requestID = uuid.NewString()
	}

	if authToken != nil {
		requestCopy.Header.Set("Authorization", authToken.(string))
	}

	if requestID != nil {
		requestCopy.Header.Set("X-Request-ID", requestID.(string))
	}

	if t.originalRoundTripper == nil {
		return http.DefaultTransport.RoundTrip(requestCopy)
	}

	return t.originalRoundTripper.RoundTrip(requestCopy)
}
