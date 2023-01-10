CREATE TABLE IF NOT EXISTS "tasks"
(
    id            SERIAL       NOT NULL,
    number        VARCHAR(40)  NOT NULL,
    user_id       INT          NOT NULL,
    time          VARCHAR(5)   NOT NULL,
    emails        VARCHAR[]    NOT NULL,
    is_container  BOOLEAN      NOT NULL,
    email_subject VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS "container_archive"
(
    id       SERIAL          NOT NULL,
    user_id  INT             NOT NULL,
    response VARCHAR(100000) NOT NULL
);

CREATE TABLE IF NOT EXISTS "bill_archive"
(
    id       SERIAL          NOT NULL,
    user_id  INT             NOT NULL,
    response VARCHAR(100000) NOT NULL
);