package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"server/apiresponse"
	"server/model"

	"github.com/gin-gonic/gin"
)

//NoteCtr is a controller for  request to note
type NoteCtr struct {
	DB *sql.DB
}

//NoteAll is a function to list up all note
func (n *NoteCtr) NoteAll(c *gin.Context) {

	noteList, err := model.NoteAll(n.DB)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "NoteAll", err.Error())
		return
	}

	if len(noteList) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": make([]int, 0),
			"error":  nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": noteList,
		"error":  nil,
	})
	return
}

//NoteAdd is a function to add note
func (n *NoteCtr) NoteAdd(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))

	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "NoteAdd", err.Error())
		return
	}

	month, err := strconv.Atoi(c.Param("month"))

	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "NoteAdd", err.Error())
		return
	}

	addedNote, err := model.NoteAdd(n.DB, year, month)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "NoteAdd", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": addedNote,
		"error":  nil,
	})
	return
}

//NoteDelete is a function to delete note
func (n *NoteCtr) NoteDelete(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))

	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "NoteAdd", err.Error())
		return
	}

	month, err := strconv.Atoi(c.Param("month"))

	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "NoteAdd", err.Error())
		return
	}

	err = model.NoteDelete(n.DB, year, month)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "NoteDelete", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": nil,
		"error":  nil,
	})
	return
}
