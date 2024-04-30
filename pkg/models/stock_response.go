package models

import "time"

type StockResponse struct {
	Ticker            string                `json:"ticker" jsonapi:"primary,stocks"`
	NewsCorrelation   float64               `json:"news_correlation" jsonapi:"attr,news_correlation"`
	SocialCorrelation float64               `json:"social_correlation" jsonapi:"attr,social_correlation"`
	NewsSentiment     []*Sentiment          `json:"news_sentiment" jsonapi:"relation,news_sentiment"`
	SocialSentiment   []*Sentiment          `json:"social_sentiment" jsonapi:"relation,social_sentiment"`
	TopContentsNews   []*TopContentResponse `json:"top_contents_news" jsonapi:"relation,top_contents_news"`
	TopContentsSocial []*TopContentResponse `json:"top_contents_social" jsonapi:"relation,top_contents_social"`
}

type Sentiment struct {
	ID            int       `json:"id" jsonapi:"primary,sentiments"`
	Date          time.Time `json:"date" jsonapi:"attr,date,iso8601"`
	DailyICI      float64   `json:"daily_ici" jsonapi:"attr,daily_ici"`
	Volume        int       `json:"volume" jsonapi:"attr,volume"`
	PositiveCount int       `json:"positive_count" jsonapi:"attr,positive_count"`
	NegativeCount int       `json:"negative_count" jsonapi:"attr,negative_count"`
	NeutralCount  int       `json:"neutral_count" jsonapi:"attr,neutral_count"`
}

type TopContentResponse struct {
	ID         int       `json:"id" jsonapi:"primary,top_contents"`
	URL        string    `json:"url" jsonapi:"attr,url"`
	PostedDate time.Time `json:"posted_date" jsonapi:"attr,posted_date,iso8601"`
	Title      string    `json:"title" jsonapi:"attr,title"`
	Summary    string    `json:"summary" jsonapi:"attr,summary"`
}
