package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xluohome/phonedata"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/search", startPage)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8022")
}

func startPage(c *gin.Context) {
	phone := c.Query("phone")
	pr, err := phonedata.Find(phone)
	if err != nil {
		panic(err)
	}
	c.JSON(200, pr)
}
