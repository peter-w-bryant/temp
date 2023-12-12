
# Kafka Commands and Flag Descriptions

## 1. Create a Topic:
```bash
kafka-topics --create --topic your_topic_name --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
```
- `--create`: Indicates that you want to create a topic.
- `--topic`: Specifies the name of the topic.
- `--bootstrap-server`: Specifies the Kafka server to connect to.
- `--partitions`: The number of partitions for the topic.
- `--replication-factor`: The replication factor for the topic.

## 2. List Created Topics:
```bash
kafka-topics --list --bootstrap-server localhost:9092
```
- `--list`: Lists all available topics.
- `--bootstrap-server`: Specifies the Kafka server to connect to.

## 3. Produce a Message:
```bash
echo "Hello, Kafka!" | kafka-console-producer --broker-list localhost:9092 --topic your_topic_name
```
- `echo "Hello, Kafka!"`: This is the message you're sending.
- `kafka-console-producer`: Kafka's command-line tool to produce messages.
- `--broker-list`: Specifies the Kafka server to connect to.
- `--topic`: Specifies the topic to send the message to.

## 4. Consume a Message:
```bash
kafka-console-consumer --bootstrap-server localhost:9092 --topic your_topic_name --from-beginning
```
- `kafka-console-consumer`: Kafka's command-line tool to consume messages.
- `--bootstrap-server`: Specifies the Kafka server to connect to.
- `--topic`: Specifies the topic to consume messages from.
- `--from-beginning`: Consumes messages from the beginning of the topic.

## 5. Delete a Topic:
```bash
kafka-topics --delete --topic your_topic_name --bootstrap-server localhost:9092
```
- `--delete`: Indicates that you want to delete a topic.
- `--topic`: Specifies the name of the topic.
- `--bootstrap-server`: Specifies the Kafka server to connect to.
