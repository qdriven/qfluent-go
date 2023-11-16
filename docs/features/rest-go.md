# rest-go

## installation

```shell
go install github.com/rest-go/rest
rest -db.url "mysql://username:password@tcp(localhost:3306)/db"
rest -addr :3000 -db.url sqlite://chinook.db
```

## Configurations

```shell
rest -config config.yml
```

```
# server listen addr
addr: :9090

# db config
db:
  url: sqlite://chinook.db

# auth config
auth:
  enabled: true
  secret: [replace-this-to-your-own-secret]
```

```shell
Database	URL Format
PostgreSQL	postgres://user:passwd@localhost:5432/db_name?search_path=schema_name
MySQL	mysql://user:passwd@tcp(127.0.0.1:3306)/db_name
SQLite	sqlite:///path/to/my.db
```

## API

[docs](https://rest-go.com/docs/guides/api)