package queue

import "github.com/Wazzymandias/go-graph/vertex"

// NewIdxMinPQ initializes and returns a new index min priority queue
func NewIdxMinPQ() *IndexMinPQ {
	return &IndexMinPQ{
		idToIdx: make(map[vertex.ID]Index),
		Vertices:    make([]vertex.Vertex, 0), idxToVertexID: make(map[Index]vertex.ID)}
}
