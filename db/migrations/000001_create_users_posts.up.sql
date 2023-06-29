CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "pic" varchar NOT NULL
);

CREATE TABLE "follow" (
  "id" bigserial PRIMARY KEY,
  "from_user" bigserial NOT NULL,
  "to_user" bigserial NOT NULL
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "body" varchar NOT NULL,
  "posted_by" bigint NOT NULL,
  "photo" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "post_like" (
  "id" bigserial PRIMARY KEY,
  "target_post" bigint NOT NULL
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "texts" varchar NOT NULL,
  "posted_by" bigint NOT NULL,
  "target_post" bigint NOT NULL
);

CREATE TABLE "comment_like" (
  "id" bigserial PRIMARY KEY,
  "target_comment" bigint NOT NULL
);

ALTER TABLE "follow" ADD FOREIGN KEY ("from_user") REFERENCES "users" ("id");

ALTER TABLE "follow" ADD FOREIGN KEY ("to_user") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("posted_by") REFERENCES "users" ("id");

ALTER TABLE "post_like" ADD FOREIGN KEY ("target_post") REFERENCES "posts" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("posted_by") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("target_post") REFERENCES "posts" ("id");

ALTER TABLE "comment_like" ADD FOREIGN KEY ("target_comment") REFERENCES "comments" ("id");