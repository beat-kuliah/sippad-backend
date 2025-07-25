-- name: CreatePengelola :one
INSERT INTO pengelola (
    nama,
    ktp,
    jabatan,
    alamat,
    kecamatan,
    kelurahan,
    kabupaten_kota,
    kode_pos,
    telepon,
    created_by,
    updated_by
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    RETURNING *;
