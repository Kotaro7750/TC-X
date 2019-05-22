package controller

import(
	"database/sql"
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"

	"server/model"
)

//UserCtr is a contoroller for request to users
type UserCtr struct {
	DB *sql.DB
}

//All is a function to list up all users
func (u * UserCtr) All (c *gin.Context)  {
	users,err := model.UserAll(u.DB)
	
	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError,resp)
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusInternalServerError,make([]*model.User, 0))
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"result": users,
		"error":nil,
	})
}

//Add is a function to add user
func (u * UserCtr) Add (c *gin.Context) {
	var user model.User 

	err := c.BindJSON(&user)
	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusBadRequest,resp)
		return
	}
	inserted,err := user.Insert(u.DB)

	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError,resp)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"result":inserted,
		"error":nil,
	})
	return
}