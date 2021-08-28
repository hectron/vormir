CREATE TABLE IF NOT EXISTS users (
  id text NOT NULL,
  name text NOT NULL,
  display_name text,
  deleted boolean NOT NULL,
  PRIMARY KEY (id)
);
