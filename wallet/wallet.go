package wallet

import (
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
