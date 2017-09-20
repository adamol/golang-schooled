package models

import (
    "database/sql"
)

type Course struct {
    Id          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    SemesterId  int    `json:"semester_id"`
    TeacherId   int    `json:"teacher_id"`
}

func (c Course) List() []*Course {
    db, _ := sql.Open("mysql", "root:root@/go_school");
    defer db.Close();

    rows, _ := db.Query("SELECT * FROM courses");
    defer rows.Close();

    var courses []*Course;

    for rows.Next() {
        course := new(Course)
        rows.Scan(
            &course.Id, &course.Name, &course.Description, &course.SemesterId, &course.TeacherId,
        );
        courses = append(courses, course)
    }

    return courses
}

func (c Course) FindById(id int) *Course {
    db, _ := sql.Open("mysql", "root:root@/go_school");
    defer db.Close();

    coursesQuery, _ := db.Prepare("SELECT * FROM courses WHERE id=?");
    defer coursesQuery.Close();
	foundCourse := coursesQuery.QueryRow(id)
    course := new(Course)
    foundCourse.Scan(
        &course.Id, &course.Name, &course.Description,
        &course.SemesterId, &course.TeacherId,
    );

    return course

    // result := models.QueryBuilder.FindById(id, "courses")
    // course := new(Course)
    // result.Scan(
    //     &course.Id, &course.Name, &course.Description,
    //     &course.SemesterId, &course.TeacherId,
    // );

    // c := new(Course)
    // models.QueryBuilder.FindById(id, "courses").Scan(
    //     &c.Id, &c.Name, &c.Description, &c.SemesterId, &c.TeacherId,
    // );
    // return c
}

func (c Course) Create(name string, description string, semester_id int, teacher_id int) {
    db, _ := sql.Open("mysql", "root:root@/go_school");
    defer db.Close();

    query, _ := db.Prepare(`
        INSERT INTO courses (name, description, semester_id, teacher_id) VALUES(?,?,?,?)
    `);
    defer query.Close();

    _, err := query.Exec(name, description, semester_id, teacher_id);
    if err != nil {
        panic(err.Error())
    }
}
