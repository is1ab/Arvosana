-- name: AddHomework :exec
INSERT INTO homework (name, semester, begin_at, end_at)
VALUES (?, ?, ?, ?);

-- name: GetAllHomeworks :many
SELECT name, semester, begin_at, end_at FROM homework;
