-- -*- coding: utf-8 -*-

--
-- バレーボールチーム
--
create table volleyball_teams (
  id          integer          primary key
, name        varchar(255)     not null unique
, short_name  varchar(10)      not null unique
);

insert into volleyball_teams (id, name, short_name)
values (51, '白鳥沢学園高校バレーボール部', '白鳥沢')
     , (52, '青葉城西高校バレーボール部'  , '青葉城西')
     , (53, '伊達工業高校バレーボール部'  , '伊達工')
     , (54, '烏野高校バレーボール部'      , '烏野')
;


/*
drop table volleyball_teams;
*/
