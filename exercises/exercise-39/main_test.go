package main

import "testing"

func TestParadise(t *testing.T) {
	got := Paradise("Heaven")
	want := "The location where i think paradise is Heaven"
	if got != want {
		t.Errorf("Error: want %s and got %s", want, got)
	}
}
