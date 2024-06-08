package middleware

import (
	"context"
	"database/sql"
	"intern_traning/internal/azuread"
	"intern_traning/internal/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

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
		_ = db.QueryRow("SELECT id FROM users WHERE oid = $1 AND deleted_at IS NULL", tokenData.ObjectID).Scan(&id)
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
		ctx = context.WithValue(ctx, "user_id", id)
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
