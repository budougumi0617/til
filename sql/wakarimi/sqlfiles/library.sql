-- -*- coding: utf-8 -*-

--
-- 本と著者と執筆のテーブル
--


-- 本
create table books (
  id      integer        primary key
, title   varchar(255)   not null
);

-- 著者
create table authors (
  id      integer        primary key
, name    varchar(255)   not null
);

-- 執筆
create table writings (
  book_id    integer       references books(id)
, author_id  integer       references authors(id)
, role       varchar(255)
, primary key (book_id, author_id)
);
create index writings_author_id_idx on writings(author_id);

-- データを挿入
insert into books (id, title)
values (31, 'デスノート')
     , (32, 'バクマン。')
     , (33, 'ヒカルの碁')
;
insert into authors(id, name)
values (71, '大場つぐみ')
     , (72, '小畑健')
     , (73, 'ほったゆみ')
;
insert into writings(book_id, author_id, role)
values (31, 71, '原作')
     , (31, 72, '漫画')
     , (32, 71, '原作')
     , (32, 72, '漫画')
     , (33, 73, '原作')
     , (33, 72, '漫画')
;



/*
drop table writings;
drop table authors;
drop table books;
*/
