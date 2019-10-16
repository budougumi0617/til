-- -*- coding: utf-8 -*-

--
-- レベル5の超能力者たち
--

create table level5_espers (
  id      integer        primary key
, name    varchar(255)   not null
, rank    integer        not null unique
);

insert into level5_espers (id, name, rank)
values (811, '一方通行', 1)
     , (812, '御坂美琴', 3)
     , (813, '麦野沈利', 4)
     , (814, '食蜂操祈', 5)
     , (815, '削板軍覇', 7)
;


create table numbers10 (
  num    integer        primary key
);

insert into numbers10 (num)
values (1), (2), (3), (4), (5), (6), (7), (8), (9), (10);


/*
drop table level5_espers;
drop table numbers10;
*/
