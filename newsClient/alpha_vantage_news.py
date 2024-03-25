import os
from confluent_kafka import Producer
import json
from datetime import datetime, timedelta
import requests
import time
from requests import Session
from dotenv import load_dotenv

load_dotenv()

config = {
    'bootstrap.servers': os.getenv('BOOTSTRAP_SERVERS'),
    'security.protocol': 'SASL_SSL',
    'sasl.mechanism': 'PLAIN',
    'sasl.username': os.getenv('CLOUD_KEY'),
    'sasl.password': os.getenv('CLOUD_SECRET')
}

def fetch_alpha_vantage_news(ticker, time_from, limit):
    """
    Fetch news data from Alpha Vantage for a specific ticker.

    Args:
        ticker (str): The stock symbol for which to fetch news.
        time_from (str): The starting date from which to fetch news.
        limit (int): The maximum number of news items to fetch.
    """
   # url = 'https://www.alphavantage.co/query?function=NEWS_SENTIMENT&tickers=AAPL&apikey=demo'
    # The URL to fetch data from Alpha Vantage API
    url = f'https://www.alphavantage.co/query?function=NEWS_SENTIMENT&tickers={ticker}&time_from={time_from}&limit={limit}&apikey={os.getenv("ALPHAVANTAGE_API_KEY")}'

    session = Session()  # Creating a new session to manage connections

    try:
        # Attempt to send a GET request to the Alpha Vantage API
       print(f"Sending request to {url} for time_from {time_from}")
       response = requests.get(url)
       #print(response.json())
        # Return the JSON response if the request was successful
    #    if response.status_code == 200:
    #        return response.json()

    except Exception as e:
        print(f"Request failed with exception {e}")
    finally:
        session.close()

def ingest_news_to_kafka(ticker, limit):
    print(f"Ingesting news for {ticker}...")
    now = datetime.now()
    yesterday = now - timedelta(days=1)
    print(yesterday.strftime("%Y%m%dT0000"))
#    The time range of the news articles you are targeting, in YYYYMMDDTHHMM format. I want todays news
    time_from = yesterday.strftime("%Y%m%dT0000")
    # Fetch news data
    news_data = fetch_alpha_vantage_news(ticker, time_from, limit)
    if news_data is not None:
        # Add the ticker to the news data
        news_data["ticker"] = ticker 
        # Produce a new message to the Kafka topic 'stocks.news.create'
        producer = Producer(config)
        producer.produce("stocks.news.create", json.dumps(news_data))
        print(f"News for {ticker} ingested successfully!")
            #producer.flush()
def ingest_news_from_file():
    # store the news in a json file
    # read from news_GOOG.json
    with open('news_GOOG.json', 'r') as f:
        data = json.load(f)
    # Produce a new message to the Kafka topic 'stocks.news.create'
    producer = Producer(config)
    producer.produce("stocks.news.create", json.dumps(data))
    print(f"News from file ingested successfully!")
    #producer.flush()
    #print(f"News for {ticker} ingested successfully!")

def main():
    # Tickers for which to fetch news
    ticker = "AAPL"
    # The maximum number of news items to fetch
    limit = 100
    ingest_news_to_kafka(ticker, limit)
    time.sleep(5)  # Sleep for 5 seconds before fetching news for the next ticker
   
if __name__ == "__main__":
    main()