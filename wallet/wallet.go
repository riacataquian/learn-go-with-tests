package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin ..
type Bitcoin int

// Stringer ..
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet ..
type Wallet struct {
	bal Bitcoin
}

// Deposit ..
func (w *Wallet) Deposit(amount Bitcoin) {
	w.bal += amount
}

// Balance ..
func (w *Wallet) Balance() Bitcoin {
	return w.bal
}

// Withdraw ..
func (w *Wallet) Withdraw(amount Bitcoin) error {
	cb := w.bal

	if cb-amount > 0 {
		w.bal -= amount
		return nil
	}

	return errors.New("Wallet has insufficient funds")
}
