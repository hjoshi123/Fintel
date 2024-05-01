import azure.functions as func
import datetime
import json
import logging
from social_media_api import run

app = func.FunctionApp()

@app.route(route="socialmedia", auth_level="anonymous")
def socialmedia(req: func.HttpRequest) -> func.HttpResponse:
    logging.info('Python HTTP trigger function processed a request.')
    sub = req.params.get('sub')
    ticker = req.params.get('ticker')
    company = req.params.get('company')
    time_from = req.params.get('time_from')

    if not sub or not ticker or not company or not time_from:
        sub = "wallstreetbets"
        ticker="NVDA"
        company = "Nvidia"
        time_from = "day"

    result = run(sub, ticker, company, time_from)
    with open('reddit_data.json', 'w') as f: 
          json.dump(result, f, indent=4)
    if result is not None:
        return func.HttpResponse(json.dumps(result))
    else:
        return func.HttpResponse(
             "Error occurred while running the function",
             status_code=500
        )