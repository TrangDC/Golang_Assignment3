package service

import (
	"context"
	"net/http"

	"github.com/techvifyharrydo/Room-Reservation/ent"
	"github.com/techvifyharrydo/Room-Reservation/internal/azuread"
	"github.com/techvifyharrydo/Room-Reservation/internal/util"
	"go.uber.org/zap"
)

type OAuthLoginCallbackRequest struct {
	Code  string `form:"code" binding:"required"`
	State string `form:"state" binding:"required"`
}

// RefreshTokenRequest is the request for refresh token.
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// AuthService is the interface for all auth services.
type AuthService interface {
	GetRedirectLoginEndpoint(ctx context.Context, r *http.Request, w http.ResponseWriter) (string, error)
	GetToken(ctx context.Context, request OAuthLoginCallbackRequest, r *http.Request) (ent.AuthenticationToken, error)
	RefreshToken(ctx context.Context, request RefreshTokenRequest) (ent.AuthenticationToken, error)
	GetAzuredID(ctx context.Context, request OAuthLoginCallbackRequest, r *http.Request) (string, error)
}

// authSvcImpl is the implementation of AuthService.
type authSvcImpl struct {
	oauthClient azuread.AzureADOAuth
	logger      *zap.Logger
}

// NewAuthService creates a new AuthService.
func NewAuthService(oAuthService azuread.AzureADOAuth, logger *zap.Logger) AuthService {
	return &authSvcImpl{
		oauthClient: oAuthService,
		logger:      logger,
	}
}

// GetRedirectLoginEndpoint returns the redirect login endpoint.
func (svc *authSvcImpl) GetRedirectLoginEndpoint(ctx context.Context, r *http.Request, w http.ResponseWriter) (string, error) {
	url, err := svc.oauthClient.GetOAuthLoginEndpoint(ctx, w, r)
	if err != nil {
		svc.logger.Error("Getting error when get oauth redirect login endpoint", zap.Error(err))
		return "", util.WrapGQLInternalError(ctx)
	}

	return url, nil
}

// GetToken returns the token.
func (svc *authSvcImpl) GetToken(ctx context.Context, request OAuthLoginCallbackRequest, r *http.Request) (ent.AuthenticationToken, error) {
	oAuthToken, err := svc.oauthClient.GetOAuthTokenFromCallback(ctx, r, request.Code, request.State)
	if err != nil {
		svc.logger.Warn("Getting error when exchange token", zap.Error(err), zap.String("code", request.Code))
		return ent.AuthenticationToken{}, util.WrapGQLBadRequestError(ctx, "Invalid callback request")
	}
	token, nonce, err := svc.oauthClient.GetToken(ctx, oAuthToken)
	if err != nil {
		svc.logger.Warn("Getting error when authorize token", zap.Error(err))
		return ent.AuthenticationToken{}, util.WrapGQLBadRequestError(ctx, "Invalid callback request for get token")
	}
	err = svc.oauthClient.ValidateNonce(ctx, r, nonce)
	if err != nil {
		svc.logger.Warn("Getting error when validate nonce", zap.Error(err))
		return ent.AuthenticationToken{}, util.WrapGQLBadRequestError(ctx, "Invalid callback request for validate nonce")
	}

	return mapTokenResponse(token), nil
}

// RefreshToken returns the token.
func (svc *authSvcImpl) RefreshToken(ctx context.Context, input RefreshTokenRequest) (ent.AuthenticationToken, error) {
	token, err := svc.oauthClient.RefreshToken(ctx, input.RefreshToken)
	if err != nil {
		svc.logger.Warn("Getting error when refresh oauth token", zap.Error(err))
		return ent.AuthenticationToken{}, util.WrapGQLBadRequestError(ctx, "Invalid refresh token")
	}

	return mapTokenResponse(token), nil
}

// GetAzuredID GetToken returns the token.
func (svc *authSvcImpl) GetAzuredID(ctx context.Context, request OAuthLoginCallbackRequest, r *http.Request) (string, error) {
	oAuthToken, err := svc.oauthClient.GetOAuthTokenFromCallback(ctx, r, request.Code, request.State)
	if err != nil {
		svc.logger.Warn("Getting error when exchange token", zap.Error(err), zap.String("code", request.Code))
		return "", util.WrapGQLBadRequestError(ctx, "Invalid callback request")
	}

	_, nonce, err := svc.oauthClient.GetToken(ctx, oAuthToken)
	if err != nil {
		svc.logger.Warn("Getting error when authorize token", zap.Error(err))
		return "", util.WrapGQLBadRequestError(ctx, "Invalid callback request for get token")
	}
	err = svc.oauthClient.ValidateNonce(ctx, r, nonce)
	if err != nil {
		svc.logger.Warn("Getting error when validate nonce", zap.Error(err))
		return "", util.WrapGQLBadRequestError(ctx, "Invalid callback request for validate nonce")
	}

	return nonce, nil
}

// mapTokenResponse maps the token response.
func mapTokenResponse(oAuthToken azuread.JwtToken) ent.AuthenticationToken {
	return ent.AuthenticationToken{
		AccessToken:  oAuthToken.AccessToken,
		RefreshToken: oAuthToken.RefreshToken,
		TokenType:    oAuthToken.TokenType,
		ExpiresAt:    oAuthToken.Expiry,
		Email:        oAuthToken.Email,
	}
}
