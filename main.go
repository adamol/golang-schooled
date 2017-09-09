package main

import (
    "log"
    "fmt"
    "strconv"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
)

type Semester struct {
    Id        int    `json:"id"`
    Type      string `json:"name"`
    Year      int    `json:"year"`
    StartWeek int    `json:"start_week"`
    EndWeek   int    `json:"end_week"`
}

var semesters = []Semester {
    Semester { Id: 1, Type: "fall", Year: 2018, StartWeek: 30, EndWeek: 50 },
    Semester { Id: 2, Type: "spring", Year: 2019, StartWeek: 2, EndWeek: 22 },
}

// A semester has-many courses
// A course belongs to a semester

// A teacher has-many courses
// A course belongs-to a teacher
type Course struct {
    Id          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    SemesterId  int    `json:"semester_id"`
    TeacherId   int    `json:"teacher_id"`
}

var courses = []Course {
    Course {
        Id: 1, Name: "Calculus", Description: "Derivatives etc",
    },
    Course {
        Id: 2, Name: "English", Description: "Poems, Book Reports, etc",
    },
    Course {
        Id: 3, Name: "Calculus 2", Description: "Integrals etc",
    },
}

// A course has-many lectures
// A lecture belongs to a course

type Lecture struct {
    Id        int      `json:"id"`
    Week      int      `json:"week"`
    Day       int      `json:"day"`
    StartTime int      `json:"start_time"`
    EndTime   int      `json:"end_time"`
    Notes     []string `json:"notes"`
    CourseId  int      `json:"course_id"`
}

var lectures = []Lecture {
    Lecture {Id: 1, Week: 1, Day: 2, StartTime: 10, EndTime: 11},
    Lecture {Id: 2, Week: 1, Day: 4, StartTime: 10, EndTime: 11},
    Lecture {Id: 3, Week: 2, Day: 2, StartTime: 10, EndTime: 11},
    Lecture {Id: 4, Week: 2, Day: 4, StartTime: 10, EndTime: 11},
    Lecture {Id: 5, Week: 1, Day: 1, StartTime: 10, EndTime: 11},
}

type User struct {
    Id           int    `json:"id"`
    Email        string `json:"email"`
    Password     string `json:"password"`
    UserableType string `json:"userable_type"` // either teacher or student
    UserableId   int    `json:"userable_id"`   // polymorphic relationship
}

// a user is something which logs in to the system
// each student/teacher in the system CAN have a login
// however they do not have to => nullable

var users = []User {
    User {
        Id: 1,
        Email: "john@teacher.com", Password: "secret",
        UserableType: "teacher", UserableId: 1,
    },
    User {
        Id: 2,
        Email: "jane@teacher.com", Password: "secret",
        UserableType: "teacher", UserableId: 2,
    },
    User {
        Id: 3,
        Email: "zikki@student.com", Password: "secret",
        UserableType: "student", UserableId: 1,
    },
    User {
        Id: 4,
        Email: "fizz@student.com", Password: "secret",
        UserableType: "student", UserableId: 3,
    },
}

type Teacher struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

var teachers = []Teacher {
    Teacher { Id: 1, Name: "John Doe" },
    Teacher { Id: 2, Name: "Jane Snow" },
}

type Student struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

var students = []Student {
    Student { Id: 1, Name: "Zikki Kiriki" },
    Student { Id: 2, Name: "Botwa Malaqua" },
    Student { Id: 3, Name: "Fizz Mcbuckton" },
    Student { Id: 4, Name: "Charlie Folargton" },
    Student { Id: 5, Name: "Zimba Bilawe" },
}

// A student has-many courses
// A course has-many students
type CourseStudent struct {
    Id        int `json:"id"`
    CourseId  int `json:"course_id"`
    StudentId int `json:"student_id"`
}

/*
teachers:
    important --> GET  courses              // list of all courses, filter theirs
    important --> GET  courses/{courseId}   // lists description, lectures
    GET  lectures/{lectureId} // click on a lecture to see notes/material
    POST courses
    POST courses/{courseId}/students
    POST courses/{courseId}/lectures
    PUT  lectures/{lectureId}
students:
    important --> GET  courses              // list of all courses, filter theirs
    important --> GET  courses/{courseId}   // lists description, lectures
    GET  lectures/{lectureId} // click on a lecture to see notes/material
*/

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}

func CourseIndex(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(courses)
}

func CourseShow(w http.ResponseWriter, r *http.Request) {
    course_id := mux.Vars(r)["student_id"]

	cid, err := strconv.Atoi(course_id)
    if err != nil {
        fmt.Println(err)
    }

    json.NewEncoder(w).Encode(courses[cid])
}

func main() {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/", Index)

    router.HandleFunc("/courses", CourseIndex);
    router.HandleFunc("/courses/{course_id:[0-9]+}", CourseShow);

    log.Fatal(http.ListenAndServe(":8084", router))
}
