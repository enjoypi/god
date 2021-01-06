package test

import (
	"testing"
	"time"

	"github.com/enjoypi/god/core"
	"github.com/stretchr/testify/require"
)

func TestNewSupervisor(t *testing.T) {
	sup := core.NewSupervisor()
	require.NotNil(t, sup)

	a := sup.Start(sampleActorType)
	require.NotNil(t, a)
	a.Post("hello")
	go func() {
		a.Post("word")
		time.Sleep(time.Second)
		core.Close()
	}()
	core.Wait()
}
