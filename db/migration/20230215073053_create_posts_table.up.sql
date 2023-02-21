CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "title" varchar,
  "link" varchar,
  "dt" varchar,
  "state" Bool,
  "img" varchar NOT NULL,
  "content" text NOT NULL,
  "created_at" timestamp
);
