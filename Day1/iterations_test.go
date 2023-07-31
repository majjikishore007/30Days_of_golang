package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t*testing.T){
	repeat:= Repeat("a",5)
	expected := "aaaaa"

	if repeat != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeat)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a" , 5)
	}
}

func ExampleRepeat(){
	a :=Repeat("a",5);
	fmt.Println(a)
	// output aaaaa
}