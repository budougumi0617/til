/* -*- coding: utf-8 -*- */

--
-- 部品表 (Bill of material)
--


-- 製品、部品
create table products (
  id      integer        primary key
, name    varchar(255)   not null unique
);

insert into products (id, name)
values (201, '製品A')
     , (202, '製品B')
     , (301, '中間品E')
     , (302, '中間品F')
     , (401, '部品J')
     , (402, '部品K')
;

-- 部品表
create table bill_of_materials (
  parent_id   integer   references products(id)
, child_id    integer   references products(id)
, count       integer   not null
, primary key(parent_id, child_id)
);
create index bill_of_materials_child_id_idx
    on bill_of_materials(child_id);

insert into bill_of_materials(parent_id, child_id, count)
values (201, 301,  2)
     , (201, 302,  3)
     , (202, 301,  1)
     , (202, 302,  4)
     , (301, 401, 10)
     , (301, 402, 20)
     , (302, 401, 25)
     , (302, 402, 15)
;


/*
drop table bill_of_materials;
drop table products;
*/
