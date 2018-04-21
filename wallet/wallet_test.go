package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	w := Wallet{}
	w.Deposit(Bitcoin(10))

	got := w.Balance()
	want := Bitcoin(10)

	if got != want {
		desc := "Accepts a a balance and sum it to the current balance"
		t.Errorf("%s: Wallet() = %s, want %s", desc, got, want)
	}
}
