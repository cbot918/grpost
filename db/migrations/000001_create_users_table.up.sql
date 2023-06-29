CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "name" varchar,
  "password" varchar NOT NULL
);