package main

type CustomerInvoice struct{}

func (c *CustomerInvoice) createInvoice(customerName string) {
	println("Customer Invoice is created. customer anme: " , customerName)
}

func (c *CustomerInvoice) makeInvoice(info string) {
	c.createInvoice(info)
}
