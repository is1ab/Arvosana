-- name: GetSubmitInfo :one
SELECT
    student.id AS student_id,
    homework.id AS homework_id,
    homework.begin_at,
    homework.end_at
FROM student
INNER JOIN homework ON student.semester = homework.semester
WHERE
    student.student_id = ? AND
    student.semester = ? AND
    homework.name = ?;

-- name: SubmitGrade :exec
INSERT INTO grade (student_id, homework_id, created_at, grade)
VALUES (?, ?, datetime('now'), ?)
