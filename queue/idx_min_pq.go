package queue

import (
	"github.com/Wazzymandias/go-graph/vertex"
	)

// MinVertexHeap provides a collection of methods
// that are relevant to any data structure that requires
// a min heap that operates on graphs.
type MinVertexHeap interface {
	Insert(v vertex.Vertex)
	ExtractMin() vertex.Vertex
	IsEmpty() bool
	DecreaseKey(vertex vertex.Vertex, value uint64)
}

// Index is simply an alias to an integer, but is
// useful for readability.
type Index = int

// IndexMinPQ is a priority queue that contains
// a mapping of vertices in addition to the array that
// represents the heap of elements.
// Keeping track of the index mappings can be useful
// for accessing values in constant time while retaining
// the properties of a min heap for other operations.
// However, it comes at the cost of maintaining
// auxiliary structures - a trade off of space vs speed.
type IndexMinPQ struct {
	idxToVertexID map[Index]vertex.ID

	idToIdx map[vertex.ID]Index

	Vertices []vertex.Vertex

	count int
}


// Insert stores a given vertex in the priority queue
// using the vertex ID as its key
func (mpq *IndexMinPQ) Insert(v vertex.Vertex) {
	var idx = mpq.count

	mpq.idxToVertexID[idx] = v.ID
	mpq.idToIdx[v.ID] = idx

	mpq.Vertices = append(mpq.Vertices, v)

	// If first element, bubble up essentially becomes noop
	mpq.bubbleUp(mpq.count)

	mpq.count++
}

// ExtractMin removes the smallest element from the
// priority queue and returns it. Additional
// operations are performed to maintain heap and set
// invariants
func (mpq *IndexMinPQ) ExtractMin() vertex.Vertex {
	if mpq.count == 0 {
		return vertex.Vertex{}
	}

	result := mpq.Vertices[0]

	// Move last element to head
	mpq.Vertices[0] = mpq.Vertices[mpq.count-1]

	// update maps to reflect extraction
	delete(mpq.idToIdx, result.ID)
	delete(mpq.idxToVertexID, mpq.count-1)

	// update index of moved element
	mpq.idToIdx[mpq.Vertices[0].ID] = 0
	mpq.idxToVertexID[0] = mpq.Vertices[0].ID

	mpq.count--
	// resize vertices array
	mpq.Vertices = mpq.Vertices[:mpq.count]
	mpq.bubbleDown(0)

	return result
}

// bubbleUp will attempt to swap values at indices up through the heap
// in order to maintain the min heap invariant.
func (mpq *IndexMinPQ) bubbleUp(index int) {
	for index > 0 && mpq.Vertices[index/2].Weight > mpq.Vertices[index].Weight {
		// update maps to reflect the switch
		mpq.idToIdx[mpq.Vertices[index/2].ID] = index
		mpq.idToIdx[mpq.Vertices[index].ID] = index / 2

		mpq.idxToVertexID[index] = mpq.Vertices[index / 2].ID
		mpq.idxToVertexID[index / 2] = mpq.Vertices[index].ID

		mpq.Vertices[index/2], mpq.Vertices[index] = mpq.Vertices[index], mpq.Vertices[index/2]

		index /= 2
	}
}

// bubbleDown will attempt to swap values at indices down through the heap
// in order to maintain the min heap invariant
func (mpq *IndexMinPQ) bubbleDown(index int) {
	for 2*(index+1) < mpq.count {
		childInd := 2 * (index + 1)

		if childInd+1 < mpq.count && mpq.Vertices[childInd].Weight > mpq.Vertices[childInd+1].Weight {
			childInd++
		}

		// if value at index is greater than value for children, swap
		if mpq.Vertices[index].Weight < mpq.Vertices[childInd].Weight {
			break
		}

		// update maps to reflect the switch
		mpq.idToIdx[mpq.Vertices[childInd].ID] = index
		mpq.idToIdx[mpq.Vertices[index].ID] = childInd

		mpq.idxToVertexID[index] = mpq.Vertices[childInd].ID
		mpq.idxToVertexID[childInd] = mpq.Vertices[index].ID

		mpq.Vertices[childInd], mpq.Vertices[index] = mpq.Vertices[index], mpq.Vertices[childInd]

		index = childInd
	}
}

// IsEmpty checks whether or not the priority queue is empty
func (mpq *IndexMinPQ) IsEmpty() bool {
	return mpq.count == 0
}

// DecreaseWeight will attempt to set the weight of a given vertex
// to the new value. It will return early if value is greater than current weight.
func (mpq *IndexMinPQ) DecreaseWeight(v vertex.Vertex, value uint64) {
	if _, exists := mpq.idToIdx[v.ID]; !exists {
		return
	}

	idx := mpq.idToIdx[v.ID]

	curWt := mpq.Vertices[idx].Weight

	if curWt < value {
		return
	}

	mpq.Vertices[mpq.idToIdx[v.ID]].Weight = value

	mpq.bubbleUp(mpq.idToIdx[v.ID])
}
