package main

type InvoiceGateway interface {
	makeInvoice(info string)
}
