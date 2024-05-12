package logger

import (
	"bitoOA/internal/config"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

type Logger struct {
	core *zap.SugaredLogger
	ctx  context.Context
}

func (l *Logger) With(key string, value interface{}) *Logger {
	return &Logger{
		core: l.core.With(key, value),
	}
}

func (l *Logger) WithError(err error) *Logger {
	return &Logger{
		core: l.core.With("error", err.Error()),
	}
}

func (l *Logger) WithService(service string) *Logger {
	return &Logger{
		core: l.core.With("service", service),
	}
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	if ctx == nil {
		ctx = context.WithValue(context.Background(), "traceId", uuid.NewString())
	} else if tId := ctx.Value("traceId"); tId == nil || tId == "" {
		tId = uuid.NewString()
		ctx = context.WithValue(ctx, "traceId", tId)
	}

	return &Logger{
		core: l.core.With("traceId", ctx.Value("traceId")),
		ctx:  ctx,
	}
}

func (l *Logger) Debug(msg string) {
	l.core.Debug(msg)
}

func (l *Logger) Info(msg string) {
	l.core.Info(msg)
}

func (l *Logger) Warn(msg string) {
	l.core.Warn(msg)
}

func (l *Logger) Error(msg string) {
	l.core.Error(msg)
}

func (l *Logger) Panic(msg string) {
	l.core.Panic(msg)
}

func New(config *config.Config) *Logger {
	zCfg := zap.NewProductionConfig()
	zCfg.EncoderConfig.LevelKey = "logLevel"
	zCfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	lvlM := map[string]zapcore.Level{
		"debug": zap.DebugLevel,
		"info":  zap.InfoLevel,
		"warn":  zap.WarnLevel,
		"error": zap.ErrorLevel,
	}

	// set log level
	if lvl, exist := lvlM[config.Log.Level]; exist {
		zCfg.Level = zap.NewAtomicLevelAt(lvl)
	} else {
		zCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		log.Println(fmt.Sprintf("invalid log level: %s use debug level as default", config.Log.Level))
	}

	logger, _ := zCfg.Build(zap.AddCallerSkip(1))
	defer logger.Sync() // flushes buffer, if any

	lgr := &Logger{
		core: logger.Sugar(),
	}

	lgr.Info("logger initialized")

	return lgr
}
