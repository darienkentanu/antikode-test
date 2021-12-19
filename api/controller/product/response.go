package product

type GetProductResponse struct {
	ID      uint    `json:"id"`
	Name    string  `json:"name"`
	Picture string  `json:"picture"`
	Price   float32 `json:"price"`
}
