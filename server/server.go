package main

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowHeaders = []string{"Origin"}

	router.Use(cors.New(config))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "bang",
		})
	})
	router.Run(":8888")
}
