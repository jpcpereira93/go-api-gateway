package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"jpcpereira93/go-api-gateway/config"
	"jpcpereira93/go-api-gateway/controllers"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatal().AnErr("Error loading configuration: %v", err)
	}

	// Set the logging level.
	zerolog.SetGlobalLevel(cfg.Logging.Level)

	// Set the logger output.
	if cfg.Logging.Output == "console" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Initialize Gin router
	router := gin.Default()

	// Load and register all controllers
	controllers.LoadControllers(router)

	// Get the configured server port
	port := strconv.Itoa(cfg.Server.Port)
	
	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Fatal().AnErr("Failed to start the server: ", err)
	}
}