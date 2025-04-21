-- SEQUENCE
CREATE SEQUENCE IF NOT EXISTS "plugins_id_seq" INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

-- TABLE
CREATE TABLE IF NOT EXISTS "plugins" (
	 "id" BIGINT DEFAULT nextval('plugins_id_seq') NOT NULL NOT NULL ,
	 CONSTRAINT "plugins_id_pk" PRIMARY KEY ("id"),
	 "pkg" VARCHAR(300) NOT NULL,
	 "ver" VARCHAR(50) NOT NULL,
	 "hash" VARCHAR(64) NOT NULL,
	 "active" BOOLEAN NOT NULL,
	 "meta" JSONB NULL,
	 "created_at" TIMESTAMPTZ NOT NULL,
	 "updated_at" TIMESTAMPTZ NOT NULL,
	 CONSTRAINT "plugins__pkg_ver__uniq" UNIQUE ("pkg","ver"),
	 CONSTRAINT "plugins__hash__uniq" UNIQUE ("hash")
);

