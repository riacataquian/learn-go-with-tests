package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	desc := "Deposit adds given N balance to the current Wallet's balance"
	t.Run(desc, func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))

		assertBal(t, w, desc, Bitcoin(10))
	})

	desc = "Withdraw subtracts given N balance to the current Wallet's balance"
	t.Run(desc, func(t *testing.T) {
		w := Wallet{bal: Bitcoin(20)}
		err := w.Withdraw(Bitcoin(10))

		assertNoErr(t, err)
		assertBal(t, w, desc, Bitcoin(10))
	})

	desc = "Withdraw without insufficient fnds should return an error"
	t.Run(desc, func(t *testing.T) {
		sb := Bitcoin(20)
		w := Wallet{bal: sb}

		err := w.Withdraw(Bitcoin(100))

		assertErr(t, err)
		assertBal(t, w, desc, sb)
	})
}

func assertBal(t *testing.T, w Wallet, desc string, want Bitcoin) {
	t.Helper()
	got := w.Balance()

	if got != want {
		t.Errorf("%s: Wallet() = %s, want %s", desc, got, want)
	}
}

func assertNoErr(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expecting nil error, got %v", err)
	}
}

func assertErr(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Fatal("expecting error, got nil")
	}

	if err != ErrInsufficient {
		t.Fatalf("expecting %v, got %v", ErrInsufficient, err)
	}
}
