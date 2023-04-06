package logcom

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	logcomapi "github.com/DRK-Blutspende-BaWueHe/logcom-api"
	"github.com/google/uuid"
)

func TestSendAndNotifyAndLog(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	}))
	defer svr.Close()

	config := Configuration{
		ServiceName: "Unknown",
		LogComURL:   svr.URL,
	}

	type testStruct struct {
		field string
	}

	Init(config)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "Authorization", "BearerToken")
	ctx = context.WithValue(ctx, "RequestID", uuid.NewString())

	t.Run("CREATE OK", func(t *testing.T) {
		err := Audit().
			Create("SUBJECT", "NAME", nil).
			WithContext(ctx).
			//WithTransactionID(uuid.New()).
			//WithBearerAuthorization("BearerToken").
			AndNotify().
			Roles("test_role").
			Message("Test notification").
			AndLog(logcomapi.Debug, "Debug log").
			Send()
		if err != nil {
			t.Errorf("expected no error")
		}
	})

	t.Run("MODIFY OK", func(t *testing.T) {
		err := Audit().
			Modify("SUBJECT", "NAME", testStruct{field: "test value of field"}, testStruct{field: "test new value of field"}).
			WithContext(ctx).
			//WithTransactionID(uuid.New()).
			//WithBearerAuthorization("BearerToken").
			AndNotify().
			Roles("test_role").
			Message("Test notification").
			AndLog(logcomapi.Debug, "Debug log").
			Send()
		if err != nil {
			t.Errorf("expected no error")
		}
	})

	t.Run("DELETE OK", func(t *testing.T) {
		err := Audit().
			Delete("SUBJECT", "NAME", testStruct{field: "test value of field"}).
			WithContext(ctx).
			//WithTransactionID(uuid.New()).
			//WithBearerAuthorization("BearerToken").
			AndNotify().
			Roles("test_role").
			Message("Test notification").
			AndLog(logcomapi.Debug, "Debug log").
			Send()
		if err != nil {
			t.Errorf("expected no error")
		}
	})

	t.Run("Custom CUSTOM_ACTION OK", func(t *testing.T) {
		err := Audit().
			CustomAction("CUSTOM_ACTION", "NAME", "SUBJECT", testStruct{field: "test value of field"}, testStruct{field: "test new value of field"}).
			WithContext(ctx).
			//WithTransactionID(uuid.New()).
			//WithBearerAuthorization("BearerToken").
			AndNotify().
			Roles("test_role").
			Message("Test notification").
			AndLog(logcomapi.Debug, "Debug log").
			Send()
		if err != nil {
			t.Errorf(fmt.Errorf("expected no error, got: %w", err).Error())
		}
	})

	t.Run("Custom DIFFERENT_ACTION OK", func(t *testing.T) {
		err := Audit().
			CustomAction("DIFFERENT_ACTION", "NAME", "SUBJECT", testStruct{field: "test value of field"}, testStruct{field: "test new value of field"}).
			WithContext(ctx).
			//WithTransactionID(uuid.New()).
			//WithBearerAuthorization("BearerToken").
			AndNotify().
			Roles("test_role").
			Message("Test notification").
			AndLog(logcomapi.Debug, "Debug log").
			Send()
		if err != nil {
			t.Errorf(fmt.Errorf("expected no error, got: %w", err).Error())
		}
	})

}
