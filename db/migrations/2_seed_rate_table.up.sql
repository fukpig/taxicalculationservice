SET statement_timeout = 60000; -- 60 seconds
SET lock_timeout = 60000; -- 60 seconds

--gopg:split
INSERT INTO rates(name,per_minute,per_km,created_at,updated_at) VALUES ('basic',20,30,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),('premium',40,50,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP);
