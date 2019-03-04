package p2

import (
	"github.com/wddpct/processors/p2"
	"testing"
)

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p2.Process()
	}
}
