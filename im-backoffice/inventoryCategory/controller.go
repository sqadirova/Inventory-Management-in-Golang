package inventoryCategory

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"im-backoffice/CONSTANTS"
	myErrors "im-backoffice/errors"
	"strconv"
)

type InventoryCategoryController struct {
	service   IInventoryCategoryService
	validator *validator.Validate
}

func NewInventoryCategoryController(service IInventoryCategoryService) InventoryCategoryController {
	return InventoryCategoryController{
		service:   service,
		validator: validator.New(),
	}
}

// GetInventoryCategoryByID godoc
// @Summary Gets an inventory category by ID.
// @Tags inventory-category
// @Accept json
// @Produce json
// @Param inventoryCategoryId path string true "Inventory Category ID"
// @Security ApiKeyAuth
// @Success 200 {object} InventoryCategoryDTO
// @Failure 400 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Router /inventory-category/{inventoryCategoryId} [get]
func (icc InventoryCategoryController) GetInventoryCategoryByID(ctx *fiber.Ctx) error {
	paramId := ctx.Params("inventoryCategoryId")

	if paramId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	inventoryCategoryId, err := uuid.FromString(paramId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	inventoryCategory, err := icc.service.FindById(ctx.Context(), inventoryCategoryId)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(inventoryCategory)
}

// GetAllInventoryCategoriesWithPagination godoc
// @Summary Get all inventory categories with pagination.
// @Tags inventory-category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page path int true "Page"
// @Param limit query int true "Limit"
// @Success 200 {array} InventoryCategoryResponse
// @Failure 400 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Router /inventory-categories/{page} [get]
func (icc InventoryCategoryController) GetAllInventoryCategoriesWithPagination(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Params("page"))

	if err != nil || page <= 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	queryParams := new(GetAllInventoryCategoriesQuery)

	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	if queryParams.Limit != 0 {
		err := icc.validator.Struct(queryParams)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).
				JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
		}
	}

	inventoryCategories, err := icc.service.FindAllWithPagination(ctx.Context(), page, queryParams)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(inventoryCategories)
}

// CreateInventoryCategory godoc
// @Summary Create a new inventory category.
// @Tags inventory-category
// @Accept json
// @Produce json
// @Param input   body  InventoryCategoryReq   true  "Inventory Category"
// @Security ApiKeyAuth
// @Success 201 {object} InventoryCategoryDTO
// @Failure 400 {object} errors.Response
// @Failure 403 {object} errors.Response
// @Router /inventory-category [post]
func (icc InventoryCategoryController) CreateInventoryCategory(ctx *fiber.Ctx) error {
	var createInventoryCategoryDTO InventoryCategoryReq

	if err := ctx.BodyParser(&createInventoryCategoryDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	err := icc.validator.Struct(createInventoryCategoryDTO)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	inventoryCategory, err := icc.service.Create(ctx.Context(), createInventoryCategoryDTO)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusCreated).JSON(inventoryCategory)
}

// UpdateInventoryCategory godoc
// @Summary Update the inventory category.
// @Tags inventory-category
// @Accept json
// @Produce json
// @Param input body InventoryCategoryReq true "Inventory Category"
// @Security ApiKeyAuth
// @Param inventoryCategoryId path string true "Inventory Category ID"
// @Success 200 {object} InventoryCategoryDTO
// @Failure 400 {object} errors.Response
// @Failure 403 {object} errors.Response
// @Router /inventory-category/{inventoryCategoryId} [put]
func (icc InventoryCategoryController) UpdateInventoryCategory(ctx *fiber.Ctx) error {
	paramId := ctx.Params("inventoryCategoryId")

	if paramId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	inventoryCategoryId, err := uuid.FromString(paramId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	var updateInventoryCategoryDTO InventoryCategoryReq

	if err := ctx.BodyParser(&updateInventoryCategoryDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	err = icc.validator.Struct(updateInventoryCategoryDTO)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	inventoryCategory, err := icc.service.Update(ctx.Context(), inventoryCategoryId, updateInventoryCategoryDTO)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(inventoryCategory)
}

// DeleteInventoryCategory godoc
// @Summary Create a new inventory category.
// @Tags inventory-category
// @Produce json
// @Param inventoryCategoryId path string true "Inventory Category ID"
// @Security ApiKeyAuth
// @Success 204 {object} nil
// @Failure 400 {object} errors.Response
// @Failure 403 {object} errors.Response
// @Router /inventory-category/{inventoryCategoryId} [delete]
func (icc InventoryCategoryController) DeleteInventoryCategory(ctx *fiber.Ctx) error {
	paramId := ctx.Params("inventoryCategoryId")

	if paramId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	inventoryCategoryId, err := uuid.FromString(paramId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	err = icc.service.Delete(ctx.Context(), inventoryCategoryId)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}

// GetAllInventoryCategories godoc
// @Summary Get all inventory categories.
// @Tags inventory-category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} InventoryCategoryDTO
// @Failure 404 {object} errors.Response
// @Router /inventory-categories [get]
func (icc InventoryCategoryController) GetAllInventoryCategories(ctx *fiber.Ctx) error {
	inventoryCategories, err := icc.service.FindAll(ctx.Context())

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(inventoryCategories)
}
