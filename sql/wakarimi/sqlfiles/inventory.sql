/* -*- coding: utf-8 -*- */

--
-- 商品と在庫
--


-- 商品
create table items (
  id        integer       primary key
, name      varchar(255)  not null unique
);

insert into items (id, name)
values (1001, '天気の子 パンフレット')
     , (1002, '天気の子 サントラCD')
     , (1003, '小説 天気の子')
;

-- 在庫
create table inventories (
  item_id   integer       primary key references items(id)
, count     integer       not null
);

insert into inventories (item_id, count)
values (1001, 10)
     , (1002, 10)
     , (1003, 10)
;



/*
drop table inventories;
drop table items;
*/