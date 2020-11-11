# how to reproduce the bug?

## first
**prepare data**
```
cd data
go run 1mEdge.go 1 3000000
```
## second
**set up dgraph**

for version one bugs, use docker-compose.yml or `v1/dgraph_v20.07.2` to setup the dgraph db.

```
cd v1
./dgraph_v20.07.2 zero 
./dgraph_v20.07.2 alpha
```

for version two bugs, use `v2/dgraph_release_v20.07.1` to setup the dgraph db.
I`make` this dgraph bin based on the `release/v20.07.1` branch.

## third
**live load the data**
v1 bugs load
```
./dgraph_v20.07.2 live -f ../data/1__3000000.rdf 
```

v2 bugs load
```
./dgraph_release_v20.07.1 live -f ../data/1__3000000.rdf 

```

## fourth
**find the bug**
```
cd code
go run q.go
```


