package main

import (
    "fmt"
    "log"
    "net/http"

    "student-management/controllers"

    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)


func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}

func main() {
    router := mux.NewRouter().StrictSlash(true);

    router.HandleFunc("/", Index);

    router.HandleFunc("/courses", controllers.CoursesController{}.Index);
    router.HandleFunc("/courses/{course_id:[0-9]+}", controllers.CoursesController{}.Show);

    // router.HandleFunc("/lectures", controllers.LecturesController{}.Index);
    // router.HandleFunc("/lectures/{lecture_id:[0-9]+}", controllers.LecturesController{}.Show);

    log.Fatal(http.ListenAndServe(":8084", router));
}

