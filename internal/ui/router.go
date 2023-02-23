package ui

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ServeRouter(fizzBuzzHandler FizzBuzzHandler) error {
	router := gin.Default()

	router.POST("/fizzbuzz", func(ctx *gin.Context) {
		fizzBuzzHandler.handleFizzBuzz(ctx)
	})

	router.GET("/stats", func(ctx *gin.Context) {
		fizzBuzzHandler.handleTopRequests(ctx)
	})

	if err := router.Run(":8088"); err != nil {
		return fmt.Errorf("an error happened with gin at runtime: %w", err)
	}

	return nil
}
