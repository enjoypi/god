package test

import (
	"testing"

	"github.com/enjoypi/god/actors"
	"github.com/stretchr/testify/require"
)

func TestActorFactory(t *testing.T) {
	a := actors.NewActor(sampleActorType)
	require.NotNil(t, a)
	require.IsType(t, (*sampleActor)(nil), a)
}
