CREATE DATABASE arch;

CREATE TABLE Users(
    user_id         SERIAL              PRIMARY KEY,
    user_name       VARCHAR (69)        NOT NULL,
    gmail           VARCHAR(124)        NOT NULL,
    age             INT                 NOT NULL ,
    password        VARCHAR(22)         NOT NULL 

);

CREATE TABLE todo (
    todo_id         SERIAL              PRIMARY KEY,
    task            VARCHAR             NOT NULL,
    created_at      TIMESTAMP           DEFAULT current_timestamp,
    is_completed    BOOLEN              DEFAULT FALSE

);

