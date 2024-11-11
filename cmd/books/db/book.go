package db

type Book struct {
	ID     int     `gorm:"primarykey;autoincrement" json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}
