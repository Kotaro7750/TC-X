package model

import (
	"database/sql"
	"errors"
	"fmt"
)

//Syubetsu is a structure of each syubetsu
type Syubetsu struct {
	Name     string `json:"name" binding:"required"`
	Syubetsu int    `json:"syubetsu" binding:"required"`
	Salary   int    `json:"salary" binding:"required"`
}

//SyubetsuAll is a function to select all data of joid from table rireki
func SyubetsuAll(db *sql.DB) ([]*Syubetsu, error) {
	tableName := "syubetsu"

	rows, err := db.Query(fmt.Sprintf("SELECT name,syubetsu,salary FROM %s ", tableName))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var syubetsuList []*Syubetsu

	for rows.Next() {
		var syubetsu Syubetsu
		if err := rows.Scan(&syubetsu.Name, &syubetsu.Syubetsu, &syubetsu.Salary); err != nil {
			return nil, err
		}
		syubetsuList = append(syubetsuList, &syubetsu)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return syubetsuList, nil
}

//SyubetsuAdd is a function to insert a syubetsu
func SyubetsuAdd(db *sql.DB, syubetsu Syubetsu) (*Syubetsu, error) {
	tableName := "syubetsu"
	_, err := db.Query(fmt.Sprintf("INSERT INTO `%s` (name,syubetsu,salary) VALUES('%s',%d,%d)", tableName, syubetsu.Name, syubetsu.Syubetsu, syubetsu.Salary))

	if err != nil {
		return nil, err
	}

	return &syubetsu, nil
}

//SyubetsuDelete is a function to delete syubetsu
func SyubetsuDelete(db *sql.DB, syubetsu int) error {
	tableName := "syubetsu"

	res, err := db.Exec(fmt.Sprintf("DELETE FROM `%s` WHERE syubetsu = '%d'", tableName, syubetsu))

	if err != nil {
		return err
	}

	if affectedRow, _ := res.RowsAffected(); affectedRow == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

//SyubetsuUpdate is a function to update syubetsu
func SyubetsuUpdate(db *sql.DB, syubetsu Syubetsu) error {
	tableName := "syubetsu"
	_, err := db.Exec(fmt.Sprintf("UPDATE `%s` SET name='%s',syubetsu=%d,salary=%d WHERE syubetsu = '%d'", tableName, syubetsu.Name, syubetsu.Syubetsu, syubetsu.Salary, syubetsu.Syubetsu))

	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}
