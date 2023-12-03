package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Sum is not correct got: %v, want: %v.\n", total, 10)
	}
}

func TestSubtract(t *testing.T) {
	total := Subtract(5, 2)
	if total != 3 {
		t.Errorf("Sum is not correct got: %v, want: %v.\n", total, 3)
	}
}

func TestDoMath(t *testing.T) {
	got := DoMath(2, 4, Add)
	want := 6
	if got != want {
		t.Errorf("Sum is not correct got: %v, want: %v.\n", got, want)
	}
}
