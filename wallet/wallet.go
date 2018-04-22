package wallet

import (
	"errors"
	"fmt"
)

// ErrInsufficient is the error to be returned for insufficient funds.
var ErrInsufficient = errors.New("cannot withdraw, insufficient funds")

// Bitcoin represents the balance of a Wallet.
type Bitcoin int

// String formats Bitcoin when printed as string.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet encapsulates a balance.
type Wallet struct {
	bal Bitcoin
}

// Deposit adds amount to the Wallet's current balance.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.bal += amount
}

// Balance returns the current balance of a Wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.bal
}

// Withdraw substracts the amount from the Wallet's current balance.
// Returns an error if funds is insufficient.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.bal {
		return ErrInsufficient
	}

	w.bal -= amount
	return nil
}
