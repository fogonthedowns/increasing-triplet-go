package main

import (
	"testing"
)

func TestFun(t *testing.T) {

	out := increasingTriplet([]int{1, 2, 3, 4, 5})
	want := true
	if out != true {
		t.Errorf("got %t, want %t", out, want)
	}

	second := increasingTriplet([]int{5, 4, 3, 2, 1})
	want = false
	if second != false {
		t.Errorf("got %t, want %t", second, want)
	}

}
