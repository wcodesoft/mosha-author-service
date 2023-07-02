package logger

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l *log.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		switch lvl {
		case logging.LevelDebug:
			l.Debugf(msg, fields)
		case logging.LevelInfo:
			l.Infof(msg, fields)
		case logging.LevelWarn:
			l.Warnf(msg, fields)
		case logging.LevelError:
			l.Errorf(msg, fields)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}
