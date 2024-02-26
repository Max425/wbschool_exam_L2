CREATE TABLE event
(
    id       SERIAL PRIMARY KEY,
    date     TIMESTAMP NOT NULL,
    user_id  TEXT      NOT NULL,
    title    TEXT      NOT NULL
);