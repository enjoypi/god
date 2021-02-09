package god

import "go.uber.org/zap"

var (
	Logger *zap.Logger
)

func init() {
	Logger, _ = zap.NewProduction()
}

func ReplaceLogger(l *zap.Logger) {
	Logger = l
}
