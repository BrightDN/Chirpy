-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, email)
VALUES(
	$1,
	CURRENT_TIMESTAMP,
	CURRENT_TIMESTAMP,
	$2)
returning *;

-- name: DeleteUsers :exec
DELETE FROM users;
