// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package sqlc

import (
	"context"
	"database/sql"
)

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM
    users
WHERE
    id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const existUser = `-- name: ExistUser :one
SELECT
    EXISTS(
        SELECT 1
        FROM users
        WHERE email = ? AND password = ?
    ) AS exists_user
`

type ExistUserParams struct {
	Email    string
	Password string
}

func (q *Queries) ExistUser(ctx context.Context, arg ExistUserParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, existUser, arg.Email, arg.Password)
	var exists_user bool
	err := row.Scan(&exists_user)
	return exists_user, err
}

const findUser = `-- name: FindUser :one
SELECT
    id,
    last_name,
    first_name,
    email,
    password,
    icon,
    created_at,
    updated_at
FROM
    users
WHERE
    id = ?
`

func (q *Queries) FindUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.LastName,
		&i.FirstName,
		&i.Email,
		&i.Password,
		&i.Icon,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getEventIDsByUserID = `-- name: GetEventIDsByUserID :many
SELECT
    ue.event_id
FROM
    user_events ue
WHERE
    ue.user_id = ?
`

func (q *Queries) GetEventIDsByUserID(ctx context.Context, userid string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getEventIDsByUserID, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var event_id string
		if err := rows.Scan(&event_id); err != nil {
			return nil, err
		}
		items = append(items, event_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGroupIDsByUserID = `-- name: GetGroupIDsByUserID :many
SELECT
    gu.group_id
FROM
    group_users gu
WHERE
    gu.user_id = ?
`

func (q *Queries) GetGroupIDsByUserID(ctx context.Context, userid string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getGroupIDsByUserID, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var group_id string
		if err := rows.Scan(&group_id); err != nil {
			return nil, err
		}
		items = append(items, group_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const upsertUser = `-- name: UpsertUser :exec
INSERT INTO
    users (
        id,
        last_name,
        first_name,
        email,
        icon,
        created_at,
        updated_at
    )
VALUES
    (
        ?,
        ?,
        ?,
        ?,
        ?,
        NOW(),
        NOW()
    ) ON DUPLICATE KEY
UPDATE
    last_name = ?,
    first_name = ?,
    email = ?,
    icon = ?,
    updated_at = NOW()
`

type UpsertUserParams struct {
	ID        string
	LastName  string
	FirstName string
	Email     string
	Icon      sql.NullString
}

func (q *Queries) UpsertUser(ctx context.Context, arg UpsertUserParams) error {
	_, err := q.db.ExecContext(ctx, upsertUser,
		arg.ID,
		arg.LastName,
		arg.FirstName,
		arg.Email,
		arg.Icon,
		arg.LastName,
		arg.FirstName,
		arg.Email,
		arg.Icon,
	)
	return err
}
