-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password) 
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: DeleteAll :exec
DELETE FROM users;

-- name: GetUserFromEmail :one
SELECT * FROM users
WHERE email = $1; 

-- name: UpdateUserInformation :one
UPDATE users
SET email = $1, hashed_password = $2
WHERE id = $3
RETURNING *; 

-- name: UpgradeUserFromId :exec
UPDATE users
SET is_chirpy_red = true
WHERE id = $1;