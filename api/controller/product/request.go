package product

type PostProductRequest struct {
	Name    string  `json:"name"`
	Picture string  `json:"picture"`
	Price   float32 `json:"price"`
}

type EditProductRequest struct {
	Name    string  `json:"name"`
	Picture string  `json:"picture"`
	Price   float32 `json:"price"`
}
