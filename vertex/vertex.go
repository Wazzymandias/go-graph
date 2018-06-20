package vertex

// ID is simply an alias to an integer, but
// is useful for readability.
type ID = int

// Vertex represents a node in a graph, associated with some
// data, weight, and a source vertex.
//
// Data and weight are used to uniquely identify the set of vertices
// of a graph.
type Vertex struct {
	Data   interface{}
	Weight uint64

	// ID is useful in situations where vertices are indexed by a unique key
	ID ID
}

// Equals uses the data of vertices for comparison.
// Two vertices are said to be equal if their data is the same.
func (v Vertex) Equals(other Vertex) bool {
	return v.Data == other.Data
}
