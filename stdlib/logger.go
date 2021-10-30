package stdlib

import (
	"strings"

	"go.uber.org/zap"
)

var (
	L *zap.Logger
)

func initializeLogger(config Config) error {
	var logger *zap.Logger
	var err error

	if strings.ToLower(config.Log.Level) == zap.DebugLevel.String() {
		logger, err = zap.NewDevelopment()
	} else {
		pc := zap.NewProductionConfig()
		if err = pc.Level.UnmarshalText([]byte(config.Log.Level)); err != nil {
			return err
		} else {
			logger, err = pc.Build()
		}
	}

	if err != nil {
		return err
	}

	L = logger
	zap.ReplaceGlobals(logger)
	zap.RedirectStdLog(logger)
	return nil
}
