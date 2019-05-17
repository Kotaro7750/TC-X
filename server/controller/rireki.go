package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

//RirekiCtr is a contoroller for request to rirekis
type RirekiCtr struct {
	DB *sql.DB
}

//RirekiAll is a function to list up all indivisual rireki of month
func (r *RirekiCtr) RirekiAll(c *gin.Context) {
}
