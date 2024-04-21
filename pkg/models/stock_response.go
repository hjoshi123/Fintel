package models

import "time"

type StockResponse struct {
	Ticker             string                `json:"ticker" jsonapi:"primary,stocks"`
	EffectiveSocialICI float64               `json:"effective_social_ici" jsonapi:"attr,effective_social_ici"`
	EffectiveNewsICI   float64               `json:"effective_news_ici" jsonapi:"attr,effective_news_ici"`
	NewsSentiment      []*Sentiment          `json:"news_sentiment" jsonapi:"relation,news_sentiment"`
	SocialSentiment    []*Sentiment          `json:"social_sentiment" jsonapi:"relation,social_sentiment"`
	TopContentsNews    []*TopContentResponse `json:"top_contents_news" jsonapi:"relation,top_contents_news"`
	TopContentsSocial  []*TopContentResponse `json:"top_contents_social" jsonapi:"relation,top_contents_social"`
}

type Sentiment struct {
	ID            int        `json:"id" jsonapi:"primary,sentiments"`
	Date          *time.Time `json:"date" jsonapi:"attr,sentiments"`
	DailyICI      float64    `json:"daily_ici" jsonapi:"attr,daily_ici"`
	Volume        int        `json:"volume" jsonapi:"attr,volume"`
	PositiveCount int        `json:"positive_count" jsonapi:"attr,positive_count"`
	NegativeCount int        `json:"negative_count" jsonapi:"attr,negative_count"`
}

type TopContentResponse struct {
	ID         int    `json:"id" jsonapi:"primary,top_contents"`
	URL        string `json:"url" jsonapi:"attr,url"`
	PostedDate string `json:"posted_date" jsonapi:"attr,posted_date"`
}