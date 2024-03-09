package auth

import (
	"github.com/go-playground/validator"
	fiber "github.com/gofiber/fiber/v2"
	"im-backoffice/CONSTANTS"
	"im-backoffice/config"
	myErrors "im-backoffice/errors"
	"im-backoffice/middleware"
)

type AuthController struct {
	service   IAuthService
	validator *validator.Validate
}

func NewAuthController(service IAuthService) AuthController {
	return AuthController{
		service:   service,
		validator: validator.New(),
	}
}

// SignIn godoc
// @Summary Sign in the user to system.
// @Tags auth
// @Accept json
// @Produce json
// @Param input   body  SignInReq   true  "Sign In Req"
// @Success 200 {object} SignInResp
// @Failure 400 {object} errors.Response
// @Failure 500 {object} errors.Response
// @Router /auth/sign-in [post]
func (ac AuthController) SignIn(ctx *fiber.Ctx) error {
	var userReqBody SignInReq

	if err := ctx.BodyParser(&userReqBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	err := ac.validator.Struct(userReqBody)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	signInResponse, err := ac.service.signIn(ctx.Context(), userReqBody)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	cookie := ac.service.setRefreshTokenToCookie(signInResponse.RefreshToken)
	ctx.Cookie(cookie)

	return ctx.Status(fiber.StatusOK).JSON(signInResponse)
}

// SignOut godoc
// @Summary Sign out the user from system.
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} SignOutDTO
// @Failure 401 {object} errors.Response
// @Failure 400 {object} errors.Response
// @Failure 403 {object} errors.Response
// @Router  /auth/sign-out [post]
func (ac AuthController) SignOut(ctx *fiber.Ctx) error {
	refreshToken := ctx.Cookies(config.Configuration.Cookie.Name)

	token := middleware.ExtractBearerToken(ctx.GetReqHeaders()["Authorization"])

	if token == "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("bearer_token_missing", CONSTANTS.LANGUAGE))
	}

	_, claims, err := middleware.ParseToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).
			JSON(myErrors.NewResponseByKey("unauthorized", CONSTANTS.LANGUAGE))
	}

	isLogout, err := ac.service.signOut(ctx.Context(), refreshToken)

	if err != nil || !isLogout {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	ctx.ClearCookie(config.Configuration.Cookie.Name)

	return ctx.Status(fiber.StatusOK).JSON(SignOutDTO{UserID: claims["sub"].(string)})
}

// GetRefreshToken godoc
// @Summary RefreshToken refreshes the given token.
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} RefreshTokenResponse
// @Failure 500 {object} errors.Response
// @Router  /auth/refresh [post]
func (ac AuthController) GetRefreshToken(ctx *fiber.Ctx) error {
	token, err := ac.service.refreshToken(ctx.Context(), ctx.Cookies(config.Configuration.Cookie.Name))

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	cookie := ac.service.setRefreshTokenToCookie(token.RefreshToken)
	ctx.Cookie(cookie)

	return ctx.Status(fiber.StatusOK).JSON(RefreshTokenResponse{AccessToken: token.AccessToken})
}
