create table if not exists orders (
    order_uid varchar(100) unique not null,
    track_number varchar(50) unique not null,
    entry varchar(50) not null,
    delivery json,
    payment json,
    items json,
    locale varchar(3),
    internal_signature varchar(50),
    customer_id varchar(50),
    delivery_service varchar(50),
    shardkey varchar(10),
    sm_id int,
    date_created date,
    "oof_hard" varchar(50)
);