package main

import (
    "fmt"
    "log"
    "net/http"
    //"database/sql"

    "student-management/controllers"

    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)


func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}

func main() {
    // insertDummyData()
    // insertMoreDummyData()

    router := mux.NewRouter().StrictSlash(true);

    router.HandleFunc("/", Index);

    coursesController := controllers.CoursesController{}
    router.HandleFunc("/courses", coursesController.Index);
    router.HandleFunc("/courses/{course_id:[0-9]+}", coursesController.Show);

    log.Fatal(http.ListenAndServe(":8084", router));
}






/**
func insertDummyData() {
    db, err := sql.Open("mysql", "root:root@/go_school");
    if err != nil {
        panic(err.Error());
    }
    defer db.Close();

    insSemester, err := db.Prepare(`
        INSERT INTO semesters (type, year, start_week, end_week) VALUES(?,?,?,?)
    `);
    if err != nil {
        panic(err.Error());
    }
    defer insSemester.Close();

    _, err = insSemester.Exec("fall", 2017, 30, 50);
    if err != nil {
        panic(err.Error());
    }

    insTeacher, err := db.Prepare("INSERT INTO teachers (name) VALUES(?)");
    if err != nil {
        panic(err.Error());
    }
    defer insTeacher.Close();

    _, err = insTeacher.Exec("John Doe");
    if err != nil {
        panic(err.Error());
    }

    insCourse, err := db.Prepare(`
        INSERT INTO courses (name, description, semester_id, teacher_id) VALUES(?,?,?,?)
    `);
    if err != nil {
        panic(err.Error());
    }
    defer insCourse.Close();

    _, err = insCourse.Exec("Calculus", "Derivates etc", 1, 1);
    if err != nil {
        panic(err.Error());
    }
}

func insertMoreDummyData() {
    db, err := sql.Open("mysql", "root:root@/go_school");
    if err != nil {
        panic(err.Error());
    }
    defer db.Close();

    insCourse, err := db.Prepare(`
        INSERT INTO courses (name, description, semester_id, teacher_id) VALUES(?,?,?,?)
    `);
    if err != nil {
        panic(err.Error());
    }
    defer insCourse.Close();

    _, err = insCourse.Exec("Calculus II", "Integrals etc", 1, 1);
    if err != nil {
        panic(err.Error());
    }
}
*/
