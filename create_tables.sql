CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    userable_type ENUM("teacher", "student"),
    userable_id INT NOT NULL,

    PRIMARY KEY(id)
);

CREATE TABLE teachers (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,

    PRIMARY KEY(id)
);

CREATE TABLE students (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,

    PRIMARY KEY(id)
);

CREATE TABLE semesters (
    id INT NOT NULL AUTO_INCREMENT,
    type ENUM("spring", "fall"),
    year INT(4) NOT NULL,
    start_week int(2) NOT NULL,
    end_week int(2) NOT NULL,

    PRIMARY KEY(id)
);

CREATE TABLE courses (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    semester_id INT NOT NULL,
    teacher_id INT NOT NULL,

    PRIMARY KEY(id),
    FOREIGN KEY (semester_id) REFERENCES semesters(id),
    FOREIGN KEY (teacher_id) REFERENCES teachers(id)
);

CREATE TABLE lectures (
    id INT NOT NULL AUTO_INCREMENT,
    week INT(2) NOT NULL,
    day ENUM('1','2','3','4','5','6','7'),
    start_time DATETIME NOT NULL,
    notes TEXT,
    end_time DATETIME NOT NULL,
    course_id INT NOT NULL,

    PRIMARY KEY(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);
