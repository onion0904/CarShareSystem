-- name: SaveEvent :exec
INSERT INTO
    events (
        id,
        users_id,
        together,
        description,
        year,
        month,
        day,
        start_date,
        end_date,
        created_at,
        updated_at
    )
VALUES
    (
        sqlc.arg(id),
        sqlc.arg(users_id),
        sqlc.arg(together),
        sqlc.arg(description),
        sqlc.arg(year),
        sqlc.arg(month),
        sqlc.arg(day),
        sqlc.arg(start_date),
        sqlc.arg(end_date),
        NOW(),
        NOW()
    ) ON DUPLICATE KEY
UPDATE
    users_id = sqlc.arg(users_id),
    together = sqlc.arg(together),
    description = sqlc.arg(description),
    year = sqlc.arg(year),
    month = sqlc.arg(month),
    day = sqlc.arg(day),
    start_date = sqlc.arg(start_date),
    end_date = sqlc.arg(end_date),
    updated_at = NOW();

-- name: DeleteEvent :exec
DELETE FROM
    events
WHERE
    id = ?;

-- name: FindEvent :one
SELECT
    *
FROM
    events
WHERE
    id = ?;

-- name: FindMonthEventID :many
SELECT
    id
FROM
    events
WHERE
    year = ? AND month = ?
