-- -*- coding: utf-8 -*-

--
-- SOS団の団員と上司
--

create table sos_brigade (
  id         integer         primary key
, name       varchar(255)    not null
, memo       varchar(255)    not null
, boss_id    integer         references sos_brigade(id)
);

insert into sos_brigade(id, name, memo, boss_id)
values (101, 'ハルヒ', '団長'    , null)
     , (102, '古泉'  , '超能力者', 101)
     , (103, 'みくる', '未来人'  , 101)
     , (104, 'キョン', '一般人'  , 102)
     , (105, '有希'  , '宇宙人'  , 101)
;

/*
drop table sos_brigade;
*/
