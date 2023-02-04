package entity

type Coffee struct {
	ID    uint    `gorm:"primary_key"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}
