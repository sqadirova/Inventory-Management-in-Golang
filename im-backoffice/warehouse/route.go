package warehouse

import (
	"github.com/gofiber/fiber/v2"
	"im-backoffice/CONSTANTS"
	"im-backoffice/middleware"
)

func RouterWarehouse(app fiber.Router, service *WarehouseServiceImpl) {
	controller := NewWarehouseController(service)

	app.Post("/warehouse",
		middleware.Protect(CONSTANTS.WORKER), controller.CreateWarehouse)

	app.Put("/warehouse/:warehouseId",
		middleware.Protect(CONSTANTS.WORKER), controller.UpdateWarehouse)

	app.Delete("/warehouse/:warehouseId",
		middleware.Protect(CONSTANTS.WORKER), controller.DeleteWarehouse)

	//app.Get("/warehouses/:page<min(1)>",
	//	middleware.Protect(CONSTANTS.WORKER), controller.GetAllWarehousesWithPagination)
	//
	//app.Get("/warehouse/:warehouseId",
	//	middleware.Protect(CONSTANTS.WORKER), controller.GetWarehouseByID)
	//
	//app.Get("/warehouses",
	//	middleware.Protect(CONSTANTS.WORKER), controller.GetAllWarehouses)
}
