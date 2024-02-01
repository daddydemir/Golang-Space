package main

func main() {
	
	agencyInvoiceService := InvoiceService{invoiceGateway: &AgencyInvoice{}}
	customerInvoiceService := InvoiceService{invoiceGateway: &CustomerInvoice{}}
	
	agencyInvoiceService.createInvoice("test")
	customerInvoiceService.createInvoice("test")

}
