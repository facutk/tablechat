-- name: CreateMessage :one
INSERT INTO messages (message)
VALUES (?)
RETURNING *;

-- name: GetMessage :one
SELECT * FROM messages
LIMIT 1;

-- name: UpdateMessage :one
UPDATE messages
SET message = ?
RETURNING *;

-- name: DeleteMessage :exec
DELETE FROM messages;

-- name: UpsertMessage :one
INSERT INTO messages (message)
VALUES (?)
ON CONFLICT DO UPDATE SET
  message = excluded.message,
  updated_at = CURRENT_TIMESTAMP
RETURNING *;