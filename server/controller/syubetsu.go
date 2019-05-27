package controller

import (
	"database/sql"
	"net/http"
	"server/apiresponse"
	"server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

//SyubetsuCtr is a controller for  request to syubetsu
type SyubetsuCtr struct {
	DB *sql.DB
}

//SyubetsuAll is a function to list up all syubetsu
func (s *SyubetsuCtr) SyubetsuAll(c *gin.Context) {

	syubetsuList, err := model.SyubetsuAll(s.DB)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "SyubetsuAll", err.Error())
		return
	}

	if len(syubetsuList) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": make([]int, 0),
			"error":  nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": syubetsuList,
		"error":  nil,
	})
	return
}

//SyubetsuAdd is a function to add syubetsu
func (s *SyubetsuCtr) SyubetsuAdd(c *gin.Context) {
	var syubetsu model.Syubetsu

	if err := c.ShouldBindJSON(&syubetsu); err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "SyubetsuAdd", err.Error())
		return
	}

	addedSyubetsu, err := model.SyubetsuAdd(s.DB, syubetsu)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "SyubetsuAdd", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": addedSyubetsu,
		"error":  nil,
	})
	return
}

//SyubetsuDelete is a function to delete syubetsu
func (s *SyubetsuCtr) SyubetsuDelete(c *gin.Context) {
	syubetsu, err := strconv.Atoi(c.Param("syubetsu"))

	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "SyubetsuDelete", err.Error())
		return
	}

	err = model.SyubetsuDelete(s.DB, syubetsu)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "SyubetsuDelete", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": nil,
		"error":  nil,
	})
	return
}

//SyubetsuUpdate is a function to update syubetsu date
func (s *SyubetsuCtr) SyubetsuUpdate(c *gin.Context) {
	requestSyubetsu, err := strconv.Atoi(c.Param("syubetsu"))

	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "SyubetsuUpdate", err.Error())
		return
	}

	var syubetsu model.Syubetsu

	if err := c.ShouldBindJSON(&syubetsu); err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "SyubetsuUpdate", err.Error())
		return
	}

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "SyubetsuUpdate", err.Error())
		return
	}

	if requestSyubetsu != syubetsu.Syubetsu {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "SyubetsuUpdate", "request Syubetsu and json syubetsu is contradicted")
		return
	}

	err = model.SyubetsuUpdate(s.DB, syubetsu)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "SyubetsuUpdate", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": nil,
		"error":  nil,
	})
	return
}
