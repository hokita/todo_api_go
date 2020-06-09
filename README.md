# todo_api_go
## setup
```shell
$ docker-compose build
$ docker-compose up
```

## Migrateion commands
### new
```shell
$ docker-compose exec web sql-migrate new --config config/db/dbconfig.yml create_[tablename]
```

### up
```shell
$ docker-compose exec web sql-migrate up --config config/db/dbconfig.yml --env "development"
```
