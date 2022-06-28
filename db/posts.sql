-- Adminer 4.8.1 PostgreSQL 14.4 (Debian 14.4-1.pgdg110+1) dump

DROP TABLE IF EXISTS "posts";
DROP SEQUENCE IF EXISTS posts_id_seq;
CREATE SEQUENCE posts_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."posts" (
    "id" bigint DEFAULT nextval('posts_id_seq') NOT NULL,
    "text" text NOT NULL,
    "time" timestamp NOT NULL,
    CONSTRAINT "posts_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "posts" ("id", "text", "time") VALUES
(1,	'post 1',	'2022-06-28 08:53:34.532015'),
(2,	'post 2',	'2022-06-28 08:53:41.342164'),
(3,	'post 3',	'2022-06-28 08:53:57.165214');

-- 2022-06-28 08:54:05.684214+00
