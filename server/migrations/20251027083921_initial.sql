-- Create "accounts" table
CREATE TABLE "accounts" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(100) NOT NULL,
  "slug" character varying(50) NOT NULL,
  "logo" text NULL,
  "domain" character varying(255) NOT NULL,
  "domain_verified" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NULL,
  "created_by" uuid NOT NULL,
  "updated_at" timestamptz NULL,
  "updated_by" uuid NULL,
  "deleted_at" timestamptz NULL,
  "deleted_by" uuid NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_accounts_deleted_at" to table: "accounts"
CREATE INDEX "idx_accounts_deleted_at" ON "accounts" ("deleted_at");
-- Create index "idx_accounts_domain" to table: "accounts"
CREATE UNIQUE INDEX "idx_accounts_domain" ON "accounts" ("domain");
-- Create index "idx_accounts_slug" to table: "accounts"
CREATE UNIQUE INDEX "idx_accounts_slug" ON "accounts" ("slug");
-- Create "apps" table
CREATE TABLE "apps" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(100) NOT NULL,
  "description" text NULL,
  "landing_url" text NOT NULL,
  "logo" text NULL,
  "branding" text NULL,
  "client_id" character varying(64) NOT NULL,
  "client_secret" character varying(255) NOT NULL,
  "redirect_uris" text[] NULL,
  "jwt_algo" character varying(10) NOT NULL,
  "jwt_secret" character varying(255) NOT NULL,
  "jwt_lifetime" character varying(20) NOT NULL,
  "refresh_token_lifetime" character varying(20) NOT NULL,
  "permanent_callback" text NOT NULL,
  "permanent_error_callback" text NOT NULL,
  "account_id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "created_by" uuid NOT NULL,
  "updated_at" timestamptz NULL,
  "updated_by" uuid NULL,
  "deleted_at" timestamptz NULL,
  "deleted_by" uuid NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_apps_client_id" to table: "apps"
CREATE UNIQUE INDEX "idx_apps_client_id" ON "apps" ("client_id");
-- Create index "idx_apps_deleted_at" to table: "apps"
CREATE INDEX "idx_apps_deleted_at" ON "apps" ("deleted_at");
-- Create "auths" table
CREATE TABLE "auths" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NULL,
  "created_at" timestamptz NULL,
  "created_by" uuid NULL,
  "updated_at" timestamptz NULL,
  "updated_by" uuid NULL,
  "deleted_at" timestamptz NULL,
  "deleted_by" uuid NULL,
  PRIMARY KEY ("id")
);
-- Create "oauths" table
CREATE TABLE "oauths" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NULL,
  "app_id" uuid NULL,
  "account_id" uuid NULL,
  "code" text NULL,
  "expires_at" timestamptz NULL,
  "created_at" timestamptz NULL,
  "created_by" uuid NULL,
  "updated_at" timestamptz NULL,
  "updated_by" uuid NULL,
  "deleted_at" timestamptz NULL,
  "deleted_by" uuid NULL,
  PRIMARY KEY ("id")
);
-- Create "passwords" table
CREATE TABLE "passwords" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NULL,
  "account_id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "created_by" uuid NULL,
  "updated_at" timestamptz NULL,
  "updated_by" uuid NULL,
  "deleted_at" timestamptz NULL,
  "deleted_by" uuid NULL,
  PRIMARY KEY ("id")
);
-- Create "sessions" table
CREATE TABLE "sessions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "ip" text NULL,
  "user_agent" text NULL,
  "refresh_token" text NULL,
  "user_id" uuid NULL,
  "app_id" uuid NULL,
  "account_id" uuid NULL,
  "created_at" timestamptz NULL,
  "created_by" uuid NULL,
  "updated_at" timestamptz NULL,
  "updated_by" uuid NULL,
  "deleted_at" timestamptz NULL,
  "deleted_by" uuid NULL,
  PRIMARY KEY ("id")
);
-- Create "users" table
CREATE TABLE "users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(100) NOT NULL,
  "email" character varying(255) NOT NULL,
  "dp" text NULL,
  "account_id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "created_by" uuid NOT NULL,
  "updated_at" timestamptz NULL,
  "updated_by" uuid NULL,
  "deleted_at" timestamptz NULL,
  "deleted_by" uuid NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");
