package test

import (
	"context"

	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type sampleActor struct {
	stdlib.SimpleActor
}

var _ stdlib.Actor = (*sampleActor)(nil)
var logger *zap.Logger

func (s *sampleActor) Handle(ctx context.Context, message def.Message, args ...interface{}) error {
	logger.Debug("handle", zap.Any("message", message))
	return nil
}

func (s *sampleActor) Initialize(v *viper.Viper) error {
	_ = s.SimpleActor.Initialize()
	logger.Info("initialize")
	return nil
}

const sampleActorType = "testActor"

func init() {
	logger, _ = zap.NewDevelopment()
	stdlib.RegisterActorCreator(sampleActorType, newSampleActor)
}

func newSampleActor() stdlib.Actor {
	return &sampleActor{}
}
