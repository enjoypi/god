package test

import (
	"testing"

	"github.com/enjoypi/god/core"
	"github.com/stretchr/testify/require"
)

func TestActorFactory(t *testing.T) {
	a := core.NewActor(sampleActorType)
	require.NotNil(t, a)
	require.IsType(t, (*sampleActor)(nil), a)
}
