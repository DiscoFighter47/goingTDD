package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := &bytes.Buffer{}
	Greet(buffer, "Zahid")
	got := buffer.String()
	want := "Hello, Zahid"
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
