package location

import (
	"github.com/gofiber/fiber/v2"
	"im-backoffice/CONSTANTS"
	"im-backoffice/middleware"
)

func RouterLocation(app fiber.Router, service *LocationServiceImpl) {
	controller := NewLocationController(service)
	app.Post("/location", middleware.Protect(CONSTANTS.WORKER), controller.CreateLocation)
	//app.Get("/location/:locationId", middleware.Protect(CONSTANTS.WORKER), controller.GetLocationById)
	//app.Get("/locations/:page<min(1)>", middleware.Protect(CONSTANTS.WORKER), controller.GetAllLocationsWithPagination)
	app.Put("/location/:locationId", middleware.Protect(CONSTANTS.WORKER), controller.UpdateLocation)
	app.Delete("/location/:locationId", middleware.Protect(CONSTANTS.WORKER), controller.DeleteLocation)
	//app.Get("/locations", middleware.Protect(CONSTANTS.WORKER), controller.GetAllLocations)
}
