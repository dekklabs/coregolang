package entities

//Product entidad de productos
type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Discount    float64 `json:"discount"`
	Image       string  `json:"image"`
	IDCategory  int64   `json:"idCategory"`
	Gallery     string  `json:"gallery"`
	CreatedAt   string  `json:"create_at"`
	UPdatedAt   string  `json:"updated_at"`
}
