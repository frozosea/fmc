CREATE TABLE IF NOT EXISTS "user"
(
    "id"       serial              NOT NULL,
    "email"    varchar(255) UNIQUE NOT NULL,
    "username" varchar(255) UNIQUE NOT NULL,
    "password" varchar(255)        NOT NULL,
    CONSTRAINT "user_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );



CREATE TABLE IF NOT EXISTS "containers"
(
    "id"          serial      NOT NULL,
    "number"      varchar(30) NOT NULL,
    "is_on_track" BOOLEAN     NOT NULL,
    "is_arrived"  BOOLEAN     NOT NULL,
    "user_id"     integer     NOT NULL,
    CONSTRAINT "containers_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );



CREATE TABLE IF NOT EXISTS "bill_numbers"
(
    "id"          serial      NOT NULL,
    "number"      varchar(70) NOT NULL,
    "is_on_track" BOOLEAN     NOT NULL,
    "is_arrived"  BOOLEAN     NOT NULL,
    "user_id"     integer     NOT NULL,
    CONSTRAINT "bill_numbers_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );

ALTER TABLE "containers"
    ADD CONSTRAINT "containers_fk0" FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "bill_numbers"
    ADD CONSTRAINT "bill_numbers_fk0" FOREIGN KEY ("user_id") REFERENCES "user" ("id");

CREATE TABLE IF NOT EXISTS "feedback"
(
    "id"      serial           NOT NULL,
    "email"   varchar(500)     NOT NULL,
    "message" varchar(1000000) NOT NULL
);


