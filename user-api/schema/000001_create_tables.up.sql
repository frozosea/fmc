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
    "number"      varchar(15) NOT NULL,
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
    "number"      varchar(35) NOT NULL,
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

CREATE OR REPLACE FUNCTION is_value_free_for_containers(_header_id integer, _value varchar) RETURNS BOOLEAN AS
$$
BEGIN
    RETURN NOT EXISTS(SELECT user_id, number
                      FROM "containers"
                      WHERE number LIKE _value
                        AND user_id != _header_id
                      LIMIT 1);
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION is_value_free_for_bill_numbers(_header_id integer, _value varchar) RETURNS BOOLEAN AS

$$
BEGIN
    RETURN NOT EXISTS(SELECT user_id, number
                      FROM "bill_numbers"
                      WHERE number LIKE _value
                        AND user_id != _header_id
                      LIMIT 1);
END
$$ LANGUAGE plpgsql;

ALTER TABLE "bill_numbers"
    ADD CONSTRAINT "bill" CHECK (is_value_free_for_bill_numbers("bill_numbers".user_id, "bill_numbers".number));
ALTER TABLE "containers"
    ADD CONSTRAINT "container" CHECK (is_value_free_for_containers(user_id, number));


