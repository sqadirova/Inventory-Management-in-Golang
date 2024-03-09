package logisticCenter

import (
	"github.com/gofiber/fiber/v2"
	"im-backoffice/CONSTANTS"
	"im-backoffice/middleware"
)

func RouterLogisticCenter(app fiber.Router, service *LogisticCenterServiceImpl) {
	controller := NewLogisticCenterController(service)
	app.Post("/logistic-center",
		middleware.Protect(CONSTANTS.WORKER), controller.CreateLogisticCenter)

	app.Put("/logistic-center/:logisticCenterId",
		middleware.Protect(CONSTANTS.WORKER), controller.UpdateLogisticCenter)

	app.Delete("/logistic-center/:logisticCenterId",
		middleware.Protect(CONSTANTS.WORKER), controller.DeleteLogisticCenter)

	app.Get("/logistic-centers/:page<min(1)>",
		middleware.Protect(CONSTANTS.WORKER), controller.GetAllLogisticCentersWithPagination)

	app.Get("/logistic-center/:logisticCenterId",
		middleware.Protect(CONSTANTS.WORKER), controller.GetLogisticCenterByID)

	app.Get("/logistic-centers",
		middleware.Protect(CONSTANTS.WORKER, CONSTANTS.INSPECTOR), controller.GetAllLogisticCenters)
}
