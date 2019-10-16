# Use PostgresQL

```bash
$ docker run --name pogre12 -p 5432:5432 -e POSTGRES_USER=dev -e POSTGRES_PASSWORD=secret -v $(pwd)/sqlfiles:/docker-entrypoint-initdb.d postgres:12
$ PGPASSWORD=secret psql -h localhost -U dev
```

# PostgresQLのコマンドの参考資料
- MySQLとPostgreSQLコマンド比較表
  - https://qiita.com/aosho235/items/c657e2fcd15fa0647471
