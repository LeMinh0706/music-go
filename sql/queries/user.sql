-- name: CreateUser :one
INSERT INTO users(
    username,
    password,
    fullname,
    email, 
    gender,
    role_id,
    avt
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;