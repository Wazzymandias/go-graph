package queue

import "testing"

func TestNewIdxMinPQ(t *testing.T) {
	pq := NewIdxMinPQ()

	if pq == nil {
		t.Fail()
	}
}