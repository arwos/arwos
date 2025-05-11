-- SEQUENCE
CREATE SEQUENCE IF NOT EXISTS "envs_id_seq" INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

-- TABLE
CREATE TABLE IF NOT EXISTS "envs" (
	 "id" BIGINT DEFAULT nextval('envs_id_seq') NOT NULL ,
	 CONSTRAINT "envs_id_pk" PRIMARY KEY ("id"),
	 "plugin_id" BIGINT NOT NULL ,
	 CONSTRAINT "envs_plugin_id_fk" FOREIGN KEY ("plugin_id") REFERENCES "plugins" ("id") ON DELETE CASCADE NOT DEFERRABLE,
	 "key" TEXT NOT NULL,
	 "value" TEXT NOT NULL,
	 "default_value" TEXT NOT NULL,
	 "desc" TEXT NOT NULL,
	 "created_at" TIMESTAMPTZ NOT NULL,
	 "updated_at" TIMESTAMPTZ NOT NULL,
	 CONSTRAINT "envs__plugin_id_key__uniq" UNIQUE ("plugin_id","key")
);

