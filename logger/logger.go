package logger

import (
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	L *zap.Logger
)

func Initialize(v *viper.Viper) error {

	var cfg config
	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}

	var logger *zap.Logger
	var err error

	if strings.ToLower(cfg.Log.Level) == zap.DebugLevel.String() {
		logger, err = zap.NewDevelopment()
	} else {
		pc := zap.NewProductionConfig()
		if err = pc.Level.UnmarshalText([]byte(cfg.Log.Level)); err != nil {
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

func CheckError(msg string, err error) {
	if err != nil {
		L.Error(msg, zap.Error(err))
	}
}
