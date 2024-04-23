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

type SocialSentiment struct {
	Source string `json:"source"`
	Feed   []struct {
		PostTitle             string `json:"post_title"`
		Body                  string `json:"body"`
		Comments              int    `json:"comments"`
		PostTime              string `json:"post_time"`
		PostURL               string `json:"post_url"`
		OverallSentimentScore struct {
			Neg      float64 `json:"neg"`
			Pos      float64 `json:"pos"`
			Neu      float64 `json:"neu"`
			Compound float64 `json:"compound"`
		} `json:"overall_sentiment_score"`
		OverallSentiment string `json:"overall_sentiment"`
		NumComments      int    `json:"num_comments"`
	} `json:"feed"`
	Items  int    `json:"items"`
	Ticker string `json:"ticker"`
}
