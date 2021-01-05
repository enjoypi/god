package core

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

type sampleActor struct {
	Actor
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
