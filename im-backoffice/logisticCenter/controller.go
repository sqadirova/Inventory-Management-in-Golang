package logisticCenter

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"im-backoffice/CONSTANTS"
	myErrors "im-backoffice/errors"
	"strconv"
)

type LogisticCenterController struct {
	service   ILogisticCenterService
	validator *validator.Validate
}

func NewLogisticCenterController(service ILogisticCenterService) LogisticCenterController {
	return LogisticCenterController{
		service:   service,
		validator: validator.New(),
	}
}

// GetLogisticCenterByID godoc
// @Summary Gets an logisticCenter by ID.
// @Tags logistic-center
// @Accept json
// @Produce json
// @Param logisticCenterId path string true "LogisticCenter ID"
// @Security ApiKeyAuth
// @Success 200 {object} LogisticCenterResponseDTO
// @Failure 400 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Router /logistic-center/{logisticCenterId} [get]
func (lcc LogisticCenterController) GetLogisticCenterByID(ctx *fiber.Ctx) error {
	paramId := ctx.Params("logisticCenterId")

	if paramId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	logisticCenterId, err := uuid.FromString(paramId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	logisticCenter, err := lcc.service.FindById(ctx.Context(), logisticCenterId)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(logisticCenter)
}

// GetAllLogisticCentersWithPagination godoc
// @Summary Get all logisticCenters with pagination.
// @Tags logistic-center
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page path int true "Page"
// @Param limit query int true "Limit"
// @Success 200 {object} LogisticCenterResponseDTO
// @Failure 400 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Router /logistic-centers/{page} [get]
func (lcc LogisticCenterController) GetAllLogisticCentersWithPagination(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Params("page"))

	if err != nil || page <= 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	queryParams := new(GetAllLogisticCentersQuery)

	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	if queryParams.Limit != 0 {
		err := lcc.validator.Struct(queryParams)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).
				JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
		}
	}

	logisticCenters, nextPage, totalPages, err := lcc.service.FindAllWithPagination(ctx.Context(), page, queryParams)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	response := ToAllLogisticCentersResponse(logisticCenters, *nextPage, *totalPages)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// CreateLogisticCenter godoc
// @Summary Create a new logisticCenter.
// @Tags logistic-center
// @Accept json
// @Produce json
// @Param input   body  LogisticCenterRequest   true  "LogisticCenter"
// @Security ApiKeyAuth
// @Success 201 {object} LogisticCenterDTO
// @Failure 400 {object} errors.Response
// @Failure 403 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Failure 409 {object} errors.Response
// @Router /logistic-center [post]
func (lcc LogisticCenterController) CreateLogisticCenter(ctx *fiber.Ctx) error {
	var createLogisticCenterDTO LogisticCenterRequest

	if err := ctx.BodyParser(&createLogisticCenterDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	err := lcc.validator.Struct(createLogisticCenterDTO)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	logisticCenter, err := lcc.service.Create(ctx.Context(), createLogisticCenterDTO)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(logisticCenter)
}

// UpdateLogisticCenter godoc
// @Summary Update the logisticCenter.
// @Tags logistic-center
// @Accept json
// @Produce json
// @Param input   body  LogisticCenterRequest   true  "LogisticCenter DTO"
// @Param logisticCenterId path string true "LogisticCenter ID"
// @Security ApiKeyAuth
// @Success 200 {object} LogisticCenterResponseDTO
// @Failure 400 {object} errors.Response
// @Failure 403 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Failure 409 {object} errors.Response
// @Router /logistic-center/{logisticCenterId} [put]
func (lcc LogisticCenterController) UpdateLogisticCenter(ctx *fiber.Ctx) error {
	paramId := ctx.Params("logisticCenterId")

	if paramId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	logisticCenterId, err := uuid.FromString(paramId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	var updateLogisticCenterDTO LogisticCenterRequest

	if err := ctx.BodyParser(&updateLogisticCenterDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	err = lcc.validator.Struct(updateLogisticCenterDTO)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	logisticCenter, err := lcc.service.Update(ctx.Context(), logisticCenterId, updateLogisticCenterDTO)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(logisticCenter)
}

// DeleteLogisticCenter godoc
// @Summary Create a new logisticCenter.
// @Tags logistic-center
// @Produce json
// @Param logisticCenterId path string true "LogisticCenter ID"
// @Security ApiKeyAuth
// @Success 204 {object} nil
// @Failure 400 {object} errors.Response
// @Failure 403 {object} errors.Response
// @Router /logistic-center/{logisticCenterId} [delete]
func (lcc LogisticCenterController) DeleteLogisticCenter(ctx *fiber.Ctx) error {
	paramId := ctx.Params("logisticCenterId")

	if paramId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	logisticCenterId, err := uuid.FromString(paramId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	err = lcc.service.Delete(ctx.Context(), logisticCenterId)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}

// GetAllLogisticCenters godoc
// @Summary Get all logisticCenters.
// @Tags logistic-center
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} LogisticCenterResponseDTO
// @Failure 404 {object} errors.Response
// @Router /logistic-centers [get]
func (lcc LogisticCenterController) GetAllLogisticCenters(ctx *fiber.Ctx) error {
	logisticCenters, err := lcc.service.FindAll(ctx.Context())

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(logisticCenters)
}
