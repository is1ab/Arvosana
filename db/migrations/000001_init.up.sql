CREATE TABLE student (
    -- https://blog.ploeh.dk/2024/06/03/youll-regret-using-natural-keys/
    id INTEGER PRIMARY KEY,
    student_id TEXT UNIQUE NOT NULL,
    semester TEXT NOT NULL
);

CREATE TABLE homework (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    semester TEXT NOT NULL,
    deadline TEXT NOT NULL,

    UNIQUE(name, semester)
);

CREATE TABLE student_grade (
    id INTEGER PRIMARY KEY,
    student_id INTEGER NOT NULL,
    homework_id INTEGER NOT NULL,
    created_at TEXT NOT NULL,
    grade REAL NOT NULL,

    FOREIGN KEY (student_id) REFERENCES student (id),
    FOREIGN KEY (homework_id) REFERENCES homework (id)
);
