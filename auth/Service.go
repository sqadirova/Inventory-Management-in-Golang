package auth

import (
	"IM-Golang/CONSTANTS"
	"IM-Golang/config"
	myErrors "IM-Golang/errors"
	"IM-Golang/keycloak"
	"context"
	"github.com/Nerzal/gocloak/v12"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

type IAuthService interface {
	signOut(ctx context.Context, refreshToken string) (bool, error)
	signIn(ctx context.Context, userReqBody SignInReq) (*SignInResp, error)
	refreshToken(ctx context.Context, refreshToken string) (*gocloak.JWT, error)
	setRefreshTokenToCookie(refreshToken string) *fiber.Cookie
}

type AuthServiceImpl struct {
	keycloak *keycloak.Keycloak
}

func GetNewAuthService(keycloak *keycloak.Keycloak) *AuthServiceImpl {
	return &AuthServiceImpl{
		keycloak: keycloak,
	}
}

func (auth *AuthServiceImpl) signOut(ctx context.Context, refreshToken string) (bool, error) {
	err := auth.keycloak.Gocloak.Logout(ctx, auth.keycloak.ClientId, auth.keycloak.ClientSecret, auth.keycloak.Realm, refreshToken)

	if err != nil {
		log.Println(err)
		return false, myErrors.NewHttpError(fiber.StatusForbidden, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	return true, nil
}

func (auth *AuthServiceImpl) signIn(ctx context.Context, userReqBody SignInReq) (*SignInResp, error) {
	jwt, err := auth.keycloak.Gocloak.Login(ctx,
		auth.keycloak.ClientId,
		auth.keycloak.ClientSecret,
		auth.keycloak.Realm,
		userReqBody.Username,
		userReqBody.Password)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("invalid_user_credentials", CONSTANTS.LANGUAGE))
	}

	return &SignInResp{AccessToken: jwt.AccessToken, RefreshToken: jwt.RefreshToken}, nil
}

func (auth *AuthServiceImpl) refreshToken(ctx context.Context, refreshToken string) (*gocloak.JWT, error) {
	token, err := auth.keycloak.Gocloak.RefreshToken(ctx, refreshToken,
		auth.keycloak.ClientId, auth.keycloak.ClientSecret, auth.keycloak.Realm)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	return token, nil
}

func (auth *AuthServiceImpl) setRefreshTokenToCookie(refreshToken string) *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = config.Configuration.Cookie.Name
	cookie.Value = refreshToken
	cookie.HTTPOnly = config.Configuration.Cookie.HttpOnly
	cookie.Secure = config.Configuration.Cookie.Secure
	cookie.SameSite = config.Configuration.Cookie.SameSite
	cookie.Expires = time.Now().Add(config.Configuration.Cookie.Expires)

	return cookie
}
