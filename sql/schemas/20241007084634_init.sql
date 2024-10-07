-- +goose Up
-- +goose StatementBegin

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "fullname" varchar NOT NULL,
  "email" varchar UNIQUE,
  "gender" int NOT NULL,
  "role_id" int NOT NULL,
  "describe" varchar NOT NULL DEFAULT 'You are great',
  "avt" varchar NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "role" (
  "id" serial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "albums" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "img_url" varchar NOT NULL,
  "user_id" bigint NOT NULL
);

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "users" ("fullname");

CREATE INDEX ON "albums" ("id");

CREATE INDEX ON "albums" ("user_id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "albums" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

INSERT INTO "role" (id, name) VALUES (1, 'Admin'),(2, 'Vip'),(3, 'User');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users" DROP CONSTRAINT IF EXISTS "users_role_id_fkey";
ALTER TABLE "albums" DROP CONSTRAINT IF EXISTS "albums_user_id_fkey";

DELETE FROM "role" WHERE id IN (1,2,3);

DROP TABLE IF EXISTS "role";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "albums";

-- +goose StatementEnd
