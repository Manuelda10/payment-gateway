package zaplogger

import (
	"context"
	"payment-gateway/internal/ports/output"
	"payment-gateway/internal/shared/requestctx"

	"go.uber.org/zap"
)

type Logger struct {
	base *zap.Logger
}

func New(base *zap.Logger) *Logger {
	return &Logger{base: base}
}

func (l *Logger) Debug(ctx context.Context, msg string, fields ...output.Field) {
	l.get(ctx).Debug(msg, toZapFields(fields)...)
}
func (l *Logger) Info(ctx context.Context, msg string, fields ...output.Field) {
	l.get(ctx).Info(msg, toZapFields(fields)...)
}
func (l *Logger) Warn(ctx context.Context, msg string, fields ...output.Field) {
	l.get(ctx).Warn(msg, toZapFields(fields)...)
}
func (l *Logger) Error(ctx context.Context, msg string, err error, fields ...output.Field) {
	zf := toZapFields(fields)
	if err != nil {
		zf = append(zf, zap.Error(err))
	}
	l.get(ctx).Error(msg, zf...)
}

func (l *Logger) get(ctx context.Context) *zap.Logger {
	if rid, ok := requestctx.RequestID(ctx); ok && rid != "" {
		return l.base.With(zap.String("request_id", rid))
	}
	return l.base
}

func toZapFields(fields []output.Field) []zap.Field {
	out := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		out = append(out, zap.Any(f.Key, f.Value))
	}
	return out
}
