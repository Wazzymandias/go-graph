package graph

import (
	"testing"
	"github.com/Wazzymandias/go-graph/vertex"
	"reflect"
)

func TestSimpleGraph_AddVertex(t *testing.T) {
	g := NewSimpleGraph()

	v := vertex.Vertex{ID:0, Data:"foo", Weight:9000}
	g.AddVertex(v)

	if _, exists := g.Vertices[v]; !exists {
		t.Fail()
	}

}

func TestSimpleGraph_AddEdge(t *testing.T) {
	g := NewSimpleGraph()

	v1 := vertex.Vertex{ID:0, Data:"foo", Weight:9000}
	v2 := vertex.Vertex{ID:1, Data:"bar", Weight:666}

	e := Edge{Weight:123456, Src:v1, Dst:v2}

	g.AddVertex(v1)
	g.AddVertex(v2)

	g.AddEdge(e)

	found, ok := g.GetEdgeWeight(v1, v2)

	if !ok || found != e.Weight {
		t.Fail()
	}

	if _, exists := g.Vertices[v1][v2]; !exists {
		t.Fail()
	}

	if _, exists := g.Vertices[v2][v1]; !exists {
		t.Fail()
	}
}

func TestSimpleGraph_DeleteEdge(t *testing.T) {
	g := NewSimpleGraph()

	v1 := vertex.Vertex{ID:0, Data:"foo", Weight:9000}
	v2 := vertex.Vertex{ID:1, Data:"bar", Weight:666}

	e := Edge{Weight:123456, Src:v1, Dst:v2}

	g.AddVertex(v1)
	g.AddVertex(v2)

	g.AddEdge(e)

	found, ok := g.GetEdgeWeight(v1, v2)

	if !ok || found != e.Weight {
		t.Fail()
	}

	g.DeleteEdge(e)


	var expectedWeight uint64 = 0

	found, ok = g.GetEdgeWeight(v1, v2)

	if ok || found != expectedWeight {
		t.Fail()
	}
}

func TestSimpleGraph_DeleteVertex(t *testing.T) {
	g := NewSimpleGraph()

	v := vertex.Vertex{ID:0, Data:"foo", Weight:9000}

	g.AddVertex(v)

	if _, exists := g.Vertices[v]; !exists {
		t.Fail()
	}

	g.DeleteVertex(v)

	if _, exists := g.Vertices[v]; exists {
		t.Fail()
	}

	if len(g.Vertices) != 0 {
		t.Fail()
	}
}

func TestSimpleGraph_GetEdgeWeight(t *testing.T) {
	g := NewSimpleGraph()

	v1 := vertex.Vertex{ID:0, Data:"foo", Weight:9000}
	v2 := vertex.Vertex{ID:1, Data:"bar", Weight:666}

	e := Edge{Weight:123456, Src:v1, Dst:v2}

	g.AddVertex(v1)
	g.AddVertex(v2)

	g.AddEdge(e)

	found, ok := g.GetEdgeWeight(v1, v2)

	if !ok || found != e.Weight {
		t.Fail()
	}
}

func TestSimpleGraph_GetNeighbors(t *testing.T) {
	g := NewSimpleGraph()

	v1 := vertex.Vertex{ID:0, Data:"foo", Weight:9000}
	v2 := vertex.Vertex{ID:1, Data:"bar", Weight:666}

	e := Edge{Weight:123456, Src:v1, Dst:v2}

	g.AddVertex(v1)
	g.AddVertex(v2)

	g.AddEdge(e)

	v1Expected := map[vertex.Vertex]uint64{v2: e.Weight}
	v2Expected := map[vertex.Vertex]uint64{v1: e.Weight}

	if !reflect.DeepEqual(g.GetNeighbors(v1), v1Expected) {
		t.Fail()
	}

	if !reflect.DeepEqual(g.GetNeighbors(v2), v2Expected) {
		t.Fail()
	}
}