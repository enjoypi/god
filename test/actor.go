package test

import (
	"math/rand"

	"github.com/enjoypi/god/types"

	"github.com/enjoypi/god/stdlib"
	"go.uber.org/zap"
)

type sampleActor struct {
	stdlib.DefaultActor
}

var _ stdlib.Actor = (*sampleActor)(nil)
var logger *zap.Logger

func (s *sampleActor) Handle(message types.Message) types.Message {
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
	stdlib.RegisterActorCreator(sampleActorType, newSampleActor)
}

func newSampleActor() stdlib.Actor {
	return &sampleActor{}
}
