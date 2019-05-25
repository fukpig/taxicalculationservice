SET statement_timeout = 60000; -- 60 seconds
SET lock_timeout = 60000; -- 60 seconds

--gopg:split

CREATE TABLE rates(
 id serial PRIMARY KEY,
 name VARCHAR (50) UNIQUE NOT NULL,
 per_minute INT NOT NULL,
 per_km INT UNIQUE NOT NULL,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL
);
