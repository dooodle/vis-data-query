# vis-data-query

This repository contains any micro-services that provide data that is used by the front end services. There may also be other use cases, and it is not decided yet to which layers usage should be constrained.

+ All these micro-services are written using the `Go` programming language. https://golang.org/pkg/database/sql/
+ When connecting to SQL databases the `sql` standard library will be used.
    + It is provided by the Go Authors.
    + It will abstract away complexities such as connection pool management.
    + It is well tested and robust and follows best practice as per the the Go Language idioms recommended by the Go Authors.
    + It supports a vast plethora of datastores - https://github.com/golang/go/wiki/SQLDrivers
         + Should new drivers be needed, that path to adding one myself is clear.
         + At the beginning on the Postgres driver is required.
         + The drivers that officially are included and also pass the necessary compatibility tests for postgres will be used (https://godoc.org/github.com/lib/pq)

The following are the lists of microservices in this repository:

+ simple

# simple 

+ `simple` is a simple REST microservice that when provided with a datastore identifier and a table name returns all the data in that table in a simple csv format in which *all* values are returned as strings, except nulls.
+ `http://<host>/mondial/<relation-name>`
    + `http://<host>/mondial/country`
```
"Albania","AL","Tirana","Albania","28750.00","2800138"
"Greece","GR","Athina","Attikis","131940.00","10816286"
"Macedonia","MK","Skopje","Macedonia","25333.00","2059794"
```
    + `http://<host>/mondial/country?h=true`

```
name,code,capital,province,area,population
"Albania","AL","Tirana","Albania","28750.00","2800138"
"Greece","GR","Athina","Attikis","131940.00","10816286"
"Macedonia","MK","Skopje","Macedonia","25333.00","2059794"
```
+ `http://<host>/mondial/names`. This provides a list of all tables available, this will URL path will change as will how the list is returned.

```
"borders"
"city"
"geo_estuary"
"country_local_name"
"city_local_name"
"city_other_name"
"country_other_name"
"citylocal_name"
"cityother_name"
"city_population"
```

# dependencies

## postgres drivers

+ `go get github.com/lib/pq`