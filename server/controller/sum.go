package controller

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"server/apiresponse"
	"server/model"
	"strconv"

	"github.com/signintech/gopdf"

	"github.com/gin-gonic/gin"
)

/*
joid,salary of each syubetsu,...,sum
*/

//SumCtr is a controller for request to summarize
type SumCtr struct {
	DB *sql.DB
}

//SummarizeAll is a function to summarize rireki of month
func (s *SumCtr) SummarizeAll(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "SummarizeAll", err.Error())
		return
	}

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "SummarizeAll", err.Error())
		return
	}

	filename := strconv.Itoa(year) + "_" + strconv.Itoa(month) + "_sum.csv"

	buf := &bytes.Buffer{}
	writer := csv.NewWriter(buf)

	//get [][]string from model
	summarization, err := model.SummarizeAll(s.DB, year, month)
	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "SummarizeAll", err.Error())
		return
	}

	for i := range summarization {
		if err := writer.Write(summarization[i]); err != nil {
			apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 12, "SummarizeAll", err.Error())
			return
		}
	}
	writer.Flush()

	if err := writer.Error(); err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 12, "SummarizeAll", err.Error())
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Data(http.StatusOK, "text/csv", buf.Bytes())
}

//SummarizeAllWithPdf is a function to summarize rireki of month and return with pdf
func (s *SumCtr) SummarizeAllWithPdf(c *gin.Context) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("aozora", "/app/AozoraMinchoRegular.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("aozora", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "您好")

	filename := "hoge.pdf"
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Data(http.StatusOK, "application/pdf", pdf.GetBytesPdf())
}
