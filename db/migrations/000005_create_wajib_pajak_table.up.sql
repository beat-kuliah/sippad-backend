CREATE TABLE "wajib_pajak" (
                               "id" BIGSERIAL PRIMARY KEY,
                               "npwpd" VARCHAR(50) NOT NULL,
                               "golongan" VARCHAR(50),
                               "no_pendaftaran" VARCHAR(50),
                               "tanggal_daftar" DATE NOT NULL,
                               "nama_jalan" VARCHAR(255),
                               "nomor" VARCHAR(50),
                               "rt" VARCHAR(10),
                               "rw" VARCHAR(10),
                               "kecamatan" VARCHAR(100),
                               "kelurahan" VARCHAR(100),
                               "kabupaten_kota" VARCHAR(100),
                               "kode_pos" VARCHAR(10),
                               "telepon" VARCHAR(50),
                               "email" VARCHAR(255),

                               "pemilik_id" BIGINT REFERENCES "pemilik" ("id"),
                               "pengelola_id" BIGINT REFERENCES "pengelola" ("id"),

                               "created_at" TIMESTAMP NOT NULL DEFAULT now(),
                               "created_by" BIGINT,
                               "updated_at" TIMESTAMP NOT NULL DEFAULT now(),
                               "updated_by" BIGINT,
                               "deleted_at" TIMESTAMP,
                               "deleted_by" BIGINT
);

CREATE INDEX "idx_wajib_pajak_deleted" ON "wajib_pajak" ("deleted_at")
    WHERE "deleted_at" IS NULL;