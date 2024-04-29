package finhistory

type HistoricalPrices struct {
	Symbol     string `json:"symbol"`
	Historical []struct {
		Date             string  `json:"date"`
		Open             float64 `json:"open"`
		High             float64 `json:"high"`
		Low              float64 `json:"low"`
		Close            float64 `json:"close"`
		AdjClose         float64 `json:"adjClose"`
		Volume           int     `json:"volume"`
		UnadjustedVolume int     `json:"unadjustedVolume"`
		Change           float64 `json:"change"`
		ChangePercent    float64 `json:"changePercent"`
		Vwap             float64 `json:"vwap"`
		Label            string  `json:"label"`
		ChangeOverTime   float64 `json:"changeOverTime"`
	} `json:"historical"`
}
