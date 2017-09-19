package models

type Semester struct {
    Id        int    `json:"id"`
    Type      string `json:"name"`
    Year      int    `json:"year"`
    StartWeek int    `json:"start_week"`
    EndWeek   int    `json:"end_week"`
}
