package location

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"im-backoffice/CONSTANTS"
	db "im-backoffice/db/sqlc"
	myErrors "im-backoffice/errors"
	"log"
)

type ILocationService interface {
	Create(ctx context.Context, createLocationDTO LocationDTOReq) (json.RawMessage, error)
	Update(ctx context.Context, locationId uuid.UUID, updateLocationDTO LocationDTOReq) (json.RawMessage, error)
	Delete(ctx context.Context, locationId uuid.UUID) error
	//FindAll(ctx context.Context) ([]json.RawMessage, error)
	//FindById(ctx context.Context, locationId uuid.UUID) (json.RawMessage, error)
	//FindAllWithPagination(ctx context.Context, page int, queryParams *GetAllLocationsQuery) ([]json.RawMessage, *int, *int, error)
}

type LocationServiceImpl struct {
	repository db.Repository
}

func GetNewLocationService(repository db.Repository) *LocationServiceImpl {
	return &LocationServiceImpl{
		repository: repository,
	}
}

func (ls *LocationServiceImpl) Create(ctx context.Context, createLocationDTO LocationDTOReq) (json.RawMessage, error) {
	location := db.ImLocation{}

	dublicateLocation, _ := ls.repository.GetLocationByWarehouseId(ctx, db.GetLocationByWarehouseIdParams{
		WarehouseID:  createLocationDTO.WarehouseID,
		LocationName: createLocationDTO.LocationName,
	})

	if dublicateLocation != location {
		log.Println("Already exist in this warehouse.")
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("duplicate_err", CONSTANTS.LANGUAGE))
	}

	location, err := ls.repository.CreateLocation(ctx, db.CreateLocationParams{
		LocationName: createLocationDTO.LocationName,
		WarehouseID:  createLocationDTO.WarehouseID,
	})

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusForbidden,
			myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	response, err := ls.repository.GetLocationById(ctx, location.ID)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	return response, nil
}

func (ls *LocationServiceImpl) Update(ctx context.Context, locationId uuid.UUID, updateLocationDTO LocationDTOReq) (json.RawMessage, error) {
	_, err := ls.repository.GetLocationById(ctx, locationId)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	emptyLocation := db.ImLocation{}

	dublicateLocation, _ := ls.repository.GetLocationByWarehouseId(ctx, db.GetLocationByWarehouseIdParams{
		WarehouseID:  updateLocationDTO.WarehouseID,
		LocationName: updateLocationDTO.LocationName,
	})

	if dublicateLocation != emptyLocation {
		log.Println("Already exist in this warehouse.")
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("duplicate_err", CONSTANTS.LANGUAGE))
	}

	location, err := ls.repository.UpdateLocation(ctx,
		db.UpdateLocationParams{
			ID:           locationId,
			LocationName: updateLocationDTO.LocationName,
			WarehouseID:  updateLocationDTO.WarehouseID,
		})

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusForbidden, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	response, err := ls.repository.GetLocationById(ctx, location.ID)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	//dto := ToLocationDto(location)

	return response, nil
}

func (ls *LocationServiceImpl) Delete(ctx context.Context, locationId uuid.UUID) error {
	_, err := ls.repository.GetLocationById(ctx, locationId)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	err = ls.repository.DeleteLocation(ctx, locationId)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusForbidden,
			myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	return nil
}

//func (ls *LocationServiceImpl) FindById(ctx context.Context, locationId uuid.UUID) (json.RawMessage, error) {
//	location, err := ls.repository.GetLocationById(ctx, locationId)
//
//	if err != nil {
//		log.Println(err)
//		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
//	}
//
//	//dto := ToLocationDto(location)
//
//	return location, nil
//}
//
//func (ls *LocationServiceImpl) FindAllWithPagination(ctx context.Context, page int, queryParams *GetAllLocationsQuery) ([]json.RawMessage, *int, *int, error) {
//	count, err := ls.repository.CountOfLocations(ctx)
//
//	if err != nil {
//		log.Println(err)
//		return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
//	}
//
//	pagination, nextPage, totalPages := util.AddPagination(page, &util.Pagination{Limit: queryParams.Limit}, count)
//
//	if totalPages == 0 {
//		return []json.RawMessage{}, &nextPage, &totalPages, nil
//	}
//
//	locations, err := ls.repository.GetLocationsWithPagination(ctx, db.GetLocationsWithPaginationParams{
//		Limit:  int32(pagination.Limit),
//		Offset: int32(pagination.Offset),
//	})
//
//	if err != nil {
//		log.Println(err)
//		return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
//	}
//
//	//dtoArr := ToLocationDtoArray(locations)
//	//response := ToLocationResponse(dtoArr, nextPage, totalPages)
//
//	return locations, &nextPage, &totalPages, nil
//}
//func (ls *LocationServiceImpl) FindAll(ctx context.Context) ([]json.RawMessage, error) {
//	locations, err := ls.repository.GetLocations(ctx)
//
//	if err != nil {
//		log.Println(err)
//		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
//	}
//
//	//dto := ToLocationDtoArray(locations)
//
//	return locations, nil
//}
