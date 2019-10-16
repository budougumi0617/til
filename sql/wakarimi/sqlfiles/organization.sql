-- -*- coding: utf-8 -*-

--
-- 会社と部署と社員
--


-- 会社
create table companies (
  id          integer       primary key
, name        varchar(255)  not null unique
);

insert into companies (id, name)
values (301, '○○会社')
;

-- 部署
create table departments (
  id          integer       primary key
, name        varchar(255)  not null unique
, company_id  integer       not null references companies(id)
);

insert into departments (id, name, company_id)
values (2001, '開発部', 301)
     , (2002, '人事部', 301)
;

-- 社員
create table employees (
  id          integer       primary key
, name        varchar(255)  not null unique
, dept_id     integer       not null references departments(id)
);

insert into employees (id, name, dept_id)
values (10001, '社員1', 2001)
     , (10002, '社員2', 2001)
     , (10003, '社員3', 2001)
     , (10004, '社員4', 2002)
     , (10005, '社員5', 2002)
     , (10006, '社員6', 2002)
;


/*
drop table employees;
drop table departments;
drop table companies;
*/
