package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"server/apiresponse"
	"server/auth"
	"server/model"

	"github.com/gin-gonic/gin"
)

//RirekiCtr is a contoroller for request to rirekis
type RirekiCtr struct {
	DB *sql.DB
}

//PersonalAll is a function to list up all indivisual rireki of month
func (r *RirekiCtr) PersonalAll(c *gin.Context) {

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "PersonalAll", err.Error())
		return
	}

	joid, err := strconv.Atoi(c.Param("joid"))
	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "PersonalAll", err.Error())
		return
	}

	ca := &auth.CertCtr{DB: r.DB}
	if authOk := ca.AuthWithToken(c, joid); !authOk {
		apiresponse.APIResponse(c, http.StatusUnauthorized, nil, 5, "PersonalAll", "Authentication Failed")
		return
	}

	rirekiList, err := model.PersonalRirekiAll(r.DB, joid, month)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "PersonalAll", err.Error())
		return
	}

	if len(rirekiList) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": make([]int, 0),
			"error":  nil,
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
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "RirekiAdd", err.Error())
		return
	}

	joid, err := strconv.Atoi(c.Param("joid"))
	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "RirekiAdd", err.Error())
		return
	}

	ca := &auth.CertCtr{DB: r.DB}
	if authOk := ca.AuthWithToken(c, joid); !authOk {
		apiresponse.APIResponse(c, http.StatusUnauthorized, nil, 5, "RirekiAdd", "Authentication Failed")
		return
	}

	var rireki model.Rireki

	if err := c.ShouldBindJSON(&rireki); err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "RirekiAdd", err.Error())
		return
	}

	if month != int(rireki.StartTime.Month()) {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "RirekiAdd", "URLParam :month not equal with JSONParam")
		return
	}

	if joid != rireki.Joid {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "RirekiAdd", "URLParam :joid not equal with JSONParam")
		return
	}

	contradictedList, err := model.IsContradicted(r.DB, rireki, month)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "RirekiAdd", err.Error())
		return
	}

	if len(contradictedList) != 0 {
		apiresponse.APIResponse(c, http.StatusBadRequest, contradictedList, 3, "RirekiAdd", "request is conflicted with database")
		return
	}

	addedRireki, err := model.RirekiAdd(r.DB, rireki, month)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "RirekiAdd", err.Error())
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
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "RirekiDelete", err.Error())
		return
	}

	joid, err := strconv.Atoi(c.Param("joid"))

	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "RirekiDelete", err.Error())
		return
	}

	ca := &auth.CertCtr{DB: r.DB}
	if authOk := ca.AuthWithToken(c, joid); !authOk {
		apiresponse.APIResponse(c, http.StatusUnauthorized, nil, 5, "RirekiDelete", "Authentication Failed")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "RirekiDelete", err.Error())
		return
	}

	err = model.RirekiDelete(r.DB, month, joid, id)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "RirekiDelete", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": nil,
		"error":  nil,
	})
	return
}

//RirekiUpdate is a function to update rireki date
func (r *RirekiCtr) RirekiUpdate(c *gin.Context) {
	month, err := strconv.Atoi(c.Param("month"))

	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "RirekiUpdate", err.Error())
		return
	}

	joid, err := strconv.Atoi(c.Param("joid"))

	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "RirekiUpdate", err.Error())
		return
	}
	ca := &auth.CertCtr{DB: r.DB}
	if authOk := ca.AuthWithToken(c, joid); !authOk {
		apiresponse.APIResponse(c, http.StatusUnauthorized, nil, 5, "RirekiUpdate", "Authentication Failed")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "RirekiUpdate", err.Error())
		return
	}

	var rireki model.Rireki

	if err := c.ShouldBindJSON(&rireki); err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "RirekiUpdate", err.Error())
		return
	}

	contradictedList, err := model.IsContradicted(r.DB, rireki, month)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "RirekiUpdate", err.Error())
		return
	}

	if len(contradictedList) != 0 {
		if len(contradictedList) >= 2 || contradictedList[0].ID != id {
			apiresponse.APIResponse(c, http.StatusBadRequest, nil, 3, "RirekiUpdate", "request is conflicted with database")
			return
		}
	}

	err = model.RirekiUpdate(r.DB, month, joid, id, rireki)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "RirekiUpdate", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": nil,
		"error":  nil,
	})
	return
}
