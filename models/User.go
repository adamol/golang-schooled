package models

type User struct {
    Id           int    `json:"id"`
    Email        string `json:"email"`
    Password     string `json:"password"`
    UserableType string `json:"userable_type"` // either teacher or student
    UserableId   int    `json:"userable_id"`   // polymorphic relationship
}

