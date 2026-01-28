package models

type CartItem struct {
	ProductID int
	Quantity  float64
}

type Cart struct {
	UserID int
	Items  []CartItem
}
