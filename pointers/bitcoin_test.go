package pointers

import "testing"

func TestBitcoin(t *testing.T) {
	b := Bitcoin(10)
	got := b.String()
	want := "10 BTC"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
