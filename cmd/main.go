package main

import (
	"github.com/google/uuid"
	"github.com/jkinyongo/tavern/internal/domain/product"
	"github.com/jkinyongo/tavern/internal/services/order"
	"github.com/jkinyongo/tavern/internal/services/tavern"
)

func main() {
	products := productInventory()
	os, err := order.NewOrderService(
		//order.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}
	tavern, err := tavern.NewTavern(
		tavern.WithOrderService(os),
	)
	if err != nil {
		panic(err)
	}
	uid, err := os.AddCustomer("Jack")
	if err != nil {
		panic(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy beverage", 4.45)
	if err != nil {
		panic(err)
	}
	peenuts, err := product.NewProduct("Peanuts", "Snacks", 2.34)
	if err != nil {
		panic(err)
	}
	wine, err := product.NewProduct("Wine", "nasty drink", 3.43)
	if err != nil {
		panic(err)
	}
	return []product.Product{
		beer, peenuts, wine,
	}
}
