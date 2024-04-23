package models

type StockSentimentInfo struct {
	PositiveCount int `json:"positive_count" jsonapi:"attr,positive_count"`
	NegativeCount int `json:"negative_count" jsonapi:"attr,negative_count"`
	NeutralCount  int `json:"neutral_count" jsonapi:"attr,neutral_count"`
}
