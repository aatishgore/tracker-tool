package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mutecomm/go-sqlcipher"
)

func openDb() *sql.DB {
	var database *sql.DB
	key := "3DD29CA851E7B56E4697F0E1F08507293D761Z05CE4D1B628663F411A8086D99"

	dbname := fmt.Sprintf("./wfh.db?_pragma_key=x'%s'&_pragma_cipher_page_size=4096", key)
	// if debug is true then don't encrypt database
	if debug {
		database, _ = sql.Open("sqlite3", "./wfh.db")
	} else {
		database, _ = sql.Open("sqlite3", dbname)
	}
	return database
}

// TODO: if more than one table create model
func logUserActivityInDB(window string, time int) {
	database := openDb()
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, window TEXT, time TEXT,timestamp DATE DEFAULT (datetime('now','localtime')))")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO people (window, time) VALUES (?, ?)")
	statement.Exec(window, time)
	defer database.Close()
}

func copyToLog() {
	database := openDb()
	rows, _ := database.Query("SELECT window, date(timestamp) as day ,sum(time) as time FROM people group by window, day")

	var (
		window string
		time   string
		day    string
	)
	for rows.Next() {
		rows.Scan(&window, &day, &time)
		if logInfo {
			logger.Println("date:" + day + "window :" + window + " time: " + time)
		}
	}
	statement, _ := database.Prepare("DELETE FROM people")
	statement.Exec()
	defer database.Close()
}
