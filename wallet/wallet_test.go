package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	w := Wallet{}
	w.Deposit(10)

	got := w.Balance()
	want := 10

	if got != want {
		desc := "Accepts a a balance and sum it to the current balance"
		t.Errorf("%s: Wallet() = %d, want %d", desc, got, want)
	}
}
