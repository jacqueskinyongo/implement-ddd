package tavern

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jkinyongo/tavern/internal/domain/product"
	"github.com/jkinyongo/tavern/internal/services/order"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)
	os, err := order.NewOrderService(
		//WithMemoryCustomerRepository(),
		order.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	uid, err := os.AddCustomer("Jack")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}
	err = tavern.Order(uid, order)
	if err != nil {
		t.Fatal(err)
	}
}

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
