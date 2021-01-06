package test

import (
	"math/rand"

	"github.com/enjoypi/god/core"
	"go.uber.org/zap"
)

type sampleActor struct {
	core.DefaultActor
}

var _ core.Actor = (*sampleActor)(nil)
var logger *zap.Logger

func (s *sampleActor) Handle(message core.Message) core.Message {
	logger.Debug("handle", zap.Any("message", message))
	return nil
}

func (s *sampleActor) Initialize() error {
	_ = s.DefaultActor.Initialize()
	logger.Info("initialize")
	return nil
}

var sampleActorType = rand.Int63()

func init() {
	logger, _ = zap.NewDevelopment()
	core.RegisterActorCreator(sampleActorType, newSampleActor)
}

func newSampleActor() core.Actor {
	return &sampleActor{}
}
