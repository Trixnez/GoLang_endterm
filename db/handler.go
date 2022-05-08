package db

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func GetAllUsers() []User {
	var users []User
	DB.Find(&users)
	return users
}

func GetUser(id uint) User {
	var user User
	if err := DB.Where("id=?", id).First(&user).Error; err != nil {
		return User{}
	}
	return user
}

func CreateUser(input CreateUserInput) {
	user := User{Name: input.Name, Surname: input.Surname, Username: input.Username}
	DB.Create(&user)
}

type CreateCardInput struct {
	Balance float32 `json:"balance" binding:"required"`
	UserId  uint    `json:"UserId" binding:"required"`
}

func GetCard(id uint) Card {
	var card Card
	if err := DB.Where("id=?", id).First(&card).Error; err != nil {
		return Card{}
	}
	return card
}

func CreateCard(input CreateCardInput) {
	card := Card{Balance: input.Balance, UserID: input.UserId}
	DB.Create(&card)
}

type CreateTypeInput struct {
	Name string `json:"name" binding:"required"`
}

func GetAllTypes() []Type {
	var types []Type
	DB.Find(&types)
	return types
}

func GetType(id uint) Type {
	var _type Type
	if err := DB.Where("id=?", id).First(&_type).Error; err != nil {
		return Type{}
	}
	return _type
}

func CreateType(input CreateTypeInput) {
	_type := Type{Name: input.Name}
	DB.Create(&_type)
}

type CreateProductInput struct {
	Name   string  `json:"name" binding:"required"`
	Price  float32 `json:"price" binding:"required"`
	Amount int     `json:"amount" binding:"required"`
	TypeID uint    `json:"TypeID" binding:"required"`
	ShopID uint    `json:"ShopID" binding:"required"`
}

func GetAllProducts() []Product {
	var products []Product
	DB.Find(&products)
	return products
}

func GetProduct(id uint) Product {
	var product Product
	if err := DB.Where("id=?", id).First(&product).Error; err != nil {
		return Product{}
	}
	return product
}

func CreateProduct(input CreateProductInput) {
	product := Product{Name: input.Name, Price: input.Price, Amount: input.Amount, TypeID: input.TypeID, ShopID: input.ShopID}
	DB.Create(&product)
}

type CreateShopInput struct {
	Name string `json:"name" binding:"required"`
}

func GetAllShops() []Shop {
	var shops []Shop
	DB.Find(&shops)
	return shops
}

func GetShop(id uint) Shop {
	var shop Shop
	if err := DB.Where("id=?", id).First(&shop).Error; err != nil {
		return Shop{}
	}
	return shop
}

func CreateShop(input CreateShopInput) {
	shop := Shop{Name: input.Name}
	DB.Create(&shop)
}
