package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content  string `json:"content" form:"content" db:"content" binding:"required,min=1"`
	ParentId string `json:"parent_id" form:"parent_id" db:"parent_id" binding:"-"`
	Ip       string `json:"ip" db:"ip"`
	Agent    string `json:"agent" db:"agent"`
}

type Ticker struct {
	ID         string  `json:"-" db:"id, primarykey, autoincrement"`
	Symbol     string  `json:"symbol" db:"symbol"`
	Price      float64 `json:"price" db:"last_price"`
	ExchangeId string  `json:"-" db:"exchange_id"`
	PairId     string  `json:"-" db:"pair_id"`
	UpdatedAt  string  `json:"updated_at" db:"updated_at"`
	CreatedAt  string  `json:"-" db:"created_at"`
}
