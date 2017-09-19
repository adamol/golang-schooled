package controllers

import (
    "fmt"
    "strconv"
    "net/http"
    "encoding/json"

    "student-management/models"

    "github.com/gorilla/mux"
)

var courses = []models.Course {
    models.Course { Id: 1, Name: "Calculus", Description: "Derivatives etc" },
    models.Course { Id: 2, Name: "English", Description: "Poems, Book Reports, etc" },
    models.Course { Id: 3, Name: "Calculus 2", Description: "Integrals etc" },
}

type CoursesController struct {}

func (controller CoursesController) Index(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(courses)
}

func (controller CoursesController) Show(w http.ResponseWriter, r *http.Request) {
    course_id := mux.Vars(r)["student_id"]

	cid, err := strconv.Atoi(course_id)
    if err != nil {
        fmt.Println(err)
    }

    var course = courses[cid]

    json.NewEncoder(w).Encode(course)
}




/**
func allCourses() {
    db, err := sql.Open("mysql", "root:root@/go_school");
    if err != nil {
        panic(err.Error());
    }
    defer db.Close();

    rows, err := db.Query("SELECT * FROM courses");
    if err != nil {
        panic(err.Error());
    }

    cols, err := rows.Columns()
    if err != nil {
        panic(err.Error());
    }

    values   := make([]sql.RawBytes, len(cols))
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    for rows.Next() {
        err = rows.Scan(scanArgs...)
        if err != nil {
            panic(err.Error());
        }

        var value string
        for i, col := range values {
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
            fmt.Println(cols[i], ": ", value)
        }
        fmt.Println("--------------------------------")
    }
    if err = rows.Err(); err != nil {
        panic(err.Error())
    }

	// var products []*Product
	// for rows.Next() {
	// 	p := new(Product)
	// 	if err := rows.Scan(&p.ID, &p.Name, &p.IsMatch, &p.Created); err != nil { ... }
	// 	products = append(products, p)
	// }
	// if err := rows.Err() { ... }
}

func findCourseById(id int) {
    db, err := sql.Open("mysql", "root:root@/go_school");
    if err != nil {
        panic(err.Error());
    }

    selCourse, err := db.Prepare("SELECT * FROM courses WHERE id=?");
    if err != nil {
        panic(err.Error());
    }
    defer selCourse.Close();

	row := selCourse.QueryRow(id)

	var course_id int;
	var name string;
	var description string;
	var semester_id int;
	var teacher_id int;
	row.Scan(&course_id, &name, &description, &semester_id, &teacher_id);
    fmt.Println("id: ", course_id, "name: ", name, "description: ", description);

    selTeacher, err := db.Prepare("SELECT * FROM teachers WHERE id=?");
    if err != nil {
        panic(err.Error());
    }
    defer selTeacher.Close();
	teacher := selTeacher.QueryRow(teacher_id)

	var tid int;
	var teacher_name string;
	teacher.Scan(&tid, &teacher_name);
    fmt.Println("teacher: ", teacher_name);

    selSemester, err := db.Prepare("SELECT * FROM semesters WHERE id=?");
    if err != nil {
        panic(err.Error());
    }
    defer selSemester.Close();
	semester := selSemester.QueryRow(teacher_id)
	var sid int;
	var semester_type string;
	var semester_year int;
	var semester_start_week int;
	var semester_end_week int;
	semester.Scan(&sid, &semester_type, &semester_year, &semester_start_week, &semester_end_week);
    fmt.Println("semester: ", semester_type, "|", semester_year);
}
*/
