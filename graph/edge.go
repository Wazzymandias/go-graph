package graph

import "github.com/Wazzymandias/go-graph/vertex"

// Edge is used to represent a relationship
// between two vertices in an undirected, weighted graph.
type Edge struct {
	Src, Dst vertex.Vertex
	Weight   uint64
}
