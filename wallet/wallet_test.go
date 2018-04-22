package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBal := func(t *testing.T, w Wallet, desc string, want Bitcoin) {
		t.Helper()
		got := w.Balance()

		if got != want {
			t.Errorf("%s: Wallet() = %s, want %s", desc, got, want)
		}
	}

	desc := "Deposit adds given N balance to the current Wallet's balance"
	t.Run(desc, func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))

		assertBal(t, w, desc, Bitcoin(10))
	})

	desc = "Withdraw subtracts given N balance to the current Wallet's balance"
	t.Run(desc, func(t *testing.T) {
		sb := Bitcoin(20)
		w := Wallet{bal: sb}

		err := w.Withdraw(Bitcoin(100))

		assertBal(t, w, desc, sb)

		if err == nil {
			t.Error("expecting error, got nil")
		}
	})
}
