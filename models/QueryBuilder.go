package models

import "database/sql"
import "strings"
import "fmt"

func list(table string) *sql.Rows {
    db, _ := sql.Open("mysql", "root:root@/go_school")
    defer db.Close()

    sql := "SELECT * FROM " + table
    rows, _ := db.Query(sql)

    return rows
}

func findById(id int, table string) *sql.Row {
    db, _ := sql.Open("mysql", "root:root@/go_school")
    defer db.Close()

    query, _ := db.Prepare("SELECT * FROM " + table + " WHERE id=?")
    defer query.Close()

	return query.QueryRow(id)
}

func insert(cols []string, vals []string, table string) {
    db, _ := sql.Open("mysql", "root:root@/go_school");
    defer db.Close();

    prepared_cols := strings.Join(cols[:], ",")

    prep_val_slice := make([]string, len(vals))
    for i := range prep_val_slice {
        prep_val_slice[i] = "?"
    }
    prepared_vals := strings.Join(prep_val_slice[:], ",")


    sql := "INSERT INTO " + table + "(" + prepared_cols + ") VALUES (" + prepared_vals + ")"
    fmt.Println(sql)
    query, _ := db.Prepare(sql);

    defer query.Close();

    _, err := query.Exec(strings.Join(vals[:], ","));
    if err != nil {
        panic(err.Error())
    }
}
