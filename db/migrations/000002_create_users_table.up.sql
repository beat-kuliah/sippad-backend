CREATE TABLE "users" (
                         "id" BIGSERIAL PRIMARY KEY,
                         "username" VARCHAR(256) UNIQUE NOT NULL,
                         "name" VARCHAR(256) NOT NULL,
                         "hashed_password" VARCHAR(256) NOT NULL,
                         "role_id" BIGINT REFERENCES "roles"("id") ON DELETE SET NULL,

                         "created_at" TIMESTAMP NOT NULL DEFAULT now(),
                         "created_by" BIGINT,
                         "updated_at" TIMESTAMP NOT NULL DEFAULT now(),
                         "updated_by" BIGINT,
                         "deleted_at" TIMESTAMP,
                         "deleted_by" BIGINT
);

CREATE INDEX "idx_users_deleted" ON "users" ("deleted_at")
    WHERE "deleted_at" IS NULL;
