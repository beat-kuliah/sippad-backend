CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "username" varchar(256) UNIQUE NOT NULL,
    "name" varchar(256) NOT NULL,
    "hashed_password" varchar(256) NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "updated_at" timestamp NOT NULL DEFAULT (now())
);