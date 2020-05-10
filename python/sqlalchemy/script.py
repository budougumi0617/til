#!/usr/bin/env python
# -*- coding: utf-8 -*-
import os
import time

from sqlalchemy.sql.elements import literal_column, and_
from sqlalchemy.sql import select, table
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import scoped_session, sessionmaker
from sqlalchemy.engine import create_engine
from sqlalchemy.orm.session import sessionmaker

Base = declarative_base()


def get_session(engine):
    session_maker = scoped_session(sessionmaker())
    session_maker.configure(bind=engine)
    return session_maker()


def get_actors(engine):
    q = select(
        [literal_column('actor_id')]
    ) \
        .select_from(table('actor')) \
        .where(
        and_(
            literal_column('first_name') == 'JOHNNY',
            literal_column('last_name') == 'LOLLOBRIGIDA',
        )
    )
    session = get_session(engine=engine)
    try:
        rsp = session.execute(q)
        res = rsp.first()
    finally:
        print('do finally')
        session.close()

    return res


def main():
    engine = create_engine(
        "mysql+mysqldb://{user}:{password}@{host}:{port}/{database}".format(
            user=os.environ["MYSQL_DB_USER"],
            database=os.environ["MYSQL_DB_DATABASE"],
            host=os.environ["MYSQL_DB_HOST"],
            port=os.environ["MYSQL_DB_PORT"],
            password=os.environ["MYSQL_DB_PASSWORD"],
            pool_pre_ping=True
        ),
    )

    actors = get_actors(engine)

    print('result = ', actors)

    print('--------------start sleep--------------')
    # SQLAlchemy が MySQL にコネクションを張ってから実際に SQL が発行されるまでに 7 秒間待つ
    # MySQL の wait_timeout 値を 5 (sec) にしているので、その間にタイムアウトが起きるはず。
    time.sleep(10)
    print('--------------end sleep--------------')

    # 一度タイムアウトが起きそうな状態で
    actors = get_actors(engine)


if __name__ == '__main__':
    main()


"""
$ docker container run --rm -d -e MYSQL_ROOT_PASSWORD=mysql -p 43306:3306 --name mysql budougumi0617/mysql-sakila:5.7
Unable to find image 'budougumi0617/mysql-sakila:5.7' locally
5.7: Pulling from budougumi0617/mysql-sakila
f7e2b70d04ae: Pull complete
df7f6307ff0a: Pull complete
e29ed02b1013: Pull complete
9cb929db392c: Pull complete
42cc77b24286: Pull complete
a6d57750cc73: Pull complete
79510826e343: Pull complete
07e462ad61e2: Pull complete
fa594cb5b94d: Pull complete
1b44278270ad: Pull complete
3edb3c323f55: Pull complete
5d16f5939e7a: Pull complete
a54c0238c354: Pull complete
deb36c105433: Pull complete
Digest: sha256:1d376804406b3f335905fc1f436808ae048a4aff57a0041de95fb686ba5e0fb2
Status: Downloaded newer image for budougumi0617/mysql-sakila:5.7
a99644c0ba6e592e6ce346051d623a1079724dd809c1faa4dfdee5c1513c30c5

$ mysql -h 127.0.0.1 --port 43306 -uroot -pmysql
mysql: [Warning] Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 2
Server version: 5.7.25 MySQL Community Server (GPL)

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> set global wait_timeout=5;
Query OK, 0 rows affected (0.00 sec)

mysql> show global variables like '%wait%';
+---------------------------------------------------+----------+
| Variable_name                                     | Value    |
+---------------------------------------------------+----------+
| innodb_lock_wait_timeout                          | 50       |
| innodb_spin_wait_delay                            | 6        |
| lock_wait_timeout                                 | 31536000 |
| performance_schema_events_waits_history_long_size | 10000    |
| performance_schema_events_waits_history_size      | 10       |
| wait_timeout                                      | 5        |
+---------------------------------------------------+----------+
6 rows in set (0.01 sec)

mysql> show database;
ERROR 1064 (42000): You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'database' at line 1
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sakila             |
| sys                |
+--------------------+
5 rows in set (0.01 sec)

mysql> use sakila;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> show tables;
+----------------------------+
| Tables_in_sakila           |
+----------------------------+
| actor                      |
| actor_info                 |
| address                    |
| category                   |
| city                       |
| country                    |
| customer                   |
| customer_list              |
| film                       |
| film_actor                 |
| film_category              |
| film_list                  |
| film_text                  |
| inventory                  |
| language                   |
| nicer_but_slower_film_list |
| payment                    |
| rental                     |
| sales_by_film_category     |
| sales_by_store             |
| staff                      |
| staff_list                 |
| store                      |
+----------------------------+
23 rows in set (0.01 sec)
```



"""