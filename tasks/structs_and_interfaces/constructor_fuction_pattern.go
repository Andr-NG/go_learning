// Write NewProduct(name string, price float64) Product
// Validate input (price must be > 0)
// Decide: return Product or *Product — why?

package main

import "errors"

type Product struct {
	ID    int
	Name  string
	Price float64
}

func NewProduct(name string, price float64) (*Product, error) {

	if price <= 0 {
		return nil, errors.New("price must be greater than 0")
	}
	
	// returning a reference to a copy of a Product instead of the copy itself
	return &Product{
		Name:  name,
		Price: price,
	}, nil

}
