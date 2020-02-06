package main

import (
	"fmt"
	"github.com/8luebottle/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Baby Tiger")
	account.Deposit(10)

	// Handle Errors
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s's Current Balance : %d", account.Owner(), account.Balance())
	fmt.Println("")
}
