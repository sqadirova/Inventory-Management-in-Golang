create schema if not exists im;

create table im.logistic_centers (
                                     id uuid primary key not null default gen_random_uuid(),
                                     logistic_center_name varchar(255) not null unique,
                                     created_at timestamptz default now(),
                                     updated_at timestamptz default now()
);

create table im.warehouses (
                               id uuid primary key not null default gen_random_uuid(),
                               warehouse_name varchar(255) not null,
                               logistic_center_id uuid not null references im.logistic_centers(id),
                               created_at timestamptz default now(),
                               updated_at timestamptz default now()
);

create table im.locations (
                              id uuid primary key not null default gen_random_uuid(),
                              location_name varchar(255) not null,
                              warehouse_id uuid not null references im.warehouses(id),
                              created_at timestamptz default now(),
                              updated_at timestamptz default now()
);

create table im.inventory_categories (
                                         id uuid primary key not null default gen_random_uuid(),
                                         inventory_category_name varchar(255) not null unique,
                                         created_at timestamptz default now(),
                                         updated_at timestamptz default now()
);

create table im.inventories (
                                id uuid primary key not null default gen_random_uuid(),
                                inventory_name varchar(255) not null,
                                inventory_rfid varchar(255) not null unique,
                                actual_qty varchar(255) not null,
                                inventory_category_id uuid not null references im.inventory_categories(id),
                                logistic_center_id uuid not null references im.logistic_centers(id),
                                warehouse_id uuid not null references im.warehouses(id),
                                location_id uuid not null references im.locations(id),
                                created_at timestamptz default now(),
                                updated_at timestamptz default now(),
                                unique (logistic_center_id,warehouse_id,location_id)
);
