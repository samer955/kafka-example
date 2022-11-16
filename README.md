# Kafka message exchange
Write and Read messages published in topics using Kafka.

First start the docker compose file with `docker compose up -d`. Kafka will be running on localhost:9092

Start the consumer:
`kafka-example\kafka-producer> go run main.go`

This will be the output : 

`*** Connected to Kafka on localhost:9092, messages will be published on test_topic ***

*** Enter a message: ***`



` which will be listening for incoming messages. 

Start the producer with `` -> Write messages in the console. The messages will be published in the topic and can be read in the consumer console.
