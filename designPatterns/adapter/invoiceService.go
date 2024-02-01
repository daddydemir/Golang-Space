package main

type InvoiceService struct {
	invoiceGateway InvoiceGateway
}

func (i *InvoiceService) createInvoice(info string)  {
	i.invoiceGateway.makeInvoice(info)
}