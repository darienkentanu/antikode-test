package outlet

type PostOutletRequest struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Address string `json:"address"`
}

type EditOutletRequest struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Address string `json:"address"`
}
