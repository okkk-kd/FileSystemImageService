CREATE SCHEMA IF NOT EXISTS store

create table store.images
(
    id         serial
        constraint images_pk
            primary key,
    name       text,
    client_id  integer,
    created_at timestamp,
    updated_at timestamp
);

alter table store.images
    owner to postgres;

create unique index images_id_uindex
    on store.images (id);

