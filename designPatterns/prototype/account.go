package main

import "fmt"

type Cloneable interface {
	clone() Account
	cloneWithIban(iban string) Account
}

var normalAccount Account = Account{
	Balance:        1000,
	Credit:         1000,
	OpenToWithdraw: true,
	OpenToPayment:  true,
	OpenToTransfer: true,
}

var negativeAccount Account = Account{
	Balance:        -500,
	Credit:         1000,
	OpenToWithdraw: true,
	OpenToPayment:  true,
	OpenToTransfer: false,
}

var frozenAccount Account = Account{
	Balance:        -1000,
	Credit:         1000,
	OpenToWithdraw: false,
	OpenToPayment:  false,
	OpenToTransfer: false,
}

type Account struct {
	Iban           string
	Balance        float32
	Credit         float32
	OpenToWithdraw bool
	OpenToPayment  bool
	OpenToTransfer bool
}

func (a Account) clone() Account {
	return a
}

func (a Account) cloneWithIban(iban string) Account {
	a.Iban = iban
	return a
}

func (a Account) println() {
	fmt.Printf("\n Iban: %s \n Balance: %f \n Credit: %f \n OpenToWithdraw: %t \n OpenToPayment: %t \n OpenToTransfer: %t \n", a.Iban, a.Balance, a.Credit, a.OpenToWithdraw, a.OpenToPayment, a.OpenToTransfer)
}
