package controllers

import "github.com/gin-gonic/gin"

// Controller defines the interface for all controllers.
type Controller interface {
	RegisterRoutes(router *gin.Engine)
}
