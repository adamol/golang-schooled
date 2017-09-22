package controllers

import (
//    "net/http"
    "student-management/models"
//    "encoding/json"
)

type LecturesController struct {}

var lectures = []models.Lecture {
    models.Lecture {Id: 1, Week: 1, Day: 2, StartTime: 10, EndTime: 11},
    models.Lecture {Id: 2, Week: 1, Day: 4, StartTime: 10, EndTime: 11},
    models.Lecture {Id: 3, Week: 2, Day: 2, StartTime: 10, EndTime: 11},
    models.Lecture {Id: 4, Week: 2, Day: 4, StartTime: 10, EndTime: 11},
    models.Lecture {Id: 5, Week: 1, Day: 1, StartTime: 10, EndTime: 11},
}

// // func (c LecturesController) Index(w http.ResponseWriter, r *http.Request) {
// //     switch r.Method {
// //     case "GET":
// //         lectures := models.Lecture{}.List()
// //
// //         json.NewEncoder(w).Encode(lectures)
// //     case "POST":
// //         models.Lecture{}.Create(
// //             r.FormValue("week"),
// //             r.FormValue("day"),
// //             r.FormValue("start_time"),
// //             r.FormValue("end_time"),
// //         )
// //     }
// // }
// //
// // func (c LecturesController) Index(w http.ResponseWriter, r *http.Request) {
// //     lecture_id, _ := strconv.Atoi(mux.Vars(r)["lecture_id"])
// //
// //     lecture   := models.Lecture{}.FindById(lecture_id)
// //
// //     json.NewEncoder(w).Encode(lecture)
// // }
