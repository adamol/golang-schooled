package controllers

import "student-management/models"

var semesters = []models.Semester {
    models.Semester { Id: 1, Type: "fall", Year: 2018, StartWeek: 30, EndWeek: 50 },
    models.Semester { Id: 2, Type: "spring", Year: 2019, StartWeek: 2, EndWeek: 22 },
}
