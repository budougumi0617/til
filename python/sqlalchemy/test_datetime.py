"""confirm how to use datetime in sqlalchemy."""

from sqlalchemy.engine import Engine, create_engine
from sqlalchemy.sql import table, insert, column, delete
from sqlalchemy.schema import Table
from datetime import datetime, timezone, timedelta
import pytest
import os


def new_engine() -> Engine:
    return create_engine(
        "mysql+mysqldb://{user}:{password}@{host}:{port}/{database}".format(
            user=os.environ["MYSQL_DB_USER"],
            database=os.environ["MYSQL_DB_DATABASE"],
            host=os.environ["MYSQL_DB_HOST"],
            port=os.environ["MYSQL_DB_PORT"],
            password=os.environ["MYSQL_DB_PASSWORD"],
        ),
        # SQLの状態を見たかったらログレベルを上げる。
        # pytest --log-cli-level=DEBUG test_datetime.py
        # echoをTrueにしているならば、sオプションでOK
        # pytest -vs test_datetime.py
        echo=True,
    )


def test_fixture(filled_engine: Engine):
    eng = filled_engine
    rsp = eng.execute('select * from til_users')
    assert rsp.rowcount == 2


@pytest.fixture(scope='session', autouse=True)
def create_tables():
    """
    Create tables.

    :return:
    """

    q = """CREATE TABLE IF NOT EXISTS til_users
(
    id int NOT NULL AUTO_INCREMENT,
    username varchar(255),
    email varchar(255),
    password char(30),
    created datetime NOT NULL,
    modified datetime NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
"""
    eng = new_engine()
    eng.execute(q)


@pytest.fixture(scope='function')
def filled_engine():
    """
    filled_engine manages database records for each test.

    データベースを利用してテストする際に利用するfixture。
    テストケース実行前にレコードを保存し、実行後に保存したレコードを全て削除する。
    :return:
    """
    # テストのたびにデータベースに事前データを登録する。
    engine = new_engine()
    jst = timezone(timedelta(hours=9), 'JST')
    # 端数が出ているとwhere句での検索が大変厳しいので、決め打ちの値にしておく。
    now = datetime(2019, 12, 1, 11, 30, tzinfo=jst)

    # prepare users
    til_users = til_users_tables()
    users = preset_til_users(now)
    engine.execute(insert(til_users), users)

    # テストを実施する。
    yield engine

    # テストのたびにprepare以降に作成したデータを削除する。
    stmt = delete(til_users).where(column("created").__eq__(now))
    engine.execute(stmt)


def til_users_tables() -> Table:
    """
    Define til_users table.

    :return: from sqlalchemy.schema.Table
    """
    return table(
        'til_users',
        column('id'),
        column('username'),
        column('email'),
        column('password'),
        column('created'),
        column('modified'),
    )


def preset_til_users(now: datetime):
    """
    Return seed data.

    :param now: レコードのcreated, modifiedの時刻として利用する。
    :return: list
    """
    return [
        {
            'id': 10001,
            'username': "user01",
            'email': "user01@example.com",
            'password': 'p@ssword',
            'created': now,
            'modified': now,
        },
        {
            'id': 10002,
            'username': "user02",
            'email': "user02@example.com",
            'password': 'foo-var',
            'created': now,
            'modified': now,
        },
    ]
