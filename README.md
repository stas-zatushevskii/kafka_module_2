# Message Filtering Service

## Project Overview

This project processes Kafka message streams and applies user-specific filtering rules.

The application is built around **Goka processors** that consume messages from Kafka topics, update user filter state, and produce processed messages into output topics.

## Project Structure

### `internal/app/constants`

Contains constant values used across the application, such as:

- Kafka topic names
- Goka group names

This package is the single source of truth for stream and processor naming.

### `internal/app/domain`

Contains domain models used in Kafka messages and processor state.

Examples include:

- input message models
- command models
- user filter models
- state models stored in Goka tables

In short, this package defines the data structures passed through Kafka topics and used during stream processing.

### `internal/app/processors`

Contains Goka processor functions that implement the stream-processing logic.

These processors are responsible for:

- reading messages from Kafka topics
- reading and updating persisted state
- applying filtering rules
- producing transformed messages to output topics

---

## Running and Testing the Project

Before starting the application, prepare the required Kafka topics.

### 1. Create Kafka topics

You need to create the following 4 topics:

- `messages`
- `messages.filtered`
- `messages.blocked`
- `users.blocked`

Use the following commands:

```bash
docker exec -i kafka_module_2-kafka-1 kafka-topics --create --if-not-exists --bootstrap-server localhost:9092 --topic messages
docker exec -i kafka_module_2-kafka-1 kafka-topics --create --if-not-exists --bootstrap-server localhost:9092 --topic messages.filtered
docker exec -i kafka_module_2-kafka-1 kafka-topics --create --if-not-exists --bootstrap-server localhost:9092 --topic messages.blocked
docker exec -i kafka_module_2-kafka-1 kafka-topics --create --if-not-exists --bootstrap-server localhost:9092 --topic users.blocked
```

---

## 2. Test Commands

### 2.1 Create filters

#### Block a user

This command creates a filter for user `1` and blocks messages from user `0`.

```bash
echo '1:{"block_users":{"user_id":1,"block_user_ids":[0]}}' | docker exec -i kafka_module_2-kafka-1 kafka-console-producer --bootstrap-server localhost:9092 --topic users.blocked --property parse.key=true --property key.separator=:
```

#### Block a word

This command creates a filter for user `1` and blocks the word `Telegram`.

```bash
echo '1:{"block_words":{"user_id":1,"block_words":["Telegram"]}}' | docker exec -i kafka_module_2-kafka-1 kafka-console-producer --bootstrap-server localhost:9092 --topic messages.blocked --property parse.key=true --property key.separator=:
```

---

### 2.2 Send test messages

#### Message from a blocked user

This message is sent from user `0` to user `1`.

Since user `1` has blocked user `0`, this message must be filtered out.

```bash
echo '0:{"user_id":0,"recipient_id":1,"message":"Hello","timestamp":"2026-03-26T10:00:00Z"}' | docker exec -i kafka_module_2-kafka-1 kafka-console-producer --bootstrap-server localhost:9092 --topic messages --property parse.key=true --property key.separator=:
```

#### Message containing a blocked word

This message is sent to user `1` and contains the blocked word `Telegram`.

It should be filtered according to the blocked words rules.

```bash
echo '0:{"user_id":0,"recipient_id":1,"message":"Telegram","timestamp":"2026-03-26T10:00:00Z"}' | docker exec -i kafka_module_2-kafka-1 kafka-console-producer --bootstrap-server localhost:9092 --topic messages --property parse.key=true --property key.separator=:
```

#### Message that should pass all filters

This message is sent from user `2` to user `1`, does not contain blocked words, and is not blocked by sender rules.

It should be delivered to the `messages.filtered` topic.

```bash
echo '2:{"user_id":2,"recipient_id":1,"message":"Hello","timestamp":"2026-03-26T10:00:00Z"}' | docker exec -i kafka_module_2-kafka-1 kafka-console-producer --bootstrap-server localhost:9092 --topic messages --property parse.key=true --property key.separator=:
```

---

## Expected Behavior

- Messages from blocked users must not be delivered.
- Messages containing blocked words must be filtered.
- Valid messages must be published to `messages.filtered`.