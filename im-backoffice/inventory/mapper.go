package inventory

import (
	"encoding/json"
	db "im-backoffice/db/sqlc"
)

func ToInventoryDTO(model db.ImInventory) InventoryDTO {
	return InventoryDTO{
		ID:                  model.ID,
		InventoryCategoryID: model.InventoryCategoryID,
		InventoryName:       model.InventoryName,
		InventoryRFID:       model.InventoryRfid,
		LogisticCenterID:    model.LogisticCenterID,
		ActualQTY:           model.ActualQty,
		WarehouseID:         model.WarehouseID,
		LocationID:          model.LocationID,
	}
}

func ToInventoryResponseArray(models []db.ImInventory) (result []InventoryDTO) {
	for _, val := range models {
		result = append(result, ToInventoryDTO(val))
	}
	return
}

func ToInventoryModel(dto InventoryDTO) db.ImInventory {
	return db.ImInventory{
		ID:                  dto.ID,
		InventoryName:       dto.InventoryName,
		InventoryRfid:       dto.InventoryRFID,
		ActualQty:           dto.ActualQTY,
		InventoryCategoryID: dto.InventoryCategoryID,
		LogisticCenterID:    dto.LogisticCenterID,
		WarehouseID:         dto.WarehouseID,
		LocationID:          dto.LocationID,
	}
}

func ToInventoryModelArray(dtoArray []InventoryDTO) (result []db.ImInventory) {
	for _, val := range dtoArray {
		result = append(result, ToInventoryModel(val))
	}
	return
}

func ToAllInventoryResponseDTO(dto []json.RawMessage, nextPage, totalPages int) AllInventoryResponseDTO {
	return AllInventoryResponseDTO{
		Inventories: dto,
		NextPage:    nextPage,
		TotalPages:  totalPages,
	}
}
