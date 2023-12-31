CREATE TABLE "accounts" (
    "id" bigserial PRIMARY KEY,
    "owner" varchar NOT NULL,
    "test_id" integer,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);
CREATE TABLE "question" (
    "question_id" bigserial PRIMARY KEY,
    "question_text" text,
    "test_name" varchar,
    "img" json,
    "created_at" timestamptz DEFAULT (now()),
    "update_at" timestamp
);
CREATE TABLE "answer" (
    "answer_id" bigserial PRIMARY KEY,
    "question_id" integer,
    "answer_text" text,
    "is_correct" boolean,
    "created_at" timestamptz DEFAULT (now()),
    "update_at" timestamp
);
CREATE TABLE "users" (
    "username" varchar PRIMARY KEY,
    "email" varchar UNIQUE not null,
    "full_name" varchar not null,
    "password_hashed" varchar not null,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "update_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);
CREATE TABLE "test_result" (
    "test_id" bigserial PRIMARY KEY,
    "username" varchar,
    "score_id" integer,
    "created_at" timestamptz DEFAULT (now()),
    "update_at" timestamp
);
CREATE TABLE "score" (
    "id" bigserial PRIMARY KEY,
    "reading_score" integer,
    "listening_score" integer,
    "total_score" integer
);
CREATE INDEX ON "question" ("question_id");
CREATE INDEX ON "answer" ("question_id");
ALTER TABLE "answer"
ADD FOREIGN KEY ("question_id") REFERENCES "question" ("question_id");
ALTER TABLE "test_result"
ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
ALTER TABLE "test_result"
ADD FOREIGN KEY ("score_id") REFERENCES "score" ("id");
ALTER TABLE "accounts"
ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");
ALTER TABLE "accounts"
ADD FOREIGN KEY ("test_id") REFERENCES "test_result" ("test_id");