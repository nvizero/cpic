CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "title" varchar,
  "link" varchar,
  "img" varchar NOT NULL,
  "content" text NOT NULL,
  "created_at" timestamp
);
