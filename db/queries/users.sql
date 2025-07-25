-- name: CreateUser :one
INSERT INTO users (
                   username,
                   name,
                   role_id,
                   hashed_password,
                   created_by,
                   updated_by
) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetUserWithRole :one
SELECT
    u.id,
    u.username,
    u.name,
    u.created_at,
    u.updated_at,
    r.id AS role_id,
    r.name AS role_name,
    r.description AS role_description
FROM users u
         LEFT JOIN roles r ON u.role_id = r.id
WHERE u.id = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateUserPassword :one
UPDATE users SET hashed_password = $1, updated_at = $2
WHERE id = $3 RETURNING *;

-- name: UpdateName :one
UPDATE users SET name = $1, updated_at = $2
WHERE id = $3 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: DeleteAllUsers :exec
DELETE FROM users;