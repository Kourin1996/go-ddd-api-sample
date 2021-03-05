CREATE TABLE IF NOT EXISTS books(
   id serial PRIMARY KEY,
   name VARCHAR (255) NOT NULL,
   description VARCHAR (1024) NOT NULL,
   price INT NOT NULL
);
