package stdlib

import (
	"github.com/enjoypi/god/logger"
	"go.uber.org/zap"
)

func Catch(f func()) {
	defer func() {
		if r := recover(); r != nil {
			logger.L.Error("panic", zap.Any("recover", r))
		}
	}()

	f()
}
