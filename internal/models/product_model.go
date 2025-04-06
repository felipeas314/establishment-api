package models

import "github.com/shopspring/decimal"

type Product struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	PhotoURL    string          `json:"photoURL"`
	Price       decimal.Decimal `json:"price"`
	CategoryID  int             `json:"category_id"`
}
