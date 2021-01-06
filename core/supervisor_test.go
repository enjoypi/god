package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewSupervisor(t *testing.T) {
	sup := NewSupervisor()
	require.NotNil(t, sup)

	a := sup.Start(sampleActorType)
	require.NotNil(t, a)
	a.Post("hello")
	go func() {
		a.Post("word")
		time.Sleep(time.Second)
		Close()
	}()
	Wait()
}
