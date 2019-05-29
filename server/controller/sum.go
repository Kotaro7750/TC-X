package controller

import (
	"bytes"
	"encoding/csv"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
joid,salary of each syubetsu,...,sum
*/

func testOfCsv() *bytes.Buffer {
	var sample = [][]string{
		{"sample_1", "sample_2", "sample_3"}, {"sample_1", "sample_2", "sample_3"},
	}

	b := &bytes.Buffer{}
	writer := csv.NewWriter(b) // utf8
	for i := 0; i < 2; i++ {
		writer.Write(sample[i])
		writer.Flush()
	}

	return b
}

//test of returning csv
func TestRetFile(c *gin.Context) {
	b := testOfCsv()
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=/summarize/test.csv")
	c.Data(http.StatusOK, "text/csv", b.Bytes())
}
