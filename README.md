Clickhouse database engine mysql
----

# RUN
```bash
# start server
$ docker-compose up -d
# insert data 
$ go run main.go
# start clickhouse client
$ docker-compose exec chs bash
$ clickhouse-client
# exchanging data with the MySQL server
> CREATE DATABASE mysql_db ENGINE = MySQL('db:3306', 'test', 'docker', 'docker')
> SHOW DATABASES
> SHOW TABLES FROM mysql_db
> SELECT * FROM mysql_db.logs