package core

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type sampleActor struct {
	ActorImpl
}

func (s *sampleActor) Handle(event Event) {
	logger.Debug(reflect.TypeOf(event).Elem().Name())
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
	defaultActorFactory.RegisterActorCreator(sampleActorType, newSampleActor)
}

func newSampleActor() Actor {
	return &sampleActor{}
}

func TestSampleActor(t *testing.T) {
	a := defaultActorFactory.new(sampleActorType)
	require.NotNil(t, a)
	require.IsType(t, (*sampleActor)(nil), a)
}
