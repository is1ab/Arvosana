-- name: AddHomework :exec
INSERT INTO homework (name, semester, begin_at, end_at)
VALUES (?, ?, ?, ?);

-- name: GetAllHomeworks :many
SELECT name, semester, begin_at, end_at FROM homework
ORDER BY begin_at DESC;

-- name: UpdateHomework :exec
UPDATE homework
SET
    name        = COALESCE(sqlc.narg(new_name), name),
    semester    = COALESCE(sqlc.narg(new_semester), semester),
    begin_at    = COALESCE(sqlc.narg(new_begin_at), begin_at),
    end_at      = COALESCE(sqlc.narg(new_end_at), end_at)
WHERE
    semester = @old_semester AND
    name = @old_name;

-- name: GetHomeworksFromSemester :many
SELECT name, begin_at, end_at FROM homework
WHERE semester = ?
ORDER BY begin_at DESC;

-- name: GetHomeworkInfo :one
SELECT semester, name, begin_at, end_at FROM homework
WHERE
    semester = ? AND
    name = ?;
