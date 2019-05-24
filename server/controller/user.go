package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"server/apiresponse"
	"server/model"
)

//UserCtr is a contoroller for request to users
type UserCtr struct {
	DB *sql.DB
}

//UserAll is a function to list up all users
func (u *UserCtr) UserAll(c *gin.Context) {
	users, err := model.UserAll(u.DB)

	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusInternalServerError, make([]*model.User, 0))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": users,
		"error":  nil,
	})
}

//UserAdd is a function to add user
func (u *UserCtr) UserAdd(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "UserAdd", err.Error())
		return
	}

	addable, err := model.IsAddable(u.DB, user.Joid)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "UserAdd", err.Error())
		return
	}

	if addable == false {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 4, "UserAdd", "that joid is already exist")
		return
	}

	added, err := model.UserAdd(u.DB,user)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "UserAdd", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": added,
		"error": nil,
	})
	return
}

//AuthUser is a function to authenticate user
func (u *UserCtr) AuthUser(c *gin.Context)  {
	var hashedPass = c.Request.Header["Authorization"][0]
	fmt.Print(hashedPass)
	c.JSON(http.StatusOK,gin.H{
		"header": c.Request.Header["Authorization"],
	})
	return
}
