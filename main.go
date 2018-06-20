package main

import (
	"github.com/Wazzymandias/go-graph/vertex"
	"github.com/Wazzymandias/go-graph/graph"
	"fmt"
)

func main() {
	g := graph.NewSimpleGraph()
	v1 := vertex.Vertex{Data: "foo", Weight: 99, ID:0}
	v2 := vertex.Vertex{Data: "bar", Weight: 13, ID:1}
	v3 := vertex.Vertex{Data: "baz", Weight: 7, ID:2}

	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddVertex(v3)

	v1v2 := graph.Edge{Weight: 1, Src: v1, Dst: v2}
	v1v3 := graph.Edge{Weight: 1000, Src: v1, Dst: v3}
	v2v3 := graph.Edge{Weight: 2, Src: v2, Dst: v3}

	g.AddEdge(v1v2)
	g.AddEdge(v1v3)
	g.AddEdge(v2v3)

	paths := graph.Dijkstra(g, v1)

	for v, path := range paths {
		fmt.Printf("Vertex ID: %d \t Cost: %d \t From ID: %d\n", v.ID, path.Weight, path.Src.ID)
	}
}
