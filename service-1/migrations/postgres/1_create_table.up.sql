
CREATE TABLE IF NOT EXISTS "posts" (
  "id" int PRIMARY KEY,
  "user_id" int not null,
  "title" varchar(500) not null,
  "body" text
);
