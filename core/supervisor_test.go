package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewSupervisor(t *testing.T) {
	sup := NewSupervisor()
	require.NotNil(t, sup)
	require.True(t, sup.Start(sampleActorType))
	go func() {
		time.Sleep(time.Second)
		Close()
	}()
	Wait()
}
