package test

import (
	"context"
	"testing"
	"time"

	"github.com/enjoypi/god/stdlib"
	"github.com/stretchr/testify/require"
)

func TestNewSupervisor(t *testing.T) {
	sup, _ := stdlib.NewSupervisor()
	require.NotNil(t, sup)

	a, _ := sup.Start(nil, sampleActorType, 0)
	require.NotNil(t, a)
	ctx := context.Background()
	a.Post(ctx, "hello")
	go func() {
		a.Post(ctx, "word")
		time.Sleep(time.Second)
		stdlib.Close()
	}()
	stdlib.Wait()
}
