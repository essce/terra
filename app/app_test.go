package app

import "testing"

func TestAdd(t *testing.T) {
	n, m := 1, 2
	v := Add(n, m)
	if v != 3 {
		t.Fatalf("unexpected response returned")
	}
}
