package queue

import (
	"testing"
	"github.com/Wazzymandias/go-graph/vertex"
)

func TestIndexMinPQ_Insert(t *testing.T) {
	pq := NewIdxMinPQ()
	v := vertex.Vertex{Weight:9000, Data:"foo", ID:0}
	pq.Insert(v)

	if pq.IsEmpty() {
		t.Fail()
	}

	found := pq.ExtractMin()

	if v != found {
		t.Fail()
	}
}

func TestIndexMinPQ_IsEmpty(t *testing.T) {
	pq := NewIdxMinPQ()
	v := vertex.Vertex{Weight:9000, Data:"foo", ID:0}
	pq.Insert(v)

	if pq.IsEmpty() {
		t.Fail()
	}

	pq.ExtractMin()

	if !pq.IsEmpty() {
		t.Fail()
	}
}

func TestIndexMinPQ_DecreaseWeight(t *testing.T) {
	pq := NewIdxMinPQ()
	v := vertex.Vertex{Weight:9000, Data:"foo", ID:0}
	pq.Insert(v)

	var targetWeight uint64 = 10
	pq.DecreaseWeight(v, targetWeight)

	found := pq.ExtractMin()

	if found.Weight != targetWeight {
		t.Fail()
	}
}

func TestIndexMinPQ_ExtractMin(t *testing.T) {
	pq := NewIdxMinPQ()
	v := vertex.Vertex{Weight:9000, Data:"foo", ID:0}
	pq.Insert(v)

	found := pq.ExtractMin()

	if found != v {
		t.Fail()
	}
}
