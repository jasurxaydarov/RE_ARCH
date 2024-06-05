CREATE DATABASE arch;

CREATE TABLE users(
    user_id         UUID                PRIMARY KEY,
    user_name       VARCHAR (69)        NOT NULL        UNIQUE,
    gmail           VARCHAR(124)        NOT NULL,
    age             INT                 NOT NULL ,
    password        VARCHAR(22)         NOT NULL 

);

CREATE TABLE todo (
    todo_id         UUID                PRIMARY KEY,
    task            VARCHAR             NOT NULL        UNIQUE,
    created_at      TIMESTAMP           DEFAULT current_timestamp,
    is_completed    BOOLEaN             DEFAULT FALSE,
    user_name         VARCHAR             REFERENCES  Users(user_name)

);



drop table todo;

SELECT * FROM users;

update users set gmail='aaaa' where user_name='Alice';