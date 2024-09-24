package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Gian")
	want := "Hello, Gian"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
