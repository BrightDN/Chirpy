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

-- name: AlterUserData :one
UPDATE users
SET
email = $1,
hashed_password = $2,
updated_at = CURRENT_TIMESTAMP
WHERE id = $3
RETURNING *;
