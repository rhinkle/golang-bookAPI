-- Adminer 4.8.1 PostgreSQL 13.4 (Debian 13.4-1.pgdg100+1) dump

DROP TABLE IF EXISTS "book";
DROP SEQUENCE IF EXISTS book_id_seq;
CREATE SEQUENCE book_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."book" (
    "ID" integer DEFAULT nextval('book_id_seq') NOT NULL,
    "Title" character varying NOT NULL,
    "Author" character varying NOT NULL,
    "Year" character varying NOT NULL,
    CONSTRAINT "book_pkey" PRIMARY KEY ("ID")
) WITH (oids = false);

INSERT INTO "book" ("ID", "Title", "Author", "Year") VALUES
(1,	'Roy Explains Loops in Go',	'Tom Scott',	'2007'),
(2,	'Cat''s Save Nine Lives',	'Cookie Cat',	'2039');

-- 2021-08-23 02:57:29.509882+00
