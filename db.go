package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func openDb() *sql.DB {
	database, _ := sql.Open("sqlite3", "./wfh.db")
	return database
}

func logToDB(window string, time float64) {
	database := openDb()
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, window TEXT, time TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO people (window, time) VALUES (?, ?)")
	statement.Exec(window, time)
	defer database.Close()
}

func copyToLog() {
	database := openDb()
	rows, _ := database.Query("SELECT window, sum(time) as time FROM people group by window")

	var window string
	var time string
	for rows.Next() {
		rows.Scan(&window, &time)
		logger.Println("window :" + window + " time: " + time)
	}
	statement, _ := database.Prepare("DELETE FROM people")
	statement.Exec()
	defer database.Close()
}
