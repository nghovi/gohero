package models

type Ticker struct {
    ID     string  `json:"-" db:"id, primarykey, autoincrement"`
    Symbol  string  `json:"symbol" db:"symbol"`
    Price  float64 `json:"price" db:"last_price"`
    ExchangeId string `json:"-" db:"exchange_id"`
    PairId string `json:"-" db:"pair_id"`
    UpdatedAt string `json:"updated_at" db:"updated_at"`
    CreatedAt string `json:"-" db:"created_at"`
}
