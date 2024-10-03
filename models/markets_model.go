package models

type Markets struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Description string  `json:"description"`
	OwnerName   string  `json:"owner_name"`
	RSType      string  `json:"rs_type"`
	Size        float64 `json:"size"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"`
}
