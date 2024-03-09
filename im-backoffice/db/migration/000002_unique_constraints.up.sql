ALTER TABLE im.warehouses
    ADD CONSTRAINT warehouse_unique UNIQUE (warehouse_name, logistic_center_id);

ALTER TABLE im.locations
    ADD CONSTRAINT location_unique UNIQUE (location_name, warehouse_id);