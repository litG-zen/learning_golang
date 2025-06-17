package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// created a simple GET API using gin
// reference site : https://gin-gonic.com/en/docs/quickstart/

func main() {
	app := gin.Default()

	app.GET("/ping", PingHandler)

	app.GET("/health", func(c *gin.Context) { //anonymous function approach
		c.JSON(http.StatusOK, gin.H{
			"message": "Sab Changa si!",
		})
	})
	app.GET("/async/api", AsyncAPIHandler)

	app.POST("/show_details", func(c *gin.Context) {
		var validator ShowPayloadBody

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

	app.Run(":" + fmt.Sprint(PORT)) // listen and serve on 0.0.0.0:8000 (for windows "localhost:8000")
}
