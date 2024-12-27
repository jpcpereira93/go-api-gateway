package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"jpcpereira93/go-api-gateway/services"
)

// Authenticate middleware to validate user authentication (replace with your own logic)
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debug().Msg("Authenticating request...")
		
		cookie, err := c.Request.Cookie("api_gw_session")
		
		if err != nil {
			log.Warn().Err(err).Msg("Missing session cookie.")

			c.AbortWithStatus(401)
			c.Error(err)
		}

		decipheredCookie, err :=  services.Decipher(cookie.Value)

		if err != nil {
			log.Warn().Err(err).Msg("Invalid session cookie.")

			c.AbortWithStatus(401)
			c.Error(err)
		}
		
		// TODO: Authenticate logic using external provider

		log.Debug().Str("Authenticated request with cookie: ", decipheredCookie)

		c.Next() // Continue processing the request when authentication is valid
	}
}
