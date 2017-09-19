package models

type CourseStudent struct {
    Id        int `json:"id"`
    CourseId  int `json:"course_id"`
    StudentId int `json:"student_id"`
}

