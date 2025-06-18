package main

import (
	"fmt"
	"net/http"
	"time"
	"web_test/auth"
	"web_test/logs"

	"github.com/gin-gonic/gin"
)

// created a simple GET API using gin
// reference site : https://gin-gonic.com/en/docs/quickstart/

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"owner":   "Lit",
	})

	log_string := fmt.Sprintf("%v %v %v", time.Now(), c.FullPath(), http.StatusOK)
	logs.Logger(log_string, false)
}

func AsyncAPIHandler(c *gin.Context) {
	fmt.Println("\n\n Async API hander called")

	token := c.GetHeader("API-KEY")

	if token != auth.API_KEY {
		c.JSON(http.StatusUnauthorized, gin.H{"message": INVALID_API_MSG})
		log_string := fmt.Sprintf("%v %v %v %v", time.Now(), c.FullPath(), http.StatusUnauthorized, INVALID_API_MSG)
		logs.Logger(log_string, true)
		return
	}

	go func() {
		// I have added this goroutine to mimick a blocking function like email-sending, heavy calculation
		// which is usually triggered in an async way, this is that attempt.
		time.Sleep(10000 * time.Millisecond)
		print("\n\n\n Sleep ended \n\n\n")
	}()

	c.JSON(http.StatusOK, gin.H{
		"message": "this is a successful response of AsyncAPI call",
	})

	log_string := fmt.Sprintf("%v %v %v %v", time.Now(), c.FullPath(), http.StatusUnauthorized, "SUCCESS")
	logs.Logger(log_string, false)
}
