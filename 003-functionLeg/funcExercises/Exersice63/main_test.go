package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Sum is not correct got: %v, want: %v\n", total, 10)
	}
}
