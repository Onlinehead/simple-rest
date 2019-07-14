SET statement_timeout = 60000; -- 60 seconds
SET lock_timeout = 60000; -- 60 seconds

--gopg:split

CREATE TABLE users (
    username    varchar PRIMARY KEY,
    birthday    integer NOT NULL
);
