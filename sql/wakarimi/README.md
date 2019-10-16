# Use PostgresQL

```bash
$ docker run --name pogre12 -p 5432:5432 -e POSTGRES_USER=dev -e POSTGRES_PASSWORD=secret -v $(pwd)/sqlfiles:/docker-entrypoint-initdb.d postgres:12
$ PGPASSWORD=secret psql -h localhost -U dev
```
