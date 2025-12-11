package order

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jkinyongo/tavern/internal/domain/customer"
	"github.com/jkinyongo/tavern/internal/domain/customer/memory"
	"github.com/jkinyongo/tavern/internal/domain/customer/mongo"
	"github.com/jkinyongo/tavern/internal/domain/product"
	prodmem "github.com/jkinyongo/tavern/internal/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.Repository
	products  product.Repository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	//Loop through all the Cfgs and apply them
	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository applies a customer repository to the OrderService
func WithCustomerRepository(cr customer.Repository) OrderConfiguration {
	// return a function that matches the orderConfiguration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, connStr string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, connStr)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	// Fetch the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}
	// Get each Product
	var products []product.Product
	var total float64

	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("Customer : %s has ordered %d products", c.GetID(), len(products))
	return total, nil
}

func (o *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}

	err = o.customers.Add(c)
	if err != nil {
		return uuid.Nil, err
	}
	return c.GetID(), nil
}
