-- name: CreateMessage :one
INSERT INTO messages (message)
VALUES (?)
RETURNING *;

-- name: GetMessage :one
SELECT * FROM messages
WHERE id = ? LIMIT 1;

-- name: UpdateMessage :one
UPDATE messages
SET message = ?
WHERE id = ?
RETURNING *;

-- name: DeleteMessage :exec
DELETE FROM messages
WHERE id = ?;