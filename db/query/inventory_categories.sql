-- name: GetInventoryCategoryById :one
SELECT *
FROM im.inventory_categories ic
WHERE ic.id = $1
LIMIT 1;

-- name: GetInventoryCategoriesWithPagination :many
SELECT *
FROM im.inventory_categories ic
ORDER BY ic.created_at DESC
LIMIT $1
OFFSET $2;

-- name: CreateInventoryCategory :one
INSERT INTO im.inventory_categories (inventory_category_name) VALUES ($1)
RETURNING *;

-- name: UpdateInventoryCategory :one
UPDATE im.inventory_categories
SET inventory_category_name=$2
WHERE id=$1
RETURNING *;

-- name: DeleteInventoryCategory :exec
DELETE FROM im.inventory_categories ic WHERE ic.id = $1;

-- name: GetInventoryCategoryByName :one
SELECT * FROM im.inventory_categories ic
WHERE ic.inventory_category_name = $1 LIMIT 1;

-- name: CountOfInventoryCategories :one
SELECT count(*)
FROM im.inventory_categories;

-- name: GetInventoryCategories :many
SELECT *
FROM im.inventory_categories ic
ORDER BY ic.created_at DESC;
