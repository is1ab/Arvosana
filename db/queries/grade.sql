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
INSERT INTO grade (student_id, homework_id, submitted_at, grade)
VALUES (?, ?, ?, ?);

-- name: GetGradeInfo :many
SELECT
    student.student_id,
    CAST(max(grade.grade) AS REAL) AS grade
FROM grade
INNER JOIN student ON grade.student_id = student.id
INNER JOIN homework ON grade.homework_id = homework.id
WHERE
    homework.name = ? AND
    homework.semester = ?
GROUP BY student.id
ORDER BY student.student_id ASC;
