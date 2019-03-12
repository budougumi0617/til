.PHONY: start stop

start:
	docker container run --rm -d -e MYSQL_ROOT_PASSWORD=mysql \
		-p 43306:3306 --name mysql_til mysql:8.0

stop:
	docker container stop mysql_til

