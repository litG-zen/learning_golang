package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// created a simple GET API using gin
// reference site : https://gin-gonic.com/en/docs/quickstart/

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"owner":   "Lit",
	})
}

func AsyncAPIHandler(c *gin.Context) {
	fmt.Println("\n\n Async API hander called")

	go func() {
		// I have added this goroutine to mimick a blocking function like email-sending, heavy calculation
		// which is usually triggered in an async way, this is that attempt.
		time.Sleep(10000 * time.Millisecond)
		print("\n\n\n Sleep ended \n\n\n")
	}()

	c.JSON(http.StatusOK, gin.H{
		"message": "this is a successful response of AsyncAPI call",
	})
}
