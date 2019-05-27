package main

import (
	"database/sql"
	"net/http"
	"server/apiresponse"
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
	syubetsuctr := &controller.SyubetsuCtr{DB: db}
	rirekictr := &controller.RirekiCtr{DB: db}

	router.Use(cors.New(config))

	//routing start

	//Authentication
	router.POST("/auth", userctr.TestAuth)

	//User
	router.GET("/user/:joid", userctr.AuthUser)
	router.POST("/user", userctr.UserAdd)

	//Syubetsu
	router.GET("/syubetsu", syubetsuctr.SyubetsuAll)
	router.POST("/syubetsu", syubetsuctr.SyubetsuAdd)
	router.DELETE("/syubetsu/:syubetsu", syubetsuctr.SyubetsuDelete)
	router.PATCH("/syubetsu/:syubetsu", syubetsuctr.SyubetsuUpdate)

	//Rireki
	router.GET("/rireki/:month/:joid", rirekictr.PersonalAll)
	router.POST("/rireki/:month/:joid", rirekictr.RirekiAdd)
	router.DELETE("/rireki/:month/:joid/:id", rirekictr.RirekiDelete)
	router.PATCH("/rireki/:month/:joid/:id", rirekictr.RirekiUpdate)

	//404
	router.NoRoute(func(c *gin.Context) {
		apiresponse.APIResponse(c, http.StatusNotFound, nil, 1, "main", "No route fo this request")
		return
	})

	router.Run(":8888")
}
