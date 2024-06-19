-- name: CreateUser :one
INSERT INTO users(id , name , created_at , updated_at)
VALUES ($1 , $2 , $3 , $4)
RETURNING *;


-- name: IssueApiKey :one
UPDATE users SET api_key = $1 WHERE id = $2
RETURNING api_key ;



-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;


