package inventory

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"im-backoffice/CONSTANTS"
	db "im-backoffice/db/sqlc"
	myErrors "im-backoffice/errors"
	"im-backoffice/util"
	"log"
)

type IInventoryService interface {
	FindById(ctx context.Context, inventoryId uuid.UUID) (json.RawMessage, error)
	FindAll(ctx context.Context, page int, queryParams *GetAllInventoriesQuery) ([]json.RawMessage, *int, *int, error)
	Create(ctx context.Context, inventoryReqDto CreateInventoryRequest) (json.RawMessage, error)
	Update(ctx context.Context, inventoryId uuid.UUID, inventoryReqDto UpdateInventoryRequest) (json.RawMessage, error)
	Delete(ctx context.Context, inventoryId uuid.UUID) error
	CheckInfoForInventory(ctx context.Context, inventoryCategoryId uuid.UUID, logisticCenterId uuid.UUID, warehouseId uuid.UUID, locationId uuid.UUID) (bool, error)
}

type InventoryServiceImpl struct {
	repository db.Repository
}

func GetNewInventoryService(repository db.Repository) *InventoryServiceImpl {
	return &InventoryServiceImpl{
		repository: repository,
	}
}

func (is *InventoryServiceImpl) FindById(ctx context.Context, inventoryId uuid.UUID) (json.RawMessage, error) {
	inventory, err := is.repository.GetOneInventoryData(ctx, inventoryId)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	return inventory, nil
}

func (is *InventoryServiceImpl) FindAll(ctx context.Context, page int, queryParams *GetAllInventoriesQuery) ([]json.RawMessage, *int, *int, error) {

	if queryParams.LogisticCenterID == uuid.Nil && queryParams.WarehouseID == uuid.Nil && queryParams.LocationID == uuid.Nil {
		count, err := is.repository.CountOfInventories(ctx, fmt.Sprintf("%%%v%%", queryParams.InventoryName))

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		pagination, nextPage, countOfPages := util.AddPagination(
			page,
			&util.Pagination{Limit: queryParams.Limit},
			count)

		if countOfPages == 0 {
			return []json.RawMessage{}, &nextPage, &countOfPages, nil
		}

		inventories, err := is.repository.GetAllInventoriesData(ctx, db.GetAllInventoriesDataParams{
			Limit:         int32(pagination.Limit),
			Offset:        int32(pagination.Offset),
			InventoryName: fmt.Sprintf("%%%v%%", queryParams.InventoryName),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		return inventories, &nextPage, &countOfPages, nil
	}

	if queryParams.LogisticCenterID != uuid.Nil && queryParams.WarehouseID == uuid.Nil && queryParams.LocationID == uuid.Nil {
		count, err := is.repository.CountOfInventoriesByLogisticCenter(ctx, db.CountOfInventoriesByLogisticCenterParams{
			LogisticCenterID: queryParams.LogisticCenterID,
			InventoryName:    fmt.Sprintf("%%%v%%", queryParams.InventoryName),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		pagination, nextPage, countOfPages := util.AddPagination(
			page,
			&util.Pagination{Limit: queryParams.Limit},
			count)

		if countOfPages == 0 {
			return []json.RawMessage{}, &nextPage, &countOfPages, nil
		}

		inventories, err := is.repository.GetInventoriesByLogisticCenter(ctx, db.GetInventoriesByLogisticCenterParams{
			LogisticCenterID: queryParams.LogisticCenterID,
			InventoryName:    fmt.Sprintf("%%%v%%", queryParams.InventoryName),
			Limit:            int32(pagination.Limit),
			Offset:           int32(pagination.Offset),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		return inventories, &nextPage, &countOfPages, nil
	}

	if queryParams.LogisticCenterID == uuid.Nil && queryParams.WarehouseID != uuid.Nil && queryParams.LocationID == uuid.Nil {
		count, err := is.repository.CountOfInventoriesByWarehouse(ctx, db.CountOfInventoriesByWarehouseParams{
			WarehouseID:   queryParams.WarehouseID,
			InventoryName: fmt.Sprintf("%%%v%%", queryParams.InventoryName),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		pagination, nextPage, countOfPages := util.AddPagination(
			page,
			&util.Pagination{Limit: queryParams.Limit},
			count)

		if countOfPages == 0 {
			return []json.RawMessage{}, &nextPage, &countOfPages, nil
		}

		inventories, err := is.repository.GetInventoriesByWarehouse(ctx, db.GetInventoriesByWarehouseParams{
			WarehouseID:   queryParams.WarehouseID,
			InventoryName: fmt.Sprintf("%%%v%%", queryParams.InventoryName),
			Limit:         int32(pagination.Limit),
			Offset:        int32(pagination.Offset),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		return inventories, &nextPage, &countOfPages, nil
	}

	if queryParams.LogisticCenterID == uuid.Nil && queryParams.WarehouseID == uuid.Nil && queryParams.LocationID != uuid.Nil {
		count, err := is.repository.CountOfInventoriesByLocation(ctx, db.CountOfInventoriesByLocationParams{
			LocationID:    queryParams.LocationID,
			InventoryName: fmt.Sprintf("%%%v%%", queryParams.InventoryName),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		pagination, nextPage, countOfPages := util.AddPagination(
			page,
			&util.Pagination{Limit: queryParams.Limit},
			count)

		if countOfPages == 0 {
			return []json.RawMessage{}, &nextPage, &countOfPages, nil
		}

		inventories, err := is.repository.GetInventoriesByLocation(ctx, db.GetInventoriesByLocationParams{
			LocationID:    queryParams.LocationID,
			InventoryName: fmt.Sprintf("%%%v%%", queryParams.InventoryName),
			Limit:         int32(pagination.Limit),
			Offset:        int32(pagination.Offset),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		return inventories, &nextPage, &countOfPages, nil
	}

	if queryParams.LogisticCenterID != uuid.Nil && queryParams.WarehouseID != uuid.Nil && queryParams.LocationID == uuid.Nil {
		count, err := is.repository.CountOfInventoriesByLogisticCenterAndWarehouse(ctx, db.CountOfInventoriesByLogisticCenterAndWarehouseParams{
			LogisticCenterID: queryParams.LogisticCenterID,
			WarehouseID:      queryParams.WarehouseID,
			InventoryName:    fmt.Sprintf("%%%v%%", queryParams.InventoryName),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		pagination, nextPage, countOfPages := util.AddPagination(
			page,
			&util.Pagination{Limit: queryParams.Limit},
			count)

		if countOfPages == 0 {
			return []json.RawMessage{}, &nextPage, &countOfPages, nil
		}

		inventories, err := is.repository.GetInventoriesByLogisticCenterAndWarehouse(ctx, db.GetInventoriesByLogisticCenterAndWarehouseParams{
			InventoryName:    fmt.Sprintf("%%%v%%", queryParams.InventoryName),
			LogisticCenterID: queryParams.LogisticCenterID,
			WarehouseID:      queryParams.WarehouseID,
			Limit:            int32(pagination.Limit),
			Offset:           int32(pagination.Offset),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		return inventories, &nextPage, &countOfPages, nil
	}

	if queryParams.LogisticCenterID != uuid.Nil && queryParams.WarehouseID == uuid.Nil && queryParams.LocationID != uuid.Nil {
		count, err := is.repository.CountOfInventoriesByLogisticCenterAndLocation(ctx, db.CountOfInventoriesByLogisticCenterAndLocationParams{
			LogisticCenterID: queryParams.LogisticCenterID,
			LocationID:       queryParams.LocationID,
			InventoryName:    fmt.Sprintf("%%%v%%", queryParams.InventoryName),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		pagination, nextPage, countOfPages := util.AddPagination(
			page,
			&util.Pagination{Limit: queryParams.Limit},
			count)

		if countOfPages == 0 {
			return []json.RawMessage{}, &nextPage, &countOfPages, nil
		}

		inventories, err := is.repository.GetInventoriesByLogisticCenterAndLocation(ctx, db.GetInventoriesByLogisticCenterAndLocationParams{
			InventoryName:    fmt.Sprintf("%%%v%%", queryParams.InventoryName),
			LogisticCenterID: queryParams.LogisticCenterID,
			LocationID:       queryParams.LocationID,
			Limit:            int32(pagination.Limit),
			Offset:           int32(pagination.Offset),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		return inventories, &nextPage, &countOfPages, nil
	}

	if queryParams.LogisticCenterID == uuid.Nil && queryParams.WarehouseID != uuid.Nil && queryParams.LocationID != uuid.Nil {
		count, err := is.repository.CountOfInventoriesByWarehouseAndLocation(ctx, db.CountOfInventoriesByWarehouseAndLocationParams{
			WarehouseID:   queryParams.WarehouseID,
			LocationID:    queryParams.LocationID,
			InventoryName: fmt.Sprintf("%%%v%%", queryParams.InventoryName),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		pagination, nextPage, countOfPages := util.AddPagination(
			page,
			&util.Pagination{Limit: queryParams.Limit},
			count)

		if countOfPages == 0 {
			return []json.RawMessage{}, &nextPage, &countOfPages, nil
		}

		inventories, err := is.repository.GetInventoriesByWarehouseAndLocation(ctx, db.GetInventoriesByWarehouseAndLocationParams{
			InventoryName: fmt.Sprintf("%%%v%%", queryParams.InventoryName),
			WarehouseID:   queryParams.WarehouseID,
			LocationID:    queryParams.LocationID,
			Limit:         int32(pagination.Limit),
			Offset:        int32(pagination.Offset),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		return inventories, &nextPage, &countOfPages, nil
	}

	if queryParams.LogisticCenterID != uuid.Nil && queryParams.WarehouseID != uuid.Nil && queryParams.LocationID != uuid.Nil {
		count, err := is.repository.CountOfInventoriesByAllFilters(ctx, db.CountOfInventoriesByAllFiltersParams{
			LogisticCenterID: queryParams.LogisticCenterID,
			WarehouseID:      queryParams.WarehouseID,
			LocationID:       queryParams.LocationID,
			InventoryName:    fmt.Sprintf("%%%v%%", queryParams.InventoryName),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		pagination, nextPage, countOfPages := util.AddPagination(
			page,
			&util.Pagination{Limit: queryParams.Limit},
			count)

		if countOfPages == 0 {
			return []json.RawMessage{}, &nextPage, &countOfPages, nil
		}

		inventories, err := is.repository.GetInventoriesByAllFilters(ctx, db.GetInventoriesByAllFiltersParams{
			InventoryName:    fmt.Sprintf("%%%v%%", queryParams.InventoryName),
			LogisticCenterID: queryParams.LogisticCenterID,
			WarehouseID:      queryParams.WarehouseID,
			LocationID:       queryParams.LocationID,
			Limit:            int32(pagination.Limit),
			Offset:           int32(pagination.Offset),
		})

		if err != nil {
			log.Println(err)
			return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		return inventories, &nextPage, &countOfPages, nil
	}

	return nil, nil, nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
}

func (is *InventoryServiceImpl) Create(ctx context.Context, inventoryReqDto CreateInventoryRequest) (json.RawMessage, error) {
	isExisted, err := is.CheckInfoForInventory(ctx, inventoryReqDto.InventoryCategoryID, inventoryReqDto.LogisticCenterID, inventoryReqDto.WarehouseID, inventoryReqDto.LocationID)

	if err != nil || !isExisted {
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey(err.Error(), CONSTANTS.LANGUAGE))
	}

	createdInventory, err := is.repository.CreateInventory(ctx, db.CreateInventoryParams{
		InventoryCategoryID: inventoryReqDto.InventoryCategoryID,
		InventoryName:       inventoryReqDto.InventoryName,
		InventoryRfid:       inventoryReqDto.InventoryRFID,
		LogisticCenterID:    inventoryReqDto.LogisticCenterID,
		ActualQty:           inventoryReqDto.ActualQTY,
		WarehouseID:         inventoryReqDto.WarehouseID,
		LocationID:          inventoryReqDto.LocationID,
	})

	var pgErr pq.PGError

	if errors.As(err, &pgErr) {
		log.Println(err)
		log.Println("In this location of warehouse already have inventory.")
		return nil, myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("duplicate_err", CONSTANTS.LANGUAGE))
	}

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	inventoryById, err := is.repository.GetOneInventoryData(ctx, createdInventory.ID)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	return inventoryById, nil
}

func (is *InventoryServiceImpl) Update(ctx context.Context, inventoryId uuid.UUID, inventoryReqDto UpdateInventoryRequest) (json.RawMessage, error) {
	_, err := is.repository.GetInventoryById(ctx, inventoryId)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	isExisted, err := is.CheckInfoForInventory(ctx, inventoryReqDto.InventoryCategoryID, inventoryReqDto.LogisticCenterID, inventoryReqDto.WarehouseID, inventoryReqDto.LocationID)

	if err != nil || !isExisted {
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey(err.Error(), CONSTANTS.LANGUAGE))
	}

	updatedInventory, err := is.repository.UpdateInventory(ctx, db.UpdateInventoryParams{
		ID:                  inventoryId,
		InventoryCategoryID: inventoryReqDto.InventoryCategoryID,
		InventoryName:       inventoryReqDto.InventoryName,
		LogisticCenterID:    inventoryReqDto.LogisticCenterID,
		ActualQty:           inventoryReqDto.ActualQTY,
		WarehouseID:         inventoryReqDto.WarehouseID,
		LocationID:          inventoryReqDto.LocationID,
	})

	var pgErr pq.PGError

	if errors.As(err, &pgErr) {
		log.Println(err)
		log.Println("In this location of warehouse already have inventory.")
		return nil, myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("duplicate_err", CONSTANTS.LANGUAGE))
	}

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	inventoryById, err := is.repository.GetOneInventoryData(ctx, updatedInventory.ID)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	return inventoryById, nil
}

func (is *InventoryServiceImpl) Delete(ctx context.Context, inventoryId uuid.UUID) error {
	_, err := is.repository.GetInventoryById(ctx, inventoryId)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	err = is.repository.DeleteInventory(ctx, inventoryId)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusForbidden, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	return nil
}

func (is *InventoryServiceImpl) CheckInfoForInventory(ctx context.Context, inventoryCategoryId uuid.UUID,
	logisticCenterId uuid.UUID, warehouseId uuid.UUID, locationId uuid.UUID) (bool, error) {
	_, err := is.repository.GetInventoryCategoryById(ctx, inventoryCategoryId)

	if err != nil {
		log.Println(err)
		return false, errors.New("not_found_inventory_category")
	}

	_, err = is.repository.GetLogisticCenterById(ctx, logisticCenterId)

	if err != nil {
		log.Println(err)
		return false, errors.New("not_found_logistic_center")
	}

	_, err = is.repository.GetOneWarehouse(ctx, warehouseId)

	if err != nil {
		log.Println(err)
		return false, errors.New("not_found_warehouse")
	}

	_, err = is.repository.GetOneLocation(ctx, locationId)

	if err != nil {
		log.Println(err)
		return false, errors.New("not_found_location")
	}

	return true, nil
}
