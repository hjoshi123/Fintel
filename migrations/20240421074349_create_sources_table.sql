CREATE TABLE IF NOT EXISTS sources(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);

INSERT INTO sources(id, name) VALUES(1, 'news'), (2, 'social');

ALTER TABLE stock_sentiment ADD COLUMN source_id serial NOT NULL;
ALTER TABLE stock_sentiment ADD CONSTRAINT fk_sources_stock_sentiment FOREIGN KEY (source_id) REFERENCES sources (id);
ALTER TABLE top_content ADD COLUMN source_id serial NOT NULL;
ALTER TABLE top_content ADD CONSTRAINT fk_sources_top_content FOREIGN KEY (source_id) REFERENCES sources(id);
