package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// created a simple GET API using gin
// reference site : https://gin-gonic.com/en/docs/quickstart/

func ping_handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"owner":   "Lit",
	})
}

func async_api_handler(c *gin.Context) {
	fmt.Println("\n\n Async API hander called")

	go func() {
		time.Sleep(10000 * time.Millisecond)
		print("\n\n\n Sleep ended \n\n\n")
	}()

	c.JSON(http.StatusOK, gin.H{
		"message": "this is a successful response of AsyncAPI call",
	})
}

type PostBodyExample struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {
	app := gin.Default()

	app.GET("/ping", ping_handler)

	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Sab Changa si!",
		})
	})
	app.GET("/async/api", async_api_handler)

	app.POST("/show_details", func(c *gin.Context) {
		var validator PostBodyExample

		// Bind Json to struct
		if err := c.ShouldBindBodyWithJSON(&validator); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "data received",
				"data":    validator,
			})
	})
	app.Run(":8000") // listen and serve on 0.0.0.0:8000 (for windows "localhost:8000")
}
