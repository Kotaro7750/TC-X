package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"server/apierror"
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
		error := apierror.Error(1, "PersonalAll", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		error := apierror.Error(1, "PersonalAll", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	rirekiList, err := model.PersonalRirekiAll(r.DB, joid, month)

	if err != nil {
		error := apierror.Error(11, "PersonalAll", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	if len(rirekiList) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"result":make([]int, 0),
			"error":nil,
		})
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
	if err != nil {
		error := apierror.Error(1, "RirekiAdd", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	var rireki model.Rireki

	if err := c.ShouldBindJSON(&rireki); err != nil {
		error := apierror.Error(1, "RirekiAdd", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	if month != int(rireki.StartTime.Month()) {
		error := apierror.Error(1, "RirekiAdd", "URLParam :month not equal with JSONParam")
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	contradictedList, err := model.IsContradicted(r.DB, rireki, month)

	if err != nil {
		error := apierror.Error(11, "RirekiAdd", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	if len(contradictedList) != 0 {
		error := apierror.Error(3, "RirekiAdd", "request is conflicted with database")
		c.JSON(http.StatusBadRequest, gin.H{
			"result": contradictedList,
			"error":  error,
		})
		return
	}

	addedRireki, err := model.RirekiAdd(r.DB, rireki, month)

	if err != nil {
		error := apierror.Error(11, "RirekiAdd", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": addedRireki,
		"error":  nil,
	})
	return
}

//RirekiDelete is a function to delete rireki with joid and id
func (r *RirekiCtr) RirekiDelete(c *gin.Context) {
	month, err := strconv.Atoi(c.Param("month"))

	if err != nil {
		error := apierror.Error(1, "RirekiDelete", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	joid, err := strconv.Atoi(c.Param("joid"))

	if err != nil {
		error := apierror.Error(1, "RirekiDelete", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		error := apierror.Error(1, "RirekiDelete", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	err = model.RirekiDelete(r.DB, month, joid, id)

	if err != nil {
		error := apierror.Error(11, "RirekiDelete", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":nil,
		"error": nil,
	})
	return
}

//RirekiUpdate is a function to update rireki date
func (r *RirekiCtr) RirekiUpdate(c *gin.Context) {
	month, err := strconv.Atoi(c.Param("month"))

	if err != nil {
		error := apierror.Error(1, "RirekiUpdate", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	joid, err := strconv.Atoi(c.Param("joid"))

	if err != nil {
		error := apierror.Error(1, "RirekiUpdate", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		error := apierror.Error(1, "RirekiUpdate", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	var rireki model.Rireki

	if err := c.ShouldBindJSON(&rireki); err != nil {
		error := apierror.Error(1, "RirekiUpdate", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	contradictedList, err := model.IsContradicted(r.DB, rireki, month)

	if err != nil {
		error := apierror.Error(11, "RirekiUpdate", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	if len(contradictedList) != 0 {
		if len(contradictedList) >= 2 || contradictedList[0].ID != id {
			error := apierror.Error(3, "RirekiUpdate", "request is conflicted with database")
			c.JSON(http.StatusBadRequest, gin.H{
				"result": contradictedList,
				"error":  error,
			})
			return
		}
	}

	err = model.RirekiUpdate(r.DB, month, joid, id, rireki)

	if err != nil {
		error := apierror.Error(11, "RirekiUpdate", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"error":  error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":nil,
		"error": nil,
	})
	return
}
