package test

import (
	"github.com/enjoypi/god/actors"
	"github.com/enjoypi/god/types"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type sampleActor struct {
	actors.SimpleActor
}

var _ actors.Actor = (*sampleActor)(nil)
var logger *zap.Logger

func (s *sampleActor) Handle(message types.Message) error {
	logger.Debug("handle", zap.Any("message", message))
	return nil
}

func (s *sampleActor) Initialize(v *viper.Viper) error {
	_ = s.SimpleActor.Initialize()
	logger.Info("initialize")
	return nil
}

var sampleActorType = "sample"

func init() {
	logger, _ = zap.NewDevelopment()
	actors.RegisterActorCreator(sampleActorType, newSampleActor)
}

func newSampleActor() actors.Actor {
	return &sampleActor{}
}
