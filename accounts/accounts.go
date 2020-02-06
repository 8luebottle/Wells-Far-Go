package accounts

import (
	"errors"
	"fmt"
)

// Account struct (Private)
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw money from your account.")

// NewAccount creates Account
func NewAccount(owner string) *Account { // Don't make it a copy. Use Pointer
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) { // a --> receiver
	fmt.Println("Gonna deposit", amount)
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) { // return nothing
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// fmt.Println(account)
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s Account\n", "Balance :", a.Balance())
}
