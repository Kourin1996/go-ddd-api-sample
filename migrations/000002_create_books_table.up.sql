CREATE TABLE IF NOT EXISTS books(
   id serial PRIMARY KEY,
   name VARCHAR (255) NOT NULL,
   description VARCHAR (1024) NOT NULL,
   price INT NOT NULL,
   user_id int,
   created_at timestamp not null default current_timestamp,
   updated_at timestamp not null default current_timestamp,
   foreign key (user_id) references users(id)
);
