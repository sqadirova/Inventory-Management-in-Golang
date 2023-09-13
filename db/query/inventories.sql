-- name: GetInventoryById :one
SELECT *
FROM im.inventories
WHERE id = $1
LIMIT 1;

-- name: GetInventories :many
SELECT *
FROM im.inventories
ORDER BY created_at DESC;

-- name: CountOfInventories :one
SELECT count(*)
FROM im.inventories WHERE inventory_name ILIKE $1;

-- name: CreateInventory :one
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
RETURNING *;

-- name: UpdateInventory :one
UPDATE im.inventories
SET    inventory_category_id=$2,
       inventory_name=$3,
       logistic_center_id=$4,
       actual_qty=$5,
       warehouse_id=$6,
       location_id=$7
WHERE  id=$1
RETURNING *;

-- name: DeleteInventory :exec
DELETE FROM im.inventories
       WHERE id = $1;

-- name: GetOneInventoryData :one
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
GROUP BY x.inventory_category;

-- name: GetAllInventoriesData :many
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
OFFSET $3;

-- name: GetInventoriesByLogisticCenter :many
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
OFFSET $4;

-- name: CountOfInventoriesByLogisticCenter :one
SELECT count(*)
FROM im.inventories i WHERE i.logistic_center_id=$1 AND i.inventory_name ILIKE $2;


-- name: GetInventoriesByWarehouse :many
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
OFFSET $4;

-- name: CountOfInventoriesByWarehouse :one
SELECT count(*)
FROM im.inventories i WHERE i.warehouse_id=$1 AND i.inventory_name ILIKE $2;


-- name: GetInventoriesByLocation :many
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
OFFSET $4;

-- name: CountOfInventoriesByLocation :one
SELECT count(*)
FROM im.inventories i WHERE i.location_id=$1 AND i.inventory_name ILIKE $2;


-- name: GetInventoriesByLogisticCenterAndWarehouse :many
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
OFFSET $5;

-- name: CountOfInventoriesByLogisticCenterAndWarehouse :one
SELECT count(*)
FROM im.inventories i WHERE i.logistic_center_id=$1 AND i.warehouse_id=$2 AND i.inventory_name ILIKE $3;


-- name: GetInventoriesByLogisticCenterAndLocation :many
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
OFFSET $5;

-- name: CountOfInventoriesByLogisticCenterAndLocation :one
SELECT count(*)
FROM im.inventories i WHERE i.logistic_center_id=$1 AND i.location_id=$2 AND i.inventory_name ILIKE $3;


-- name: GetInventoriesByWarehouseAndLocation :many
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
OFFSET $5;

-- name: CountOfInventoriesByWarehouseAndLocation :one
SELECT count(*)
FROM im.inventories i WHERE i.warehouse_id=$1 AND i.location_id=$2 AND i.inventory_name ILIKE $3;


-- name: GetInventoriesByAllFilters :many
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
OFFSET $6;

-- name: CountOfInventoriesByAllFilters :one
SELECT count(*)
FROM im.inventories i WHERE i.logistic_center_id=$1 AND i.warehouse_id=$2 AND i.location_id=$3 AND i.inventory_name ILIKE $4;

