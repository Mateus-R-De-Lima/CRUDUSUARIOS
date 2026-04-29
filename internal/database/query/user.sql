-- name: CreateUser :exec
INSERT INTO users (id, first_name, last_name, biography)
VALUES ($1, $2, $3, $4);

-- name: GetUser :one
SELECT id, first_name, last_name, biography
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, first_name, last_name, biography
FROM users;

-- name: UpdateUser :exec
UPDATE users
SET first_name = $2,
    last_name = $3,
    biography = $4
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;