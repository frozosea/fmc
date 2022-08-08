CREATE TABLE IF NOT EXISTS "tasks"
(
    id            SERIAL       NOT NULL,
    number        VARCHAR(40)  NOT NULL,
    user_id       INT          NOT NULL,
    country       VARCHAR(10)  NOT NULL,
    time          VARCHAR(5)   NOT NULL,
    emails        VARCHAR[]    NOT NULL,
    is_container  BOOLEAN      NOT NULL,
    email_subject VARCHAR(255) NOT NULL
);