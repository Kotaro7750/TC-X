package model

import (
	"database/sql"
	"fmt"
	"strconv"
)

//SummarizeAll is a function to summarize year_month_rireki and ruturn [][]string. [[joid,salary of each syubetsu,...,sum],..[]]
func SummarizeAll(db *sql.DB, year int, month int) ([][]string, error) {
	//get []{joid,name} and make [index][joid,name]
	userList, err := UserListString(db)
	if err != nil {
		return nil, err
	}
	iJoidMap := make(map[string]int, 0)
	for i := 0; i < len(userList); i++ {
		iJoidMap[userList[i][0]] = i
	}

	//get []{syubetsu,name,salary} and make [index][syubetsu,name,salary]
	syubetsuList, err := SyubetsuListString(db)
	if err != nil {
		return nil, err
	}
	jSyubetsuMap := make(map[string]int, 0)
	for i := 0; i < len(syubetsuList); i++ {
		jSyubetsuMap[syubetsuList[i][0]] = i
	}

	moneySum := make([][]float64, len(userList))
	for i := range moneySum {
		moneySum[i] = make([]float64, len(syubetsuList)+1)
	}

	for i := 0; i < len(userList); i++ {
		for j := 0; j < len(syubetsuList)+1; j++ {
			moneySum[i][j] = 0
		}
	}
	//select joid,syubetsu,startTime,endTime and row => [[joid,syubetsu,time],...[]]
	rirekiList, err := RirekiListString(db, year, month)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	fmt.Print(rirekiList)

	for n := range rirekiList {
		i := iJoidMap[rirekiList[n][0]]
		j := jSyubetsuMap[rirekiList[n][1]]
		time, _ := strconv.Atoi(rirekiList[n][2])
		hourSalary, _ := strconv.Atoi(syubetsuList[j][2])
		moneySum[i][j] += float64(time) * float64(hourSalary) / 60
	}

	for i := 0; i < len(userList); i++ {
		personalSum := 0.0
		for j := 0; j < len(syubetsuList); j++ {
			personalSum += moneySum[i][j]
		}
		moneySum[i][len(syubetsuList)] = personalSum
	}
	fmt.Print(moneySum)

	summarize := make([][]string, len(userList)+1)
	for i := range summarize {
		summarize[i] = make([]string, len(syubetsuList)+3)
	}

	for i := 0; i < len(userList)+1; i++ {
		for j := 0; j < len(syubetsuList)+3; j++ {
			summarize[i][j] = ""
		}
	}
	summarize[0][0] = "Joid"
	summarize[0][1] = "名前"

	for j := 2; j < len(syubetsuList)+2; j++ {
		summarize[0][j] = syubetsuList[j-2][1]
	}
	summarize[0][len(syubetsuList)+2] = "合計"

	for i := 1; i < len(userList)+1; i++ {
		summarize[i][0] = userList[i-1][0]
		summarize[i][1] = userList[i-1][1]
		for j := 0; j < len(syubetsuList)+1; j++ {
			summarize[i][j+2] = strconv.Itoa(int(moneySum[i-1][j])) + "円"
		}
	}
	fmt.Print(summarize)

	/*
		sum [][]string. sum(0,0) is "joid/syubetsu". sum(0,j) is syubetsu_map[j]. sum(i,0) is joid_map[i]. sum(i,j) is salary of joid_map[i] & syubetsu_map[j]
		for line of row{
			joid := line[0]
			syubetsu := line[1]
			startTime := line[2]
			endTime := line[3]

			time := endTime - startTime in minutes format
			salary := time * syubetsu_salary_map[syubetsu]
			sum[joid_map[joid]][syubetsu_map[syubetsu]] += salary
		}
	*/

	return summarize, nil
}
