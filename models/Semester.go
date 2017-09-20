package models

import "database/sql"

type Semester struct {
    Id        int    `json:"id"`
    Type      string `json:"name"`
    Year      int    `json:"year"`
    StartWeek int    `json:"start_week"`
    EndWeek   int    `json:"end_week"`
}

func (s Semester) FindById(id int) *Semester {
    db, _ := sql.Open("mysql", "root:root@/go_school");
    defer db.Close();

    semestersQuery, _ := db.Prepare("SELECT * FROM semesters WHERE id=?");
    defer semestersQuery.Close();
	foundSemester := semestersQuery.QueryRow(id)
    semester := new(Semester)
	foundSemester.Scan(
	    &semester.Id, &semester.Type, &semester.Year,
	    &semester.StartWeek, &semester.EndWeek,
	);

    return semester
}
