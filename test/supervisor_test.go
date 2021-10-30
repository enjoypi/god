package test

import (
	"testing"
	"time"

	"github.com/enjoypi/god/stdlib"
	"github.com/stretchr/testify/require"
)

func TestNewSupervisor(t *testing.T) {
	sup := stdlib.NewSupervisor()
	require.NotNil(t, sup)

	a := sup.Start(sampleActorType)
	require.NotNil(t, a)
	a.Post("hello")
	go func() {
		a.Post("word")
		time.Sleep(time.Second)
		stdlib.Close()
	}()
	stdlib.Wait()
}
