package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := add(2,2)
	expected := 4

	if (sum != expected) {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
func ExampleAdd() {
	sum := add(1, 5)
	fmt.Println(sum)
	// Output: 6
}