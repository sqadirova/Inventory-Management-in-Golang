package warehouse

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"im-backoffice/CONSTANTS"
	myErrors "im-backoffice/errors"
)

type WarehouseController struct {
	service   IWarehouseService
	validator *validator.Validate
}

func NewWarehouseController(service IWarehouseService) WarehouseController {
	return WarehouseController{
		service:   service,
		validator: validator.New(),
	}
}

// CreateWarehouse godoc
// @Summary Create a new warehouse
// @Tags warehouse
// @Accept json
// @Produce json
// @Param input   body  WarehouseReq   true  "Warehouse"
// @Security ApiKeyAuth
// @Success 201 {object} WarehouseRes
// @Failure 400 {object} errors.Response
// @Failure 403 {object} errors.Response
// @Failure 409 {object} errors.Response
// @Router /warehouse [post]
func (wc WarehouseController) CreateWarehouse(ctx *fiber.Ctx) error {
	var createWarehouseDTO WarehouseReq

	if err := ctx.BodyParser(&createWarehouseDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	err := wc.validator.Struct(createWarehouseDTO)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	warehouse, err := wc.service.Create(ctx.Context(), createWarehouseDTO)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusCreated).JSON(warehouse)
}

// UpdateWarehouse godoc
// @Summary Update the warehouse.
// @Tags warehouse
// @Accept json
// @Produce json
// @Param input body  WarehouseReq   true  "Warehouse DTO"
// @Security ApiKeyAuth
// @Param warehouseId path string true "Warehouse ID"
// @Success 200 {object} WarehouseRes
// @Failure 400 {object} errors.Response
// @Failure 403 {object} errors.Response
// @Router /warehouse/{warehouseId} [put]
func (wc WarehouseController) UpdateWarehouse(ctx *fiber.Ctx) error {
	paramId := ctx.Params("warehouseId")

	if paramId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	warehouseId, err := uuid.FromString(paramId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	var updateWarehouseDTO WarehouseReq

	if err := ctx.BodyParser(&updateWarehouseDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	err = wc.validator.Struct(updateWarehouseDTO)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	warehouse, err := wc.service.Update(ctx.Context(), warehouseId, updateWarehouseDTO)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(warehouse)
}

// DeleteWarehouse godoc
// @Summary Create a new warehouse
// @Tags warehouse
// @Produce json
// @Param warehouseId path string true "Warehouse ID"
// @Security ApiKeyAuth
// @Success 204 {object} nil
// @Failure 400 {object} errors.Response
// @Failure 403 {object} errors.Response
// @Router /warehouse/{warehouseId} [delete]
func (wc WarehouseController) DeleteWarehouse(ctx *fiber.Ctx) error {
	paramId := ctx.Params("warehouseId")

	if paramId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	warehouseId, err := uuid.FromString(paramId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	err = wc.service.Delete(ctx.Context(), warehouseId)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}

//
//// GetAllWarehousesWithPagination godoc
//// @Summary Get all warehouses with pagination.
//// @Tags warehouse
//// @Accept json
//// @Produce json
//// @Security ApiKeyAuth
//// @Param page path int true "Page"
//// @Param limit query int true "Limit"
//// @Success 200 {array} WarehouseResponse
//// @Failure 400 {object} errors.Response
//// @Failure 404 {object} errors.Response
//// @Router /warehouses/{page} [get]
//func (wc WarehouseController) GetAllWarehousesWithPagination(ctx *fiber.Ctx) error {
//	page, err := strconv.Atoi(ctx.Params("page"))
//
//	if err != nil || page <= 0 {
//		return ctx.Status(fiber.StatusBadRequest).
//			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
//	}
//
//	queryParams := new(GetAllWarehousesQuery)
//
//	if err := ctx.QueryParser(queryParams); err != nil {
//		return ctx.Status(fiber.StatusBadRequest).
//			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
//	}
//
//	if queryParams.Limit != 0 {
//		err := wc.validator.Struct(queryParams)
//
//		if err != nil {
//			return ctx.Status(fiber.StatusBadRequest).
//				JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
//		}
//	}
//
//	warehouses, nextPage, totalPages, err := wc.service.FindAllWithPagination(ctx.Context(), page, queryParams)
//
//	if err != nil {
//		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
//	}
//
//	response := ToWarehouseResponse(warehouses, *nextPage, *totalPages)
//
//	return ctx.Status(fiber.StatusOK).JSON(response)
//}
//
//// GetWarehouseByID godoc
//// @Summary Gets a warehouse by ID.
//// @Tags warehouse
//// @Accept json
//// @Produce json
//// @Param warehouseId path string true "Warehouse ID"
//// @Security ApiKeyAuth
//// @Success 200 {object} WarehouseRes
//// @Failure 400 {object} errors.Response
//// @Failure 404 {object} errors.Response
//// @Router /warehouse/{warehouseId} [get]
//func (wc WarehouseController) GetWarehouseByID(ctx *fiber.Ctx) error {
//	paramId := ctx.Params("warehouseId")
//
//	if paramId == "" {
//		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
//	}
//
//	warehouseId, err := uuid.FromString(paramId)
//
//	if err != nil {
//		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
//	}
//
//	warehouse, err := wc.service.FindById(ctx.Context(), warehouseId)
//
//	if err != nil {
//		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
//	}
//
//	return ctx.Status(fiber.StatusOK).JSON(warehouse)
//}
//
//// GetAllWarehouses godoc
//// @Summary Get all warehouses.
//// @Tags warehouse
//// @Accept json
//// @Produce json
//// @Security ApiKeyAuth
//// @Success 200 {array} WarehouseRes
//// @Failure 404 {object} errors.Response
//// @Router /warehouses [get]
//func (wc WarehouseController) GetAllWarehouses(ctx *fiber.Ctx) error {
//	warehouses, err := wc.service.FindAll(ctx.Context())
//
//	if err != nil {
//		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
//	}
//
//	return ctx.Status(fiber.StatusOK).JSON(warehouses)
//}
