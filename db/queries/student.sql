-- name: AddStudent :exec
INSERT INTO student (student_id, semester)
VALUES (?, ?);

-- name: UpdateStudent :exec
UPDATE student
SET
    student_id = @new_student_id,
    semester = @new_semester
WHERE
    student_id = @old_studend_id AND semester = @old_semester;


-- name: DeleteStudent :exec
PRAGMA foreign_keys = ON;
DELETE FROM student
WHERE student_id = ? AND semester = ?;

-- name: GetAllStudents :many
SELECT student_id, semester FROM student;

-- name: GetStudentsBySemester :many
SELECT student_id FROM student
WHERE semester = ?;

-- name: GetStudentInfo :many
SELECT
    homework.name,
    grade.submitted_at,
    CAST(max(grade.grade) AS REAL) AS grade
FROM grade
INNER JOIN student ON grade.student_id = student.id
INNER JOIN homework ON grade.homework_id = homework.id
WHERE
    student.student_id = ? AND
    student.semester = ?
GROUP BY homework.id
ORDER BY grade.submitted_at DESC;
