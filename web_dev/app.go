package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"web_test/logs"

	"github.com/gin-gonic/gin"
)

// created a simple GET API using gin
// reference site : https://gin-gonic.com/en/docs/quickstart/

func main() {
	// Setting DEBUG mode ON
	gin.SetMode(gin.DebugMode)

	app := gin.Default()

	// Defining folder to load templates from
	app.LoadHTMLGlob("templates/*")

	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "rendering HTML",
		})

		log_string := fmt.Sprintf("%v %v %v", time.Now(), c.FullPath(), http.StatusOK)
		logs.Logger(log_string, false)
	})

	app.GET("/ping", PingHandler)

	app.GET("/health", func(c *gin.Context) { //anonymous function approach

		greetints := []string{
			"Waah kya scene hai",
			"Sab Changa si!",
			"Bindassss",
			"Bhai lagi padi hai!",
			"Rest kar lijie bhrata",
		}

		c.JSON(http.StatusOK, gin.H{
			"message": greetints[rand.Intn(len(greetints))],
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
