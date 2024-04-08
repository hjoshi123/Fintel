CREATE TABLE IF NOT EXISTS stock_sentiment(
	id SERIAL PRIMARY KEY,
	ticker text NOT NULL,
	daily_ici float NOT NULL,
	chatter integer not null,
	created_at timestamp default CURRENT_TIMESTAMP,
	updated_at timestamp default CURRENT_TIMESTAMP
);
