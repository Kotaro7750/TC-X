package apiresponse

import (
	"server/apierror"

	"github.com/gin-gonic/gin"
)

//Result is a interface for ignore type assertion
type Result interface{}

//APIResponse is a function to response
func APIResponse(c *gin.Context, status int, result Result, errorCode int, errorResource string, errMessage string) {
	error := apierror.Error(errorCode, errorResource, errMessage)
	c.JSON(status, gin.H{
		"result": result,
		"error":  error,
	})
}
