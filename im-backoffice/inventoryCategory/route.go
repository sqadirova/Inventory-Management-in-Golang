package inventoryCategory

import (
	"github.com/gofiber/fiber/v2"
	"im-backoffice/CONSTANTS"
	"im-backoffice/middleware"
)

func RouterInventoryCategory(app fiber.Router, service *InventoryCategoryServiceImpl) {
	controller := NewInventoryCategoryController(service)
	app.Get("/inventory-category/:inventoryCategoryId",
		middleware.Protect(CONSTANTS.WORKER), controller.GetInventoryCategoryByID)

	app.Get("/inventory-categories/:page<min(1)>",
		middleware.Protect(CONSTANTS.WORKER), controller.GetAllInventoryCategoriesWithPagination)

	app.Post("/inventory-category",
		middleware.Protect(CONSTANTS.WORKER), controller.CreateInventoryCategory)

	app.Put("/inventory-category/:inventoryCategoryId",
		middleware.Protect(CONSTANTS.WORKER), controller.UpdateInventoryCategory)

	app.Delete("/inventory-category/:inventoryCategoryId",
		middleware.Protect(CONSTANTS.WORKER), controller.DeleteInventoryCategory)

	app.Get("/inventory-categories",
		middleware.Protect(CONSTANTS.WORKER), controller.GetAllInventoryCategories)

}
