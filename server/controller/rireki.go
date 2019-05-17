package controller

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"server/model"

	"github.com/gin-gonic/gin"
)

//RirekiCtr is a contoroller for request to rirekis
type RirekiCtr struct {
	DB *sql.DB
}

//PersonalAll is a function to list up all indivisual rireki of month
func (r *RirekiCtr) PersonalAll(c *gin.Context) {
	joid, err := strconv.Atoi(c.Param("joid"))
	rirekiList, err := model.PersonalRirekiAll(r.DB, joid, 5)

	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(rirekiList) == 0 {
		c.JSON(http.StatusOK, make([]int, 0))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": rirekiList,
		"error":  nil,
	})
	return
}
