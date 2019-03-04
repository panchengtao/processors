package p1

import (
	"github.com/wddpct/processors/p1"
	"testing"
)

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p1.Process()
	}
}
