import re
import sys
import praw
import time
import json
import pprint
import operator
import datetime
import os
from dotenv import load_dotenv
from praw.models import MoreComments
from nltk.tokenize import sent_tokenize
from iexfinance.stocks import Stock as IEXStock

# to add the path for Python to search for files to use my edited version of vaderSentiment
sys.path.append('../sentimentService/')
from vader_sentiment import SentimentIntensityAnalyzer

load_dotenv()

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
def run(sub, ticker, post_limit,company, time_from):
   try:
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
         if not post.clicked:
            title = post.title
            body = post.selftext
            post_url = post.url
            comments = post.comments.list()[:100]
            for c in comments:
               if isinstance(c, MoreComments):
                  comments.remove(c)
            comments = [c.body for c in comments]
            all_text = {"title": title, "body": body, "comments": comments}
            sentiment, score = analyze_sentiment(all_text)
            # convet utc to human readable time
            post_time = datetime.datetime.fromtimestamp(post.created_utc).strftime('%Y-%m-%d %H:%M:%S')
            reddit_data_dict["feed"].append({
               "post_title": title,
               "body": body,
               "comments": comments,
               "post_time": post_time,
               "post_url": post_url,
               "overall_sentiment": sentiment,
               "overall_sentiment_score": score
            })
            ct+=1
            print(f"Processed {ct} posts")
      reddit_data_dict["items"] = len(reddit_data_dict["feed"])
      reddit_data_dict["overall_sentiment"], reddit_data_dict["overall_sentiment_score"] = calculate_ticker_sentiment(reddit_data_dict)
      
      return reddit_data_dict
   except Exception as e:
      print(f"Error in run: {e}")
      return None
         
def analyze_sentiment(post):
   try:
      analyzer = SentimentIntensityAnalyzer()
      pos_count = 0
      neg_count = 0
      neutral_count = 0 
      for sentence in sent_tokenize(post["title"] + " " + post["body"]):
         sentiment = analyzer.polarity_scores(sentence)
         if (sentiment["compound"] > .005) or (sentiment["pos"] > abs(sentiment["neg"])):
            pos_count += 1
         elif (sentiment["compound"] < -.005) or (abs(sentiment["neg"]) > sentiment["pos"]):
            neg_count += 1
         else:
            neutral_count += 1
      for comment in post["comments"]:
         for sentence in sent_tokenize(comment):
            sentiment = analyzer.polarity_scores(sentence)
            if (sentiment["compound"] > .005) or (sentiment["pos"] > abs(sentiment["neg"])):
               pos_count += 1
            elif (sentiment["compound"] < -.005) or (abs(sentiment["neg"]) > sentiment["pos"]):
               neg_count += 1
            else:
               neutral_count += 1
      sentiment_counts = {"bullish": pos_count, "bearish": neg_count, "neutral": neutral_count}
      max_sentiment = max(sentiment_counts, key=sentiment_counts.get)
      return max_sentiment, sentiment_counts[max_sentiment]
   except Exception as e:
      print(f"Error in analyze_sentiment: {e}")
      return None, None

def calculate_ticker_sentiment(data):
   try:
      total_count = 0
      ticker_sentiment = {"bullish": 0, "bearish": 0, "neutral": 0}
      for post in data["feed"]:
         sentiment = post["overall_sentiment"]
         ticker_sentiment[sentiment] += 1
         total_count += 1
      overall_sentiment = max(ticker_sentiment, key=ticker_sentiment.get)
      for key, value in ticker_sentiment.items():
         ticker_sentiment[key] = (value / total_count) * 100
      return ticker_sentiment, overall_sentiment
   except Exception as e:
      print(f"Error in calculate_ticker_sentiment: {e}")
      return None, None

if __name__ == "__main__":
   try:
      limit = 50
      sub = "wallstreetbets"
      ticker="NVDA"
      company = "Nvdia"
      time_from = "week"
      res = run(sub, ticker,limit,company, time_from)
      with open('reddit_data.json', 'w') as f: 
         json.dump(res, f, indent=4)
   except Exception as e:
      print(f"Error in main: {e}")