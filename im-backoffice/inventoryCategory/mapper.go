package inventoryCategory

import (
	"github.com/gofrs/uuid"
	db "im-backoffice/db/sqlc"
)

func ToInventoryCategoryDto(model db.ImInventoryCategory) InventoryCategoryDTO {
	return InventoryCategoryDTO{
		ID:                    model.ID.String(),
		InventoryCategoryName: model.InventoryCategoryName,
	}
}

func ToInventoryCategoryModel(dto InventoryCategoryDTO) db.ImInventoryCategory {
	return db.ImInventoryCategory{
		ID:                    uuid.FromStringOrNil(dto.ID),
		InventoryCategoryName: dto.InventoryCategoryName,
	}
}

func ToInventoryCategoryDtoArray(models []db.ImInventoryCategory) (result []InventoryCategoryDTO) {
	for _, val := range models {
		result = append(result, ToInventoryCategoryDto(val))
	}

	return
}

func ToInventoryCategoryModelArray(dtoArray []InventoryCategoryDTO) (result []db.ImInventoryCategory) {
	for _, val := range dtoArray {
		result = append(result, ToInventoryCategoryModel(val))
	}
	return
}

func ToInventoryCategoryResponse(dto []InventoryCategoryDTO, nextPage, totalPages int) InventoryCategoryResponse {
	return InventoryCategoryResponse{
		InventoryCategories: dto,
		NextPage:            nextPage,
		TotalPages:          totalPages,
	}
}

//func ToInventoryCategoryResponseArray(models []db.ImInventoryCategory,nextPage) (result []InventoryCategoryResponse) {
//	for _, val := range models {
//		result = append(resuflt, ToInventoryCategoryDto(val))
//	}
//
//	return
//}
