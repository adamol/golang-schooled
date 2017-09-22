package models

import "strconv"

type Course struct {
    Id          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    SemesterId  int    `json:"semester_id"`
    TeacherId   int    `json:"teacher_id"`
}

func (self Course) List() []*Course {
    rows := list("courses")
    defer rows.Close()

    var courses []*Course;

    for rows.Next() {
        c := new(Course)
        rows.Scan(
            &c.Id, &c.Name, &c.Description, &c.SemesterId, &c.TeacherId,
        );
        courses = append(courses, c)
    }

    return courses
}

func (self Course) FindById(id int) *Course {
    c := new(Course)
    findById(id, "courses").Scan(
        &c.Id, &c.Name, &c.Description, &c.SemesterId, &c.TeacherId,
    );
    return c
}

func (c Course) Create(name string, description string, semester_id int, teacher_id int) {
    cols := []string{"name", "description", "semester_id", "teacher_id"}

    sid := strconv.Itoa(semester_id)
    tid  := strconv.Itoa(teacher_id)
    vals := []string{name, description, sid, tid}

    insert(cols, vals, "courses")
    // db, _ := sql.Open("mysql", "root:root@/go_school");
    // defer db.Close();

    // query, _ := db.Prepare(`
    //     INSERT INTO courses (name, description, semester_id, teacher_id) VALUES(?,?,?,?)
    // `);
    // defer query.Close();

    // _, err := query.Exec(name, description, semester_id, teacher_id);
    // if err != nil {
    //     panic(err.Error())
    // }
}
