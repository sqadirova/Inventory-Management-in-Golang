package user

import (
	"IM-Golang/CONSTANTS"
	"IM-Golang/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouterUser(app fiber.Router, service *UserServiceImpl) {
	controller := NewUserController(service)

	app.Get("/users", middleware.Protect(CONSTANTS.ADMIN), controller.GetAllUsers)
	app.Post("/user", middleware.Protect(CONSTANTS.ADMIN), controller.CreateUser)
	app.Put("/user/:id", middleware.Protect(CONSTANTS.ADMIN), controller.UpdateUser)
	app.Get("/user/me", middleware.Protect(), controller.GetUserMe)
	app.Get("/user/roles", middleware.Protect(CONSTANTS.ADMIN), controller.GetAllRoles)
	app.Get("/user/roles/:id", middleware.Protect(CONSTANTS.ADMIN), controller.GetOneRole)
}
