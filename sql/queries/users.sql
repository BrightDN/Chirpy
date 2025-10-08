-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, email, hashed_password)
VALUES(
	$1,
	CURRENT_TIMESTAMP,
	CURRENT_TIMESTAMP,
	$2,
	$3)
returning *;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUser :one
SELECT *
FROM users
WHERE email = $1;