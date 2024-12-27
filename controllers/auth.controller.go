package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"jpcpereira93/go-api-gateway/config"
	"jpcpereira93/go-api-gateway/middleware"
	"jpcpereira93/go-api-gateway/services"
)

var authCookieName string

func init() {
	// Load configuration
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatal().AnErr("Error loading configuration: %v", err)
	}

	authCookieName = cfg.Auth.CookieName

	RegisterController(&AuthController{})
}

// AuthController handles authentication related routes.
type AuthController struct{}

// Login credentials struct.
type LoginCredentials struct {
	Email     string    `form:"email"`
	Password  string    `form:"password"`
}

// RegisterRoutes registers the routes for AuthController.
func (c *AuthController) RegisterRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth") // Parent route for nested routes
	
	authGroup.POST("/login", c.Login)
	authGroup.POST("/logout", middleware.Authenticate(), c.Login)
}

// Handles the /auth/login route.
func (a *AuthController) Login(ctx *gin.Context) {
	var loginCredentials LoginCredentials

   if err := ctx.BindJSON(&loginCredentials); err != nil {
	log.Error().Err(err)
    ctx.Error(err)
   }

   // TODO: External Authentication validation using given credentials
   
 
   data, err := json.Marshal(loginCredentials)

   if err != nil {
	log.Error().Err(err)
	ctx.Error(err)
   }

   cookieValue := services.Cipher(data)

   ctx.SetCookie(authCookieName, cookieValue, 60, "/", "", false, true)
}

// Handles the /auth/logout route.
func (a *AuthController) Logout(ctx *gin.Context) {
	// TODO: External logout request

   // Delete the cookie.
   ctx.SetCookie(authCookieName, "", -1, "/", "", false, true)
}
