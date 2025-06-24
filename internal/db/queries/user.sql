-- name: CreateUser :one
INSERT INTO users (username, chat_id, created_at)
VALUES ($1, $2, NOW())
RETURNING *;

-- name: GetUserByChatId :one
SELECT * FROM users
WHERE chat_id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: UpdateChatId :one
UPDATE users
SET chat_id = $2,
    updated_at = NOW()
WHERE username = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE chat_id = $1;