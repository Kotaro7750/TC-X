package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"
)

//Rireki is a structure of each rireki
type Rireki struct {
	ID        int       `json:"id"`
	Joid      int       `json:"joid"`
	Syubetsu  int       `json:"syubetsu" binding:"required"`
	About     string    `json:"about" binding:"required"`
	StartTime time.Time `json:"startTime" binding:"required"`
	EndTime   time.Time `json:"endTime" binding:"required"`
}

//PersonalRirekiAll is a function to select all data of joid from table rireki
//table should decided by client ,but month is hard coded now
func PersonalRirekiAll(db *sql.DB, joid int, year int, month int) ([]*Rireki, error) {
	tableName := strconv.Itoa(year) + "_" + strconv.Itoa(month) + "_rireki"

	rows, err := db.Query(fmt.Sprintf("SELECT id,joid,syubetsu,about,start_time,end_time FROM %s WHERE joid=%d", tableName, joid))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rirekiList []*Rireki

	for rows.Next() {
		var rireki Rireki
		if err := rows.Scan(&rireki.ID, &rireki.Joid, &rireki.Syubetsu, &rireki.About, &rireki.StartTime, &rireki.EndTime); err != nil {
			return nil, err
		}
		loc, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			return nil, err
		}
		startTime, _ := time.ParseInLocation(`2006-01-02T15:04:05`, rireki.StartTime.Format("2006-01-02T15:04:05"), loc)
		endTime, _ := time.ParseInLocation(`2006-01-02T15:04:05`, rireki.EndTime.Format("2006-01-02T15:04:05"), loc)

		rireki.StartTime = startTime
		rireki.EndTime = endTime
		rirekiList = append(rirekiList, &rireki)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rirekiList, nil
}

//RirekiAdd is a function to insert a rireki to table rireki of month
//table should decided by client ,but month is hard coded now
func RirekiAdd(db *sql.DB, rireki Rireki, year int, month int) (*Rireki, error) {
	tableName := strconv.Itoa(year) + "_" + strconv.Itoa(month) + "_rireki"

	formatedStartTime := rireki.StartTime.Format("2006-01-02 15:04:05")
	formatedEndTime := rireki.EndTime.Format("2006-01-02 15:04:05")

	_, err := db.Query(fmt.Sprintf("INSERT INTO `%s` (joid,syubetsu,about,start_time,end_time) VALUES(%d,%d,\"%s\",cast('%s' as DATETIME),cast('%s' as DATETIME))", tableName, rireki.Joid, rireki.Syubetsu, rireki.About, formatedStartTime, formatedEndTime))

	if err != nil {
		return nil, err
	}

	return &rireki, nil
}

//IsContradicted is a function to decide if given rireki's is contradicted
func IsContradicted(db *sql.DB, rireki Rireki, year int, month int) ([]*Rireki, error) {
	tableName := strconv.Itoa(year) + "_" + strconv.Itoa(month) + "_rireki"

	formatedStartTime := rireki.StartTime.Format("2006-01-02 15:04:05")
	formatedEndTime := rireki.EndTime.Format("2006-01-02 15:04:05")

	rows, err := db.Query(fmt.Sprintf("SELECT id,joid,syubetsu,about,start_time,end_time FROM `%s` WHERE (joid = '%d') AND ((`start_time` BETWEEN cast('%s' as DATETIME) AND cast('%s' as DATETIME)) OR (`end_time` BETWEEN cast('%s' as DATETIME) AND cast('%s' as DATETIME)))", tableName, rireki.Joid, formatedStartTime, formatedEndTime, formatedStartTime, formatedEndTime))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contradictedList []*Rireki

	for rows.Next() {
		var contradictedRireki Rireki
		if err := rows.Scan(&contradictedRireki.ID, &contradictedRireki.Joid, &contradictedRireki.Syubetsu, &contradictedRireki.About, &contradictedRireki.StartTime, &contradictedRireki.EndTime); err != nil {
			return nil, err
		}
		contradictedList = append(contradictedList, &contradictedRireki)
	}

	if len(contradictedList) != 0 {
		return contradictedList, nil
	}

	return nil, nil
}

//RirekiDelete is a function to delete rireki
func RirekiDelete(db *sql.DB, year int, month int, joid int, id int) error {
	tableName := strconv.Itoa(year) + "_" + strconv.Itoa(month) + "_rireki"

	res, err := db.Exec(fmt.Sprintf("DELETE FROM `%s` WHERE (id = '%d' AND joid = '%d')", tableName, id, joid))

	if err != nil {
		return err
	}

	if affectedRow, _ := res.RowsAffected(); affectedRow == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

//RirekiUpdate is a function to update rireki
func RirekiUpdate(db *sql.DB, year int, month int, joid int, id int, rireki Rireki) error {
	tableName := strconv.Itoa(year) + "_" + strconv.Itoa(month) + "_rireki"

	formatedStartTime := rireki.StartTime.Format("2006-01-02 15:04:05")
	formatedEndTime := rireki.EndTime.Format("2006-01-02 15:04:05")

	_, err := db.Exec(fmt.Sprintf("UPDATE `%s` SET joid=%d,syubetsu=%d,about='%s',start_time=cast('%s' as DATETIME),end_time=cast('%s' as DATETIME) WHERE joid = '%d' AND id = '%d'", tableName, rireki.Joid, rireki.Syubetsu, rireki.About, formatedStartTime, formatedEndTime, joid, id))

	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}

//RirekiListString is a function to return [[joid,syubetsu,time]...[]] from rireki
func RirekiListString(db *sql.DB, year int, month int) ([][3]string, error) {
	tableName := strconv.Itoa(year) + "_" + strconv.Itoa(month) + "_rireki"
	rows, err := db.Query(fmt.Sprintf("SELECT joid,syubetsu,start_time,end_time FROM %s", tableName))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rirekiList [][3]string
	for rows.Next() {
		var rireki Rireki
		if err := rows.Scan(&rireki.Joid, &rireki.Syubetsu, &rireki.StartTime, &rireki.EndTime); err != nil {
			return nil, err
		}
		loc, _ := time.LoadLocation("Asia/Tokyo")
		startTime, _ := time.ParseInLocation(`2006-01-02T15:04:05`, rireki.StartTime.Format("2006-01-02T15:04:05"), loc)
		endTime, _ := time.ParseInLocation(`2006-01-02T15:04:05`, rireki.EndTime.Format("2006-01-02T15:04:05"), loc)

		time := endTime.Sub(startTime).Minutes()
		rirekiList = append(rirekiList, [3]string{strconv.Itoa(rireki.Joid), strconv.Itoa(rireki.Syubetsu), strconv.Itoa(int(time))})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rirekiList, nil
}
