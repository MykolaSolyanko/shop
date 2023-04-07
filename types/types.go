package types

type Product struct {
	Name string
	Price float64
	Quantity int
}

type Customer struct {
	Name string
	Email string
	Cart map[string] int
}

type Order struct {
	Product string
	Quantity int
}
