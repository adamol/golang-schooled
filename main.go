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

    coursesController := controllers.CoursesController{}
    router.HandleFunc("/courses", coursesController.Index);
    router.HandleFunc("/courses/{course_id:[0-9]+}", coursesController.Show);

    log.Fatal(http.ListenAndServe(":8084", router));
}

