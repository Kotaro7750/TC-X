package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	//"time"

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
	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	rirekiList, err := model.PersonalRirekiAll(r.DB, joid, month)

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

//RirekiAdd is a function to add rireki of month
func (r *RirekiCtr) RirekiAdd(c *gin.Context) {
	month, err := strconv.Atoi(c.Param("month"))
	if  err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	var rireki model.Rireki

	if err := c.ShouldBindJSON(&rireki); err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if month != int(rireki.StartTime.Month()) {
		resp := errors.New("month")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	contradictedList,err := model.IsContradicted(r.DB,rireki,month);

	if  err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(contradictedList) != 0 {
		c.JSON(http.StatusBadRequest, contradictedList)
		return
	}

	addedRireki, err := model.RirekiAdd(r.DB, rireki, month)

	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": addedRireki,
		"error":  nil,
	})
	return
}

//RirekiDelete is a function to delete rireki with joid and id
func (r *RirekiCtr) RirekiDelete(c *gin.Context)  {
	month, err := strconv.Atoi(c.Param("month"))

	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	joid, err := strconv.Atoi(c.Param("joid"))

	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	err = model.RirekiDelete(r.DB,month,joid,id)

	if err != nil {
		resp := err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"error":resp,})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  nil,
	})
	return
}

//RirekiUpdate is a function to update rireki date
func (r *RirekiCtr) RirekiUpdate(c *gin.Context)  {
	month, err := strconv.Atoi(c.Param("month"))

	if err != nil {
		fmt.Print(err)
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	joid, err := strconv.Atoi(c.Param("joid"))

	if err != nil {
		fmt.Print(err)
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Print(err)
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	var rireki model.Rireki

	if err := c.ShouldBindJSON(&rireki); err != nil {
		fmt.Print(err)
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	contradictedList,err := model.IsContradicted(r.DB,rireki,month);

	if err !=nil {
		fmt.Print(err)
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(contradictedList) != 0  {
		if len(contradictedList) >= 2 || contradictedList[0].ID != id {
			c.JSON(http.StatusBadRequest, contradictedList)
			return
		}
	}

	if  err != nil {
		fmt.Print(err)
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}


	err = model.RirekiUpdate(r.DB,month,joid,id,rireki)

	if err != nil {
		fmt.Print(err)
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  nil,
	})
	return
}