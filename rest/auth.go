package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"intern_traning/internal/util"
	"intern_traning/service"
	"net/http"
	"net/url"
)

// AuthController is the interface for all auth controllers.
type AuthController interface {
	OAuthLogin(c *gin.Context)
	OAuthCallback(c *gin.Context)
	RefreshToken(c *gin.Context)
}

// authCtrlImpl is the implementation of AuthController.
type authCtrlImpl struct {
	authService       service.AuthService
	logger            *zap.Logger
	clientRedirectURL string
}

// NewAuthController creates a new AuthController.
func NewAuthController(authService service.AuthService, clientRedirectURL string, logger *zap.Logger) AuthController {
	return &authCtrlImpl{
		authService:       authService,
		clientRedirectURL: clientRedirectURL,
		logger:            logger,
	}
}

// OAuthLogin handles the OAuth login.
func (i authCtrlImpl) OAuthLogin(c *gin.Context) {
	redirectUrl, err := i.authService.GetRedirectLoginEndpoint(c, c.Request, c.Writer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusFound, redirectUrl)
}

// OAuthCallback handles the OAuth callback.
func (i authCtrlImpl) OAuthCallback(c *gin.Context) {
	ctx := c.Request.Context()
	if err := c.Query("error"); len(err) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	r := service.OAuthLoginCallbackRequest{}
	if err := c.ShouldBindQuery(&r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, util.WrapGQLBadRequestError(ctx, "Invalid request to create token"))
		return
	}
	token, err := i.authService.GetToken(ctx, r, c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, util.WrapGQLBadRequestError(ctx, "Invalid request to create token"))
		return
	}
	queryParams := url.Values{
		"accessToken":  {token.AccessToken},
		"refreshToken": {token.RefreshToken},
		"expiresAt":    {fmt.Sprintf("%d", token.ExpiresAt.Unix())},
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("%s?%s", i.clientRedirectURL, queryParams.Encode()))
}

// RefreshToken handles the refresh token.
func (i authCtrlImpl) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()

	r := service.RefreshTokenRequest{}
	if err := c.ShouldBindJSON(&r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, util.WrapGQLBadRequestError(ctx, "Invalid request to refresh token with invalid body"))
		return
	}

	token, err := i.authService.RefreshToken(ctx, r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, util.WrapGQLBadRequestError(ctx, "Invalid request to refresh token"))
		return
	}

	c.JSON(http.StatusOK, token)
}
