package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	}
	t.Run("Deposite", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposite(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		assertBalance(t, got, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(30),
		}
		err := wallet.Withdraw(Bitcoin(10))
		if err != nil {
			t.Errorf("didn't expected any error but got '%s'", err)
		}
		got := wallet.Balance()
		want := Bitcoin(20)
		assertBalance(t, got, want)
	})

	t.Run("Withdraw insufficient balance", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}
		err := wallet.Withdraw(Bitcoin(30))
		if err != ErrInsufficiantBalance {
			t.Errorf("want '%s' but didn't get any", ErrInsufficiantBalance)
		}
	})
}
