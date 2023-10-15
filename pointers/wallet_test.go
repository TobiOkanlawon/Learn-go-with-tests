package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet *Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but did not get one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposits", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, &wallet, want)
	})

	t.Run("withdraw accurately", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		if err != nil {
			t.Fatalf("Did not expect an error but got %q", err)
		}

		want := Bitcoin(10)

		assertBalance(t, &wallet, want)
	})

	t.Run("rejects overdraft withdrawal", func(t *testing.T) {
		wallet := &Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(20))

		assertError(t, err, ErrInsufficientFunds)
	})
}

func TestBitcoin(t *testing.T) {
	got := Bitcoin(10).String()
	want := "10 BTC"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
