package model

import (
	"database/sql"
	"errors"
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

	if len(userList) == 0 && len(syubetsuList) == 0 {
		err := errors.New("length of user or syubetsu is zero")
		return nil, err
	}

	moneySum, err := moneySum(db, year, month, userList, iJoidMap, syubetsuList, jSyubetsuMap)
	if err != nil {
		return nil, err
	}

	//sum personal total
	for i := 0; i < len(userList); i++ {
		personalSum := 0.0
		for j := 0; j < len(syubetsuList); j++ {
			personalSum += moneySum[i][j]
		}
		moneySum[i][len(syubetsuList)] = personalSum
	}

	//prepare table
	summarize := make([][]string, len(userList)+1)
	for i := range summarize {
		summarize[i] = make([]string, len(syubetsuList)+3)
	}

	for i := 0; i < len(userList)+1; i++ {
		for j := 0; j < len(syubetsuList)+3; j++ {
			summarize[i][j] = ""
		}
	}

	//set row and column name
	summarize[0][0] = "Joid"
	summarize[0][1] = "名前"

	for j := 2; j < len(syubetsuList)+2; j++ {
		summarize[0][j] = syubetsuList[j-2][1]
	}
	summarize[0][len(syubetsuList)+2] = "合計"

	//put summarize from moneySum
	for i := 1; i < len(userList)+1; i++ {
		summarize[i][0] = userList[i-1][0]
		summarize[i][1] = userList[i-1][1]
		for j := 0; j < len(syubetsuList)+1; j++ {
			summarize[i][j+2] = strconv.Itoa(int(moneySum[i-1][j])) + "円"
		}
	}
	fmt.Print(summarize)
	return summarize, nil
}

func moneySum(db *sql.DB, year int, month int, userList [][2]string, iJoidMap map[string]int, syubetsuList [][3]string, jSyubetsuMap map[string]int) ([][]float64, error) {
	//prepare summarization of money
	moneySum := make([][]float64, len(userList))
	for i := range moneySum {
		moneySum[i] = make([]float64, len(syubetsuList)+1)
	}

	for i := 0; i < len(userList); i++ {
		for j := 0; j < len(syubetsuList)+1; j++ {
			moneySum[i][j] = 0
		}
	}

	//get all rireki of year,month
	rirekiList, err := RirekiListString(db, year, month)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	//summarize rireki
	for n := range rirekiList {
		i := iJoidMap[rirekiList[n][0]]
		j := jSyubetsuMap[rirekiList[n][1]]
		time, _ := strconv.Atoi(rirekiList[n][2])
		hourSalary, _ := strconv.Atoi(syubetsuList[j][2])
		moneySum[i][j] += float64(time) * float64(hourSalary) / 60
	}

	//sum personal total
	for i := 0; i < len(userList); i++ {
		personalSum := 0.0
		for j := 0; j < len(syubetsuList); j++ {
			personalSum += moneySum[i][j]
		}
		moneySum[i][len(syubetsuList)] = personalSum
	}

	return moneySum, nil

}
