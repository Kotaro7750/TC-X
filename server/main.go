package main

import (
	"database/sql"
	"net/http"
	"os"
	"server/apiresponse"
	"server/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//clientURL := os.Getenv("CLIENT_URL")
	clientURL := os.Getenv("CLIENT_URL")
	listenPort := os.Getenv("LISTEN_PORT")
	dbURL := os.Getenv("DB_URL")
	router := gin.Default()

	//CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{clientURL}
	config.AllowHeaders = []string{"*"}

	//DB
	//db, err := sql.Open("mysql", "root:root@([db]:3306)/TCX?parseTime=true")
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	userctr := &controller.UserCtr{DB: db}
	syubetsuctr := &controller.SyubetsuCtr{DB: db}
	rirekictr := &controller.RirekiCtr{DB: db}
	notectr := &controller.NoteCtr{DB: db}
	sumctr := &controller.SumCtr{DB: db}

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
	router.GET("/rireki/:year/:month/:joid", rirekictr.PersonalAll)
	router.POST("/rireki/:year/:month/:joid", rirekictr.RirekiAdd)
	router.DELETE("/rireki/:year/:month/:joid/:id", rirekictr.RirekiDelete)
	router.PATCH("/rireki/:year/:month/:joid/:id", rirekictr.RirekiUpdate)

	//Note
	router.GET("/note", notectr.NoteAll)
	router.POST("/note/:year/:month", notectr.NoteAdd)
	router.DELETE("/note/:year/:month", notectr.NoteDelete)

	router.GET("/sum/:year/:month", sumctr.SummarizeAll)

	//404
	router.NoRoute(func(c *gin.Context) {
		apiresponse.APIResponse(c, http.StatusNotFound, nil, 1, "main", "No route fo this request")
		return
	})

	router.Run(":" + listenPort)
}
