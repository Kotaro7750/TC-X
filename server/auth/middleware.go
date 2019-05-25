package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"server/apiresponse"
	"strings"

	"github.com/gin-gonic/gin"
)

//CertCtr is a controller for certification
type CertCtr struct {
	DB *sql.DB
}

//AuthWithToken is a function to authorize user with token and joid
func (ca *CertCtr) AuthWithToken(c *gin.Context, joid int) bool {
	var authHeader = c.Request.Header["Authorization"]

	if authHeader == nil {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 2, "AuthWithToken", "Authorization header does not exist")
		return false
	}
	protocol := strings.Split(authHeader[0], " ")[0]
	token := strings.Split(authHeader[0], " ")[1]

	if protocol != "Bearer" || token == "" {
		apiresponse.APIResponse(c, http.StatusBadRequest, nil, 2, "AuthWithToken", "Authorization header is not bearer")
		return false
	}
	//select record with joid and token
	rows, err := ca.DB.Query(fmt.Sprintf("SELECT joid FROM `users` WHERE joid = '%d' AND token = '%s'", joid, token))
	if err != nil {
		fmt.Print(err.Error())
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "AuthWithToken", err.Error())
		return false
	}
	defer rows.Close()
	hit := 0
	for rows.Next() {
		hit = hit + 1
	}

	if err := rows.Err(); err != nil {
		apiresponse.APIResponse(c, http.StatusInternalServerError, nil, 11, "AuthWithToken", err.Error())
		return false
	}

	if hit != 1 {
		apiresponse.APIResponse(c, http.StatusForbidden, nil, 5, "RirekiAdd", "UnAuthorized")
		return false
	}

	return true
}
