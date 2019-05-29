package model

import "database/sql"

//SummarizeAll is a function to summarize year_month_rireki and ruturn [][]string. [[joid,salary of each syubetsu,...,sum],..[]]
func SummarizeAll(db *sql.DB, year int, month int) {
	//get []{joid,name}
	//get []{syubetsu,name,salary}
	//select joid,syubetsu,

}
