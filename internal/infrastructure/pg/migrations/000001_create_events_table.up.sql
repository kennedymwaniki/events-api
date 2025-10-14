BEGIN;

CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
	title   varchar(255) NOT NULL,
	description  varchar(1000) NOT NULL,
	location  varchar(255) NOT NULL,
	date  varchar(50) NOT NULL,
	time  time NOT NULL,
	capacity  integer NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMIT;