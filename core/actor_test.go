package core

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

type sampleActor struct {
	ActorImpl
}

func (s *sampleActor) Handle(event Event) {
	panic("implement me")
}

func (s *sampleActor) Impl() *ActorImpl {
	return &s.ActorImpl
}

func (s sampleActor) Initialize() error {
	return nil
}

func (s sampleActor) Terminate() {
}

var sampleActorType = rand.Int63()

func init() {
	defaultActorFactory.RegisterActorCreator(sampleActorType, newSampleActor)
}

func newSampleActor() Actor {
	return &sampleActor{}
}

func TestNewActor(t *testing.T) {
	a := defaultActorFactory.new(sampleActorType)
	require.NotNil(t, a)
}
