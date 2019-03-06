package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2, 4)
	want := 6

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func ExampleAdd() {
	sum := Add(4, 6)
	fmt.Println(sum)
	// Output: 10
}
