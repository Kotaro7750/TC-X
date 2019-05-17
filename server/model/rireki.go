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
	Syubetsu  int
	About     string
	StartTime time.Time
	EndTime   time.Time
}

//PersonalRirekiAll is a function to select all data of joid from table rireki
//table should decided by client ,but month is hard coded now
func PersonalRirekiAll(db *sql.DB, joid int, month int) ([]*Rireki, error) {
	tableName := strconv.Itoa(month) + "_rireki"
	fmt.Print(tableName)

	rows, err := db.Query(fmt.Sprintf("SELECT joid,syubetsu,about,start_time,end_time FROM %s WHERE joid=%d", tableName, joid))

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	var rirekiList []*Rireki

	for rows.Next() {
		var rireki Rireki
		if err := rows.Scan(&rireki.Joid, &rireki.Syubetsu, &rireki.About, &rireki.StartTime, &rireki.EndTime); err != nil {
			fmt.Print(err)
			return nil, err
		}
		rirekiList = append(rirekiList, &rireki)
	}

	if err := rows.Err(); err != nil {
		fmt.Print(err)
		return nil, err
	}

	return rirekiList, nil
}
