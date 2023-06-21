package main

import "fmt"

/*
Ürün ve Müşteriler olacak, ürünün fiyatı azalınca müşteriye bildirim gidecek
*/

type Product struct {
	price         float32
	discountAlert []User
}

func (p *Product) addDiscountAlertUser(user User) {
	p.discountAlert = append(p.discountAlert, user)
	fmt.Println(user.name, " ürün fiyatı azalınca bilgilendirileceksiniz")
}

func (p *Product) removeDiscountAlertUser(user User) {
	for i := 0; i < len(p.discountAlert); i++ {
		if p.discountAlert[i] == user {
			p.discountAlert = append(p.discountAlert[:i], p.discountAlert[i+1:]...)
			break
		}
	}
}

func (p *Product) updatePrice(newPrice float32) {
	if p.price > newPrice {
		p.price = newPrice
		// save to db
		for _, u := range p.discountAlert {
			u.createPublisherMessage()
		}
	} else {
		p.price = newPrice
	}
}

type User struct {
	name string
}

func (u *User) createPublisherMessage() {
	// sendNotification
	fmt.Println(u.name + " favorine eklediğin ürün şu anda indirimde.")
}
