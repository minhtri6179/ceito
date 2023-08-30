CREATE TABLE "question" (
    "question_id" integer PRIMARY KEY,
    "question_text" text,
    "answer_id" integer,
    "created_at" timestamptz DEFAULT (now()),
    "update_at" timestamp
);
CREATE TABLE "answer" (
    "answer_id" integer PRIMARY KEY,
    "question_id" integer,
    "answer_text" text,
    "is_correct" boolean,
    "created_at" timestamptz DEFAULT (now()),
    "update_at" timestamp
);
CREATE TABLE "users" (
    "username" varchar PRIMARY KEY,
    "email" varchar,
    "full_name" varchar,
    "password_hashed" varchar,
    "created_at" timestamptz DEFAULT (now()),
    "update_at" timestamp
);
CREATE TABLE "test" (
    "test_id" integer PRIMARY KEY,
    "username" varchar,
    "created_at" timestamptz DEFAULT (now()),
    "update_at" timestamp
);
CREATE TABLE "score" (
    "score_id" integer PRIMARY KEY,
    "test_id" integer,
    "reading_score" integer,
    "listening_score" integer,
    "total_score" integer
);
CREATE INDEX ON "question" ("question_id");
CREATE INDEX ON "answer" ("question_id");
ALTER TABLE "answer"
ADD FOREIGN KEY ("question_id") REFERENCES "question" ("question_id");
ALTER TABLE "test"
ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
ALTER TABLE "score"
ADD FOREIGN KEY ("test_id") REFERENCES "test" ("test_id");