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

CREATE TABLE IF NOT EXISTS "feedback"
(
    "id"      SERIAL           NOT NULL,
    "email"   VARCHAR(500)     NOT NULL,
    "message" VARCHAR(1000000) NOT NULL
);

CREATE TABLE IF NOT EXISTS "company"
(
    "id"                       SERIAL           NOT NULL,
    "user_id"                  INT              NOT NULL,
    "company_full_name"        varchar(500)     NOT NULL,
    "company_abbreviated_name" varchar(500)     NOT NULL,
    "inn"                      varchar(500)     NOT NULL,
    "ogrn"                     varchar(500)     NOT NULL,
    "legal_address"            varchar(500)     NOT NULL,
    "post_address"             varchar(500)     NOT NULL,
    "work_email"               varchar(1000000) NOT NULL,
    CONSTRAINT "company_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );

ALTER TABLE "company"
    ADD CONSTRAINT "company_fk0" FOREIGN KEY ("user_id") REFERENCES "user" ("id");

CREATE TABLE IF NOT EXISTS "balance"
(
    "id"      SERIAL NOT NULL,
    "user_id" INT    NOT NULL,
    "value"   FLOAT  NOT NULL,
    CONSTRAINT "balance_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );

ALTER TABLE "balance"
    ADD CONSTRAINT "balance_fk0" FOREIGN KEY ("user_id") REFERENCES "user" ("id");

CREATE TABLE IF NOT EXISTS "balance_transaction"
(
    "id"         SERIAL                                             NOT NULL,
    "user_id"    INT                                                NOT NULL,
    "value"      FLOAT                                              NOT NULL,
    "type"       INT                                                NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT "balance_tr_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );

ALTER TABLE "balance_transaction"
    ADD CONSTRAINT "balance_fk0" FOREIGN KEY ("user_id") REFERENCES "user" ("id");

CREATE TABLE IF NOT EXISTS "numbers"
(
    "id"               SERIAL  NOT NULL,
    "number"           VARCHAR NOT NULL,
    "user_id"          INT     NOT NULL,
    "days_on_tracking" INT     NOT NULL,
    "is_container"     BOOLEAN NOT NULL,
    CONSTRAINT "numbers_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );
ALTER TABLE "numbers"
    ADD CONSTRAINT "numbers_fk0" FOREIGN KEY ("user_id") REFERENCES "user" ("id");


CREATE TABLE IF NOT EXISTS "number_transaction"
(
    "id"             SERIAL NOT NULL,
    "transaction_id" INT    NOT NULL,
    "number_id"      INT    NOT NULL,
    CONSTRAINT "number_tr_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );


ALTER TABLE "number_transaction"
    ADD CONSTRAINT "numbers_tr_fk0" FOREIGN KEY ("number_id") REFERENCES "numbers" ("id");
ALTER TABLE "number_transaction"
    ADD CONSTRAINT "number_tr_fk1" FOREIGN KEY ("transaction_id") REFERENCES "balance_transaction" ("id");

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


