-- Create "accounts" table
CREATE TABLE "accounts" (
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
-- Create "apps" table
CREATE TABLE "apps" (
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
-- Create "auths" table
CREATE TABLE "auths" (
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
-- Create "email_verification_requests" table
CREATE TABLE "email_verification_requests" (
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
-- Create "oauths" table
CREATE TABLE "oauths" (
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
-- Create "reset_password_requests" table
CREATE TABLE "reset_password_requests" (
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
-- Create "users" table
CREATE TABLE "users" (
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
