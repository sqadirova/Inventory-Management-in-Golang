// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: inventories.sql

package db

import (
	"context"
	"encoding/json"

	"github.com/gofrs/uuid"
)

const countOfInventories = `-- name: CountOfInventories :one
SELECT count(*)
FROM im.inventories WHERE inventory_name ILIKE $1
`

func (q *Queries) CountOfInventories(ctx context.Context, inventoryName string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countOfInventories, inventoryName)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countOfInventoriesByAllFilters = `-- name: CountOfInventoriesByAllFilters :one
SELECT count(*)
FROM im.inventories i WHERE i.logistic_center_id=$1 AND i.warehouse_id=$2 AND i.location_id=$3 AND i.inventory_name ILIKE $4
`

type CountOfInventoriesByAllFiltersParams struct {
	LogisticCenterID uuid.UUID `json:"logistic_center_id"`
	WarehouseID      uuid.UUID `json:"warehouse_id"`
	LocationID       uuid.UUID `json:"location_id"`
	InventoryName    string    `json:"inventory_name"`
}

func (q *Queries) CountOfInventoriesByAllFilters(ctx context.Context, arg CountOfInventoriesByAllFiltersParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countOfInventoriesByAllFilters,
		arg.LogisticCenterID,
		arg.WarehouseID,
		arg.LocationID,
		arg.InventoryName,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countOfInventoriesByLocation = `-- name: CountOfInventoriesByLocation :one
SELECT count(*)
FROM im.inventories i WHERE i.location_id=$1 AND i.inventory_name ILIKE $2
`

type CountOfInventoriesByLocationParams struct {
	LocationID    uuid.UUID `json:"location_id"`
	InventoryName string    `json:"inventory_name"`
}

func (q *Queries) CountOfInventoriesByLocation(ctx context.Context, arg CountOfInventoriesByLocationParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countOfInventoriesByLocation, arg.LocationID, arg.InventoryName)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countOfInventoriesByLogisticCenter = `-- name: CountOfInventoriesByLogisticCenter :one
SELECT count(*)
FROM im.inventories i WHERE i.logistic_center_id=$1 AND i.inventory_name ILIKE $2
`

type CountOfInventoriesByLogisticCenterParams struct {
	LogisticCenterID uuid.UUID `json:"logistic_center_id"`
	InventoryName    string    `json:"inventory_name"`
}

func (q *Queries) CountOfInventoriesByLogisticCenter(ctx context.Context, arg CountOfInventoriesByLogisticCenterParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countOfInventoriesByLogisticCenter, arg.LogisticCenterID, arg.InventoryName)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countOfInventoriesByLogisticCenterAndLocation = `-- name: CountOfInventoriesByLogisticCenterAndLocation :one
SELECT count(*)
FROM im.inventories i WHERE i.logistic_center_id=$1 AND i.location_id=$2 AND i.inventory_name ILIKE $3
`

type CountOfInventoriesByLogisticCenterAndLocationParams struct {
	LogisticCenterID uuid.UUID `json:"logistic_center_id"`
	LocationID       uuid.UUID `json:"location_id"`
	InventoryName    string    `json:"inventory_name"`
}

func (q *Queries) CountOfInventoriesByLogisticCenterAndLocation(ctx context.Context, arg CountOfInventoriesByLogisticCenterAndLocationParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countOfInventoriesByLogisticCenterAndLocation, arg.LogisticCenterID, arg.LocationID, arg.InventoryName)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countOfInventoriesByLogisticCenterAndWarehouse = `-- name: CountOfInventoriesByLogisticCenterAndWarehouse :one
SELECT count(*)
FROM im.inventories i WHERE i.logistic_center_id=$1 AND i.warehouse_id=$2 AND i.inventory_name ILIKE $3
`

type CountOfInventoriesByLogisticCenterAndWarehouseParams struct {
	LogisticCenterID uuid.UUID `json:"logistic_center_id"`
	WarehouseID      uuid.UUID `json:"warehouse_id"`
	InventoryName    string    `json:"inventory_name"`
}

func (q *Queries) CountOfInventoriesByLogisticCenterAndWarehouse(ctx context.Context, arg CountOfInventoriesByLogisticCenterAndWarehouseParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countOfInventoriesByLogisticCenterAndWarehouse, arg.LogisticCenterID, arg.WarehouseID, arg.InventoryName)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countOfInventoriesByWarehouse = `-- name: CountOfInventoriesByWarehouse :one
SELECT count(*)
FROM im.inventories i WHERE i.warehouse_id=$1 AND i.inventory_name ILIKE $2
`

type CountOfInventoriesByWarehouseParams struct {
	WarehouseID   uuid.UUID `json:"warehouse_id"`
	InventoryName string    `json:"inventory_name"`
}

func (q *Queries) CountOfInventoriesByWarehouse(ctx context.Context, arg CountOfInventoriesByWarehouseParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countOfInventoriesByWarehouse, arg.WarehouseID, arg.InventoryName)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countOfInventoriesByWarehouseAndLocation = `-- name: CountOfInventoriesByWarehouseAndLocation :one
SELECT count(*)
FROM im.inventories i WHERE i.warehouse_id=$1 AND i.location_id=$2 AND i.inventory_name ILIKE $3
`

type CountOfInventoriesByWarehouseAndLocationParams struct {
	WarehouseID   uuid.UUID `json:"warehouse_id"`
	LocationID    uuid.UUID `json:"location_id"`
	InventoryName string    `json:"inventory_name"`
}

func (q *Queries) CountOfInventoriesByWarehouseAndLocation(ctx context.Context, arg CountOfInventoriesByWarehouseAndLocationParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countOfInventoriesByWarehouseAndLocation, arg.WarehouseID, arg.LocationID, arg.InventoryName)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createInventory = `-- name: CreateInventory :one
INSERT INTO im.inventories
(
    inventory_category_id,
    inventory_name,
    inventory_rfid,
    logistic_center_id,
    actual_qty,
    warehouse_id,
    location_id
)
VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING id, inventory_name, inventory_rfid, actual_qty, inventory_category_id, logistic_center_id, warehouse_id, location_id, created_at, updated_at
`

type CreateInventoryParams struct {
	InventoryCategoryID uuid.UUID `json:"inventory_category_id"`
	InventoryName       string    `json:"inventory_name"`
	InventoryRfid       string    `json:"inventory_rfid"`
	LogisticCenterID    uuid.UUID `json:"logistic_center_id"`
	ActualQty           string    `json:"actual_qty"`
	WarehouseID         uuid.UUID `json:"warehouse_id"`
	LocationID          uuid.UUID `json:"location_id"`
}

func (q *Queries) CreateInventory(ctx context.Context, arg CreateInventoryParams) (ImInventory, error) {
	row := q.db.QueryRowContext(ctx, createInventory,
		arg.InventoryCategoryID,
		arg.InventoryName,
		arg.InventoryRfid,
		arg.LogisticCenterID,
		arg.ActualQty,
		arg.WarehouseID,
		arg.LocationID,
	)
	var i ImInventory
	err := row.Scan(
		&i.ID,
		&i.InventoryName,
		&i.InventoryRfid,
		&i.ActualQty,
		&i.InventoryCategoryID,
		&i.LogisticCenterID,
		&i.WarehouseID,
		&i.LocationID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteInventory = `-- name: DeleteInventory :exec
DELETE FROM im.inventories
       WHERE id = $1
`

func (q *Queries) DeleteInventory(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteInventory, id)
	return err
}

const getAllInventoriesData = `-- name: GetAllInventoriesData :many
SELECT jsonb_build_object(
               'inventory_category', x.inventory_category,
               'inventories', jsonb_agg(
                       jsonb_build_object(
                               'inventory_id', a.id,
                               'inventory_name', a.inventory_name,
                               'inventory_rfid', a.inventory_rfid,
                               'actual_qty', a.actual_qty,
                               'logistic_center', y.logistic_center,
                               'warehouse',z.warehouse,
                               'location',t.location
                           )
                   )
           )
FROM      im.inventories a
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('inventory_category_id', ic.id, 'inventory_category_name', ic.inventory_category_name) AS inventory_category
    FROM   im.inventory_categories ic
    WHERE  ic.id = a.inventory_category_id ) AS x
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('logistic_center_id', lc.id, 'logistic_center_name', lc.logistic_center_name) AS logistic_center
    FROM   im.logistic_centers lc
    WHERE  lc.id = a.logistic_center_id ) AS y
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('warehouse_id', w.id, 'warehouse_name', w.warehouse_name) AS warehouse
    FROM   im.warehouses w
    WHERE  w.id = a.warehouse_id ) AS z
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('location_id', l.id, 'location_name', l.location_name) AS location
    FROM   im.locations l
    WHERE  l.id = a.location_id ) AS t
                        ON  true
WHERE a.inventory_name ILIKE $1
GROUP BY x.inventory_category
LIMIT $2
OFFSET $3
`

type GetAllInventoriesDataParams struct {
	InventoryName string `json:"inventory_name"`
	Limit         int32  `json:"limit"`
	Offset        int32  `json:"offset"`
}

func (q *Queries) GetAllInventoriesData(ctx context.Context, arg GetAllInventoriesDataParams) ([]json.RawMessage, error) {
	rows, err := q.db.QueryContext(ctx, getAllInventoriesData, arg.InventoryName, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []json.RawMessage
	for rows.Next() {
		var jsonb_build_object json.RawMessage
		if err := rows.Scan(&jsonb_build_object); err != nil {
			return nil, err
		}
		items = append(items, jsonb_build_object)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventories = `-- name: GetInventories :many
SELECT id, inventory_name, inventory_rfid, actual_qty, inventory_category_id, logistic_center_id, warehouse_id, location_id, created_at, updated_at
FROM im.inventories
ORDER BY created_at DESC
`

func (q *Queries) GetInventories(ctx context.Context) ([]ImInventory, error) {
	rows, err := q.db.QueryContext(ctx, getInventories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ImInventory
	for rows.Next() {
		var i ImInventory
		if err := rows.Scan(
			&i.ID,
			&i.InventoryName,
			&i.InventoryRfid,
			&i.ActualQty,
			&i.InventoryCategoryID,
			&i.LogisticCenterID,
			&i.WarehouseID,
			&i.LocationID,
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

const getInventoriesByAllFilters = `-- name: GetInventoriesByAllFilters :many
SELECT jsonb_build_object(
               'inventory_category', x.inventory_category,
               'inventories', jsonb_agg(
                       jsonb_build_object(
                               'inventory_id', a.id,
                               'inventory_name', a.inventory_name,
                               'inventory_rfid', a.inventory_rfid,
                               'actual_qty', a.actual_qty,
                               'logistic_center', y.logistic_center,
                               'warehouse',z.warehouse,
                               'location',t.location
                           )
                   )
           )
FROM      im.inventories a
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('inventory_category_id', ic.id, 'inventory_category_name', ic.inventory_category_name) AS inventory_category
    FROM   im.inventory_categories ic
    WHERE  ic.id = a.inventory_category_id ) AS x
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('logistic_center_id', lc.id, 'logistic_center_name', lc.logistic_center_name) AS logistic_center
    FROM   im.logistic_centers lc
    WHERE  lc.id = a.logistic_center_id ) AS y
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('warehouse_id', w.id, 'warehouse_name', w.warehouse_name) AS warehouse
    FROM   im.warehouses w
    WHERE  w.id = a.warehouse_id ) AS z
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('location_id', l.id, 'location_name', l.location_name) AS location
    FROM   im.locations l
    WHERE  l.id = a.location_id ) AS t
                        ON  true
WHERE a.inventory_name ILIKE $1 AND
      a.logistic_center_id=$2 AND
      a.warehouse_id=$3 AND
      a.location_id=$4
GROUP BY x.inventory_category
LIMIT $5
OFFSET $6
`

type GetInventoriesByAllFiltersParams struct {
	InventoryName    string    `json:"inventory_name"`
	LogisticCenterID uuid.UUID `json:"logistic_center_id"`
	WarehouseID      uuid.UUID `json:"warehouse_id"`
	LocationID       uuid.UUID `json:"location_id"`
	Limit            int32     `json:"limit"`
	Offset           int32     `json:"offset"`
}

func (q *Queries) GetInventoriesByAllFilters(ctx context.Context, arg GetInventoriesByAllFiltersParams) ([]json.RawMessage, error) {
	rows, err := q.db.QueryContext(ctx, getInventoriesByAllFilters,
		arg.InventoryName,
		arg.LogisticCenterID,
		arg.WarehouseID,
		arg.LocationID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []json.RawMessage
	for rows.Next() {
		var jsonb_build_object json.RawMessage
		if err := rows.Scan(&jsonb_build_object); err != nil {
			return nil, err
		}
		items = append(items, jsonb_build_object)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoriesByLocation = `-- name: GetInventoriesByLocation :many
SELECT jsonb_build_object(
               'inventory_category', x.inventory_category,
               'inventories', jsonb_agg(
                       jsonb_build_object(
                               'inventory_id', a.id,
                               'inventory_name', a.inventory_name,
                               'inventory_rfid', a.inventory_rfid,
                               'actual_qty', a.actual_qty,
                               'logistic_center', y.logistic_center,
                               'warehouse',z.warehouse,
                               'location',t.location
                           )
                   )
           )
FROM      im.inventories a
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('inventory_category_id', ic.id, 'inventory_category_name', ic.inventory_category_name) AS inventory_category
    FROM   im.inventory_categories ic
    WHERE  ic.id = a.inventory_category_id ) AS x
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('logistic_center_id', lc.id, 'logistic_center_name', lc.logistic_center_name) AS logistic_center
    FROM   im.logistic_centers lc
    WHERE  lc.id = a.logistic_center_id ) AS y
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('warehouse_id', w.id, 'warehouse_name', w.warehouse_name) AS warehouse
    FROM   im.warehouses w
    WHERE  w.id = a.warehouse_id ) AS z
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('location_id', l.id, 'location_name', l.location_name) AS location
    FROM   im.locations l
    WHERE  l.id = a.location_id ) AS t
                        ON  true
WHERE a.inventory_name ILIKE $1 AND
      a.location_id=$2
GROUP BY x.inventory_category
LIMIT $3
OFFSET $4
`

type GetInventoriesByLocationParams struct {
	InventoryName string    `json:"inventory_name"`
	LocationID    uuid.UUID `json:"location_id"`
	Limit         int32     `json:"limit"`
	Offset        int32     `json:"offset"`
}

func (q *Queries) GetInventoriesByLocation(ctx context.Context, arg GetInventoriesByLocationParams) ([]json.RawMessage, error) {
	rows, err := q.db.QueryContext(ctx, getInventoriesByLocation,
		arg.InventoryName,
		arg.LocationID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []json.RawMessage
	for rows.Next() {
		var jsonb_build_object json.RawMessage
		if err := rows.Scan(&jsonb_build_object); err != nil {
			return nil, err
		}
		items = append(items, jsonb_build_object)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoriesByLogisticCenter = `-- name: GetInventoriesByLogisticCenter :many
SELECT jsonb_build_object(
               'inventory_category', x.inventory_category,
               'inventories', jsonb_agg(
                       jsonb_build_object(
                               'inventory_id', a.id,
                               'inventory_name', a.inventory_name,
                               'inventory_rfid', a.inventory_rfid,
                               'actual_qty', a.actual_qty,
                               'logistic_center', y.logistic_center,
                               'warehouse',z.warehouse,
                               'location',t.location
                           )
                   )
           )
FROM      im.inventories a
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('inventory_category_id', ic.id, 'inventory_category_name', ic.inventory_category_name) AS inventory_category
    FROM   im.inventory_categories ic
    WHERE  ic.id = a.inventory_category_id ) AS x
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('logistic_center_id', lc.id, 'logistic_center_name', lc.logistic_center_name) AS logistic_center
    FROM   im.logistic_centers lc
    WHERE  lc.id = a.logistic_center_id ) AS y
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('warehouse_id', w.id, 'warehouse_name', w.warehouse_name) AS warehouse
    FROM   im.warehouses w
    WHERE  w.id = a.warehouse_id ) AS z
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('location_id', l.id, 'location_name', l.location_name) AS location
    FROM   im.locations l
    WHERE  l.id = a.location_id ) AS t
                        ON  true
WHERE a.inventory_name ILIKE $1 AND
      a.logistic_center_id=$2
GROUP BY x.inventory_category
LIMIT $3
OFFSET $4
`

type GetInventoriesByLogisticCenterParams struct {
	InventoryName    string    `json:"inventory_name"`
	LogisticCenterID uuid.UUID `json:"logistic_center_id"`
	Limit            int32     `json:"limit"`
	Offset           int32     `json:"offset"`
}

func (q *Queries) GetInventoriesByLogisticCenter(ctx context.Context, arg GetInventoriesByLogisticCenterParams) ([]json.RawMessage, error) {
	rows, err := q.db.QueryContext(ctx, getInventoriesByLogisticCenter,
		arg.InventoryName,
		arg.LogisticCenterID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []json.RawMessage
	for rows.Next() {
		var jsonb_build_object json.RawMessage
		if err := rows.Scan(&jsonb_build_object); err != nil {
			return nil, err
		}
		items = append(items, jsonb_build_object)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoriesByLogisticCenterAndLocation = `-- name: GetInventoriesByLogisticCenterAndLocation :many
SELECT jsonb_build_object(
               'inventory_category', x.inventory_category,
               'inventories', jsonb_agg(
                       jsonb_build_object(
                               'inventory_id', a.id,
                               'inventory_name', a.inventory_name,
                               'inventory_rfid', a.inventory_rfid,
                               'actual_qty', a.actual_qty,
                               'logistic_center', y.logistic_center,
                               'warehouse',z.warehouse,
                               'location',t.location
                           )
                   )
           )
FROM      im.inventories a
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('inventory_category_id', ic.id, 'inventory_category_name', ic.inventory_category_name) AS inventory_category
    FROM   im.inventory_categories ic
    WHERE  ic.id = a.inventory_category_id ) AS x
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('logistic_center_id', lc.id, 'logistic_center_name', lc.logistic_center_name) AS logistic_center
    FROM   im.logistic_centers lc
    WHERE  lc.id = a.logistic_center_id ) AS y
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('warehouse_id', w.id, 'warehouse_name', w.warehouse_name) AS warehouse
    FROM   im.warehouses w
    WHERE  w.id = a.warehouse_id ) AS z
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('location_id', l.id, 'location_name', l.location_name) AS location
    FROM   im.locations l
    WHERE  l.id = a.location_id ) AS t
                        ON  true
WHERE a.inventory_name ILIKE $1 AND
      a.logistic_center_id=$2 AND
      a.location_id=$3
GROUP BY x.inventory_category
LIMIT $4
OFFSET $5
`

type GetInventoriesByLogisticCenterAndLocationParams struct {
	InventoryName    string    `json:"inventory_name"`
	LogisticCenterID uuid.UUID `json:"logistic_center_id"`
	LocationID       uuid.UUID `json:"location_id"`
	Limit            int32     `json:"limit"`
	Offset           int32     `json:"offset"`
}

func (q *Queries) GetInventoriesByLogisticCenterAndLocation(ctx context.Context, arg GetInventoriesByLogisticCenterAndLocationParams) ([]json.RawMessage, error) {
	rows, err := q.db.QueryContext(ctx, getInventoriesByLogisticCenterAndLocation,
		arg.InventoryName,
		arg.LogisticCenterID,
		arg.LocationID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []json.RawMessage
	for rows.Next() {
		var jsonb_build_object json.RawMessage
		if err := rows.Scan(&jsonb_build_object); err != nil {
			return nil, err
		}
		items = append(items, jsonb_build_object)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoriesByLogisticCenterAndWarehouse = `-- name: GetInventoriesByLogisticCenterAndWarehouse :many
SELECT jsonb_build_object(
               'inventory_category', x.inventory_category,
               'inventories', jsonb_agg(
                       jsonb_build_object(
                               'inventory_id', a.id,
                               'inventory_name', a.inventory_name,
                               'inventory_rfid', a.inventory_rfid,
                               'actual_qty', a.actual_qty,
                               'logistic_center', y.logistic_center,
                               'warehouse',z.warehouse,
                               'location',t.location
                           )
                   )
           )
FROM      im.inventories a
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('inventory_category_id', ic.id, 'inventory_category_name', ic.inventory_category_name) AS inventory_category
    FROM   im.inventory_categories ic
    WHERE  ic.id = a.inventory_category_id ) AS x
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('logistic_center_id', lc.id, 'logistic_center_name', lc.logistic_center_name) AS logistic_center
    FROM   im.logistic_centers lc
    WHERE  lc.id = a.logistic_center_id ) AS y
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('warehouse_id', w.id, 'warehouse_name', w.warehouse_name) AS warehouse
    FROM   im.warehouses w
    WHERE  w.id = a.warehouse_id ) AS z
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('location_id', l.id, 'location_name', l.location_name) AS location
    FROM   im.locations l
    WHERE  l.id = a.location_id ) AS t
                        ON  true
WHERE a.inventory_name ILIKE $1 AND
      a.logistic_center_id=$2 AND
      a.warehouse_id=$3
GROUP BY x.inventory_category
LIMIT $4
OFFSET $5
`

type GetInventoriesByLogisticCenterAndWarehouseParams struct {
	InventoryName    string    `json:"inventory_name"`
	LogisticCenterID uuid.UUID `json:"logistic_center_id"`
	WarehouseID      uuid.UUID `json:"warehouse_id"`
	Limit            int32     `json:"limit"`
	Offset           int32     `json:"offset"`
}

func (q *Queries) GetInventoriesByLogisticCenterAndWarehouse(ctx context.Context, arg GetInventoriesByLogisticCenterAndWarehouseParams) ([]json.RawMessage, error) {
	rows, err := q.db.QueryContext(ctx, getInventoriesByLogisticCenterAndWarehouse,
		arg.InventoryName,
		arg.LogisticCenterID,
		arg.WarehouseID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []json.RawMessage
	for rows.Next() {
		var jsonb_build_object json.RawMessage
		if err := rows.Scan(&jsonb_build_object); err != nil {
			return nil, err
		}
		items = append(items, jsonb_build_object)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoriesByWarehouse = `-- name: GetInventoriesByWarehouse :many
SELECT jsonb_build_object(
               'inventory_category', x.inventory_category,
               'inventories', jsonb_agg(
                       jsonb_build_object(
                               'inventory_id', a.id,
                               'inventory_name', a.inventory_name,
                               'inventory_rfid', a.inventory_rfid,
                               'actual_qty', a.actual_qty,
                               'logistic_center', y.logistic_center,
                               'warehouse',z.warehouse,
                               'location',t.location
                           )
                   )
           )
FROM      im.inventories a
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('inventory_category_id', ic.id, 'inventory_category_name', ic.inventory_category_name) AS inventory_category
    FROM   im.inventory_categories ic
    WHERE  ic.id = a.inventory_category_id ) AS x
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('logistic_center_id', lc.id, 'logistic_center_name', lc.logistic_center_name) AS logistic_center
    FROM   im.logistic_centers lc
    WHERE  lc.id = a.logistic_center_id ) AS y
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('warehouse_id', w.id, 'warehouse_name', w.warehouse_name) AS warehouse
    FROM   im.warehouses w
    WHERE  w.id = a.warehouse_id ) AS z
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('location_id', l.id, 'location_name', l.location_name) AS location
    FROM   im.locations l
    WHERE  l.id = a.location_id ) AS t
                        ON  true
WHERE a.inventory_name ILIKE $1 AND
      a.warehouse_id=$2
GROUP BY x.inventory_category
LIMIT $3
OFFSET $4
`

type GetInventoriesByWarehouseParams struct {
	InventoryName string    `json:"inventory_name"`
	WarehouseID   uuid.UUID `json:"warehouse_id"`
	Limit         int32     `json:"limit"`
	Offset        int32     `json:"offset"`
}

func (q *Queries) GetInventoriesByWarehouse(ctx context.Context, arg GetInventoriesByWarehouseParams) ([]json.RawMessage, error) {
	rows, err := q.db.QueryContext(ctx, getInventoriesByWarehouse,
		arg.InventoryName,
		arg.WarehouseID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []json.RawMessage
	for rows.Next() {
		var jsonb_build_object json.RawMessage
		if err := rows.Scan(&jsonb_build_object); err != nil {
			return nil, err
		}
		items = append(items, jsonb_build_object)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoriesByWarehouseAndLocation = `-- name: GetInventoriesByWarehouseAndLocation :many
SELECT jsonb_build_object(
               'inventory_category', x.inventory_category,
               'inventories', jsonb_agg(
                       jsonb_build_object(
                               'inventory_id', a.id,
                               'inventory_name', a.inventory_name,
                               'inventory_rfid', a.inventory_rfid,
                               'actual_qty', a.actual_qty,
                               'logistic_center', y.logistic_center,
                               'warehouse',z.warehouse,
                               'location',t.location
                           )
                   )
           )
FROM      im.inventories a
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('inventory_category_id', ic.id, 'inventory_category_name', ic.inventory_category_name) AS inventory_category
    FROM   im.inventory_categories ic
    WHERE  ic.id = a.inventory_category_id ) AS x
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('logistic_center_id', lc.id, 'logistic_center_name', lc.logistic_center_name) AS logistic_center
    FROM   im.logistic_centers lc
    WHERE  lc.id = a.logistic_center_id ) AS y
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('warehouse_id', w.id, 'warehouse_name', w.warehouse_name) AS warehouse
    FROM   im.warehouses w
    WHERE  w.id = a.warehouse_id ) AS z
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('location_id', l.id, 'location_name', l.location_name) AS location
    FROM   im.locations l
    WHERE  l.id = a.location_id ) AS t
                        ON  true
WHERE a.inventory_name ILIKE $1 AND
      a.warehouse_id=$2 AND
      a.location_id=$3
GROUP BY x.inventory_category
LIMIT $4
OFFSET $5
`

type GetInventoriesByWarehouseAndLocationParams struct {
	InventoryName string    `json:"inventory_name"`
	WarehouseID   uuid.UUID `json:"warehouse_id"`
	LocationID    uuid.UUID `json:"location_id"`
	Limit         int32     `json:"limit"`
	Offset        int32     `json:"offset"`
}

func (q *Queries) GetInventoriesByWarehouseAndLocation(ctx context.Context, arg GetInventoriesByWarehouseAndLocationParams) ([]json.RawMessage, error) {
	rows, err := q.db.QueryContext(ctx, getInventoriesByWarehouseAndLocation,
		arg.InventoryName,
		arg.WarehouseID,
		arg.LocationID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []json.RawMessage
	for rows.Next() {
		var jsonb_build_object json.RawMessage
		if err := rows.Scan(&jsonb_build_object); err != nil {
			return nil, err
		}
		items = append(items, jsonb_build_object)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoryById = `-- name: GetInventoryById :one
SELECT id, inventory_name, inventory_rfid, actual_qty, inventory_category_id, logistic_center_id, warehouse_id, location_id, created_at, updated_at
FROM im.inventories
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetInventoryById(ctx context.Context, id uuid.UUID) (ImInventory, error) {
	row := q.db.QueryRowContext(ctx, getInventoryById, id)
	var i ImInventory
	err := row.Scan(
		&i.ID,
		&i.InventoryName,
		&i.InventoryRfid,
		&i.ActualQty,
		&i.InventoryCategoryID,
		&i.LogisticCenterID,
		&i.WarehouseID,
		&i.LocationID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOneInventoryData = `-- name: GetOneInventoryData :one
SELECT jsonb_build_object(
               'inventory_category', x.inventory_category,
               'inventory', jsonb_agg(
                       jsonb_build_object(
                               'inventory_id', a.id,
                               'inventory_name', a.inventory_name,
                               'inventory_rfid', a.inventory_rfid,
                               'actual_qty', a.actual_qty,
                               'logistic_center', y.logistic_center,
                               'warehouse',z.warehouse,
                               'location',t.location
                           )
                   )
           )
FROM      im.inventories a
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('inventory_category_id', ic.id, 'inventory_category_name', ic.inventory_category_name) AS inventory_category
    FROM   im.inventory_categories ic
    WHERE  ic.id = a.inventory_category_id ) AS x
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('logistic_center_id', lc.id, 'logistic_center_name', lc.logistic_center_name) AS logistic_center
    FROM   im.logistic_centers lc
    WHERE  lc.id = a.logistic_center_id ) AS y
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('warehouse_id', w.id, 'warehouse_name', w.warehouse_name) AS warehouse
    FROM   im.warehouses w
    WHERE  w.id = a.warehouse_id ) AS z
                        ON  true
              LEFT JOIN lateral
    (
    SELECT jsonb_build_object('location_id', l.id, 'location_name', l.location_name) AS location
    FROM   im.locations l
    WHERE  l.id = a.location_id ) AS t
                        ON  true
WHERE a.id = $1
GROUP BY x.inventory_category
`

func (q *Queries) GetOneInventoryData(ctx context.Context, id uuid.UUID) (json.RawMessage, error) {
	row := q.db.QueryRowContext(ctx, getOneInventoryData, id)
	var jsonb_build_object json.RawMessage
	err := row.Scan(&jsonb_build_object)
	return jsonb_build_object, err
}

const updateInventory = `-- name: UpdateInventory :one
UPDATE im.inventories
SET    inventory_category_id=$2,
       inventory_name=$3,
       logistic_center_id=$4,
       actual_qty=$5,
       warehouse_id=$6,
       location_id=$7
WHERE  id=$1
RETURNING id, inventory_name, inventory_rfid, actual_qty, inventory_category_id, logistic_center_id, warehouse_id, location_id, created_at, updated_at
`

type UpdateInventoryParams struct {
	ID                  uuid.UUID `json:"id"`
	InventoryCategoryID uuid.UUID `json:"inventory_category_id"`
	InventoryName       string    `json:"inventory_name"`
	LogisticCenterID    uuid.UUID `json:"logistic_center_id"`
	ActualQty           string    `json:"actual_qty"`
	WarehouseID         uuid.UUID `json:"warehouse_id"`
	LocationID          uuid.UUID `json:"location_id"`
}

func (q *Queries) UpdateInventory(ctx context.Context, arg UpdateInventoryParams) (ImInventory, error) {
	row := q.db.QueryRowContext(ctx, updateInventory,
		arg.ID,
		arg.InventoryCategoryID,
		arg.InventoryName,
		arg.LogisticCenterID,
		arg.ActualQty,
		arg.WarehouseID,
		arg.LocationID,
	)
	var i ImInventory
	err := row.Scan(
		&i.ID,
		&i.InventoryName,
		&i.InventoryRfid,
		&i.ActualQty,
		&i.InventoryCategoryID,
		&i.LogisticCenterID,
		&i.WarehouseID,
		&i.LocationID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
