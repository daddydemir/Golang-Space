package main

func main() {

	macbookPro := Product{
		price: 5000,
	}

	var demir User
	var darth User

	demir.name = "Demir"
	darth.name = "Darth"

	macbookPro.addDiscountAlertUser(demir)
	macbookPro.addDiscountAlertUser(darth)

	macbookPro.updatePrice(200)
}
