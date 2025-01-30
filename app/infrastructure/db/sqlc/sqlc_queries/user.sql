-- name: UserFindById :one
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
    id = ?;

-- name: UpsertUser :exec
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
        sqlc.arg(id),
        sqlc.arg(last_name),
        sqlc.arg(first_name),
        sqlc.arg(email),
        sqlc.arg(icon),
        NOW(),
        NOW()
    ) ON DUPLICATE KEY
UPDATE
    last_name = sqlc.arg(last_name),
    first_name = sqlc.arg(first_name),
    email = sqlc.arg(email),
    icon = sqlc.arg(icon),
    updated_at = NOW();

-- name: UserFindAll :many
SELECT
    *
FROM
    users;

-- name: DeleteUser :exec
DELETE FROM
    users
WHERE
    id = ?;

-- name: ExistUser :one
SELECT
    EXISTS(
        SELECT 1
        FROM users
        WHERE email = ? AND password = ?
    ) AS exists_user;

-- name: FindAllGroupID :many
SELECT
    gu.group_id
FROM
    group_users gu
WHERE
    gu.user_id = ?;

-- name: UserFindByEmail :one
SELECT
    *
FROM
    users
WHERE
    email = ?;

-- name: UserFindName :one
SELECT
    CONCAT(first_name, ' ', last_name) AS full_name
FROM
    users
WHERE
    id = ?;
