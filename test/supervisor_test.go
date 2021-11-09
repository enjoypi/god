package test

import (
	"testing"
	"time"

	"github.com/enjoypi/god/stdlib"
	"github.com/stretchr/testify/require"
)

func TestNewSupervisor(t *testing.T) {
	sup, _ := stdlib.NewSupervisor()
	require.NotNil(t, sup)

	a, _ := sup.Start(nil, sampleActorType)
	require.NotNil(t, a)
	a.Post("hello")
	go func() {
		a.Post("word")
		time.Sleep(time.Second)
		stdlib.Close()
	}()
	stdlib.Wait()
}
