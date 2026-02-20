package main

import "testing"

func TestAdd(t *testing.T) { //This calls the struct type testing.'T' from the testing package so that it can perform the test
	got := Add(9, 3)
	want := 12
	if got != want {
		t.Errorf("error- want %d and got %d", want, got) // here we can
	}
}
