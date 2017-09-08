package main

import (
    "log"
    "fmt"
    "strconv"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
)

type Student struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

var students = []Student {
    Student { Id: 1, Name: "John Doe" },
    Student { Id: 2, Name: "Jane Snow" },
    Student { Id: 3, Name: "Fizz Mcbuckton" },
    Student { Id: 4, Name: "Charlie Folargton" },
    Student { Id: 5, Name: "Zimba Bilawe" },
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}

func StudentsIndex(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(students)
}

func StudentsShow(w http.ResponseWriter, r *http.Request) {
    student_id := mux.Vars(r)["student_id"]

	sid, err := strconv.Atoi(student_id)
    if err != nil {
        fmt.Println(err)
    }

    json.NewEncoder(w).Encode(students[sid])
}

func main() {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/", Index)
    router.HandleFunc("/students", StudentsIndex)
    router.HandleFunc("/students/{student_id}", StudentsShow)

    log.Fatal(http.ListenAndServe(":8084", router))
}
