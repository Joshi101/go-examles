package main

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"prics"`
}

var Products []Product

func GetProducts() ([]Product, error) {
	Products = []Product{{1, "Pen", 1, 100.0},
		{2, "Pencil", 2, 10.00},
		{3, "NoteBook", 2, 100.0}}

	return Products, nil

}
