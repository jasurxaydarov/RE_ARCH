CREATE DATABASE iqro;
CREATE EXTENSION IF NOT EXISTS  "uuid-ossp";


CREATE TABLE sections (
    section_id      SERIAL              PRIMARY KEY,
    section_name    VARCHAR(128)        NOT NULL,
    icon            VARCHAR(512)        NOT NULL,
    description     VARCHAR(1024)       NOT NULL
);


CREATE TABLE categories(
    category_id      SERIAL             PRIMARY KEY,
    category_name    VARCHAR(128)       NOT NULL,
    icon             VARCHAR(512)       NOT NULL,
    description      VARCHAR(1024)      NOT NULL,
    section_id       INT                REFERENCES  sections(section_id)
);

CREATE TABLE mentors(
    mentor_id       SERIAL              PRIMARY KEY,
    fullname        VARCHAR(128)        NOT NULL
);

CREATE TABLE courses(
    course_id       SERIAL              PRIMARY KEY,
    course_name     VARCHAR(128)        NOT NULL,
    description     VARCHAR(1024)       NOT NULL,
    price           DECIMAL             NOT NULL,
    is_free         BOOLEAN             DEFAULT FALSE,
    category_id     INT                 REFERENCES categories(category_id),
    mentor_id       INT                 REFERENCES  mentors(mentor_id)

);


CREATE TABLE units (
    unit_id         UUID                PRIMARY KEY   DEFAULT uuid_generate_v4(),
    unit_name       VARCHAR(64)         NOT NULL,
    description     VARCHAR(1024)       NOT NULL,
    course_id       INT                 REFERENCES courses(course_id)
);

CREATE TABLE videos(
    video_id        UUID                PRIMARY KEY   DEFAULT uuid_generate_v4(),
    title           VARCHAR(512)        NOT NULL,
    video_url       VARCHAR(512)        NOT NULL,
    is_viewed       BOOLEAN             DEFAULT FALSE,
    viewed_line     FLOAT               DEFAULT 0,
    state_order     INTEGER             DEFAULT 1,
    unit_id         UUID                REFERENCES units(unit_id)

);


CREATE TABLE documents(
    document_id     UUID                PRIMARY KEY   DEFAULT uuid_generate_v4(),
    title           VARCHAR(128)        NOT NULL,
    video_id        UUID                REFERENCES videos(video_id)
);


CREATE TABLE subtitles(
    subtitle_id     UUID               PRIMARY KEY   DEFAULT uuid_generate_v4(),
    subtitle        VARCHAR(256)       NOT NULL,
    information     TEXT               NOT NULL,
    document_id     UUID               REFERENCES   documents(document_id)  
);


CREATE TYPE 
    usr_type
AS
    ENUM(
        'viewer',
        'user'
    );

CREATE TABLE users(

    user_id         UUID                PRIMARY KEY     DEFAULT uuid_generate_v4(),
    username        VARCHAR(25)         NOT NULL,
    password        VARCHAR(60)         NOT NULL,
    user_type       usr_type            DEFAULT 'viewer',
    fullname        VARCHAR(128),
    gmail           VARCHAR(128)        NOT NULL,
    phone_number    VARCHAR(15),
    date_of_birth   DATE,
    created_at      TIMESTAMP           DEFAULT current_timestamp,
    deleted_at      TIMESTAMP           DEFAULT NULL,
);

CREATE TABLE purchased_courses(
    purchase_id     SERIAL              PRIMARY KEY,
    user_id         UUID                REFERENCES users(user_id),
    course_id       INTEGER             REFERENCES courses(course_id),
    purchase_date   TIMESTAMP           DEFAULT current_timestamp

);