package log// Factory is the default logging wrapper that can create
import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logger instances either for a given Context or context-less.
type Factory struct {
	logger *zap.Logger
}

// NewFactory creates a new Factory.
func NewFactory(logger *zap.Logger) Factory {
	return Factory{logger: logger}
}

// GetLogger get logger
func (b Factory) GetLogger() *zap.Logger {
	return b.logger
}

// Bg creates a context-unaware logger.
func (b Factory) Bg() *zap.Logger {
	return b.logger
}

// For returns a context-aware Logger. If the context
// contains an OpenTracing span, all logging calls are also
// echo-ed into the span.
func (b Factory) For(ctx context.Context) *zap.Logger {
	rID := ctx.Value("_requestID")
	if rID != nil {
		requestID, ok := rID.(string)
		if ok {
			return b.logger.With(zap.String("request_id", requestID))
		}
	}
	return b.Bg()
}

// With creates a child logger, and optionally adds some context fields to that logger.
func (b Factory) With(fields ...zapcore.Field) Factory {
	return Factory{logger: b.logger.With(fields...)}
}
