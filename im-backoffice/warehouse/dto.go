package warehouse

import (
	"encoding/json"
	"github.com/gofrs/uuid"
)

type WarehouseDTO struct {
	ID               string    `json:"id"`
	WarehouseName    string    `json:"warehouse_name"`
	LogisticCenterId uuid.UUID `json:"logistic_center_id"`
}

type WarehouseReq struct {
	WarehouseName    string    `json:"warehouse_name" validate:"required"`
	LogisticCenterId uuid.UUID `json:"logistic_center_id" validate:"required"`
}

type WarehouseRes struct {
	ID              uuid.UUID         `json:"warehouse_id"`
	WarehouseName   string            `json:"warehouse_name"`
	LogisticCenters LogisticCenterDto `json:"logistic_centers"`
}

type WarehouseResponse struct {
	Warehouses []json.RawMessage `json:"warehouses"`
	NextPage   int               `json:"next_page"`
	TotalPages int               `json:"total_pages"`
}

type GetAllWarehousesQuery struct {
	Limit int64 `query:"limit" validate:"gte=10"`
	// in feature adding search fields with validation
}

type LogisticCenterDto struct {
	ID                 string `json:"logistic_center_id"`
	LogisticCenterName string `json:"logistic_center_name"`
}
