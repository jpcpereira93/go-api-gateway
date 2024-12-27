package controllers

import "github.com/gin-gonic/gin"

// Registry holds all registered controllers.
var Registry []Controller

// RegisterController registers a controller in the global registry.
func RegisterController(controller Controller) {
	Registry = append(Registry, controller)
}

// LoadControllers registers all controllers in the registry with the Gin router.
func LoadControllers(router *gin.Engine) {
	for _, controller := range Registry {
		controller.RegisterRoutes(router)
	}
}
