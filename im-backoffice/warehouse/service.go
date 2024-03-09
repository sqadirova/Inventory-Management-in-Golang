package warehouse

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

type IWarehouseService interface {
	Create(ctx context.Context, createWarehouseDTO WarehouseReq) (json.RawMessage, error)
	Update(ctx context.Context, warehouseId uuid.UUID, updateWarehouseDTO WarehouseReq) (json.RawMessage, error)
	Delete(ctx context.Context, warehouseId uuid.UUID) error
	//FindAll(ctx context.Context) ([]json.RawMessage, error)
	//FindById(ctx context.Context, warehouseId uuid.UUID) (json.RawMessage, error)
	//FindAllWithPagination(ctx context.Context, page int, queryParams *GetAllWarehousesQuery) ([]json.RawMessage, *int, *int, error)
}

type WarehouseServiceImpl struct {
	repository db.Repository
}

func GetNewWarehouseService(repository db.Repository) *WarehouseServiceImpl {
	return &WarehouseServiceImpl{
		repository: repository,
	}
}

func (ws *WarehouseServiceImpl) Create(ctx context.Context, createWarehouseDTO WarehouseReq) (json.RawMessage, error) {
	warehouse := db.ImWarehouse{}

	dublicateWarehouse, _ := ws.repository.GetWarehouseByLogisticCenterId(ctx, db.GetWarehouseByLogisticCenterIdParams{
		LogisticCenterID: createWarehouseDTO.LogisticCenterId,
		WarehouseName:    createWarehouseDTO.WarehouseName,
	})

	if dublicateWarehouse != warehouse {
		log.Println("Data already exist in this logisticCenter.")
		return nil, myErrors.NewHttpError(fiber.StatusConflict, myErrors.NewResponseByKey("duplicate_err", CONSTANTS.LANGUAGE))
	}

	warehouse, err := ws.repository.CreateWarehouse(ctx, db.CreateWarehouseParams{
		WarehouseName:    createWarehouseDTO.WarehouseName,
		LogisticCenterID: createWarehouseDTO.LogisticCenterId,
	})

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusForbidden, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	createdWarehouse, err := ws.repository.GetWarehouseById(ctx, warehouse.ID)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusForbidden, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	return createdWarehouse, nil
}

func (ws *WarehouseServiceImpl) Update(ctx context.Context, warehouseId uuid.UUID, updateWarehouseDTO WarehouseReq) (json.RawMessage, error) {
	_, err := ws.repository.GetWarehouseById(ctx, warehouseId)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	updatedWarehouse, err := ws.repository.UpdateWarehouse(ctx,
		db.UpdateWarehouseParams{
			ID:               warehouseId,
			WarehouseName:    updateWarehouseDTO.WarehouseName,
			LogisticCenterID: updateWarehouseDTO.LogisticCenterId})

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusForbidden,
			myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	warehouse, err := ws.repository.GetWarehouseById(ctx, updatedWarehouse.ID)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusForbidden, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	//dto := ToWarehouseDto(updatedWarehouse)

	return warehouse, nil
}

func (ws *WarehouseServiceImpl) Delete(ctx context.Context, warehouseId uuid.UUID) error {
	_, err := ws.repository.GetWarehouseById(ctx, warehouseId)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	err = ws.repository.DeleteWarehouse(ctx, warehouseId)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusForbidden,
			myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	return nil
}

//func (ws *WarehouseServiceImpl) FindById(ctx context.Context, warehouseId uuid.UUID) (json.RawMessage, error) {
//	warehouse, err := ws.repository.GetWarehouseById(ctx, warehouseId)
//
//	if err != nil {
//		log.Println(err)
//		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
//	}
//
//	//dto := ToWarehouseDto(warehouse)
//
//	return warehouse, nil
//}
//
//func (ws *WarehouseServiceImpl) FindAllWithPagination(ctx context.Context, page int, queryParams *GetAllWarehousesQuery) ([]json.RawMessage, *int, *int, error) {
//	count, err := ws.repository.CountOfWarehouses(ctx)
//
//	if err != nil {
//		log.Println(err)
//		return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
//	}
//
//	pagination, nextPage, totalPages := util.AddPagination(page, &util.Pagination{Limit: queryParams.Limit}, count)
//
//	if totalPages == 0 {
//		return []json.RawMessage{}, nil, nil, nil
//	}
//
//	warehouses, err := ws.repository.GetWarehousesWithPagination(ctx, db.GetWarehousesWithPaginationParams{
//		Limit:  int32(pagination.Limit),
//		Offset: int32(pagination.Offset),
//	})
//
//	if err != nil {
//		log.Println(err)
//		return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
//	}
//
//	//dtoArr := ToWarehouseDtoArray(warehouses)
//	//response := ToWarehouseResponse(dtoArr, nextPage, totalPages)
//
//	return warehouses, &nextPage, &totalPages, nil
//}
//
//func (ws *WarehouseServiceImpl) FindAll(ctx context.Context) ([]json.RawMessage, error) {
//	warehouses, err := ws.repository.GetWarehouses(ctx)
//
//	if err != nil {
//		log.Println(err)
//		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
//	}
//
//	//dto := ToWarehouseDtoArray(warehouses)
//
//	return warehouses, nil
//}
