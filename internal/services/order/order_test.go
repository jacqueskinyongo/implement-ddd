package order

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jkinyongo/tavern/internal/domain/customer"
	"github.com/jkinyongo/tavern/internal/domain/product"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy beverage", 4.45)
	if err != nil {
		t.Fatal(err)
	}
	peenuts, err := product.NewProduct("Peanuts", "Snacks", 2.34)
	if err != nil {
		t.Fatal(err)
	}
	wine, err := product.NewProduct("Wine", "nasty drink", 3.43)
	if err != nil {
		t.Fatal(err)
	}
	return []product.Product{
		beer, peenuts, wine,
	}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}
	cust, err := customer.NewCustomer("Jack")
	if err != nil {
		t.Error(err)
	}
	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
