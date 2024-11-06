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
    max(grade.grade) AS grade
FROM homework
CROSS JOIN student
LEFT JOIN grade ON
    homework.id = grade.homework_id AND
    student.id = grade.student_id
WHERE
    homework.semester = ? AND
    homework.name = ?
GROUP BY student.id
ORDER BY student.student_id ASC;

-- name: GetStudentSubmitHistory :many
SELECT
    grade.grade,
    grade.submitted_at
FROM grade
INNER JOIN student ON grade.student_id = student.id
INNER JOIN homework ON grade.homework_id = homework.id
WHERE
    homework.semester = ? AND
    homework.name = ? AND
    student.student_id = ?
ORDER BY grade.submitted_at DESC;
