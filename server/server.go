package main

import (
	"database/sql"
	"server/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	//CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowHeaders = []string{"*"}

	//DB
	db, err := sql.Open("mysql", "root:root@([db]:3306)/TCX?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	userctr := &controller.UserCtr{DB: db}
	rirekictr := &controller.RirekiCtr{DB: db}

	router.Use(cors.New(config))

	//routing start

	router.GET("/user/:joid", userctr.AuthUser)
	router.POST("/user", userctr.UserAdd)

	router.POST("/auth", userctr.TestAuth)

	router.GET("/rireki/:month/:joid", rirekictr.PersonalAll)
	router.POST("/rireki/:month/:joid", rirekictr.RirekiAdd)
	router.DELETE("/rireki/:month/:joid/:id", rirekictr.RirekiDelete)
	router.PATCH("/rireki/:month/:joid/:id", rirekictr.RirekiUpdate)

	router.Run(":8888")
}
