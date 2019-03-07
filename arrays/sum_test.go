package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	checkSum := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%d' want '%d' given '%v'", got, want, numbers)
		}
	}

	t.Run("Collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		checkSum(t, got, want, numbers)
	})

	t.Run("Empty slice", func(t *testing.T) {
		numbers := []int{}
		got := Sum(numbers)
		want := 0
		checkSum(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}
	t.Run("Make the sums of some slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{0, 9})
		want := []int{6, 9}
		checkSums(t, got, want)
	})

	t.Run("Empty slice", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{})
		want := []int{6, 0}
		checkSums(t, got, want)
	})
}
