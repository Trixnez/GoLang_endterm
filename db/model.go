package db

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username" gorm:"unique"`
	Card     Card
}

type Card struct {
	ID      uint    `json:"id" gorm:"primary_key"`
	Balance float32 `json:"balance"`
	UserID  uint    `json:"UserID" gorm:"<-:create"`
}

type Type struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Products []Product
}

type Product struct {
	ID     uint    `json:"id" gorm:"primary_key"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Amount int     `json:"amount"`
	TypeID uint    `json:"TypeID"`
	ShopID uint    `json:"ShopID"`
}

type Shop struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Products []Product
}
