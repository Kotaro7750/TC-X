package model

import (
	"database/sql"
	"fmt"
	"strconv"
)

//NoteAll is a function to select all data of joid from table rireki
func NoteAll(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SHOW TABLES like '%_rireki'")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var noteList []string

	for rows.Next() {
		var noteName string
		if err := rows.Scan(&noteName); err != nil {
			return nil, err
		}
		noteList = append(noteList, noteName)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return noteList, nil
}

//NoteAdd is a function to insert a note
func NoteAdd(db *sql.DB, year int, month int) (string, error) {
	tableName := strconv.Itoa(year) + "_" + strconv.Itoa(month) + "_rireki"
	_, err := db.Query(fmt.Sprintf("CREATE TABLE `%s`(id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,joid INT NOT NULL,syubetsu INT NOT NULL,about TEXT NOT NULL,start_time DATETIME NOT NULL,end_time DATETIME NOT NULL)", tableName))

	if err != nil {
		return "", err
	}

	return tableName, nil
}

//NoteDelete is a function to delete note
func NoteDelete(db *sql.DB, year int, month int) error {
	tableName := strconv.Itoa(year) + "_" + strconv.Itoa(month) + "_rireki"

	_, err := db.Exec(fmt.Sprintf("DROP TABLE `%s`", tableName))

	if err != nil {
		return err
	}

	return nil
}
