package main

import (
	"fmt"
	"github.com/8luebottle/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Baby Tiger")
	account.Deposit(10)
	fmt.Println(account)
}
