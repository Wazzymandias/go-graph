package graph

import "testing"

func TestNewSimpleGraph(t *testing.T) {
	g := NewSimpleGraph()

	if g == nil {
		t.Fail()
	}
}
