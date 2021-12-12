# bookstore_oauth-api
OAuth API with DDD

### Set database for the project
Create a file `docker-compose.yml` with the content:
```yaml
version: '3'

services:
  cassandra-1:
    image: cassandra:latest
    container_name: cassandra-1
    ports:
      - 9042:9042
    environment:
      - "MAX_HEAP_SIZE=256M"
      - "HEAP_NEWSIZE=128M"
      - "CASSANDRA_SEEDS=cassandra-1,cassandra-2"

  cassandra-2:
    image: cassandra:latest
    container_name: cassandra-2
    environment:
      - "MAX_HEAP_SIZE=256M"
      - "HEAP_NEWSIZE=128M"
      - "CASSANDRA_SEEDS=cassandra-1,cassandra-2"
    depends_on:
      - cassandra-1

  cassandra-3:
    image: cassandra:latest
    container_name: cassandra-3
    environment:
      - "MAX_HEAP_SIZE=256M"
      - "HEAP_NEWSIZE=128M"
      - "CASSANDRA_SEEDS=cassandra-1,cassandra-2"
    depends_on:
      - cassandra-2
```

Access to container `cql` session:
```shell
docker exec -it cassandra-1 cqlsh
```

Create new keyspace:
```shell
CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy','replication_factor':1};
```

Check the new keyspace:
```shell
describe keyspaces;
```

Use the new keyspace:
```shell
USE oauth;
```

Describe tables:
```shell
describe tables;
```

Create `access_tokens` table:
```cql
CREATE TABLE access_tokens(access_token VARCHAR PRIMARY KEY, user_id BIGINT, client_id BIGINT, expires BIGINT);
```

```cql
SELECT * FROM access_tokens;
```
