-- name: AddHomework :exec
INSERT INTO homework (name, semester, deadline)
VALUES (?, ?, ?);

-- name: GetAllHomeworks :many
SELECT name, semester, deadline FROM homework;
