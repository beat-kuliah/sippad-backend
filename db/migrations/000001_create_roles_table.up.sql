CREATE TABLE "roles" (
                         "id" BIGSERIAL PRIMARY KEY,
                         "name" VARCHAR(100) UNIQUE NOT NULL,
                         "description" TEXT,
                         "created_at" TIMESTAMP NOT NULL DEFAULT now(),
                         "created_by" BIGINT,
                         "updated_at" TIMESTAMP NOT NULL DEFAULT now(),
                         "updated_by" BIGINT,
                         "deleted_at" TIMESTAMP,
                         "deleted_by" BIGINT
);

CREATE INDEX "idx_roles_deleted" ON "roles" ("deleted_at")
    WHERE "deleted_at" IS NULL;
