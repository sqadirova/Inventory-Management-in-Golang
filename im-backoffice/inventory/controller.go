package inventory

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"im-backoffice/CONSTANTS"
	myErrors "im-backoffice/errors"
	"log"
	"strconv"
)

type InventoryController struct {
	service   IInventoryService
	validator *validator.Validate
}

func NewInventoryController(service IInventoryService) InventoryController {
	return InventoryController{
		service:   service,
		validator: validator.New(),
	}
}

// GetAllInventories godoc
// @Summary Get all inventories.
// @Tags inventory
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page path int true "Page"
// @Param limit query int true "Limit"
// @Param q query string false "Search text (length min=3)"
// @Success 200 {array} InventoryResponse
// @Failure 400 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Router /inventories/{page} [get]
func (ic InventoryController) GetAllInventories(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Params("page"))

	if err != nil || page <= 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	queryParams := new(GetAllInventoriesQuery)

	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	if queryParams.InventoryName != "" || queryParams.Limit < 10 {
		err := ic.validator.Struct(queryParams)

		if err != nil {
			log.Println(err)
			return ctx.Status(fiber.StatusBadRequest).
				JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
		}
	}

	inventories, nextPage, totalPages, err := ic.service.FindAll(ctx.Context(), page, queryParams)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	response := ToAllInventoryResponseDTO(inventories, *nextPage, *totalPages)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// GetInventoryByID godoc
// @Summary Gets an inventory by ID.
// @Tags inventory
// @Accept json
// @Produce json
// @Param inventoryId path string true "Inventory ID"
// @Security ApiKeyAuth
// @Success 200 {object} InventoryResponse
// @Failure 400 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Router /inventory/{inventoryId} [get]
func (ic InventoryController) GetInventoryByID(ctx *fiber.Ctx) error {
	paramId := ctx.Params("inventoryId")

	if paramId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	inventoryId, err := uuid.FromString(paramId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	inventory, err := ic.service.FindById(ctx.Context(), inventoryId)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(inventory)
}

// CreateInventory godoc
// @Summary Create a new inventory.
// @Tags inventory
// @Accept json
// @Produce json
// @Param input   body  CreateInventoryRequest   true  "Inventory"
// @Security ApiKeyAuth
// @Success 201 {object} InventoryResponse
// @Failure 400 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Failure 500 {object} errors.Response
// @Router /inventory [post]
func (ic InventoryController) CreateInventory(ctx *fiber.Ctx) error {
	var createInventoryRequest CreateInventoryRequest

	if err := ctx.BodyParser(&createInventoryRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	err := ic.validator.Struct(createInventoryRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	inventory, err := ic.service.Create(ctx.Context(), createInventoryRequest)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusCreated).JSON(inventory)
}

// UpdateInventory godoc
// @Summary Update the inventory.
// @Tags inventory
// @Accept json
// @Produce json
// @Param input   body  UpdateInventoryRequest   true  "Inventory DTO"
// @Security ApiKeyAuth
// @Param inventoryId path string true "Inventory ID"
// @Success 200 {object} InventoryResponse
// @Failure 400 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Failure 500 {object} errors.Response
// @Router /inventory/{inventoryId} [put]
func (ic InventoryController) UpdateInventory(ctx *fiber.Ctx) error {
	paramId := ctx.Params("inventoryId")

	if paramId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	inventoryId, err := uuid.FromString(paramId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	var updateInventoryRequest UpdateInventoryRequest

	if err = ctx.BodyParser(&updateInventoryRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	err = ic.validator.Struct(updateInventoryRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	inventory, err := ic.service.Update(ctx.Context(), inventoryId, updateInventoryRequest)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(inventory)
}

// DeleteInventory godoc
// @Summary Delete inventory by id.
// @Tags inventory
// @Produce json
// @Param inventoryId path string true "Inventory ID"
// @Security ApiKeyAuth
// @Success 400 {object} errors.Response
// @Failure 400 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Failure 403 {object} errors.Response
// @Router /inventory/{inventoryId} [delete]
func (ic InventoryController) DeleteInventory(ctx *fiber.Ctx) error {
	paramId := ctx.Params("inventoryId")

	if paramId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	//in feature we will add delete inventory physically

	return ctx.Status(fiber.StatusBadRequest).JSON(
		fiber.Map{"key": "delete_inventory_msg",
			"message": "Contact Administrator to delete inventory",
		})
}
