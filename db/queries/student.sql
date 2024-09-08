-- name: AddStudent :exec
INSERT INTO student (student_id, semester)
VALUES (?, ?);

-- name: GetAllStudents :many
SELECT student_id, semester FROM student;

-- name: GetStudentsBySemester :many
SELECT student_id FROM student
WHERE semester = ?;
