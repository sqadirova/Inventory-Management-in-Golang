// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: inventory_categories.sql

package db

import (
	"context"

	"github.com/gofrs/uuid"
)

const countOfInventoryCategories = `-- name: CountOfInventoryCategories :one
SELECT count(*)
FROM im.inventory_categories
`

func (q *Queries) CountOfInventoryCategories(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countOfInventoryCategories)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createInventoryCategory = `-- name: CreateInventoryCategory :one
INSERT INTO im.inventory_categories (inventory_category_name) VALUES ($1)
RETURNING id, inventory_category_name, created_at, updated_at
`

func (q *Queries) CreateInventoryCategory(ctx context.Context, inventoryCategoryName string) (ImInventoryCategory, error) {
	row := q.db.QueryRowContext(ctx, createInventoryCategory, inventoryCategoryName)
	var i ImInventoryCategory
	err := row.Scan(
		&i.ID,
		&i.InventoryCategoryName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteInventoryCategory = `-- name: DeleteInventoryCategory :exec
DELETE FROM im.inventory_categories ic WHERE ic.id = $1
`

func (q *Queries) DeleteInventoryCategory(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteInventoryCategory, id)
	return err
}

const getInventoryCategories = `-- name: GetInventoryCategories :many
SELECT id, inventory_category_name, created_at, updated_at
FROM im.inventory_categories ic
ORDER BY ic.created_at DESC
`

func (q *Queries) GetInventoryCategories(ctx context.Context) ([]ImInventoryCategory, error) {
	rows, err := q.db.QueryContext(ctx, getInventoryCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ImInventoryCategory
	for rows.Next() {
		var i ImInventoryCategory
		if err := rows.Scan(
			&i.ID,
			&i.InventoryCategoryName,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoryCategoriesWithPagination = `-- name: GetInventoryCategoriesWithPagination :many
SELECT id, inventory_category_name, created_at, updated_at
FROM im.inventory_categories ic
ORDER BY ic.created_at DESC
LIMIT $1
OFFSET $2
`

type GetInventoryCategoriesWithPaginationParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetInventoryCategoriesWithPagination(ctx context.Context, arg GetInventoryCategoriesWithPaginationParams) ([]ImInventoryCategory, error) {
	rows, err := q.db.QueryContext(ctx, getInventoryCategoriesWithPagination, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ImInventoryCategory
	for rows.Next() {
		var i ImInventoryCategory
		if err := rows.Scan(
			&i.ID,
			&i.InventoryCategoryName,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoryCategoryById = `-- name: GetInventoryCategoryById :one
SELECT id, inventory_category_name, created_at, updated_at
FROM im.inventory_categories ic
WHERE ic.id = $1
LIMIT 1
`

func (q *Queries) GetInventoryCategoryById(ctx context.Context, id uuid.UUID) (ImInventoryCategory, error) {
	row := q.db.QueryRowContext(ctx, getInventoryCategoryById, id)
	var i ImInventoryCategory
	err := row.Scan(
		&i.ID,
		&i.InventoryCategoryName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getInventoryCategoryByName = `-- name: GetInventoryCategoryByName :one
SELECT id, inventory_category_name, created_at, updated_at FROM im.inventory_categories ic
WHERE ic.inventory_category_name = $1 LIMIT 1
`

func (q *Queries) GetInventoryCategoryByName(ctx context.Context, inventoryCategoryName string) (ImInventoryCategory, error) {
	row := q.db.QueryRowContext(ctx, getInventoryCategoryByName, inventoryCategoryName)
	var i ImInventoryCategory
	err := row.Scan(
		&i.ID,
		&i.InventoryCategoryName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateInventoryCategory = `-- name: UpdateInventoryCategory :one
UPDATE im.inventory_categories
SET inventory_category_name=$2
WHERE id=$1
RETURNING id, inventory_category_name, created_at, updated_at
`

type UpdateInventoryCategoryParams struct {
	ID                    uuid.UUID `json:"id"`
	InventoryCategoryName string    `json:"inventory_category_name"`
}

func (q *Queries) UpdateInventoryCategory(ctx context.Context, arg UpdateInventoryCategoryParams) (ImInventoryCategory, error) {
	row := q.db.QueryRowContext(ctx, updateInventoryCategory, arg.ID, arg.InventoryCategoryName)
	var i ImInventoryCategory
	err := row.Scan(
		&i.ID,
		&i.InventoryCategoryName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}