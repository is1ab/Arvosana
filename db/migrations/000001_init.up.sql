CREATE TABLE student (
    id TEXT UNIQUE NOT NULL,
    semester TEXT NOT NULL
);

CREATE TABLE homework (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    created_at INTEGER DEFAULT current_timestamp,
    deadline INTEGER NOT NULL
);

CREATE TABLE student_grade (
    student_id TEXT NOT NULL,
    homework_id INTEGER NOT NULL,
    grade INTEGER NOT NULL,

    FOREIGN KEY (student_id) REFERENCES student (id),
    FOREIGN KEY (homework_id) REFERENCES homework (id),

    UNIQUE(student_id, homework_id)
);
