CREATE TABLE users (
   id UUID PRIMARY KEY,
   login TEXT NOT NULL UNIQUE,
   password TEXT NOT NULL,
   role TEXT NOT NULL
);
