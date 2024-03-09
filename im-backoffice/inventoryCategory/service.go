package inventoryCategory

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"im-backoffice/CONSTANTS"
	db "im-backoffice/db/sqlc"
	myErrors "im-backoffice/errors"
	"im-backoffice/util"
	"log"
)

type IInventoryCategoryService interface {
	FindById(ctx context.Context, id uuid.UUID) (*InventoryCategoryDTO, error)
	FindAllWithPagination(ctx context.Context, page int, queryParams *GetAllInventoryCategoriesQuery) (*InventoryCategoryResponse, error)
	Create(ctx context.Context, createInventoryCatDTO InventoryCategoryReq) (*InventoryCategoryDTO, error)
	Update(ctx context.Context, inventoryCategoryId uuid.UUID, updateInventoryCategory InventoryCategoryReq) (*InventoryCategoryDTO, error)
	Delete(ctx context.Context, inventoryCategoryId uuid.UUID) error
	FindAll(ctx context.Context) (*[]InventoryCategoryDTO, error)
}

type InventoryCategoryServiceImpl struct {
	repository db.Repository
}

func GetNewInventoryCategoryService(repository db.Repository) *InventoryCategoryServiceImpl {
	return &InventoryCategoryServiceImpl{
		repository: repository,
	}
}

func (ics *InventoryCategoryServiceImpl) FindById(ctx context.Context, id uuid.UUID) (*InventoryCategoryDTO, error) {
	inventoryCategory, err := ics.repository.GetInventoryCategoryById(ctx, id)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	dto := ToInventoryCategoryDto(inventoryCategory)

	return &dto, nil
}

func (ics *InventoryCategoryServiceImpl) FindAllWithPagination(ctx context.Context, page int, queryParams *GetAllInventoryCategoriesQuery) (*InventoryCategoryResponse, error) {
	count, err := ics.repository.CountOfInventoryCategories(ctx)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	pagination, nextPage, totalPages := util.AddPagination(page, &util.Pagination{Limit: queryParams.Limit}, count)

	if totalPages == 0 {
		return &InventoryCategoryResponse{
				InventoryCategories: []InventoryCategoryDTO{},
				NextPage:            0,
				TotalPages:          0,
			},
			nil
	}

	inventoryCategories, err := ics.repository.GetInventoryCategoriesWithPagination(ctx, db.GetInventoryCategoriesWithPaginationParams{
		Limit:  int32(pagination.Limit),
		Offset: int32(pagination.Offset),
	})

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	dto := ToInventoryCategoryDtoArray(inventoryCategories)
	response := ToInventoryCategoryResponse(dto, nextPage, totalPages)

	return &response, nil
}

func (ics *InventoryCategoryServiceImpl) Create(ctx context.Context, createInventoryCategory InventoryCategoryReq) (*InventoryCategoryDTO, error) {
	inventoryCategory := db.ImInventoryCategory{}

	existedInventoryCategory, _ := ics.repository.GetInventoryCategoryByName(ctx, createInventoryCategory.InventoryCategoryName)

	if existedInventoryCategory != inventoryCategory {
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("already_exists", CONSTANTS.LANGUAGE))
	}

	inventoryCategory, err := ics.repository.CreateInventoryCategory(ctx, createInventoryCategory.InventoryCategoryName)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusForbidden, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	dto := ToInventoryCategoryDto(inventoryCategory)

	return &dto, nil
}

func (ics *InventoryCategoryServiceImpl) Update(ctx context.Context, inventoryCategoryId uuid.UUID,
	updateInventoryCategory InventoryCategoryReq) (*InventoryCategoryDTO, error) {
	_, err := ics.repository.GetInventoryCategoryById(ctx, inventoryCategoryId)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	inventoryCategory, err := ics.repository.UpdateInventoryCategory(ctx,
		db.UpdateInventoryCategoryParams{
			ID:                    inventoryCategoryId,
			InventoryCategoryName: updateInventoryCategory.InventoryCategoryName})

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusForbidden, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	dto := ToInventoryCategoryDto(inventoryCategory)

	return &dto, nil
}

func (ics *InventoryCategoryServiceImpl) Delete(ctx context.Context, inventoryCategoryId uuid.UUID) error {
	_, err := ics.repository.GetInventoryCategoryById(ctx, inventoryCategoryId)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	err = ics.repository.DeleteInventoryCategory(ctx, inventoryCategoryId)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusForbidden,
			myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	return nil
}

func (ics *InventoryCategoryServiceImpl) FindAll(ctx context.Context) (*[]InventoryCategoryDTO, error) {
	inventoryCategories, err := ics.repository.GetInventoryCategories(ctx)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	dto := ToInventoryCategoryDtoArray(inventoryCategories)

	return &dto, nil
}
