package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"jpcpereira93/go-api-gateway/config"
	"jpcpereira93/go-api-gateway/middleware"
)

var dogsApiBaseUrl string

func init() {
	// Load configuration
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatal().AnErr("Error loading configuration: %v", err)
	}

	dogsApiBaseUrl = cfg.DogsApi.BaseUrl
	
	if len(dogsApiBaseUrl) == 0 {
		log.Fatal().Msg("Missing Dogs API Proxy Url configuration.")
	}

	RegisterController(&DogsController{})
}

// DogsController handles `dog` routes.
type DogsController struct{}

// RegisterRoutes registers the routes for TestController.
func (c *DogsController) RegisterRoutes(router *gin.Engine) {
	router.Any(
		"/dogs/*any", 
		middleware.Authenticate(), 
		middleware.ProxyWeb(dogsApiBaseUrl), 
	)
}