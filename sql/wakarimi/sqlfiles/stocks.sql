-- -*- coding: utf-8 -*-

--
-- とある企業の株価の推移
--

create table aapl_stock_prices (
  date      date            primary key
, price     decimal(10, 2)  not null
);

insert into aapl_stock_prices (date, price)
values ('2019-08-19', 210.35)
     , ('2019-08-20', 210.36)
     , ('2019-08-21', 212.64)
     , ('2019-08-22', 212.46)
     , ('2019-08-23', 202.64)
;


/*
drop table aapl_stock_prices;
*/
