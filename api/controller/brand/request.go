package brand

type PostBrandRequest struct {
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Banner string `json:"banner"`
}

type EditBrandRequest struct {
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Banner string `json:"banner"`
}
