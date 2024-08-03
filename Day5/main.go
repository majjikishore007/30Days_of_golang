package main

import (
	"errors"
	"fmt"
)
type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC",b)
}

type  Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amout Bitcoin){
	fmt.Printf("addres in function of %v\n",&w.balance)
	w.balance+= amout
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

    if amount > w.balance {
        return ErrInsufficientFunds
    }

    w.balance -= amount
    return nil
}

func (w *Wallet) Balance() Bitcoin{
	// this can be skipped becase the pointer are automatically dereferenced
	return (*w).balance
}

func main()  {
	
}