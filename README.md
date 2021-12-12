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
