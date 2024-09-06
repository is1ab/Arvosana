-- name: AddStudent :exec
INSERT INTO student (id, semester)
VALUES (?, ?);
