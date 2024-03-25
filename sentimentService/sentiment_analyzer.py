from transformers import BertTokenizer, BertForSequenceClassification
from summarizer import Summarizer
import torch

def load_model_and_tokenizer():
    # Load FinBERT tokenizer and model
    tokenizer = BertTokenizer.from_pretrained('yiyanghkust/finbert-tone', do_lower_case=True)
    model = BertForSequenceClassification.from_pretrained('yiyanghkust/finbert-tone')
    return tokenizer, model

def summarize_article(text, ratio=0.2):
    # Initialize the BERT-based summarizer
    summarizer = Summarizer()
    
    # Summarize the article
    summary = summarizer(text, ratio=ratio)
    return summary

def analyze_sentiment(text, tokenizer, model):
    inputs = tokenizer(text, return_tensors="pt", max_length=512, truncation=True)
    outputs = model(**inputs)
    logits = outputs.logits
    probabilities = torch.softmax(logits, dim=1).detach().numpy()[0]
    positive_prob = probabilities[1]  # probability of being positive
    return positive_prob

def analyze_all_articles(articles, threshold=0.3):
    tokenizer, model = load_model_and_tokenizer()

    sentiment_scores = []

    for idx, article in enumerate(articles, 1):
        try:
            headline = article.get("text")
            if not headline:
                print(f"Skipping article {idx}: No headline provided.")
                continue
            
            sentiment_score = analyze_sentiment(headline, tokenizer, model)
            summary = summarize_article(headline)
            sentiment_scores.append(sentiment_score)
        except Exception as e:
            print(f"Error analyzing sentiment for headline '{headline}': {e}")
            sentiment_scores.append(0)

    # Determine overall sentiment score
    overall_sentiment_score = sum(sentiment_scores) / len(sentiment_scores) if sentiment_scores else None
    
    # Determine overall sentiment label
    if overall_sentiment_score is not None:
        overall_sentiment_label = "Positive" if overall_sentiment_score >= threshold else "Negative"
    else:
        overall_sentiment_label = "Unknown"

    return overall_sentiment_score, overall_sentiment_label, summary

