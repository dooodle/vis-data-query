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

`simple` is a simple microservice that when provided with a datastore identifier and a table name returns all the data in that table in a simple csv format in which *all* values are returned as strings, except nulls.

for example.

```
"Albania","AL","Tirana","Albania","28750.00","2800138"
"Greece","GR","Athina","Attikis","131940.00","10816286"
"Macedonia","MK","Skopje","Macedonia","25333.00","2059794"
"Serbia","SRB","Beograd","Serbia","77474.00","7120666"
"Montenegro","MNE","Podgorica","Montenegro","14026.00","620029"
"Kosovo","KOS","Prishtine","Kosovo","10887.00","1733872"
"Andorra","AND","Andorra la Vella","Andorra","450.00","78115"
"France","F","Paris","Île-de-France","547030.00","64933400"
"Spain","E","Madrid","Madrid","504750.00","46815916"
"Austria","A","Wien","Wien","83850.00","8499759"
"Czech Republic","CZ","Praha","Praha","78703.00","10562214"
"Germany","D","Berlin","Berlin","356910.00","80219695"
"Hungary","H","Budapest","Budapest","93030.00","9937628"
"Italy","I","Roma","Lazio","301230.00","59433744"
"Liechtenstein","FL","Vaduz","Liechtenstein","160.00","36636"
"Cyprus","CY","Lefkosia","Cyprus","9251.00","840407"
"Gaza Strip","GAZA",,,"365.00","1760037"
```

# dependencies

## postgres drivers

+ `go get github.com/lib/pq`