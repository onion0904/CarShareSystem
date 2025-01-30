-- name: UpdateGroup :exec
UPDATE
    `groups`
SET
    name = sqlc.arg(name),
    icon = sqlc.arg(icon),
    updated_at = NOW()
WHERE
    id = sqlc.arg(id);

-- name: SaveGroup :exec
INSERT INTO
    `groups` (
        id,
        name,
        icon,
        created_at,
        updated_at
    )
VALUES
    (
        sqlc.arg(id),
        sqlc.arg(name),
        sqlc.arg(icon),
        NOW(),
        NOW()
    ) ON DUPLICATE KEY
UPDATE
    name = sqlc.arg(name),
    icon = sqlc.arg(icon),
    updated_at = NOW();

-- name: AddGroupIDToUser :exec
INSERT INTO
    group_users (group_id, user_id)
VALUES
    (
        sqlc.arg(groupID),
        sqlc.arg(userID)
    )
ON DUPLICATE KEY UPDATE
    group_id = sqlc.arg(groupID), 
    user_id = sqlc.arg(userID);

-- name: DeleteGroup :exec
DELETE FROM
    `groups`
WHERE
    id = sqlc.arg(groupID);

-- name: FindAllUserID :many
SELECT
    user_id
FROM
    group_users
WHERE
    group_id = sqlc.arg(groupID);

-- name: FindGroup :one
SELECT
    id,
    name,
    icon,
    created_at,
    updated_at
FROM
    `groups`
WHERE
    id = sqlc.arg(groupID);

-- name: FindGroupName :one
SELECT
    name
FROM
    `groups`
WHERE
    id = sqlc.arg(groupID);
