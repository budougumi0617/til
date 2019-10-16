/* -*- coding: utf-8 -*- */

--
-- 兵団メンバー
--

create table members (
  id      integer     primary key
, name    text        not null
, height  integer     not null
, gender  char(1)     not null
);

insert into members(id, name, height, gender)
values (101, 'エレン',   170, 'M')
     , (102, 'ミカサ',   170, 'F')
     , (103, 'アルミン', 163, 'M')
     , (104, 'ジャン',   175, 'M')
     , (105, 'サシャ',   168, 'F')
     , (106, 'コニー',   158, 'M');


/*
drop table members;
*/
