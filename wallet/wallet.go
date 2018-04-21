package wallet

// Wallet ..
type Wallet struct {
	bal int
}

// Deposit ..
func (w Wallet) Deposit(amount int) {
	w.bal += amount
}

// Balance ..
func (w Wallet) Balance() int {
	return w.bal
}
