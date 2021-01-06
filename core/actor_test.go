package core

import (
	"math/rand"
	"testing"

	"go.uber.org/zap"

	"github.com/stretchr/testify/require"
)

type sampleActor struct {
	ActorImpl
}

var _ Actor = (*sampleActor)(nil)

func (s *sampleActor) Handle(message Message) Message {
	logger.Debug("handle", zap.Any("message", message))
	return nil
}

func (s *sampleActor) Impl() *ActorImpl {
	return &s.ActorImpl
}

func (s *sampleActor) Initialize() error {
	logger.Debug("initialize")
	return nil
}

func (s *sampleActor) Terminate() {
	logger.Debug("terminate")
}

var sampleActorType = rand.Int63()

func init() {
	logger, _ = zap.NewDevelopment()
	defaultActorFactory.RegisterActorCreator(sampleActorType, newSampleActor)
}

func newSampleActor() Actor {
	return &sampleActor{}
}

func TestSampleActor(t *testing.T) {
	a := defaultActorFactory.new(sampleActorType)
	require.NotNil(t, a)
	require.IsType(t, (*sampleActor)(nil), a)
	a.Post("hello")
	Wait()
}
