-- name: AddHomework :exec
INSERT INTO homework (name, deadline)
VALUES (?, ?);

-- name: GetAllHomeworks :many
SELECT name, created_at, deadline FROM homework;
