CREATE TABLE "cities"
(
    "id"           serial NOT NULL,
    "ru_full_name" TEXT   NOT NULL UNIQUE,
    "en_full_name" TEXT   NOT NULL UNIQUE,
    "country_id"   INT    NOT NULL UNIQUE,
    CONSTRAINT "cities_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );

CREATE TABLE "countries"
(
    "id"           serial NOT NULL,
    "ru_full_name" TEXT   NOT NULL UNIQUE,
    "en_full_name" TEXT   NOT NULL UNIQUE,
    CONSTRAINT "countries_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );


CREATE TABLE "price"
(
    "id"           serial NOT NULL,
    "from_city_id" bigint NOT NULL,
    "usd_price"    bigint NOT NULL,
    "container_id" bigint NOT NULL,
    "contact_id"   bigint NOT NULL,
    "to_city_id"   bigint NOT NULL,
    "from_date"    DATE   NOT NULL,
    "expires_date" DATE   NOT NULL
) WITH (
      OIDS= FALSE
    );



CREATE TABLE "containers"
(
    "id"   serial NOT NULL,
    "type" TEXT   NOT NULL,
    CONSTRAINT "container_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );



CREATE TABLE "company"
(
    "id"           serial NOT NULL,
    "url"          TEXT   NOT NULL,
    "email"        TEXT   NOT NULL,
    "name"         TEXT   NOT NULL,
    "phone_number" TEXT   NOT NULL,
    CONSTRAINT "contact_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );



ALTER TABLE "cities"
    ADD CONSTRAINT "cities_fk1" FOREIGN KEY ("country_id") REFERENCES "countries" ("id");
ALTER TABLE "price"
    ADD CONSTRAINT "price_fk1" FOREIGN KEY ("from_city_id") REFERENCES "cities" ("id");
ALTER TABLE "price"
    ADD CONSTRAINT "price_fk2" FOREIGN KEY ("container_id") REFERENCES "containers" ("id");
ALTER TABLE "price"
    ADD CONSTRAINT "price_fk3" FOREIGN KEY ("contact_id") REFERENCES "company" ("id");
ALTER TABLE "price"
    ADD CONSTRAINT "price_fk4" FOREIGN KEY ("to_city_id") REFERENCES "cities" ("id");
