package middleware

import (
	"net/http"
	"strings"
	"trec/internal/azuread"
	"trec/internal/util"

	"github.com/gin-gonic/gin"
)

// AuthenticateMiddleware is a middleware to authenticate user
func AuthenticateMiddleware(oauthClient azuread.AzureADOAuth) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		// Skip pre-flight request
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		token := ParseBearerTokenFromRequest(c.Request)
		if len(token) == 0 || oauthClient.VerifyAccessToken(ctx, token) != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.WrapGQLUnauthorizedError(ctx))
			return
		}
		c.Next()
	}
}

// ParseBearerTokenFromRequest parses the bearer token from request
func ParseBearerTokenFromRequest(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) > 6 && strings.ToUpper(authHeader[0:6]) == "BEARER" {
		// Default jwt token
		return authHeader[7:]
	}

	return ""
}
