package main

import (
	"context"
	"errors"
	"log"
	"unary/proto/message"
	"unary/proto/service"
)

var products = []message.Response{
	{
		ProductName:  "Smartphone",
		Brand:        "XPhone",
		Model:        "X10",
		Price:        2999.00,
		Availability: true,
		Feature: &message.Feature{
			ScreenSize:      "6.5 inches",
			Resolution:      "1080x2340 pixels",
			Processor:       "Snapdragon 865",
			Ram:             "8 GB",
			InternalMemory:  "128 GB",
			Camera:          "Rear: 64 MP, Front: 32 MP",
			BatteryCapacity: "5000 mAh",
		},
	},
	{
		ProductName:  "Laptop",
		Brand:        "TechBook",
		Model:        "Pro 2024",
		Price:        1799.00,
		Availability: true,
		Feature: &message.Feature{
			ScreenSize:      "15.6 inches",
			Resolution:      "1920x1080 pixels",
			Processor:       "Intel Core i7-1165G7",
			Ram:             "16 GB",
			InternalMemory:  "512 GB SSD",
			Camera:          "Front: 8 MP",
			BatteryCapacity: "Up to 8 hours",
		},
	},
}

type Server struct {
	service.UnimplementedUnaryServiceServer
}

func (s Server) GetProduct(ctx context.Context, request *message.Request) (*message.Response, error) {

	log.Printf("REQUEST: %v \n", request.Values)

	if request.Values > 1 {
		return nil, errors.New("product not found")
	} else {
		return &products[request.Values], nil
	}
}
