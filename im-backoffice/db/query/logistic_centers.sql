-- name: CountLogisticCenters :one
SELECT count(*)
FROM   im.logistic_centers;

-- name: CreateLogisticCenter :one
INSERT INTO im.logistic_centers
(logistic_center_name)
VALUES
    ($1)
RETURNING *;

-- name: UpdateLogisticCenter :one
UPDATE im.logistic_centers
SET    logistic_center_name=$2
WHERE  id=$1
RETURNING *;

-- name: DeleteLogisticCenter :exec
DELETE FROM im.logistic_centers
WHERE  id = $1;

-- name: GetLogisticCentersWithPagination :many
WITH warehouses AS
         (
             SELECT s.logistic_center_id as logisticCenterId,
                    jsonb_build_object(
                            'warehouse_id',s.id,
                            'warehouse_name', s.warehouse_name,
                            'locations',  COALESCE(
                                            jsonb_agg(
                                            jsonb_build_object(
                                                    'location_id', c.id,
                                                    'location_name', c.location_name
                                            ))
                                            FILTER (WHERE c.id IS NOT NULL),'[]')
                    ) AS warehouses_list
             FROM im.warehouses AS s
                      FULL OUTER JOIN im.locations AS c
                                      ON c.warehouse_id = s.id
             GROUP BY s.id,s.warehouse_name
         )
SELECT jsonb_build_object(
               'logistic_center_id', logCen.id,
               'logistic_center_name', logCen.logistic_center_name,
               'warehouses', COALESCE(json_agg(w.warehouses_list) FILTER ( WHERE w.warehouses_list IS NOT NULL),'[]')
       )
FROM warehouses AS w
         FULL JOIN im.logistic_centers AS logCen
                   ON logCen.id = w.logisticCenterId
GROUP BY logCen.logistic_center_name,logCen.id
ORDER BY logCen.created_at DESC
LIMIT $1
    OFFSET $2;

-- name: GetAllLogisticCenters :many
WITH warehouses AS
         (
             SELECT s.logistic_center_id as logisticCenterId,
                    jsonb_build_object(
                            'warehouse_id',s.id,
                            'warehouse_name', s.warehouse_name,
                            'locations',  COALESCE(
                                            jsonb_agg(
                                            jsonb_build_object(
                                                    'location_id', c.id,
                                                    'location_name', c.location_name
                                            ))
                                            FILTER (WHERE c.id IS NOT NULL),'[]')
                    ) AS warehouses_list
             FROM im.warehouses AS s
                      FULL OUTER JOIN im.locations AS c
                                      ON c.warehouse_id = s.id
             GROUP BY s.id,s.warehouse_name
         )
SELECT jsonb_build_object(
               'logistic_center_id', logCen.id,
               'logistic_center_name', logCen.logistic_center_name,
               'warehouses', COALESCE(json_agg(w.warehouses_list) FILTER ( WHERE w.warehouses_list IS NOT NULL),'[]')
       )
FROM warehouses AS w
         FULL JOIN im.logistic_centers AS logCen
                   ON logCen.id = w.logisticCenterId
GROUP BY logCen.logistic_center_name,logCen.id
ORDER BY logCen.created_at DESC;


-- name: GetLogisticCenterById :one
WITH warehouses AS
         (
             SELECT s.logistic_center_id as logisticCenterId,
                    jsonb_build_object(
                            'warehouse_id',s.id,
                            'warehouse_name', s.warehouse_name,
                            'locations',  COALESCE(
                                            jsonb_agg(
                                            jsonb_build_object(
                                                    'location_id', c.id,
                                                    'location_name', c.location_name
                                            ))
                                            FILTER (WHERE c.id IS NOT NULL),'[]')
                    ) AS warehouses_list
             FROM im.warehouses AS s
                      FULL OUTER JOIN im.locations AS c
                                      ON c.warehouse_id = s.id
             GROUP BY s.id,s.warehouse_name
         )
SELECT jsonb_build_object(
               'logistic_center_id', logCen.id,
               'logistic_center_name', logCen.logistic_center_name,
               'warehouses', COALESCE(json_agg(w.warehouses_list) FILTER ( WHERE w.warehouses_list IS NOT NULL),'[]')
       )
FROM warehouses AS w
         FULL JOIN im.logistic_centers AS logCen
                   ON logCen.id = w.logisticCenterId
WHERE logCen.id=$1
GROUP BY logCen.logistic_center_name,logCen.id;


-- name: GetOneLogisticCenterData :one
SELECT * FROM im.logistic_centers lc WHERE lc.id=$1 LIMIT 1;

-- name: GetLogisticCenterByName :one
SELECT * FROM im.logistic_centers lc WHERE lc.logistic_center_name=$1 LIMIT 1;
