# Kafka Protobuf Console
Kafka Protobuf Console is a CLI which provides the following functionalities:
* Produce protobuf message to Kafka using Json input
* Consume protobuf message from Kafka and displayed the message in Json format
* Produce protobuf messages continuously to Kafka with random field values (WIP)

## Installation
### From source
1. Clone the repository.
2. The Golang version installed should be higher than v1.11.
2. Run ```make install```.

### Compiled binary
WIP

## Quick Start
Example scripts can be found within the ```examples``` directory. For convenience, a docker compose file has also been included so that the user can try using the console with minimal set up.

## Kafka Protobuf Console Producer
### Usage
```go
Usage:
  kafka-protobuf-console produce [flags]

Flags:
  -d, --descriptor string   File descriptor path
  -h, --help                help for produce
  -n, --name string         Fully qualified Proto message name
  -t, --topic string        Destination Kafka topic

Global Flags:
  -b, --brokers strings   Comma separated Kafka brokers address
```

## Kafka Protobuf Console Consumer
### Usage