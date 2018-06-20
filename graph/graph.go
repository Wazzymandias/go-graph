package graph

import "github.com/Wazzymandias/go-graph/vertex"

// Graph interface represents basic methods
// that can be used to compose Graph data structures
// of various complexities
type Graph interface {
	AddVertex(v vertex.Vertex)
	AddEdge(e Edge)
	GetNeighbors(src vertex.Vertex) map[vertex.Vertex]uint64
	DeleteVertex(v vertex.Vertex)
	DeleteEdge(e Edge)
}

// NewSimpleGraph initializes and returns a reference to a new
// Simple Graph
func NewSimpleGraph() *SimpleGraph {
	return &SimpleGraph{Vertices: make(map[vertex.Vertex]map[vertex.Vertex]uint64)}
}
