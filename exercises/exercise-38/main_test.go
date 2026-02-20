package main

import "testing"

func TestDoMath(t *testing.T) {
	got := DoMath(12, 43, Add)
	want := 55
	if got != want {
		t.Errorf("error- want is %v got is %v", got, want)
	}
}
