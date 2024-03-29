package test

import (
	"testing"

	"github.com/enjoypi/god/stdlib"
	"github.com/stretchr/testify/require"
)

func TestActorFactory(t *testing.T) {
	a := stdlib.NewActor(sampleActorType, 0)
	require.NotNil(t, a)
	require.IsType(t, (*sampleActor)(nil), a)
}
