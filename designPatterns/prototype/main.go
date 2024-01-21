package main

import "fmt"

func main() {

	var normalAccount Account = Account{
		Iban:           "1234",
		Balance:        1000,
		Credit:         1000,
		OpenToWithdraw: true,
		OpenToPayment:  true,
		OpenToTransfer: true,
	}

	//normalAccount.println()

	var negativeAccount Account = normalAccount.clone()
	negativeAccount.Iban = "2345"
	negativeAccount.Balance = -300
	negativeAccount.OpenToTransfer = false

	//normalAccount.println()
	//negativeAccount.println()

	frozen := frozenAccount.clone()
	frozen.Iban = "1234567"

	fmt.Println("Frozen ")
	frozen.println()

	frozen2 := frozenAccount.cloneWithIban("5823705828129841")
	fmt.Println("Frozen ")
	frozen2.println()

}
