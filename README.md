# FinTel

Fintel is a project designed to harness the vast amounts of unstructured data from financial news sources and social media platforms like Reddit. This project aims to aggregate and analyze the discussions and opinions expressed about stocks across various companies, providing real-time sentiment analysis. The tool categorizes sentiments into positive, negative, and neutral. We calculate the ICI score from both news and social media data. The ICI score is a measure of the overall sentiment of the stock which is calculated as:

`ln((1 +positive sentiment count) / (1- negative sentiment count))`

Using this ICI score, we calculate the pearson correlation coefficient between the ICI score and the stock price. This correlation coefficient is used to determine if the stock price follows the sentiment of the news or social media data.

## Frontend Repo

https://github.com/rakshitha31/Fintel-Frontend

## Architecture

![FinTel](<./img/Big%20Data%20(1).png>)

Our backend architecture consists of 3 components:

- Backend API
- Backend Pubsub
- Microservices to fetch data from news and social media sources.

The microservices are responsible to collect data from news and social media and insert into the Kafka topics (`stocks.create.news.`, `stocks.create.social.`). This was done to avoid the synchronous nature of the API calls and to make the system more scalable.
As the data is inserted into the Kafka topics, the backend pubsub service consumes the data and processes it. The processed data is then inserted into the database. The backend API is responsible for fetching the data from the database and serving it to the frontend.

For pubsub, we used Kafka with Sarama library in Go. The pubsub creates consumer group based on regex pattern i.e. stocks.create.news. and stocks.create.social. will be in a single consumer group. We also used Kafka's transactional feature to ensure that data is not written to the queue if Kafka is not ready and consumer doesn't consume the data if the data is not committed.

We created a Helm chart in `dochart` folder to deploy our backend and frontend services to DigitalOcean Kubernetes cluster. We used DigitalOcean's managed Kubernetes service to deploy our services. For the JSON Api, we used the JSONAPI spec to create the API endpoints to ensure our response is standardized.

## Technologies Used

- Apache Kafka
- K8S
- Helm
- Go
- DigitalOcean

## Jira Board

https://bigdataarch.atlassian.net/jira/software/projects/BIG/boards/1/backlog?epics=visible&issueParent=10002

## Link to the project

https://finsent.hjoshi.me/

## Team Members

- Hemant Joshi
- Nagarakshitha Bangalore Ramu
- Saurabh Shetty
- Matthew McDermott
