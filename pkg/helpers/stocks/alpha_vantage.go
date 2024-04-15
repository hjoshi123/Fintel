package helpers

type AlphaVantageNewsResponse struct {
	Items                    string      `json:"items,omitempty"`
	SentimentScoreDefinition string      `json:"sentiment_score_definition,omitempty"`
	RelevanceScoreDefinition string      `json:"relevance_score_definition,omitempty"`
	Feed                     []StockFeed `json:"feed,omitempty"`
	Ticker                   string      `json:"ticker,omitempty"`
}

type StockFeed struct {
	Title                string `json:"title,omitempty"`
	URL                  string `json:"url,omitempty"`
	TimePublished        string `json:"time_published,omitempty"`
	Summary              string `json:"summary,omitempty"`
	BannerImage          string `json:"banner_image,omitempty"`
	Source               string `json:"source,omitempty"`
	CategoryWithinSource string `json:"category_within_source,omitempty"`
	SourceDomain         string `json:"source_domain,omitempty"`
	Topics               []struct {
		Topic          string `json:"topic,omitempty"`
		RelevanceScore string `json:"relevance_score,omitempty"`
	} `json:"topics,omitempty"`
	OverallSentimentScore float64 `json:"overall_sentiment_score,omitempty"`
	OverallSentimentLabel string  `json:"overall_sentiment_label,omitempty"`
	TickerSentiment       []struct {
		Ticker               string `json:"ticker,omitempty"`
		RelevanceScore       string `json:"relevance_score,omitempty"`
		TickerSentimentScore string `json:"ticker_sentiment_score,omitempty"`
		TickerSentimentLabel string `json:"ticker_sentiment_label,omitempty"`
	} `json:"ticker_sentiment,omitempty"`
}
