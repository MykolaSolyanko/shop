package repository

/*
Implement a shopping application in Go that allows customers to make purchases.
Instead of using a traditional database, use a map as a cache to store product information and customer data.

The application should have the following features:
Customers should be able to browse a list of available products and view their details, including the name, price, and quantity in stock.
Customers should be able to add products to their shopping cart and view the contents of their cart.
Customers should be able to checkout and complete their purchases, reducing the quantity of the purchased items in stock and updating
their customer information.

The application should enforce a maximum quantity limit for each product
and should prevent customers from purchasing more than the available quantity.

To implement the application, create a Product struct with the following fields:
Name string: the name of the product.
Price float64: the price of the product.
Quantity int: the number of units of the product in stock.

Create a map to store the product information, with the product name as the key and the Product struct as the value.

Create a Customer struct with the following fields:
Name string: the name of the customer.
Email string: the email address of the customer.
Cart map[string]int: a map to store the products in the customer's shopping cart, with the product name as the key and the quantity as the value.

Create a map to store the customer information, with the customer email as the key and the Customer struct as the value.

Implement the following methods for the application:
ViewProducts(): displays a list of available products with their details.
ViewProductDetails(productName string): displays the details of a specific product.
AddToCart(customerEmail string, productName string, quantity int): adds a product to a customer's cart.
ViewCart(customerEmail string): displays the contents of a customer's cart.
Checkout(customerEmail string): completes a customer's purchase and updates the product quantities and customer information.
To use the map-based cache in the application, store the product and customer information in a map instead of a traditional database.
*/

import (
	"errors"

	"github.com/MykolaSolyanko/shop/types"
)

type Repository struct {
	products  map[string]types.Product
	customers map[string]types.Customer
}

func New() *Repository {
	return &Repository{
		products:  make(map[string]types.Product),
		customers: make(map[string]types.Customer),
	}
}

var (
	ErrProductNotFound = errors.New("product not found")
	ErrCustomerNotFound = errors.New("customer not found")
	ErrNotEnoughQuantity = errors.New("not enough quantity")
)


func (r *Repository) ViewProducts() []types.Product {
	products := make([]types.Product, 0, len(r.products))

	for _, product := range r.products {
		products = append(products, product)
	}

	return products
}

func (r *Repository) ViewProductDetails(productName string) (types.Product, error) {
	product, ok := r.products[productName]
	if !ok {
		return types.Product{}, ErrProductNotFound
	}

	return product, nil
}

func (r *Repository) AddToCart(customerEmail string, productName string, quantity int) error {
	customer, ok := r.customers[customerEmail]
	if !ok {
		return ErrCustomerNotFound
	}

	product, ok := r.products[productName]
	if !ok {
		return ErrProductNotFound
	}

	if quantity > product.Quantity {
		return ErrNotEnoughQuantity
	}

	customer.Cart[productName] = quantity
	product.Quantity -= quantity

	r.customers[customerEmail] = customer
	r.products[productName] = product

	return nil
}

func (r *Repository) ViewCart(customerEmail string) ([]types.Order, error) {
	customer, ok := r.customers[customerEmail]
	if !ok {
		return nil, ErrCustomerNotFound
	}

	orders := make([]types.Order, 0, len(customer.Cart))

	for productName, quantity := range customer.Cart {
		orders = append(orders, types.Order{
			Product:  productName,
			Quantity: quantity,
		})
	}

	return orders, nil
}