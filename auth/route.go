package auth

import (
	"IM-Golang/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouterAuth(app fiber.Router, service *AuthServiceImpl) {
	controller := NewAuthController(service)

	app.Post("/auth/sign-in", middleware.Common, controller.SignIn)
	app.Post("/auth/sign-out", middleware.Common, controller.SignOut)
	app.Post("/auth/refresh", middleware.Common, controller.GetRefreshToken)
}
