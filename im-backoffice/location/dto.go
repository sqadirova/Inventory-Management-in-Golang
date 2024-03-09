package location

import (
	"github.com/gofrs/uuid"
)

type LocationDTO struct {
	ID           string       `json:"location_id"`
	LocationName string       `json:"location_name"`
	Warehouse    WarehouseDto `json:"warehouse"`
}

type LocationDTOReq struct {
	LocationName string    `json:"location_name" validate:"required"`
	WarehouseID  uuid.UUID `json:"warehouse_id" validate:"required"`
}

type LocationResponse struct {
	Locations  []LocationDTO `json:"locations"`
	NextPage   int           `json:"next_page"`
	TotalPages int           `json:"total_pages"`
}

type GetAllLocationsQuery struct {
	Limit int64 `query:"limit" validate:"gte=10"`
	// in feature adding search fields with validation
}

type WarehouseDto struct {
	ID            string `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
}
