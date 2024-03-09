package location

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"im-backoffice/CONSTANTS"
	myErrors "im-backoffice/errors"
)

type LocationController struct {
	service   ILocationService
	validator *validator.Validate
}

func NewLocationController(service ILocationService) LocationController {
	return LocationController{
		service:   service,
		validator: validator.New(),
	}
}

// CreateLocation godoc
// @Summary Create a new location.
// @Tags location
// @Accept json
// @Produce json
// @Param input  body  LocationDTOReq true "Location"
// @Security ApiKeyAuth
// @Success 201 {object} LocationDTO
// @Failure 400 {object} errors.Response
// @Failure 403 {object} errors.Response
// @Router /location [post]
func (lc LocationController) CreateLocation(ctx *fiber.Ctx) error {
	var createLocationDTO LocationDTOReq

	if err := ctx.BodyParser(&createLocationDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	err := lc.validator.Struct(createLocationDTO)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	location, err := lc.service.Create(ctx.Context(), createLocationDTO)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusCreated).JSON(location)
}

// UpdateLocation godoc
// @Summary Update the location.
// @Tags location
// @Accept json
// @Produce json
// @Param input body LocationDTOReq   true  "Location DTO"
// @Security ApiKeyAuth
// @Param locationId path string true "Location ID"
// @Success 200 {object} LocationDTO
// @Failure 400 {object} errors.Response
// @Failure 403 {object} errors.Response
// @Router /location/{locationId} [put]
func (lc LocationController) UpdateLocation(ctx *fiber.Ctx) error {
	paramId := ctx.Params("locationId")

	if paramId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	locationId, err := uuid.FromString(paramId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	var updateLocationDTO LocationDTOReq

	if err := ctx.BodyParser(&updateLocationDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	err = lc.validator.Struct(updateLocationDTO)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	location, err := lc.service.Update(ctx.Context(), locationId, updateLocationDTO)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(location)
}

// DeleteLocation godoc
// @Summary Delete location by id.
// @Tags location
// @Produce json
// @Param locationId path string true "Location ID"
// @Security ApiKeyAuth
// @Success 204 {object} nil
// @Failure 400 {object} errors.Response
// @Failure 403 {object} errors.Response
// @Router /location/{locationId} [delete]
func (lc LocationController) DeleteLocation(ctx *fiber.Ctx) error {
	paramId := ctx.Params("locationId")

	if paramId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	locationId, err := uuid.FromString(paramId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	err = lc.service.Delete(ctx.Context(), locationId)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}

//
//// GetAllLocationsWithPagination godoc
//// @Summary Get all locations with pagination.
//// @Tags location
//// @Accept json
//// @Produce json
//// @Security ApiKeyAuth
//// @Param page path int true "Page"
//// @Param limit query int true "Limit"
//// @Success 200 {array} LocationResponse
//// @Failure 400 {object} errors.Response
//// @Failure 404 {object} errors.Response
//// @Router /locations/{page} [get]
//func (lc LocationController) GetAllLocationsWithPagination(ctx *fiber.Ctx) error {
//	page, err := strconv.Atoi(ctx.Params("page"))
//
//	if err != nil || page <= 0 {
//		return ctx.Status(fiber.StatusBadRequest).
//			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
//	}
//
//	queryParams := new(GetAllLocationsQuery)
//
//	if err := ctx.QueryParser(queryParams); err != nil {
//		return ctx.Status(fiber.StatusBadRequest).
//			JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
//	}
//
//	if queryParams.Limit != 0 {
//		err := lc.validator.Struct(queryParams)
//
//		if err != nil {
//			return ctx.Status(fiber.StatusBadRequest).
//				JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
//		}
//	}
//
//	locations, err := lc.service.FindAllWithPagination(ctx.Context(), page, queryParams)
//
//	if err != nil {
//		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
//	}
//
//	return ctx.Status(fiber.StatusOK).JSON(locations)
//}
//
//// GetLocationById godoc
//// @Summary Gets a location by ID.
//// @Tags location
//// @Accept json
//// @Produce json
//// @Param locationId path string true "Location ID"
//// @Security ApiKeyAuth
//// @Success 200 {object} LocationDTO
//// @Failure 400 {object} errors.Response
//// @Failure 404 {object} errors.Response
//// @Router /location/{locationId} [get]
//func (lc LocationController) GetLocationById(ctx *fiber.Ctx) error {
//	paramId := ctx.Params("locationId")
//
//	if paramId == "" {
//		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
//	}
//
//	locationId, err := uuid.FromString(paramId)
//
//	if err != nil {
//		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
//	}
//
//	location, err := lc.service.FindById(ctx.Context(), locationId)
//
//	if err != nil {
//		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
//	}
//
//	return ctx.Status(fiber.StatusOK).JSON(location)
//}
//
//// GetAllLocations godoc
//// @Summary Get all locations.
//// @Tags location
//// @Accept json
//// @Produce json
//// @Security ApiKeyAuth
//// @Success 200 {array} LocationDTO
//// @Failure 404 {object} errors.Response
//// @Router /locations [get]
//func (lc LocationController) GetAllLocations(ctx *fiber.Ctx) error {
//	locations, err := lc.service.FindAll(ctx.Context())
//
//	if err != nil {
//		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
//	}
//
//	return ctx.Status(fiber.StatusOK).JSON(locations)
//}
