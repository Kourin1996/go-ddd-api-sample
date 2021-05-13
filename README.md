# go-ddd-api-sample

A sample code of REST API with Domain Driven Development

* echo
* go-pg

## How to run

```
git clone git@github.com:Kourin1996/go-crud-api-sample.git

```

## How to start

```
go build ./api/main.go
export POSTGRESQL_URL='postgres://postgres:postgres@localhost:5432/test?sslmode=disable'
POSTGRES_URL=postgres://postgres:postgres@localhost:5432/test?sslmode=disable migrate -database ${POSTGRESQL_URL} -path migrations up
./main
```
