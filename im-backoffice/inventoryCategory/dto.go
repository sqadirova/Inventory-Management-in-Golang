package inventoryCategory

type InventoryCategoryDTO struct {
	ID                    string `json:"id"`
	InventoryCategoryName string `json:"inventory_category_name"`
}

type InventoryCategoryReq struct {
	InventoryCategoryName string `json:"inventory_category_name"  validate:"required"`
}

type InventoryCategoryResponse struct {
	InventoryCategories []InventoryCategoryDTO `json:"inventory_categories"`
	NextPage            int                    `json:"next_page"`
	TotalPages          int                    `json:"total_pages"`
}

type GetAllInventoryCategoriesQuery struct {
	Limit int64 `query:"limit" validate:"gte=10"`
	// in feature adding search fields with validation
}
