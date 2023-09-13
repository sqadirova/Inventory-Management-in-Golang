-- name: GetLocationById :one
select jsonb_build_object(
               'location_id', a.id,
               'location_name', a.location_name,
               'warehouse', y.warehouse)
from im.locations a
         left join lateral (
    select jsonb_build_object('warehouse_id', w.id, 'warehouse_name', w.warehouse_name) as warehouse
    from im.warehouses w where w.id = a.warehouse_id
    ) as y on true
where a.id = $1
limit 1;

-- name: GetLocationsWithPagination :many
select jsonb_build_object(
               'location_id', a.id,
               'location_name', a.location_name,
               'warehouse', y.warehouse)
from im.locations a
         left join lateral (
    select jsonb_build_object('warehouse_id', w.id, 'warehouse_name', w.warehouse_name) as warehouse
    from im.warehouses w where w.id = a.warehouse_id
    ) as y on true
order by a.created_at desc
limit $1
offset $2;

-- name: GetLocations :many
select jsonb_build_object(
               'location_id', a.id,
               'location_name', a.location_name,
               'warehouse', y.warehouse)
from im.locations a
         left join lateral (
    select jsonb_build_object('warehouse_id', w.id, 'warehouse_name', w.warehouse_name) as warehouse
    from im.warehouses w where w.id = a.warehouse_id
    ) as y on true
order by a.created_at desc;


-- name: CreateLocation :one
INSERT INTO im.locations
    (location_name,warehouse_id)
VALUES
    ($1,$2)
RETURNING *;

-- name: UpdateLocation :one
UPDATE im.locations
SET    location_name=$2,
       warehouse_id=$3
WHERE  id=$1
RETURNING *;

-- name: DeleteLocation :exec
DELETE FROM im.locations
WHERE  id = $1;

-- name: GetLocationByName :one
SELECT *
FROM im.locations
WHERE location_name = $1
LIMIT 1;

-- name: CountOfLocations :one
SELECT count(*)
FROM im.locations;

-- name: GetLocationByWarehouseId :one
SELECT * FROM im.locations l WHERE l.warehouse_id=$1 AND l.location_name ILIKE $2;

-- name: GetOneLocation :one
SELECT * FROM im.locations l WHERE l.id=$1;

