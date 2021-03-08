CREATE TABLE IF NOT EXISTS users(
  id serial PRIMARY KEY,
  name VARCHAR (255) NOT NULL,
  password VARCHAR (255) NOT NULL,
  bio VARCHAR(1024) NOT NULL,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp
);
