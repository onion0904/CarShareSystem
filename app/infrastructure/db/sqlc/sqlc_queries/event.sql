-- name: GetEventByID :one
SELECT * FROM events WHERE id = ?;

-- name: InsertEvent :exec
INSERT INTO events (id, users_id, together, description, date, important) 
VALUES (?, ?, ?, ?, ?, ?);

-- name: DeleteEventByID :exec
DELETE FROM events WHERE id = ?;
