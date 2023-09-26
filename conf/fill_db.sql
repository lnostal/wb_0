
-- CREATE TABLE IF NOT EXISTS testtb (
--                                       user_id serial primary key,
--                                       username varchar (50) unique not null,
--                                       password varchar (50) not null,
--                                       created timestamp not null
-- );
--
-- CREATE TABLE IF NOT EXISTS ttt (
--                                       order_id serial primary key,
--                                       name varchar (50) unique not null,
--                                       created timestamp not null
-- );


--
-- нормализация для слабых духом
--

create table if not exists deliveries (
    name varchar(100)
);


create table if not exists payments (
    transaction varchar(100) unique not null primary key,
    request_id bigint,
    currency varchar(4),
    provider varchar(20),
    amount int,
    payment_dt timestamp, --??
    bank varchar(15), --??
    delivery_cost int,
    goods_total int, -- итоговая цена ???
    custom_fee int
);


create table if not exists items (
    id serial primary key,
    chrt_id int,
    track_number varchar(50),
    price int,
    rid varchar(50),
    name varchar(50),
    sale int,
    size varchar(10),
    total_price int,
    nm_id int,
    brand varchar(50),
    statu int
);

create table if not exists orders (
    order_uid varchar(100) unique not null,
    track_number varchar(50) unique not null,
    entry varchar(50) not null,
    delivery json, --- ?
    locale varchar(3),
    internal_signature varchar(50),
    customer_id varchar(50),
    delivery_service varchar(50),
    shardkey varchar(10),
    sm_id int,
    date_created date,
    "oof_hard" varchar(50)
);