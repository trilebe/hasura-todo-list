CREATE TABLE "public"."task" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "description" text NOT NULL, "is_completed" boolean NOT NULL DEFAULT false, PRIMARY KEY ("id") );
CREATE EXTENSION IF NOT EXISTS pgcrypto;
