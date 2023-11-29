CREATE TABLE IF NOT EXISTS "roles" (
  "role_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
  "role_name" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS "users" (
  "user_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
  "fk_role_id" int4,
  "username" varchar UNIQUE NOT NULL,
  "display_name" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS "user_informations" (
  "user_information_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
  "fk_user_id" int4,
  "profile_image_path" text,
  "email" varchar,
  "bio" text,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS "quizzes" (
  "quiz_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
  "fk_user_id" int4,
  "fk_quiz_category_id" int4,
  "title" text NOT NULL,
  "description" text,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS "quiz_categories" (
  "quiz_category_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
  "category" text,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS "quiz_questions" (
  "quiz_question_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
  "fk_quiz_id" int4,
  "question" text,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS "quiz_question_choices" (
  "quiz_question_choice_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
  "fk_quiz_question_id" int4,
  "question" text,
  "is_correct" bool,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS "quiz_histories" (
  "quiz_history_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
  "fk_quiz_id" int4,
  "fk_quiz_creator_user_id" int4,
  "fk_quiz_participant_user_id" int4,
  "fk_quiz_category_id" int4,
  "score" int,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS "admins" (
  "admin_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
  "fk_role_id" int4,
  "username" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS "quiz_tags" (
  "quiz_tag_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
  "fk_quiz_id" int4,
  "fk_tag_id" int4,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS "tags" (
  "tag_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
  "tag_name" text,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS "reports" (
  "report_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
  "fk_user_id" int4,
  "fk_report_status_id" int4,
  "report_content" text,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS "report_statuses" (
  "report_status_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
  "status" text,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "used_flg" bool DEFAULT true
);

ALTER TABLE "users" ADD FOREIGN KEY ("fk_role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "quizzes" ADD FOREIGN KEY ("fk_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "quizzes" ADD FOREIGN KEY ("fk_quiz_category_id") REFERENCES "quiz_categories" ("quiz_category_id");

ALTER TABLE "quiz_questions" ADD FOREIGN KEY ("fk_quiz_id") REFERENCES "quizzes" ("quiz_id");

ALTER TABLE "quiz_question_choices" ADD FOREIGN KEY ("fk_quiz_question_id") REFERENCES "quiz_questions" ("quiz_question_id");

ALTER TABLE "quiz_histories" ADD FOREIGN KEY ("fk_quiz_creator_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "quiz_histories" ADD FOREIGN KEY ("fk_quiz_participant_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "quiz_histories" ADD FOREIGN KEY ("fk_quiz_id") REFERENCES "quizzes" ("quiz_id");

ALTER TABLE "quiz_histories" ADD FOREIGN KEY ("fk_quiz_category_id") REFERENCES "quiz_categories" ("quiz_category_id");

ALTER TABLE "admins" ADD FOREIGN KEY ("fk_role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "quiz_tags" ADD FOREIGN KEY ("fk_quiz_id") REFERENCES "quizzes" ("quiz_id");

ALTER TABLE "quiz_tags" ADD FOREIGN KEY ("fk_tag_id") REFERENCES "tags" ("tag_id");

ALTER TABLE "reports" ADD FOREIGN KEY ("fk_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "reports" ADD FOREIGN KEY ("fk_report_status_id") REFERENCES "report_statuses" ("report_status_id");

ALTER TABLE "user_informations" ADD FOREIGN KEY ("fk_user_id") REFERENCES "users" ("user_id")
