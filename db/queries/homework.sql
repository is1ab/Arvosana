-- name: AddHomework :exec
INSERT INTO homework (name, semester, begin_at, end_at)
VALUES (?, ?, ?, ?);

-- name: GetAllHomeworks :many
SELECT name, semester, begin_at, end_at FROM homework
ORDER BY begin_at DESC;

-- name: GetHomeworksFromSemester :many
SELECT name, begin_at, end_at FROM homework
WHERE semester = ?
ORDER BY begin_at DESC;
