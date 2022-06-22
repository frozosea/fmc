CREATE TABLE "cities"
(
    "id"        serial NOT NULL,
    "unlocode"  TEXT   NOT NULL,
    "full_name" TEXT   NOT NULL UNIQUE,
    CONSTRAINT "cities_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );



CREATE TABLE "lines"
(
    "id"        serial NOT NULL,
    "scac"      TEXT   NOT NULL UNIQUE,
    "full_name" TEXT   NOT NULL UNIQUE,
    "image_url" TEXT   NOT NULL UNIQUE,
    CONSTRAINT "lines_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );



CREATE TABLE "price"
(
    "id"           serial NOT NULL,
    "line_id"      bigint NOT NULL,
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



CREATE TABLE "container"
(
    "id"        serial NOT NULL,
    "full_name" TEXT   NOT NULL,
    CONSTRAINT "container_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );



CREATE TABLE "contact"
(
    "id"           serial NOT NULL,
    "url"          TEXT   NOT NULL,
    "email"        TEXT   NOT NULL,
    "agent_name"   TEXT   NOT NULL,
    "phone_number" TEXT   NOT NULL,
    CONSTRAINT "contact_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );



ALTER TABLE "price"
    ADD CONSTRAINT "price_fk0" FOREIGN KEY ("line_id") REFERENCES "lines" ("id");
ALTER TABLE "price"
    ADD CONSTRAINT "price_fk1" FOREIGN KEY ("from_city_id") REFERENCES "cities" ("id");
ALTER TABLE "price"
    ADD CONSTRAINT "price_fk2" FOREIGN KEY ("container_id") REFERENCES "container" ("id");
ALTER TABLE "price"
    ADD CONSTRAINT "price_fk3" FOREIGN KEY ("contact_id") REFERENCES "contact" ("id");
ALTER TABLE "price"
    ADD CONSTRAINT "price_fk4" FOREIGN KEY ("to_city_id") REFERENCES "cities" ("id");







