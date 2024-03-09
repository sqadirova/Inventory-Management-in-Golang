package warehouse

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	db "im-backoffice/db/sqlc"
)

func ToWarehouseDto(model db.ImWarehouse) WarehouseDTO {
	return WarehouseDTO{
		ID:               model.ID.String(),
		WarehouseName:    model.WarehouseName,
		LogisticCenterId: model.LogisticCenterID,
	}
}

func ToWarehouseResponse(dto []json.RawMessage, nextPage, totalPages int) WarehouseResponse {
	return WarehouseResponse{
		Warehouses: dto,
		NextPage:   nextPage,
		TotalPages: totalPages,
	}
}

func ToWarehouseModel(dto WarehouseDTO) db.ImWarehouse {
	return db.ImWarehouse{
		ID:               uuid.FromStringOrNil(dto.ID),
		WarehouseName:    dto.WarehouseName,
		LogisticCenterID: dto.LogisticCenterId,
	}
}

func ToWarehouseDtoArray(models []db.ImWarehouse) (result []WarehouseDTO) {
	for _, val := range models {
		result = append(result, ToWarehouseDto(val))
	}

	return
}

func ToWarehouseModelArray(dtoArray []WarehouseDTO) (result []db.ImWarehouse) {
	for _, val := range dtoArray {
		result = append(result, ToWarehouseModel(val))
	}
	return
}
