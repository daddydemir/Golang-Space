package main

type AgencyInvoice struct{}

func (a *AgencyInvoice) createInvoice(agencies []string) {
	println("Agency Invoice is created. agency count:", len(agencies))
}

func (a *AgencyInvoice) makeInvoice(info string) {
	agencies := make([]string, 1)
	agencies = append(agencies , info)
	
	a.createInvoice(agencies)
}
