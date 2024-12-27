package middleware

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Proxy web requests for the given target.
func ProxyWeb(targetURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the target URL
		target, err := url.Parse(targetURL)

		if err != nil {
			log.Err(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid proxy target URL"})
			return
		}

		// Create a reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(target)
		
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = target.Host
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = target.Path + c.Param("any")
		}

		// Serve the proxy
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
