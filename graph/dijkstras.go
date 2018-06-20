package graph

import (
	"github.com/Wazzymandias/go-graph/queue"
	"github.com/Wazzymandias/go-graph/vertex"
	)

// Paths represents the predecessor and weight of a
// given vertex.
type Paths map[vertex.Vertex]Path

// Path holds weight and source information, associated
// with a given vertex. Weight can be a function of several
// factors:
//   - Weight of a destination vertex
//   - Weight of an edge
//   - Weight of source vertex
// The way path is currently being used is with
// weight = dstVertex.Weight + edgeWeight
// However, src vertex is provided if instead it's required that:
// weight = srcVertex.Weight + edgeWeight + dstVertex.Weight
type Path struct {
	Weight uint64
	Src    vertex.Vertex
}

// Dijkstra takes a graph and source node and returns the
// set of shortest paths from src to each vertex.
// Note: Current implementation inserts vertex into priority queue
// using its ID as the key. It is assumed that each
// vertex has a unique ID in the set of all vertices.
func Dijkstra(g *SimpleGraph, src vertex.Vertex) Paths {
	// Go doesn't have builtin sets - a boolean map is the closest facsimile
	visited := make(map[vertex.Vertex]bool)

	weights := Paths{src: {Weight: 0, Src: src}}

	pq := queue.NewIdxMinPQ()
	pq.Insert(src)

	for !pq.IsEmpty() {
		current := pq.ExtractMin()

		if visited[current] {
			continue
		}

		for nbr := range g.GetNeighbors(current) {

			edgeWeight, _ := g.GetEdgeWeight(current, nbr)

			if _, exists := weights[nbr]; !exists {

				weights[nbr] = Path{Src: current, Weight: edgeWeight + nbr.Weight}

				pq.Insert(nbr)

			} else if weights[nbr].Weight > edgeWeight+nbr.Weight {

				weights[nbr] = Path{Src: current, Weight: edgeWeight + nbr.Weight}
				pq.DecreaseWeight(nbr, edgeWeight + nbr.Weight)

			}
		}

		visited[current] = true
	}

	return weights
}
