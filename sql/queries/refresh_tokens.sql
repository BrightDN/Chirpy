-- name: CreateRefreshToken :one

INSERT INTO refresh_tokens(token, updated_at, created_at, expires_at, revoked_at, user_id)
VALUES (
	$1,
	CURRENT_TIMESTAMP,
	CURRENT_TIMESTAMP,
	CURRENT_TIMESTAMP + INTERVAL '60 DAY',
	NULL,
	$2)
RETURNING *;

-- name: GetUserFromToken :one

SELECT *
FROM refresh_tokens
WHERE token = $1;

-- name: RevokeToken :exec
UPDATE refresh_tokens
SET revoked_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE token = $1
  AND revoked_at IS NULL;
