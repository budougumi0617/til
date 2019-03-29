.PHONY: setup init history status up mysql.start mysql.stop

setup:
	@if [ -z `which sql-migrate 2> /dev/null` ]; then \
		go get github.com/rubenv/sql-migrate/...; \
	fi

init:
	mysql -h 127.0.0.1 --port 43306 -uroot < database.sql

status: setup
	sql-migrate status

up: setup
	sql-migrate up

hisotry:
	mysql -h 127.0.0.1 --port 43306 -uroot -D sqlx_development -e "SELECT * FROM gorp_migrations;"

mysql.start:
	docker run --rm -d -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -p 43306:3306 --name sqlx_mysql mysql:5.7

mysql.stop:
	docker stop sqlx_mysql
