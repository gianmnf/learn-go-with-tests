package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Gian")

	got := buffer.String()
	want := "Hello, Gian"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
