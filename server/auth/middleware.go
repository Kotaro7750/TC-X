package auth

import (
	"fmt"
	"net/http"
	"server/apiresponse"

	"github.com/gin-gonic/gin"
)

func aba() {
	fmt.Print(":sfdsfsf")
}

//AuthWithToken is a function to authorize user with token and joid
func AuthWithToken(c *gin.Context, joid int, token string) bool {
	//select record with joid and token
	//if exist -> return true
	//if not -> apiresponse with error and return false
	var authNg bool
	if authNg {
		apiresponse.APIResponse(c, http.StatusForbidden, nil, 1, "RirekiAdd", "UnAuthorized")
		return false
	}

	return true
}
