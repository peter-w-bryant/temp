
# Kafka and Zookeeper with Docker

This guide provides instructions for setting up a Kafka server with Zookeeper using Docker and Docker Compose.

## Prerequisites

- Docker
- Docker Compose

Ensure Docker and Docker Compose are installed on your system.

## Setup

1. Create a `docker-compose.yml` file with the following content:

    ```yaml
    version: '3'
    services:
      zookeeper:
        image: wurstmeister/zookeeper
        ports:
          - "2181:2181"
      kafka:
        image: wurstmeister/kafka
        ports:
          - "9092:9092"
        environment:
          KAFKA_ADVERTISED_HOST_NAME: kafka
          KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
        volumes:
          - /var/run/docker.sock:/var/run/docker.sock
        depends_on:
          - zookeeper
    ```

2. Run the Docker Compose command to start Kafka and Zookeeper:

    ```bash
    docker-compose up -d
    ```

## Testing the Setup

1. **Check Container Status**

    Check the status of the containers using:

    ```bash
    docker ps
    ```

2. **Log Monitoring**

    Check the logs for Kafka and Zookeeper:

    ```bash
    docker logs [container_name_or_id]
    ```

3. **Create a Kafka Topic**

    Create a topic named `test-topic`:

    ```bash
    docker exec -it [kafka_container_name_or_id] kafka-topics.sh --create --topic test-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
    ```

4. **Produce and Consume Messages**

    - Produce a message:

      ```bash
      docker exec -it [kafka_container_name_or_id] kafka-console-producer.sh --broker-list localhost:9092 --topic test-topic
      ```

    - Consume the message:

      ```bash
      docker exec -it [kafka_container_name_or_id] kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test-topic --from-beginning
      ```

5. **Check Zookeeper Logs**

    ```bash
    docker logs [zookeeper_container_name_or_id]
    ```

## Stopping the Services

To stop and remove the containers, run:

```bash
docker-compose down
```
