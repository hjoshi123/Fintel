ALTER TABLE stock_sentiment ADD COLUMN info jsonb;
ALTER TABLE stock_sentiment ADD COLUMN price FLOAT NOT NULL DEFAULT 0;
