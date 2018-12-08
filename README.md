# vis-data-query

This repository contains any micro-services that provide data that is used by the front end services. There may also be other use cases, and it is not decided yet to which layers usage should be constrained.

+ All these micro-services are written using the `Go` programming language. https://golang.org/pkg/database/sql/
+ When connecting to SQL databases the the `sql` standard library will be used.
    + It is provided by the Go Authors.
    + It will abstract away complexities such as connection pool management.
    + It is well tested and robust and follows best practice as per the the Go Language idioms recommended by the Go Authors.
    + It supports a vast plethora of datastores - https://github.com/golang/go/wiki/SQLDrivers
         + Should new drivers be needed, that path to adding one myself is clear.
         + at the beginning on the Postgres driver is required.
         + the drivers that officially are included and pass the necessary compatibility tests for postgres will be used (https://godoc.org/github.com/lib/pq)

The following are the lists of microservices in this repository:

+ simple



# simple 

`simple` is a simple microservice that when provided with a datastore identifier and a table name returns all the data in that table in csv format.

# dependencies

## post gres drivers

+ `go get github.com/lib/pq`