package main

import "testing"

func TestUser(t *testing.T) {
	md := &MockDatastore{
		Users: map[int]User{
			2: {ID: 2, First: "jenny"},
		},
	}
	s := &Service{
		ds: md,
	}
	u, err := s.GetUser(2)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if u.First != "jenny" {
		t.Errorf("got: %s, want %s", u.First, "jenny")
	}
}
