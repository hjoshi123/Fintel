package models

type PriceUpdate struct {
	Ticker   string `json:"ticker"`
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}
