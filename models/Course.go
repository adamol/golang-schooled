package models

type Course struct {
    Id          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    SemesterId  int    `json:"semester_id"`
    TeacherId   int    `json:"teacher_id"`
}
