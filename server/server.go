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
	router.GET("/users", userctr.All)
	router.POST("/users", userctr.Add)

	router.GET("/rireki/:joid", rirekictr.PersonalAll)
	router.POST("/rireki", rirekictr.RirekiAdd)

	router.Run(":8888")
}
