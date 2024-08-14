package main

import (
	"bytes"
	"testing"
)


func TestGreet(t *testing.T)  {
	buffer := bytes.Buffer{}
	Greet(&buffer,"majji")

	got := buffer.String()
	want := "Hello majji"

	if got != want{
		t.Fatalf("got %q  want %q",got,want)
	}
}