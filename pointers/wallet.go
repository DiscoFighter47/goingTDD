package pointers

import "errors"

// ErrInsufficiantBalance represents insufficiant balance error
var ErrInsufficiantBalance = errors.New("Insufficient balance")

// Wallet holdes information of a wallet
type Wallet struct {
	balance Bitcoin
}

// Deposite deposites an amount into the wallet
func (w *Wallet) Deposite(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the balance of the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw withdraws an amount from the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficiantBalance
	}
	w.balance -= amount
	return nil
}
