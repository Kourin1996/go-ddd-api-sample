CREATE TABLE IF NOT EXISTS users(
  id serial PRIMARY KEY,
  username VARCHAR (255) NOT NULL UNIQUE,
  password VARCHAR (255) NOT NULL,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp,
  deleted_at timestamp default null
);
