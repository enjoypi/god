package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestActorFactory(t *testing.T) {
	a := defaultActorFactory.new(sampleActorType)
	require.NotNil(t, a)
	require.IsType(t, (*sampleActor)(nil), a)
}
