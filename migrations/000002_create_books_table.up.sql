CREATE TABLE IF NOT EXISTS books(
   id serial PRIMARY KEY,
   name VARCHAR (255) NOT NULL,
   description VARCHAR (1024) NOT NULL,
   price INT NOT NULL,
   created_at timestamp not null default current_timestamp,
   updated_at timestamp not null default current_timestamp,
   deleted_at timestamp default null
);
