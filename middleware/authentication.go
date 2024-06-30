package middleware

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"trec/internal/azuread"
	"trec/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type BodyQuery struct {
	OperationName string `json:"operationName"`
}

type Payload struct {
	UserID   uuid.UUID
	ForOwner bool
	ForTeam  bool
	ForAll   bool
}

// AuthenticateMiddleware is a middleware to authenticate user
func AuthenticateMiddleware(oauthClient azuread.AzureADOAuth, db *sql.DB, logger *zap.Logger) func(c *gin.Context) {
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
		var id uuid.UUID
		var deleted_at any
		err = db.QueryRow("SELECT id, deleted_at FROM users WHERE oid = $1", tokenData.ObjectID).Scan(&id, &deleted_at)
		if err != nil && err != sql.ErrNoRows {
			logger.Error("", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.WrapGQLUnauthorizedError(ctx))
			return
		}
		if deleted_at != nil {
			logger.Error(fmt.Sprintf("User %s is deleted", tokenData.ObjectID))
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.WrapGQLUnauthorizedError(ctx))
			return
		}
		if id == uuid.Nil {
			_, err = db.Query("WITH upsert AS ( UPDATE users SET name = $2, work_email = $3 WHERE oid = $1 RETURNING * ) "+
				"INSERT INTO users (oid, name, work_email) SELECT $1, $2, $3 WHERE NOT EXISTS ( SELECT 1 FROM upsert );", tokenData.ObjectID,
				tokenData.Name,
				tokenData.PreferredUsername)
			if err != nil {
				logger.Error("", zap.Error(err))
				c.AbortWithStatusJSON(http.StatusUnauthorized, util.WrapGQLUnauthorizedError(ctx))
				return
			}
			_ = db.QueryRow("SELECT id FROM users WHERE oid = $1 AND deleted_at IS NULL", tokenData.ObjectID).Scan(&id)
		}
		ctx = context.WithValue(ctx, Payload{}, &Payload{
			UserID: id,
		})
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func ValidPermission(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}
		var operationName string
		var userId uuid.UUID
		var forOwner, forTeam, forAll bool
		payload := c.Request.Context().Value(Payload{}).(*Payload)
		userId = payload.UserID
		if c.Request.Method == http.MethodPost && c.Request.URL.Path == "/graphql" {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				log.Printf("Error reading request body: %v", err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			var bodyQuery BodyQuery
			_ = json.Unmarshal(bodyBytes, &bodyQuery)
			operationName = bodyQuery.OperationName
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
		id := uuid.UUID{}
		err := db.QueryRow("SELECT id FROM permissions WHERE operation_name LIKE '%' || $1 || '%'", operationName).Scan(&id)
		if err != nil && err != sql.ErrNoRows {
			log.Printf("Error reading user permissions: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error reading user permissions"})
			return
		}
		if id != uuid.Nil {
			err = db.QueryRow("SELECT for_owner, for_team, for_all FROM entity_permissions WHERE entity_id = $1 AND permission_id = $2 AND entity_type = 'user'", userId, id).Scan(&forOwner, &forTeam, &forAll)
			if err != nil {
				if err == sql.ErrNoRows {
					c.AbortWithStatusJSON(http.StatusForbidden, util.WrapGQLPermissionError(ctx))
					return
				} else {
					log.Printf("Error reading user permissions: %v", err)
					c.JSON(http.StatusInternalServerError, gin.H{"message": "Error reading user permissions"})
					return
				}
			}
		}
		ctx = context.WithValue(ctx, Payload{}, &Payload{
			UserID:   userId,
			ForOwner: forOwner,
			ForTeam:  forTeam,
			ForAll:   forAll,
		})
		c.Request = c.Request.WithContext(ctx)
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
