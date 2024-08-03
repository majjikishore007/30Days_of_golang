package main

import (
	"testing"
)

func TestWallet(t *testing.T)  {
    t.Run("Deposit",func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        want :=Bitcoin(10)
        assertBalance(t,wallet,want)
        
    })
    t.Run("Withdraw",func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(10)}
        wallet.Withdraw(Bitcoin(8))
        want :=Bitcoin(2)
        assertBalance(t,wallet,want)
    })
    t.Run("Withdraw insufficent funds",func(t *testing.T) {
        startingBalance := Bitcoin(20)
        wallet := Wallet{startingBalance}
        err := wallet.Withdraw(Bitcoin(100))

        assertBalance(t, wallet, startingBalance)
        assertError(t,err,ErrInsufficientFunds.Error())
    })
    
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin){
    t.Helper()
    got := wallet.Balance()
    if got != want {
        t.Errorf("got %d and want %d",got,want)
    }
}
func assertError(t *testing.T, got error,want string)  {
    t.Helper()
    if got == nil {
        t.Fatal("didn't get an error but wanted one")
    }

    if got.Error() != want {
        t.Errorf("got %q, want %q", got, want)
    }
}