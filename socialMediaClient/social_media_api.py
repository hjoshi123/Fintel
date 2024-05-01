import re
import sys
import praw
import time
import json
import pprint
import operator
import datetime
import os
import pandas as pd
from dotenv import load_dotenv
from praw.models import MoreComments
from confluent_kafka import Producer
from nltk.tokenize import sent_tokenize
from pyfin_sentiment.model import SentimentModel
from iexfinance.stocks import Stock as IEXStock
load_dotenv()
# to add the path for Python to search for files to use my edited version of vaderSentiment
sys.path.append('../sentimentService/')
from sentiment_analyzer import SentimentAnalyzerService

config = {
    'bootstrap.servers': os.getenv('BOOTSTRAP_SERVERS'),
    'security.protocol': 'SASL_SSL',
    'sasl.mechanism': 'PLAIN',
    'sasl.username': os.getenv('CLOUD_KEY'),
    'sasl.password': os.getenv('CLOUD_SECRET')
}

def setup(sub):
   if sub == "":
      sub = "wallstreetbets"

   # create a reddit instance
   try:
      reddit = praw.Reddit(client_id=os.getenv("CLIENT_ID"), client_secret=os.getenv("CLIENT_SECRET"),
                           username=os.getenv("REDDIT_USERNAME"), password=os.getenv("REDDIT_PASSWORD"),
                           user_agent=os.getenv("REDDIT_USER_AGENT"))
      # create an instance of the subreddit
      subreddit = reddit.subreddit(sub)
   except Exception as e:
      print("Error: ", e)
   return subreddit



   # reddit_data_dict should be like this:
   # {
      # items:"length of feed array",
      # source: "reddit",      
      #feed:[
         # {
            # post_title: "title1",
            # body: "body",
            # comments: ["comment1", "comment2", "comment3"],
            # post_url: "url"
            # overall_sentiment: "sentiment"
            # overall_sentiment_score: "score"
            
         # },
         # {
            # post_title: "title2",
            # body: "body",
            # comments: ["comment1", "comment2", "comment3"],
            # post_url: "url"
            # overall_sentiment: "sentiment"
            #overall_sentiment_score: "score"
         # },..
     # ]
   # }
def run(sub, ticker,company, time_from):
   try:
      sentiment_analyzer_service = SentimentAnalyzerService()
      reddit_data_dict = {}
      reddit_data_dict["source"]="reddit"
      reddit_data_dict["feed"] = []
      subreddit = setup(sub)
      new_posts_ticker = list(subreddit.search(ticker, time_filter=time_from, sort='top'))
      new_posts_company = list(subreddit.search(company, time_filter=time_from, sort='top'))
      new_posts = new_posts_ticker + new_posts_company
      print(f"Processing posts from r/{sub} for {ticker} and {company}...")
      ct=0
      for post in new_posts:
         if not post.clicked and post.num_comments >= 10:
            title = post.title
            body = post.selftext
            post_url = post.url
            comments = []
            for c in post.comments[:150]:
               if isinstance(c, MoreComments):
                  continue
               comments.append(c.body)
            all_text = [{"title": title, "body": body, "comments": comments}]
            sentiment, label = sentiment_analyzer_service.analyze_all_articles(all_text)
            post_time = datetime.datetime.fromtimestamp(post.created_utc).strftime('%Y-%m-%d %H:%M:%S')
            ct+=1
            print(f"Processed {ct} posts")
            reddit_data_dict["feed"].append({
                  "post_title": title,
                  "body": body,
                  "comments": len(comments),
                  "post_time": post_time,
                  "post_url": post_url,
                  "overall_sentiment_score": sentiment,
                  "overall_sentiment": label,
                  "num_comments": post.num_comments,
            })
      reddit_data_dict["ticker"] = ticker
      reddit_data_dict["num_posts"] = len(new_posts)
      return reddit_data_dict
   except Exception as e:
      print(f"Error in run: {e}")
      return None
   
def ingest_reddit_data_to_kafka(data):
   try:
      producer = Producer(config)
      producer.produce("stocks.social.create", json.dumps(data))
      print(f"Reddit data ingested successfully!")
      producer.flush()
   except Exception as e:
      print(f"Error in ingest_reddit_data_to_kafka: {e}")


def ingest_reddit_data_to_kafka_from_file():
   try:
      os.chdir("data")
      for file in os.listdir():
         if file.endswith('.json'):  # make sure the file is a JSON file
            with open(file, 'r') as f:
               file_contents = json.load(f)
               producer = Producer(config)
               producer.produce("stocks.social.create", json.dumps(file_contents))
               print(f"Reddit data ingested successfully!")
               producer.flush()
   except Exception as e:
      print(f"Error in ingest_reddit_data_to_kafka: {e}")
      
if __name__ == "__main__":
   try:
      sub = "wallstreetbets"
      time_from = "month"
      ticker = "TSLA"
      company="Tesla"
      #res = run(sub, ticker,company, time_from)
      # read from a csv file
      tickers = pd.read_csv("../stocks/stocklist.csv")
         
   except Exception as e:
      print(f"Error in main: {e}")