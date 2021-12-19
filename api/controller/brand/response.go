package brand

type GetBrandResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Banner string `json:"banner"`
}
