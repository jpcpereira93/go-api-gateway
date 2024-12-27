package controllers

import (
	"github.com/gin-gonic/gin"

	"jpcpereira93/go-api-gateway/middleware"
)

func init() {
	RegisterController(&TestController{})
}

// TestController handles `test` routes.
type TestController struct{}

// RegisterRoutes registers the routes for TestController.
func (c *TestController) RegisterRoutes(router *gin.Engine) {
	router.GET("/test", middleware.Authenticate(), func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Test"})
	})
}