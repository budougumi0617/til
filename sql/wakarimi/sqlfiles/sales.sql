/* -*- coding: utf-8 -*- */

--
-- 商品と価格
--

-- 商品
create table sales_items (
  id           serial        primary key
, name         varchar(255)  not null unique
);

-- 価格
create table sales_prices (
  item_id      integer    not null references sales_items(id)
, start_date   date       not null
, price        integer    not null
, primary key (item_id, start_date)
);


--
-- ダミーデータを作成
--

\set n_products    20000
\set n_months      100
\set ratio         0.1

-- ダミーの商品
insert into sales_items (name)
select 'sales item #'||t.i
from generate_series(1, :n_products) as t(i);

-- 価格テーブルにすべての商品をダミー価格で登録
insert into sales_prices (item_id, start_date, price)
select t.i, '2000-01-01', random() * 1000 + 1000
from generate_series(1, :n_products) as t(i);

-- 複数のダミー価格を登録
with months(date) as (
    select ('2000-01-01'::date + (t.i||'month')::interval)::date
    from generate_series(1, :n_months - 1) as t(i)
)
insert into sales_prices(item_id, start_date, price)
select p.id, m.date, random() * random() * 1000 + 1000
from sales_items p cross join months m
where random() < 0.1;


/*
drop table sales_prices;
drop table sales_items;
*/
