package structs

type Book struct {
	ID       string  `json:"id" validate:"required"`
	Title    string  `json:"title" validate:"required"`
	Author   string  `json:"author" validate:"required"`
	Quantity uint    `json:"quantity" validate:"required"`
	Price    float64 `json:"price" validate:"required,gt=0"`
}
