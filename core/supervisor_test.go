package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSupervisor(t *testing.T) {
	sup := NewSupervisor()
	require.NotNil(t, sup)
	require.True(t, sup.Start(sampleActorType))
}
