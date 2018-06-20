package graph

import "github.com/Wazzymandias/go-graph/vertex"

// SimpleGraph represents an undirected, weighted graph
// with no parallel edges or loops.
type SimpleGraph struct {
	// Vertices are a set of vertices, each with
	// their own set of neighbors. Since this is a simple graph,
	// if an edge exists between Vertex A and Vertex B,
	// Vertex B will appear in Vertex A's set and vice versa.
	Vertices map[vertex.Vertex]map[vertex.Vertex]uint64
	count    int
}

// AddVertex adds a given vertex to the graph, if it doesn't
// already exist. Attempting to add an already existing
// vertex to the Graph will essentially be a noop.
func (g *SimpleGraph) AddVertex(v vertex.Vertex) {
	_, exists := g.Vertices[v]

	if !exists {
		g.Vertices[v] = make(map[vertex.Vertex]uint64)
		g.count++
	}
}

// AddEdge creates an association with a given weight between two vertices.
// Each vertex in the edge adds the other to its set of neighbors.
func (g *SimpleGraph) AddEdge(e Edge) {
	_, srcExists := g.Vertices[e.Src]

	if !srcExists {
		g.Vertices[e.Src] = make(map[vertex.Vertex]uint64)
	}

	_, dstExists := g.Vertices[e.Dst]

	if !dstExists {
		g.Vertices[e.Dst] = make(map[vertex.Vertex]uint64)
	}

	g.Vertices[e.Src][e.Dst] = e.Weight
	g.Vertices[e.Dst][e.Src] = e.Weight
}

// GetNeighbors returns the set of vertices adjacent to a given vertex.
func (g *SimpleGraph) GetNeighbors(src vertex.Vertex) map[vertex.Vertex]uint64 {
	return g.Vertices[src]
}

// DeleteVertex removes a given vertex from the Graph,
// and deletes it from the set of neighbors for every other
// vertex in the graph.
func (g *SimpleGraph) DeleteVertex(v vertex.Vertex) {
	delete(g.Vertices, v)

	for v := range g.Vertices {
		_, hasNeighbors := g.Vertices[v]

		if !hasNeighbors {
			continue
		}

		_, exists := g.Vertices[v][v]

		if exists {
			delete(g.Vertices[v], v)
		}
	}

	g.count--
}

// DeleteEdge updates the set of neighbors
// for each vertex in the edge. Each vertex attempts to
// remove the other from its set of neighbors, if it exists.
func (g *SimpleGraph) DeleteEdge(e Edge) {
	delete(g.Vertices[e.Src], e.Dst)
	delete(g.Vertices[e.Dst], e.Src)
}

// GetEdgeWeight returns the cost associated with a given edge between
// two vertices and a boolean denoting existence.
// Edge weights are assumed to be non negative integers.
// If an edge does not exist between the specified vertices, a false boolean
// value will be returned. Note that since edges may exist with 0 weight value,
// the boolean should be checked for existence.
func (g *SimpleGraph) GetEdgeWeight(src, dst vertex.Vertex) (uint64, bool) {
	if _, exists := g.Vertices[src]; !exists {
		return 0, false
	}

	if _, exists := g.Vertices[src][dst]; !exists {
		return 0, false
	}

	return g.Vertices[src][dst], true
}
