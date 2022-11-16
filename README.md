# Kafka message exchange
Write and Read messages published in topics using Kafka.

First start the docker compose file with `docker compose up -d`. Kafka will be running on localhost:9092

Start the consumer:
`kafka-example\kafka-consumer> go run main.go`

It will be listening for incoming messages. 

Start the producer:
`kafka-example\kafka-producer> go run main.go`

Write messages in the producer console. The messages will be published in the topic and can be read in the consumer console.
