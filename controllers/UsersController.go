package controllers

import "student-management/models"

var users = []models.User {
    models.User {
        Id: 1,
        Email: "john@teacher.com", Password: "secret",
        UserableType: "teacher", UserableId: 1,
    },
    models.User {
        Id: 2,
        Email: "jane@teacher.com", Password: "secret",
        UserableType: "teacher", UserableId: 2,
    },
    models.User {
        Id: 3,
        Email: "zikki@student.com", Password: "secret",
        UserableType: "student", UserableId: 1,
    },
    models.User {
        Id: 4,
        Email: "fizz@student.com", Password: "secret",
        UserableType: "student", UserableId: 3,
    },
}
