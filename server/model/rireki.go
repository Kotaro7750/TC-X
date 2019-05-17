package model

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

//Rireki is a structure of each rireki
type Rireki struct {
	Joid      int
	Syubetsu  int `json:"syubetsu" binding:"required"`
	About     string `json:"about" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
}

//PersonalRirekiAll is a function to select all data of joid from table rireki
//table should decided by client ,but month is hard coded now
func PersonalRirekiAll(db *sql.DB, joid int, month int) ([]*Rireki, error) {
	tableName := strconv.Itoa(month) + "_rireki"

	rows, err := db.Query(fmt.Sprintf("SELECT joid,syubetsu,about,start_time,end_time FROM %s WHERE joid=%d", tableName, joid))

	if err != nil {
		return nil, err
	}

	var rirekiList []*Rireki

	for rows.Next() {
		var rireki Rireki
		if err := rows.Scan(&rireki.Joid, &rireki.Syubetsu, &rireki.About, &rireki.StartTime, &rireki.EndTime); err != nil {
			return nil, err
		}
		rirekiList = append(rirekiList, &rireki)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rirekiList, nil
}

//RirekiAdd is a function to insert a rireki to table rireki of month
//table should decided by client ,but month is hard coded now
func RirekiAdd(db *sql.DB,rireki Rireki,month int)(*Rireki,error)  {
	tableName := strconv.Itoa(month) + "_rireki"

	formatedStartTime := rireki.StartTime.Format("2006-01-02 15:04:05")
	formatedEndTime := rireki.EndTime.Format("2006-01-02 15:04:05")
	
	_, err := db.Query(fmt.Sprintf("INSERT INTO `%s` (joid,syubetsu,about,start_time,end_time) VALUES(%d,%d,\"%s\",cast('%s' as DATETIME),cast('%s' as DATETIME))", tableName, rireki.Joid,rireki.Syubetsu,rireki.About,formatedStartTime,formatedEndTime))

	if err != nil {
		return nil, err
	}

	return &rireki, nil
}
