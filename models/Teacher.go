package models

import "database/sql"

type Teacher struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

func (t Teacher) FindById(id int) *Teacher {
    db, _ := sql.Open("mysql", "root:root@/go_school");
    defer db.Close();

    teachersQuery, _ := db.Prepare("SELECT * FROM teachers WHERE id=?");
    defer teachersQuery.Close();
	foundTeacher := teachersQuery.QueryRow(id)
    teacher := new(Teacher)
	foundTeacher.Scan(&teacher.Id, &teacher.Name);

    return teacher
}

