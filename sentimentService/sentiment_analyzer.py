"""Simple spacy & vaderSentiment based analyzer."""

# import necessary libraries:
import spacy, string, en_core_web_sm
import pandas as pd
from vaderSentiment import vaderSentiment

# pip install spacy vaderSentiment

class SentimentAnalyzerService:
    def __init__(self):
        self.english = spacy.load("en_core_web_sm")  # load spacy lang
        self.analyzer = vaderSentiment.SentimentIntensityAnalyzer()  # create analyzer using vaderSentiment

    def analyze_all_articles(self, article_dicts):
        for article in article_dicts:
            if isinstance(article, dict):
                print("Performing sentiment analysis on reddit posts and comments....")

                # analyze & store title
                title_sent = self.analyze(article.get("title", ""))
             
                # analyze & store description
                body_sent = self.analyze(article.get("body", ""))
                avg_body_sentiment = self.calculate_average_sentiment(body_sent)
            
                # analyze & store actual article text
                comments_sent = [self.analyze(comment) for comment in article.get("comments", [])]
                average_comment_sentiment = self.calculate_average_sentiment(comments_sent[0])
           

                # calculate the average sentiment scores
                total_sentiment, label = self.get_ticker_average_sentiment(title_sent[0], avg_body_sentiment, average_comment_sentiment)
                

        return total_sentiment, label

    def analyze(self, text):
        if isinstance(text, list):
            # If text is a list, analyze each comment separately and return a list of results
            return [self.analyze(comment) for comment in text]
        else:
            # If text is not a list, proceed as before
            english = spacy.load("en_core_web_sm") 
            result = english(text)
            sentences = [str(s) for s in result.sents]  # go thru sentences
            analyzer = vaderSentiment.SentimentIntensityAnalyzer()  # create analyzer using vaderSentiment
            sentiment = [analyzer.polarity_scores(str(s)) for s in sentences]
            return sentiment
    
    def calculate_average_sentiment(self, sentiment_list):
        print("Calculating the average sentiment score for the comments....")
        # If sentiment_list is empty, return a neutral sentiment
        if not sentiment_list:
            return {'neg': 0.0, 'pos': 0.0, 'neu': 0.0, 'compound': 0.0}
        # Calculate the average sentiment score
        avg_comment_sentiment = {
            'neg': sum(score['neg'] for score in sentiment_list) / len(sentiment_list),
            'pos': sum(score['pos'] for score in sentiment_list) / len(sentiment_list),
            'neu': sum(score['neu'] for score in sentiment_list) / len(sentiment_list),
            'compound': sum(score['compound'] for score in sentiment_list) / len(sentiment_list),
        }
        return avg_comment_sentiment
    
    def get_ticker_average_sentiment(self, title, body, avg_comments):
        # Calculate the average of each sentiment score across title, body, and comments
        print("Calculating the average sentiment score for the ticker....")
        total_sentiment = {
            'neg': (title['neg'] + body['neg'] + avg_comments['neg'])/3,
            'pos': (title['pos'] + body['pos'] + avg_comments['pos'])/3,
            'neu': (title['neu'] + body['neu'] + avg_comments['neu'])/3,
            'compound': (title['compound'] + body['compound'] + avg_comments['compound'])/3,
        }
        
        
        if total_sentiment['compound'] >= 0.0:
            label = 'bullish'
        elif total_sentiment['compound'] < 0.0:
            label = 'bearish'
        else:
            label = 'neutral'
        return total_sentiment, label
        
        
        
        