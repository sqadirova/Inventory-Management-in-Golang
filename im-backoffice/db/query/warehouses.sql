-- name: GetWarehouseById :one
SELECT jsonb_build_object(
               'warehouse_id', a.id,
               'warehouse_name', a.warehouse_name,
               'logisticCenter', y.logisticCenter)
FROM im.warehouses a
         LEFT JOIN LATERAL (
    SELECT jsonb_build_object('logistic_center_id', l.id, 'logistic_center_name', l.logistic_center_name) AS logisticCenter
    FROM im.logistic_centers l WHERE l.id = a.logistic_center_id
    ) AS y ON true
WHERE a.id = $1
LIMIT 1;

-- name: GetWarehousesWithPagination :many
SELECT jsonb_build_object(
               'warehouse_id', a.id,
               'warehouse_name', a.warehouse_name,
               'logisticCenter', y.logisticCenter)
FROM im.warehouses a
         LEFT JOIN LATERAL (
    SELECT jsonb_build_object('logistic_center_id', l.id, 'logistic_center_name', l.logistic_center_name) AS logisticCenter
    FROM im.logistic_centers l WHERE l.id = a.logistic_center_id
    ) AS y ON true
ORDER  BY a.created_at DESC
LIMIT $1
    OFFSET $2;

-- name: GetWarehouses :many
SELECT jsonb_build_object(
               'warehouse_id', a.id,
               'warehouse_name', a.warehouse_name,
               'logisticCenters', y.logisticCenter)
FROM im.warehouses a
         LEFT JOIN LATERAL (
    SELECT jsonb_build_object('logistic_center_id', l.id, 'logistic_center_name', l.logistic_center_name) AS logisticCenter
    FROM im.logistic_centers l WHERE l.id = a.logistic_center_id
    ) AS y ON true
ORDER BY a.created_at DESC;

-- name: CountOfWarehouses :one
SELECT count(*)
FROM im.warehouses;

-- name: CreateWarehouse :one
INSERT INTO im.warehouses
(warehouse_name,logistic_center_id)
VALUES
    ($1,$2)
RETURNING *;

-- name: UpdateWarehouse :one
UPDATE im.warehouses
SET    warehouse_name=$2,
       logistic_center_id=$3
WHERE  id=$1
RETURNING *;

-- name: DeleteWarehouse :exec
DELETE FROM im.warehouses
WHERE id = $1;

-- name: GetWarehouseByName :one
SELECT *
FROM im.warehouses
WHERE warehouse_name = $1
LIMIT 1;

-- name: GetWarehouseByLogisticCenterId :one
SELECT * FROM im.warehouses w WHERE w.logistic_center_id=$1 AND w.warehouse_name ILIKE $2;

-- name: GetOneWarehouse :one
SELECT * FROM im.warehouses w WHERE w.id=$1;

