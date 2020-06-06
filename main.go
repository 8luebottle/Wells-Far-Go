package main

import (
	"fmt"

	"github.com/8luebottle/Wells-Far-Go/accounts"
)

func main() {
	account := accounts.NewAccount("Baby Tiger")
	account.Deposit(10)
	fmt.Println(account)
}
