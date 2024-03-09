package inventory

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"im-backoffice/inventoryCategory"
	"im-backoffice/logisticCenter"
)

type InventoryDTO struct {
	ID                  uuid.UUID `json:"id"`
	InventoryCategoryID uuid.UUID `json:"inventory_category_id"  validate:"required"`
	InventoryName       string    `json:"inventory_name"  validate:"required"`
	InventoryRFID       string    `json:"inventory_rfid" validate:"required"`
	LogisticCenterID    uuid.UUID `json:"logistic_center_id" validate:"required"`
	ActualQTY           string    `json:"actual_qty"`
	WarehouseID         uuid.UUID `json:"warehouse_id" validate:"required"`
	LocationID          uuid.UUID `json:"location_id" validate:"required"`
}

type CreateInventoryRequest struct {
	InventoryCategoryID uuid.UUID `json:"inventory_category_id"  validate:"required"`
	InventoryName       string    `json:"inventory_name"  validate:"required"`
	InventoryRFID       string    `json:"inventory_rfid" validate:"required"`
	LogisticCenterID    uuid.UUID `json:"logistic_center_id" validate:"required"`
	ActualQTY           string    `json:"actual_qty" validate:"required"`
	WarehouseID         uuid.UUID `json:"warehouse_id" validate:"required"`
	LocationID          uuid.UUID `json:"location_id" validate:"required"`
}

type UpdateInventoryRequest struct {
	InventoryCategoryID uuid.UUID `json:"inventory_category_id" validate:"required"`
	InventoryName       string    `json:"inventory_name" validate:"required"`
	LogisticCenterID    uuid.UUID `json:"logistic_center_id" validate:"required"`
	ActualQTY           string    `json:"actual_qty" validate:"required"`
	WarehouseID         uuid.UUID `json:"warehouse_id" validate:"required"`
	LocationID          uuid.UUID `json:"location_id" validate:"required"`
}

type InventoryResponse struct {
	InventoryCategory inventoryCategory.InventoryCategoryDTO `json:"inventory_category"`
	Inventories       []InventoryInfoDto                     `json:"inventories"`
}

type GetInventoryDTO struct {
}

type GetAllInventoriesQuery struct {
	InventoryName    string    `query:"inventory_name" validate:"min=3"`
	LogisticCenterID uuid.UUID `query:"logistic_center_id"`
	WarehouseID      uuid.UUID `query:"warehouse_id"`
	LocationID       uuid.UUID `query:"location_id"`
	Limit            int64     `query:"limit" validate:"gte=10,required"`
}

type AllInventoryResponseDTO struct {
	Inventories []json.RawMessage `json:"inventories"`
	NextPage    int               `json:"next_page"`
	TotalPages  int               `json:"total_pages"`
}

type WarehouseDto struct {
	WarehouseID   uuid.UUID `json:"warehouse_id"`
	WarehouseName string    `json:"warehouse_name"`
}

type LocationDto struct {
	LocationID   string `json:"location_id"`
	LocationName string `json:"location_name"`
}

type InventoryInfoDto struct {
	InventoryID    string                           `json:"inventory_id"`
	ActualQTY      string                           `json:"actual_qty"`
	InventoryName  string                           `json:"inventory_name"`
	InventoryRFID  string                           `json:"inventory_rfid"`
	LogisticCenter logisticCenter.LogisticCenterRes `json:"logistic_center"`
	Warehouse      WarehouseDto                     `json:"warehouse"`
	Location       LocationDto                      `json:"location"`
}
