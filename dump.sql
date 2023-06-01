-- Adminer 4.8.1 PostgreSQL 15.2 (Debian 15.2-1.pgdg110+1) dump

\connect "glt";

CREATE SEQUENCE location_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;
CREATE TABLE "public"."location" (
    "id" integer DEFAULT nextval('location_id_seq') NOT NULL,
    "user_id" integer,
    "time" timestamp,
    "lat" double precision,
    "lon" double precision,
    CONSTRAINT "location_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE SEQUENCE users_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;
CREATE TABLE "public"."users" (
    "id" integer DEFAULT nextval('users_id_seq') NOT NULL,
    "login" text,
    "password" text,
    CONSTRAINT "users_login_key" UNIQUE ("login"),
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE SEQUENCE vehicules_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;
CREATE TABLE "public"."vehicules" (
    "id" integer DEFAULT nextval('vehicules_id_seq') NOT NULL,
    "name" text,
    "type" text,
    CONSTRAINT "vehicules_name_key" UNIQUE ("name"),
    CONSTRAINT "vehicules_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE TYPE "usage_enum" AS ENUM ('take', 'leave');

CREATE SEQUENCE vehicules_usage_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;
CREATE TABLE "public"."vehicules_usage" (
    "id" integer DEFAULT nextval('vehicules_usage_id_seq') NOT NULL,
    "user_id" integer,
    "vehicules_id" integer,
    "usage" usage_enum,
    "time" timestamp,
    CONSTRAINT "vehicules_usage_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


ALTER TABLE ONLY "public"."location" ADD CONSTRAINT "location_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."vehicules_usage" ADD CONSTRAINT "vehicules_usage_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."vehicules_usage" ADD CONSTRAINT "vehicules_usage_vehicules_id_fkey" FOREIGN KEY (vehicules_id) REFERENCES vehicules(id) NOT DEFERRABLE;

-- 2023-05-30 12:28:10.149249+00