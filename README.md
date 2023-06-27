# mosha-author-service

Author microservice used in Mosha.

## Database

The main database used in the service is ScyllaDB. It is a Cassandra compatible database. The database is used to store
the authors and their information. To run the database using docker:

```bash
docker run --name scylla --rm -it -p 9042:9042 -d scylladb/scylla --smp 2  
```

It's possible to verify the status of the database using the following command:

```bash
docker exec -it scylla nodetool status
```

Before using it's necessary to create a keyspace otherwise it will not be possible to connect to the database. To create
the keyspace enter the `cqlsh` using the following command:

```bash
docker exec -it scylla cqlsh
```

In the `cqlsh` enter the following command to create the keyspace:

```cql
CREATE KEYSPACE authorservice WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1}; 
```