-- name: AddHomework :exec
INSERT INTO homework (name, created_at, deadline)
VALUES (?, datetime('now', 'localtime'), ?);

-- name: GetAllHomeworks :many
SELECT name, created_at, deadline FROM homework;
