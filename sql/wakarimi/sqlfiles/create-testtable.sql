/* -*- coding: utf-8 -*- */

--
-- 基本操作の確認に使うテスト用テーブル
--

-- テスト用テーブルを作成
create table testtable1 (
  id     integer   primary key
, name   text      not null
, age    integer
);


/*
-- テスト用データを作成
insert into testtable1(id, name, age)
values (101, 'Alice', 20)
     , (102, 'Bob'  , 25)
     , (103, 'Cathy', 22)
;
*/


/*
drop table testtable1;
*/
