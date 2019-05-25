package controller

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"server/apiresponse"
	"server/auth"
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

	added, err := model.UserAdd(u.DB, user)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "UserAdd", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": added,
		"error":  nil,
	})
	return
}

//AuthUser is a function to authenticate user
func (u *UserCtr) AuthUser(c *gin.Context) {
	joid, err := strconv.Atoi(c.Param("joid"))
	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 1, "AuthUser", err.Error())
		return
	}

	var authHeader = c.Request.Header["Authorization"]

	if authHeader == nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 2, "AuthUser", "Authorization header does not exist")
		return
	}
	protocol := strings.Split(authHeader[0], " ")[0]
	hashedPass := strings.Split(authHeader[0], " ")[1]

	if protocol != "Bearer" || hashedPass == "" {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 2, "AuthUser", "Authorization header is not bearer")
		return
	}

	//TODO: hashedPass validation

	hitUser, err := model.AuthUser(u.DB, joid, hashedPass)

	if err != nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 11, "AuthUser", err.Error())
		return
	}

	if hitUser == nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 5, "AuthUser", "sign in failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": hitUser,
		"error":  nil,
	})
	return
}

func (u *UserCtr) TestAuth(c *gin.Context) {
	ca := &auth.CertCtr{DB: u.DB}
	authOk := ca.AuthWithToken(c, 63)

	c.JSON(http.StatusOK, gin.H{
		"result": authOk,
		"error":  nil,
	})
	return

}
