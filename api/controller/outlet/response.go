package outlet

type GetOutletResponse struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Picture   string  `json:"picture"`
	Address   string  `json:"address"`
	Longitute string  `json:"longitude"`
	Latitute  string  `json:"latitute"`
	Distance  float32 `json:"distance"`
}
