package middleware

import (
	"database/sql"
	"net/http"
	"strings"
	"trec/internal/azuread"
	"trec/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AuthenticateMiddleware is a middleware to authenticate user
func AuthenticateMiddleware(oauthClient azuread.AzureADOAuth, db *sql.DB) func(c *gin.Context) {
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
		tokenData, err := oauthClient.DecodeToken(ctx, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.WrapGQLUnauthorizedError(ctx))
			return
		}
		err = db.QueryRow("WITH upsert AS ( UPDATE users SET name = $2, work_email = $3 WHERE oid = $1 RETURNING * ) "+
			"INSERT INTO users (oid, name, work_email) SELECT $1, $2, $3 WHERE NOT EXISTS ( SELECT 1 FROM upsert );").Scan(uuid.MustParse(tokenData.ObjectID), tokenData.Name, tokenData.PreferredUsername)
		if err != nil {
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
