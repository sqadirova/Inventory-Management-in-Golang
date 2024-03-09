package logisticCenter

import (
	"encoding/json"
	"github.com/gofrs/uuid"
)

type LogisticCenterDTO struct {
	ID                 string `json:"logistic_center_id"`
	LogisticCenterName string `json:"logistic_center_name"  validate:"required"`
}

type LogisticCenterResponseDTO struct {
	ID                 string               `json:"id"`
	LogisticCenterName string               `json:"logistic_center_name"`
	Warehouses         []LGWarehouseInfoDTO `json:"warehouses"`
}

type LogisticCenterRes struct {
	ID                 string `json:"id"`
	LogisticCenterName string `json:"logistic_center_name"`
}

type LogisticCenterRequest struct {
	LogisticCenterName string `json:"logistic_center_name"  validate:"required"`
}

type AllLogisticCentersResponse struct {
	LogisticCenters []json.RawMessage `json:"logistic_centers"`
	NextPage        int               `json:"next_page"`
	TotalPages      int               `json:"total_pages"`
}

type GetAllLogisticCentersQuery struct {
	Limit int64 `query:"limit" validate:"gte=10"`
	// in feature adding search fields with validation
}

type LGWarehouseInfoDTO struct {
	ID            uuid.UUID           `json:"id"`
	WarehouseName string              `json:"warehouse_name"`
	Locations     []LGLocationInfoDTO `json:"locations"`
}

type LGLocationInfoDTO struct {
	ID           uuid.UUID `json:"id"`
	LocationName string    `json:"location_name"`
}
