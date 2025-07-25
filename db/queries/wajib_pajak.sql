-- name: CreateWajibPajak :one
INSERT INTO wajib_pajak (
    npwpd,
    golongan,
    no_pendaftaran,
    tanggal_daftar,
    nama_jalan,
    nomor,
    rt,
    rw,
    kecamatan,
    kelurahan,
    kabupaten_kota,
    kode_pos,
    telepon,
    email,
    pemilik_id,
    pengelola_id,
    created_by,
    updated_by
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
             $11, $12, $13, $14, $15, $16, $17, $18
         ) RETURNING *;

-- name: GetWajibPajakDetail :one
SELECT
    wp.id AS wajib_pajak_id,
    wp.npwpd,
    wp.golongan,
    wp.no_pendaftaran,
    wp.tanggal_daftar,
    wp.nama_jalan,
    wp.nomor,
    wp.rt,
    wp.rw,
    wp.kecamatan AS wp_kecamatan,
    wp.kelurahan AS wp_kelurahan,
    wp.kabupaten_kota AS wp_kabupaten_kota,
    wp.kode_pos AS wp_kode_pos,
    wp.telepon AS wp_telepon,
    wp.email AS wp_email,

    -- Data Pemilik
    p.id AS pemilik_id,
    p.nama AS pemilik_nama,
    p.ktp AS pemilik_ktp,
    p.jabatan AS pemilik_jabatan,
    p.alamat AS pemilik_alamat,
    p.kecamatan AS pemilik_kecamatan,
    p.kelurahan AS pemilik_kelurahan,
    p.kabupaten_kota AS pemilik_kabupaten_kota,
    p.kode_pos AS pemilik_kode_pos,
    p.telepon AS pemilik_telepon,

    -- Data Pengelola
    pg.id AS pengelola_id,
    pg.nama AS pengelola_nama,
    pg.ktp AS pengelola_ktp,
    pg.jabatan AS pengelola_jabatan,
    pg.alamat AS pengelola_alamat,
    pg.kecamatan AS pengelola_kecamatan,
    pg.kelurahan AS pengelola_kelurahan,
    pg.kabupaten_kota AS pengelola_kabupaten_kota,
    pg.kode_pos AS pengelola_kode_pos,
    pg.telepon AS pengelola_telepon

FROM wajib_pajak wp
         LEFT JOIN pemilik p ON wp.pemilik_id = p.id
         LEFT JOIN pengelola pg ON wp.pengelola_id = pg.id
WHERE wp.id = $1;
