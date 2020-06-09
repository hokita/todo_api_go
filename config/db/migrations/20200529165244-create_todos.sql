-- +migrate Up
CREATE TABLE IF NOT EXISTS todos (
	id INTEGER NOT NULL,
	content TEXT NOT NULL,
	done boolean NOT NULL,
	date DATE NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS todos;
