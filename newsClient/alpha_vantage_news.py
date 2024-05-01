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

def fetch_alpha_vantage_news(ticker, time_from,time_to, limit):
    """
    Fetch news data from Alpha Vantage for a specific ticker.

    Args:
        ticker (str): The stock symbol for which to fetch news.
        time_from (str): The starting date from which to fetch news.
        limit (int): The maximum number of news items to fetch.
    """
   # url = 'https://www.alphavantage.co/query?function=NEWS_SENTIMENT&tickers=AAPL&apikey=demo'
    # The URL to fetch data from Alpha Vantage API
    url = f'https://www.alphavantage.co/query?function=NEWS_SENTIMENT&tickers={ticker}&time_from={time_from}&time_to={time_to}&limit=1000&sort=LATEST&apikey={os.getenv("ALPHAVANTAGE_API_KEY")}'

    session = Session()  # Creating a new session to manage connections

    try:
        # Attempt to send a GET request to the Alpha Vantage API
       print(f"Sending request to {url} for time_from {time_from}")
       response = requests.get(url)
       #print(response.json())
        # Return the JSON response if the request was successful
       if response.status_code == 200:
           return response.json()

    except Exception as e:
        print(f"Request failed with exception {e}")
    finally:
        session.close()

def ingest_news_to_kafka(ticker, limit):
    print(f"Ingesting news for {ticker}...")
    # now = datetime.now()
    # thirty_days_ago = now - timedelta(days=30)
    # time_to = now.strftime("%Y%m%dT0000")
    # time_from = thirty_days_ago.strftime("%Y%m%dT0000")
    time_from="20240401T0000"
    time_to="20240411T0000"
    # Fetch news data
    news_data = fetch_alpha_vantage_news(ticker, time_from,time_to, limit)
    # put the news in a json file
    with open('msft.json', 'w') as f:
        json.dump(news_data, f)        

    # if news_data is not None:
    #     # Add the ticker to the news data
    #     news_data["ticker"] = ticker 
    #     # Produce a new message to the Kafka topic 'stocks.news.create'
    #     producer = Producer(config)
    #     producer.produce("stocks.news.create", json.dumps(news_data))
    #     print(f"News for {ticker} ingested successfully!")
    #     producer.flush()
def ingest_news_from_file(ticker):
    # store the news in a json file
    # read from news_GOOG.json
    # loop through the .json files in the data folder
    os.chdir("data")
    for file in os.listdir():
        if file.endswith('.json'):  # make sure the file is a JSON file
            with open(file, 'r') as f:
                file_contents = json.load(f)
            producer = Producer(config)
            producer.produce("stocks.news.create", json.dumps(file_contents))
            print(f"News from file {file} ingested successfully!")
            producer.flush()
            print(f"News for {ticker} ingested successfully!")

def main():
    # Tickers for which to fetch news
    ticker = "MSFT"
    # The maximum number of news items to fetch
    ingest_news_to_kafka(ticker,1000)
    #ingest_news_from_file("MSFT")
    time.sleep(5)  # Sleep for 5 seconds before fetching news for the next ticker
   
if __name__ == "__main__":
    main()