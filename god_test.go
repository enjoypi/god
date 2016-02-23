package god

import (
	"testing"
)

func TestStartNode(t *testing.T) {
	StartNode("0.0.0.0:3724")
	Console().Run()

}

func BenchmarkStartNode(b *testing.B) {
}
