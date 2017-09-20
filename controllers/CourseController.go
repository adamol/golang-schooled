package controllers

import (
    "strconv"
    "net/http"
    "encoding/json"

    "student-management/models"

    "github.com/gorilla/mux"
)

type CoursesController struct {}

func (controller CoursesController) Index(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        courses := models.Course{}.List()

        json.NewEncoder(w).Encode(courses)
    case "POST":
        semester_id, _ := strconv.Atoi(r.FormValue("semester_id"))
        teacher_id , _ := strconv.Atoi(r.FormValue("teacher_id"))

        models.Course{}.Create(
            r.FormValue("name"),
            r.FormValue("description"),
            semester_id,
            teacher_id,
        )

    }
}

func (controller CoursesController) Show(w http.ResponseWriter, r *http.Request) {
    course_id, _ := strconv.Atoi(mux.Vars(r)["course_id"])

    course   := models.Course{}.FindById(course_id)
    teacher  := models.Teacher{}.FindById(course.TeacherId)
    semester := models.Semester{}.FindById(course.SemesterId)

    json.NewEncoder(w).Encode(struct {
        Id int
        Name string
        Description string
        Teacher *models.Teacher
        Semester *models.Semester
    } {course.Id, course.Name, course.Description, teacher, semester})
}

