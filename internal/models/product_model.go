package models

import "github.com/shopspring/decimal"

type Product struct {
	ID          int
	Name        string
	Description string
	PhotoURL    string
	Price       decimal.Decimal
}
