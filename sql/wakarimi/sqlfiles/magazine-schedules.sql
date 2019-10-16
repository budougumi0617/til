-- -*- coding: utf-8 -*-

--
-- 週刊マンガ雑誌の発売日（2019-08-01〜2019-08-07）
--

create table magazine_schedules (
  date      date            not null
, title     varchar(255)    not null
, primary key (date, title)
);

insert into magazine_schedules(date, title)
values ('2019-07-01', '少年ジャンプ')
     , ('2019-07-01', 'ヤングマガジン')
     , ('2019-07-01', 'スピリッツ')
     , ('2019-07-03', '少年サンデー')
     , ('2019-07-03', '少年マガジン')
     , ('2019-07-04', '少年チャンピオン')
     , ('2019-07-04', 'ヤングジャンプ')
     , ('2019-07-04', 'ヤングサンデー')
     , ('2019-07-04', 'モーニング')
     , ('2019-07-05', '漫画TIMES')
;


create table calendar201907w1 (
  date       date      primary key
);

insert into calendar201907w1 (date)
values ('2019-07-01')
     , ('2019-07-02')
     , ('2019-07-03')
     , ('2019-07-04')
     , ('2019-07-05')
     , ('2019-07-06')
     , ('2019-07-07')
;


/*
drop table magazine_schedules;
drop table calendar201907w1;
*/
