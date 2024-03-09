package logisticCenter

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"im-backoffice/CONSTANTS"
	db "im-backoffice/db/sqlc"
	myErrors "im-backoffice/errors"
	"im-backoffice/util"
	"log"
)

type ILogisticCenterService interface {
	FindById(ctx context.Context, logisticCenterId uuid.UUID) (json.RawMessage, error)
	FindAllWithPagination(ctx context.Context, page int, queryParams *GetAllLogisticCentersQuery) ([]json.RawMessage, *int, *int, error)
	FindAll(ctx context.Context) ([]json.RawMessage, error)
	Create(ctx context.Context, createLogisticCenterDTO LogisticCenterRequest) (*LogisticCenterDTO, error)
	Update(ctx context.Context, logisticCenterId uuid.UUID, updateLogisticCenterDTO LogisticCenterRequest) (json.RawMessage, error)
	Delete(ctx context.Context, logisticCenterId uuid.UUID) error
}

type LogisticCenterServiceImpl struct {
	repository db.Repository
}

func GetNewLogisticCenterService(repository db.Repository) *LogisticCenterServiceImpl {
	return &LogisticCenterServiceImpl{
		repository: repository,
	}
}

func (lcs *LogisticCenterServiceImpl) Create(ctx context.Context, createLogisticCenterDTO LogisticCenterRequest) (*LogisticCenterDTO, error) {
	createdLogisticCenter, err := lcs.repository.CreateLogisticCenter(ctx, createLogisticCenterDTO.LogisticCenterName)

	var pgErr pq.PGError

	if errors.As(err, &pgErr) {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("duplicate_err", CONSTANTS.LANGUAGE))
	}

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusForbidden, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	logisticCenter := ToLogisticCenterDto(createdLogisticCenter)

	return &logisticCenter, nil
}

func (lcs *LogisticCenterServiceImpl) Update(ctx context.Context, logisticCenterId uuid.UUID, updateLogisticCenterDTO LogisticCenterRequest) (json.RawMessage, error) {
	logisticCenter, err := lcs.repository.GetOneLogisticCenterData(ctx, logisticCenterId)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	logisticCenter.LogisticCenterName = updateLogisticCenterDTO.LogisticCenterName

	updatedLogisticCenter, err := lcs.repository.UpdateLogisticCenter(ctx, db.UpdateLogisticCenterParams{
		ID:                 logisticCenter.ID,
		LogisticCenterName: logisticCenter.LogisticCenterName,
	})

	var pgErr pq.PGError

	if errors.As(err, &pgErr) {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("duplicate_err", CONSTANTS.LANGUAGE))
	}

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusForbidden, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	logisticCenterById, err := lcs.repository.GetLogisticCenterById(ctx, updatedLogisticCenter.ID)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusForbidden, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	return logisticCenterById, nil
}

func (lcs *LogisticCenterServiceImpl) Delete(ctx context.Context, logisticCenterId uuid.UUID) error {
	_, err := lcs.repository.GetOneLogisticCenterData(ctx, logisticCenterId)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	err = lcs.repository.DeleteLogisticCenter(ctx, logisticCenterId)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusForbidden, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	return nil
}

func (lcs *LogisticCenterServiceImpl) FindById(ctx context.Context, logisticCenterId uuid.UUID) (json.RawMessage, error) {
	logisticCenter, err := lcs.repository.GetLogisticCenterById(ctx, logisticCenterId)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	return logisticCenter, nil
}

func (lcs *LogisticCenterServiceImpl) FindAllWithPagination(ctx context.Context, page int, queryParams *GetAllLogisticCentersQuery) ([]json.RawMessage, *int, *int, error) {
	count, err := lcs.repository.CountLogisticCenters(ctx)

	if err != nil {
		log.Println(err)
		return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	pagination, nextPage, totalPages := util.AddPagination(page, &util.Pagination{Limit: queryParams.Limit}, count)

	if totalPages == 0 {
		return []json.RawMessage{}, &nextPage, &totalPages, nil
	}

	logisticCenters, err := lcs.repository.GetLogisticCentersWithPagination(ctx, db.GetLogisticCentersWithPaginationParams{
		Limit:  int32(pagination.Limit),
		Offset: int32(pagination.Offset),
	})

	if err != nil {
		log.Println(err)
		return nil, nil, nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	return logisticCenters, &nextPage, &totalPages, nil
}

func (lcs *LogisticCenterServiceImpl) FindAll(ctx context.Context) ([]json.RawMessage, error) {
	logisticCenters, err := lcs.repository.GetAllLogisticCenters(ctx)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	return logisticCenters, nil
}
