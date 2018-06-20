package vertex

import "testing"

func TestVertex_Equals(t *testing.T) {
	v1 := Vertex{Weight: 0, Data: "foo"}
	v2 := Vertex{Weight: 321, Data: "foo"}
	v3 := Vertex{Weight: 1, Data: "bar"}

	if !v1.Equals(v2) {
		t.Fail()
	}

	if v1.Equals(v3) {
		t.Fail()
	}
}
