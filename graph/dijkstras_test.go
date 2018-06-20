package graph

import (
		"github.com/Wazzymandias/go-graph/vertex"
	"testing"
	"fmt"
)

// TODO randomized graphs
//        (v2)
//     /         \
//    1           2
//   /             \
// (v1) -- 1000 -- (v3)
func TestDijkstra(t *testing.T) {
	g := NewSimpleGraph()
	v1 := vertex.Vertex{Data: "foo", Weight: 99, ID:0}
	v2 := vertex.Vertex{Data: "bar", Weight: 13, ID:1}
	v3 := vertex.Vertex{Data: "baz", Weight: 7, ID:2}

	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddVertex(v3)

	v1v2 := Edge{Weight: 1, Src: v1, Dst: v2}
	v1v3 := Edge{Weight: 1000, Src: v1, Dst: v3}
	v2v3 := Edge{Weight: 2, Src: v2, Dst: v3}

	g.AddEdge(v1v2)
	g.AddEdge(v1v3)
	g.AddEdge(v2v3)

	// final graph:
	//
	// (v1: 99) --1-- (v2: 13)
	//  \              |
	//   \             2
	//    \            |
	//     --1000--(v3: 7)
	// Cost/Weight = edgeWeight + vertex.Weight
	// Therefore, cost of shortest path to:
	//   - v1 -> 0  (path to self is 0)
	//   - v2 -> 14 (edgeWeight: 1, vertex.Weight: 13)
	//   - v3 -> 9  (edgeWeight: 2, vertex.Weight: 7)
	paths := Dijkstra(g, v1)

	fmt.Println(paths)

	if paths[v1].Weight != 0 {
		t.Fail()
	}

	if paths[v2].Weight != 14 {
		t.Fail()
	}

	if paths[v3].Weight != 9 {
		t.Fail()
	}
}
