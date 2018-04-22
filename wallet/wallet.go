package wallet

import (
	"errors"
	"fmt"
)

// ErrInsufficient ..
var ErrInsufficient = errors.New("cannot withdraw, insufficient funds")

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
	if amount > w.bal {
		return ErrInsufficient
	}

	w.bal -= amount
	return nil
}
