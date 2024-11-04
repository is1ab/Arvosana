CREATE TABLE student (
    -- https://blog.ploeh.dk/2024/06/03/youll-regret-using-natural-keys/
    id INTEGER PRIMARY KEY,
    student_id TEXT NOT NULL,
    semester SEMESTER_TEXT NOT NULL,

    UNIQUE(student_id, semester)
);

CREATE TABLE homework (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    semester SEMESTER_TEXT NOT NULL,
    begin_at DATETIME_TEXT NOT NULL,
    end_at DATETIME_TEXT NOT NULL,

    UNIQUE(name, semester)
);

CREATE TABLE grade (
    id INTEGER PRIMARY KEY,
    student_id INTEGER NOT NULL,
    homework_id INTEGER NOT NULL,
    submitted_at DATETIME_TEXT NOT NULL,
    grade REAL NOT NULL,

    FOREIGN KEY (student_id) REFERENCES student (id) ON DELETE CASCADE,
    FOREIGN KEY (homework_id) REFERENCES homework (id) ON DELETE CASCADE
);
