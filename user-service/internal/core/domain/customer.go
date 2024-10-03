package domain

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserName string
	Type     string

	CustomerId string
}

type Customer struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        string
	MobileNumber string
	Address      string

	CountryId string
}

type Country struct {
	gorm.Model
	Name string
	Code string
}

type Category struct {
	gorm.Model
	Name        string
	Description string
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Image       []string
	StockQty    int
	// Code 		string
	// Cost float64

	CategoryID int

	UserId         int
	MerChantId     int
	ProductModelID int
}

type MerChant struct {
	gorm.Model
	Name        string
	Description string
	Email       string
	Mobile      string
	Address     string
}

type ProductModel struct {
	gorm.Model
	Name        string
	Description string
}

type Cart struct {
	gorm.Model
	TotalAmount float64
	Items       []CartItem

	UserID int
}

type CartItem struct {
	gorm.Model
	Quantity int

	CartID    int
	ProductID int
}

type Order struct {
	gorm.Model
	Status          string
	PaymentMethod   string
	ShippingMethod  string
	ShippingAddress string
	DeliveryDate    string
	TotalAmount     float64
	Items           []OrderDetail

	UserID int
}

type OrderDetail struct {
	gorm.Model
	Quantity int
	Price    float64

	ProductID int
}
