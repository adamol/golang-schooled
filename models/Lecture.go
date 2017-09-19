package models

type Lecture struct {
    Id        int      `json:"id"`
    Week      int      `json:"week"`
    Day       int      `json:"day"`
    StartTime int      `json:"start_time"`
    EndTime   int      `json:"end_time"`
    Notes     []string `json:"notes"`
    CourseId  int      `json:"course_id"`
}


