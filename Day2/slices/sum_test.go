package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection sum of fixed size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)

		want := 15
		if got != want {
			t.Errorf("got %d , want %d and given %v", got, want, numbers)
		}
	})
	t.Run("collection with no size ", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)

		want := 6
		if got != want {
			t.Errorf("got %d , want %d and given %v", got, want, numbers)
		}

	})
	t.Run("variable functions ", func(t *testing.T) {
		got := SumAll([]int{0, 2}, []int{1, 2, 3})
		want := []int{2, 6}
		// want := "bob"
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v and given", got, want)
		}

	})
	t.Run("Sum of the tails leaving one element ", func(t *testing.T) {
		got := SumOfTails([]int{}, []int{0, 9, 5, 6})
		want := []int{0, 20}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
