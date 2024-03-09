package inventory

import (
	"github.com/gofiber/fiber/v2"
	"im-backoffice/CONSTANTS"
	"im-backoffice/middleware"
)

func RouterInventory(app fiber.Router, service *InventoryServiceImpl) {
	controller := NewInventoryController(service)

	app.Post("/inventory",
		middleware.Protect(CONSTANTS.WORKER), controller.CreateInventory)

	app.Put("/inventory/:inventoryId",
		middleware.Protect(CONSTANTS.WORKER), controller.UpdateInventory)

	app.Delete("/inventory/:inventoryId",
		middleware.Protect(CONSTANTS.WORKER), controller.DeleteInventory)

	app.Get("/inventories/:page<min(1)>",
		middleware.Protect(CONSTANTS.WORKER, CONSTANTS.INSPECTOR), controller.GetAllInventories)

	app.Get("/inventory/:inventoryId",
		middleware.Protect(CONSTANTS.WORKER, CONSTANTS.INSPECTOR), controller.GetInventoryByID)
}
