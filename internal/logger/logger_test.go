package logger

import (
	"bitoOA/internal/config"
	"context"
	"testing"
)

func TestLogger(t *testing.T) {
	cfg := config.Config{
		Log: struct{ Level string }{Level: "debug"},
	}
	log := New(&cfg)
	log.Debug("debug message")

	ctx := context.WithValue(context.Background(), "traceId", "testTraceId")
	log = log.WithContext(ctx)
	log.Debug("debug with ctx")
}
