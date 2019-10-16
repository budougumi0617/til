/* -*- coding: utf-8 -*- */

--
-- 生徒と成績
--

-- 生徒
create table students (
  id        integer    primary key
, name      text       not null
, gender    char(1)    not null
, class     text       not null
);

insert into students (id, name, gender, class)
values (201, 'さくら ももこ'  , 'F', '3-4')
     , (202, 'はなわ かずひこ', 'M', '3-4')
     , (203, 'ほなみ たまえ'  , 'F', '3-4')
     , (204, 'まるお すえお'  , 'M', '3-4')
;

-- テストの点数
create table test_scores (
  student_id  integer  not null references students(id)
, subject     text     not null  -- 教科（国語、算数、理科、社会）
, score       integer  not null  -- 点数
, primary key (student_id, subject)
);

insert into test_scores (student_id, subject, score)
values -- まるこ
       (201, '国語',  60)
     , (201, '算数',  40)
     , (201, '理科',  40)
     , (201, '社会',  50)
       -- はなわくん
     , (202, '国語',  60)
     , (202, '算数',  70)
     , (202, '理科',  50)
     , (202, '社会',  70)
       -- たまちゃん
     , (203, '国語',  80)
     , (203, '算数',  80)
     , (203, '理科',  70)
     , (203, '社会', 100)
       -- まるおくん
     , (204, '国語',  80)
     , (204, '算数',  90)
     , (204, '理科', 100)
     , (204, '社会', 100)
;

/*
drop table test_scores;
drop table students;
*/
