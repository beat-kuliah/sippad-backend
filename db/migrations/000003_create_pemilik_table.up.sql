CREATE TABLE "pemilik" (
                           "id" BIGSERIAL PRIMARY KEY,
                           "nama" VARCHAR(255) NOT NULL,
                           "ktp" VARCHAR(50) NOT NULL,
                           "jabatan" VARCHAR(100),
                           "alamat" TEXT,
                           "kecamatan" VARCHAR(100),
                           "kelurahan" VARCHAR(100),
                           "kabupaten_kota" VARCHAR(100),
                           "kode_pos" VARCHAR(10),
                           "telepon" VARCHAR(50),

                           "created_at" TIMESTAMP NOT NULL DEFAULT now(),
                           "created_by" BIGINT,
                           "updated_at" TIMESTAMP NOT NULL DEFAULT now(),
                           "updated_by" BIGINT,
                           "deleted_at" TIMESTAMP,
                           "deleted_by" BIGINT
);

CREATE INDEX "idx_pemilik_deleted" ON "pemilik" ("deleted_at")
    WHERE "deleted_at" IS NULL;