package plugin

import (
	"fmt"
	"plugin"
	"testing"
)

func TestLoadAndRunP1(t *testing.T) {
	p, err := plugin.Open("../main/processorutil.so")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("P1Process")
	if err != nil {
		panic(err)
	}
	fmt.Println(f.(func()string)())
}

func BenchmarkLoadAndRunP1(b *testing.B) {
	p, err := plugin.Open("../main/processorutil.so")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("P2Process")
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		f.(func()string)()
	}
}
