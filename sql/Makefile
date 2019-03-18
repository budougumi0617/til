.PHONY: start stop mysql

start:
	docker container run --rm -d -e MYSQL_ALLOW_EMPTY_PASSWORD=yes \
		-p 43306:3306 --name mysql_til kakakakakku/mysql-world-database:8.0

stop:
	docker container stop mysql_til

mysql:
	mysql -h 127.0.0.1 --port 43306 -uroot -D world

