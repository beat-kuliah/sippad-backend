// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: roles.sql

package db

import (
	"context"
	"database/sql"
)

const createRole = `-- name: CreateRole :one
INSERT INTO roles (
    name,
    description,
    created_by,
    updated_by
) VALUES ($1, $2, $3, $4) RETURNING id, name, description, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

type CreateRoleParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	CreatedBy   sql.NullInt64  `json:"created_by"`
	UpdatedBy   sql.NullInt64  `json:"updated_by"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.db.QueryRowContext(ctx, createRole,
		arg.Name,
		arg.Description,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}
