package main

type Coffee struct {
	ID    uint    `gorm:"primary_key"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

var AllCoffees = []Coffee{
	{1, "Cappucino", "Cappucino", 34},
	{2, "Caffè Americano", "Americano", 28},
	{3, "Caffè Misto", "Brewed", 30},
	{4, "Pistachio Latte", "Latte", 30},
}
