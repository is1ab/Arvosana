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
    grade.grade
FROM homework
CROSS JOIN student
LEFT JOIN grade ON
    homework.id = grade.homework_id AND
    student.id = grade.student_id
WHERE
    student.student_id = ? AND
    student.semester = ? AND
    homework.begin_at < @before
GROUP BY homework.id
ORDER BY homework.begin_at DESC;
